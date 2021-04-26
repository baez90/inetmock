// Code generated by go-enum
// DO NOT EDIT!

package endpoint

import (
	"fmt"
	"strings"
)

const (
	// NetProtoUDP is a NetProto of type UDP.
	NetProtoUDP NetProto = iota
	// NetProtoTCP is a NetProto of type TCP.
	NetProtoTCP
)

const _NetProtoName = "UDPTCP"

var _NetProtoNames = []string{
	_NetProtoName[0:3],
	_NetProtoName[3:6],
}

// NetProtoNames returns a list of possible string values of NetProto.
func NetProtoNames() []string {
	tmp := make([]string, len(_NetProtoNames))
	copy(tmp, _NetProtoNames)
	return tmp
}

var _NetProtoMap = map[NetProto]string{
	0: _NetProtoName[0:3],
	1: _NetProtoName[3:6],
}

// String implements the Stringer interface.
func (x NetProto) String() string {
	if str, ok := _NetProtoMap[x]; ok {
		return str
	}
	return fmt.Sprintf("NetProto(%d)", x)
}

var _NetProtoValue = map[string]NetProto{
	_NetProtoName[0:3]:                  0,
	strings.ToLower(_NetProtoName[0:3]): 0,
	_NetProtoName[3:6]:                  1,
	strings.ToLower(_NetProtoName[3:6]): 1,
}

// ParseNetProto attempts to convert a string to a NetProto
func ParseNetProto(name string) (NetProto, error) {
	if x, ok := _NetProtoValue[name]; ok {
		return x, nil
	}
	return NetProto(0), fmt.Errorf("%s is not a valid NetProto, try [%s]", name, strings.Join(_NetProtoNames, ", "))
}

// MarshalText implements the text marshaller method
func (x NetProto) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *NetProto) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseNetProto(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}