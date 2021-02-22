// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: audit/v1/dns_details.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DNSOpCode int32

const (
	//buf:lint:ignore ENUM_ZERO_VALUE_SUFFIX
	DNSOpCode_DNS_OP_CODE_QUERY  DNSOpCode = 0
	DNSOpCode_DNS_OP_CODE_STATUS DNSOpCode = 2
	DNSOpCode_DNS_OP_CODE_NOTIFY DNSOpCode = 4
	DNSOpCode_DNS_OP_CODE_UPDATE DNSOpCode = 5
)

// Enum value maps for DNSOpCode.
var (
	DNSOpCode_name = map[int32]string{
		0: "DNS_OP_CODE_QUERY",
		2: "DNS_OP_CODE_STATUS",
		4: "DNS_OP_CODE_NOTIFY",
		5: "DNS_OP_CODE_UPDATE",
	}
	DNSOpCode_value = map[string]int32{
		"DNS_OP_CODE_QUERY":  0,
		"DNS_OP_CODE_STATUS": 2,
		"DNS_OP_CODE_NOTIFY": 4,
		"DNS_OP_CODE_UPDATE": 5,
	}
)

func (x DNSOpCode) Enum() *DNSOpCode {
	p := new(DNSOpCode)
	*p = x
	return p
}

func (x DNSOpCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DNSOpCode) Descriptor() protoreflect.EnumDescriptor {
	return file_audit_v1_dns_details_proto_enumTypes[0].Descriptor()
}

func (DNSOpCode) Type() protoreflect.EnumType {
	return &file_audit_v1_dns_details_proto_enumTypes[0]
}

