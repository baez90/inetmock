//go:generate go-enum -f $GOFILE --lower --marshal --names

package endpoint

import (
	"errors"
	"net"
	"strings"

	"go.uber.org/multierr"
)

/* ENUM(
UDP,
TCP
)
*/
type NetProto int

type Uplink struct {
	Proto      NetProto
	Listener   net.Listener
	PacketConn net.PacketConn
}

func (u Uplink) Addr() net.Addr {
	if u.Listener != nil {
		return u.Listener.Addr()
	}
	if u.PacketConn != nil {
		return u.PacketConn.LocalAddr()
	}
	return nil
}

func (u Uplink) Close() (err error) {
	if u.Listener != nil {
		err = multierr.Append(err, u.Listener.Close())
	}
	if u.PacketConn != nil {
		err = multierr.Append(err, u.PacketConn.Close())
	}
	return
}

type ListenerSpec struct {
	Name      string
	Protocol  string
	Address   string `mapstructure:"listenAddress"`
	Port      uint16
	Endpoints map[string]Spec
}

func (l ListenerSpec) Uplink() (uplink Uplink, err error) {
	switch l.Protocol {
	case "udp", "udp4", "udp6":
		uplink.Proto = NetProtoUDP
		uplink.PacketConn, err = net.ListenUDP(l.Protocol, &net.UDPAddr{
			IP:   net.ParseIP(l.Address),
			Port: int(l.Port),
		})
	case "tcp", "tcp4", "tcp6":
		uplink.Proto = NetProtoTCP
		uplink.Listener, err = net.ListenTCP(l.Protocol, &net.TCPAddr{
			IP:   net.ParseIP(l.Address),
			Port: int(l.Port),
		})
	default:
		err = errors.New("protocol not supported")
	}
	return
}

type HandlerReference string

func (h HandlerReference) ToLower() HandlerReference {
	return HandlerReference(strings.ToLower(string(h)))
}

type Spec struct {
	HandlerRef HandlerReference `mapstructure:"handler"`
	TLS        bool
	Handler    ProtocolHandler `mapstructure:"-"`
	Options    map[string]interface{}
}
