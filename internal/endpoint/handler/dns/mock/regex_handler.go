package mock

import (
	"net"

	"github.com/miekg/dns"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/inetmock/inetmock/pkg/audit"
	"gitlab.com/inetmock/inetmock/pkg/audit/details"
	"gitlab.com/inetmock/inetmock/pkg/logging"
	"go.uber.org/zap"
)

type regexHandler struct {
	handlerName  string
	routes       []resolverRule
	fallback     ResolverFallback
	auditEmitter audit.Emitter
	logger       logging.Logger
}

func (rh *regexHandler) AddRule(rule resolverRule) {
	rh.routes = append(rh.routes, rule)
}

func (rh *regexHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	timer := prometheus.NewTimer(requestDurationHistogram.WithLabelValues(rh.handlerName))
	defer func() {
		timer.ObserveDuration()
	}()

	rh.recordRequest(r, w.LocalAddr(), w.RemoteAddr())

	m := new(dns.Msg)
	m.Compress = false
	m.SetReply(r)

	if r.Opcode == dns.OpcodeQuery {
		rh.handleQuery(m)
	}
	if err := w.WriteMsg(m); err != nil {
		rh.logger.Error(
			"Failed to write DNS response message",
			zap.Error(err),
		)
	}
}

func (rh *regexHandler) handleQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			totalHandledRequestsCounter.WithLabelValues(rh.handlerName).Inc()
			for _, rule := range rh.routes {
				if !rule.pattern.MatchString(q.Name) {
					continue
				}
				m.Authoritative = true
				answer := &dns.A{
					Hdr: dns.RR_Header{
						Name:   q.Name,
						Rrtype: dns.TypeA,
						Class:  dns.ClassINET,
						Ttl:    60,
					},
					A: rule.response,
				}
				m.Answer = append(m.Answer, answer)
				rh.logger.Info(
					"matched DNS rule",
					zap.String("pattern", rule.pattern.String()),
					zap.String("response", rule.response.String()),
				)
				return
			}
			rh.handleFallbackForMessage(m, q)
		default:
			unhandledRequestsCounter.WithLabelValues(rh.handlerName).Inc()
			rh.logger.Warn(
				"Unhandled DNS question type - no response will be sent",
				zap.Uint16("question_type", q.Qtype),
			)
		}
	}
}

func (rh *regexHandler) handleFallbackForMessage(m *dns.Msg, q dns.Question) {
	fallbackIP := rh.fallback.GetIP()
	answer := &dns.A{
		Hdr: dns.RR_Header{
			Name:   q.Name,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    60,
		},
		A: fallbackIP,
	}
	rh.logger.Info(
		"Falling back to generated IP",
		zap.String("response", fallbackIP.String()),
	)
	m.Authoritative = true
	m.Answer = append(m.Answer, answer)
}

func (rh *regexHandler) recordRequest(m *dns.Msg, localAddr, remoteAddr net.Addr) {
	dnsDetails := &details.DNS{
		OPCode: details.DNSOpCode(m.Opcode),
	}

	for _, q := range m.Question {
		dnsDetails.Questions = append(dnsDetails.Questions, details.DNSQuestion{
			RRType: details.ResourceRecordType(q.Qtype),
			Name:   q.Name,
		})
	}

	ev := audit.Event{
		Transport:       guessTransportFromAddr(localAddr),
		Application:     audit.AppProtocol_DNS,
		ProtocolDetails: dnsDetails,
	}

	ev.SetSourceIPFromAddr(remoteAddr)
	ev.SetDestinationIPFromAddr(localAddr)

	rh.auditEmitter.Emit(ev)
}

func guessTransportFromAddr(addr net.Addr) audit.TransportProtocol {
	switch addr.(type) {
	case *net.TCPAddr:
		return audit.TransportProtocol_TCP
	case *net.UDPAddr:
		return audit.TransportProtocol_UDP
	default:
		return audit.TransportProtocol_UNKNOWN_TRANSPORT
	}
}