func (x DNSOpCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DNSOpCode.Descriptor instead.
func (DNSOpCode) EnumDescriptor() ([]byte, []int) {
	return file_audit_v1_dns_details_proto_rawDescGZIP(), []int{0}
}

type ResourceRecordType int32

const (
	ResourceRecordType_RESOURCE_RECORD_TYPE_UNSPECIFIED ResourceRecordType = 0
	ResourceRecordType_RESOURCE_RECORD_TYPE_A           ResourceRecordType = 1
	ResourceRecordType_RESOURCE_RECORD_TYPE_NS          ResourceRecordType = 2
	ResourceRecordType_RESOURCE_RECORD_TYPE_CNAME       ResourceRecordType = 5
	ResourceRecordType_RESOURCE_RECORD_TYPE_SOA         ResourceRecordType = 6
	ResourceRecordType_RESOURCE_RECORD_TYPE_PTR         ResourceRecordType = 12
	ResourceRecordType_RESOURCE_RECORD_TYPE_HINFO       ResourceRecordType = 13
	ResourceRecordType_RESOURCE_RECORD_TYPE_MINFO       ResourceRecordType = 14
	ResourceRecordType_RESOURCE_RECORD_TYPE_MX          ResourceRecordType = 15
	ResourceRecordType_RESOURCE_RECORD_TYPE_TXT         ResourceRecordType = 16
	ResourceRecordType_RESOURCE_RECORD_TYPE_RP          ResourceRecordType = 17
	ResourceRecordType_RESOURCE_RECORD_TYPE_AAAA        ResourceRecordType = 28
	ResourceRecordType_RESOURCE_RECORD_TYPE_SRV         ResourceRecordType = 33
	ResourceRecordType_RESOURCE_RECORD_TYPE_NAPTR       ResourceRecordType = 35
)

// Enum value maps for ResourceRecordType.
var (
	ResourceRecordType_name = map[int32]string{
		0:  "RESOURCE_RECORD_TYPE_UNSPECIFIED",
		1:  "RESOURCE_RECORD_TYPE_A",
		2:  "RESOURCE_RECORD_TYPE_NS",
		5:  "RESOURCE_RECORD_TYPE_CNAME",
		6:  "RESOURCE_RECORD_TYPE_SOA",
		12: "RESOURCE_RECORD_TYPE_PTR",
		13: "RESOURCE_RECORD_TYPE_HINFO",
		14: "RESOURCE_RECORD_TYPE_MINFO",
		15: "RESOURCE_RECORD_TYPE_MX",
		16: "RESOURCE_RECORD_TYPE_TXT",
		17: "RESOURCE_RECORD_TYPE_RP",
		28: "RESOURCE_RECORD_TYPE_AAAA",
		33: "RESOURCE_RECORD_TYPE_SRV",
		35: "RESOURCE_RECORD_TYPE_NAPTR",
	}
	ResourceRecordType_value = map[string]int32{
		"RESOURCE_RECORD_TYPE_UNSPECIFIED": 0,
		"RESOURCE_RECORD_TYPE_A":           1,
		"RESOURCE_RECORD_TYPE_NS":          2,
		"RESOURCE_RECORD_TYPE_CNAME":       5,
		"RESOURCE_RECORD_TYPE_SOA":         6,
		"RESOURCE_RECORD_TYPE_PTR":         12,
		"RESOURCE_RECORD_TYPE_HINFO":       13,
		"RESOURCE_RECORD_TYPE_MINFO":       14,
		"RESOURCE_RECORD_TYPE_MX":          15,
		"RESOURCE_RECORD_TYPE_TXT":         16,
		"RESOURCE_RECORD_TYPE_RP":          17,
		"RESOURCE_RECORD_TYPE_AAAA":        28,
		"RESOURCE_RECORD_TYPE_SRV":         33,
		"RESOURCE_RECORD_TYPE_NAPTR":       35,
	}
)

func (x ResourceRecordType) Enum() *ResourceRecordType {
	p := new(ResourceRecordType)
	*p = x
	return p
}

func (x ResourceRecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceRecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_audit_v1_dns_details_proto_enumTypes[1].Descriptor()
}

func (ResourceRecordType) Type() protoreflect.EnumType {
	return &file_audit_v1_dns_details_proto_enumTypes[1]
}

func (x ResourceRecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceRecordType.Descriptor instead.
func (ResourceRecordType) EnumDescriptor() ([]byte, []int) {
	return file_audit_v1_dns_details_proto_rawDescGZIP(), []int{1}
}

type DNSQuestionEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ResourceRecordType `protobuf:"varint,1,opt,name=type,proto3,enum=inetmock.audit.v1.ResourceRecordType" json:"type,omitempty"`
	Name string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DNSQuestionEntity) Reset() {
	*x = DNSQuestionEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_dns_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSQuestionEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSQuestionEntity) ProtoMessage() {}

func (x *DNSQuestionEntity) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_dns_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSQuestionEntity.ProtoReflect.Descriptor instead.
func (*DNSQuestionEntity) Descriptor() ([]byte, []int) {
	return file_audit_v1_dns_details_proto_rawDescGZIP(), []int{0}
}

func (x *DNSQuestionEntity) GetType() ResourceRecordType {
	if x != nil {
		return x.Type
	}
	return ResourceRecordType_RESOURCE_RECORD_TYPE_UNSPECIFIED
}

func (x *DNSQuestionEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DNSDetailsEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opcode    DNSOpCode            `protobuf:"varint,1,opt,name=opcode,proto3,enum=inetmock.audit.v1.DNSOpCode" json:"opcode,omitempty"`
	Questions []*DNSQuestionEntity `protobuf:"bytes,2,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *DNSDetailsEntity) Reset() {
	*x = DNSDetailsEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_dns_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSDetailsEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSDetailsEntity) ProtoMessage() {}

func (x *DNSDetailsEntity) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_dns_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSDetailsEntity.ProtoReflect.Descriptor instead.
func (*DNSDetailsEntity) Descriptor() ([]byte, []int) {
	return file_audit_v1_dns_details_proto_rawDescGZIP(), []int{1}
}

func (x *DNSDetailsEntity) GetOpcode() DNSOpCode {
	if x != nil {
		return x.Opcode
	}
	return DNSOpCode_DNS_OP_CODE_QUERY
}

func (x *DNSDetailsEntity) GetQuestions() []*DNSQuestionEntity {
	if x != nil {
		return x.Questions
	}
	return nil
}

var File_audit_v1_dns_details_proto protoreflect.FileDescriptor

var file_audit_v1_dns_details_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x69, 0x6e,
	0x65, 0x74, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x22,
	0x62, 0x0a, 0x11, 0x44, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x65, 0x74, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x10, 0x44, 0x4e, 0x53, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x70, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x65, 0x74, 0x6d,
	0x6f, 0x63, 0x6b, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x4e, 0x53,
	0x4f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x42,
	0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x65, 0x74, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2a, 0x6a, 0x0a, 0x09, 0x44, 0x4e, 0x53, 0x4f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x44, 0x4e, 0x53, 0x5f, 0x4f, 0x50, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x4e, 0x53, 0x5f, 0x4f, 0x50,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x44, 0x4e, 0x53, 0x5f, 0x4f, 0x50, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f,
	0x54, 0x49, 0x46, 0x59, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x4e, 0x53, 0x5f, 0x4f, 0x50,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x2a, 0xc4,
	0x03, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4e, 0x53, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x41,
	0x10, 0x06, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52,
	0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x54, 0x52, 0x10, 0x0c,
	0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x0d,
	0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x0e,
	0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x58, 0x10, 0x0f, 0x12, 0x1c, 0x0a,
	0x18, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x58, 0x54, 0x10, 0x10, 0x12, 0x1b, 0x0a, 0x17, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x50, 0x10, 0x11, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x41, 0x41, 0x41, 0x10, 0x1c, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x52, 0x56, 0x10, 0x21, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x41,
	0x50, 0x54, 0x52, 0x10, 0x23, 0x42, 0x7a, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x62, 0x61, 0x65, 0x7a, 0x39, 0x30, 0x2e, 0x69, 0x6e, 0x65, 0x74, 0x6d,
	0x6f, 0x63, 0x6b, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x11, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x65, 0x74, 0x6d,
	0x6f, 0x63, 0x6b, 0x2f, 0x69, 0x6e, 0x65, 0x74, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x15, 0x49, 0x4e, 0x65, 0x74,
	0x4d, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_audit_v1_dns_details_proto_rawDescOnce sync.Once
	file_audit_v1_dns_details_proto_rawDescData = file_audit_v1_dns_details_proto_rawDesc
)

func file_audit_v1_dns_details_proto_rawDescGZIP() []byte {
	file_audit_v1_dns_details_proto_rawDescOnce.Do(func() {
		file_audit_v1_dns_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_audit_v1_dns_details_proto_rawDescData)
	})
	return file_audit_v1_dns_details_proto_rawDescData
}

var file_audit_v1_dns_details_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_audit_v1_dns_details_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_audit_v1_dns_details_proto_goTypes = []interface{}{
	(DNSOpCode)(0),            // 0: inetmock.audit.v1.DNSOpCode
	(ResourceRecordType)(0),   // 1: inetmock.audit.v1.ResourceRecordType
	(*DNSQuestionEntity)(nil), // 2: inetmock.audit.v1.DNSQuestionEntity
	(*DNSDetailsEntity)(nil),  // 3: inetmock.audit.v1.DNSDetailsEntity
}
var file_audit_v1_dns_details_proto_depIdxs = []int32{
	1, // 0: inetmock.audit.v1.DNSQuestionEntity.type:type_name -> inetmock.audit.v1.ResourceRecordType
	0, // 1: inetmock.audit.v1.DNSDetailsEntity.opcode:type_name -> inetmock.audit.v1.DNSOpCode
	2, // 2: inetmock.audit.v1.DNSDetailsEntity.questions:type_name -> inetmock.audit.v1.DNSQuestionEntity
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_audit_v1_dns_details_proto_init() }
func file_audit_v1_dns_details_proto_init() {
	if File_audit_v1_dns_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audit_v1_dns_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSQuestionEntity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_audit_v1_dns_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSDetailsEntity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_audit_v1_dns_details_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_audit_v1_dns_details_proto_goTypes,
		DependencyIndexes: file_audit_v1_dns_details_proto_depIdxs,
		EnumInfos:         file_audit_v1_dns_details_proto_enumTypes,
		MessageInfos:      file_audit_v1_dns_details_proto_msgTypes,
	}.Build()
	File_audit_v1_dns_details_proto = out.File
	file_audit_v1_dns_details_proto_rawDesc = nil
	file_audit_v1_dns_details_proto_goTypes = nil
	file_audit_v1_dns_details_proto_depIdxs = nil
}
