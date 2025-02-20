// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: limiter_module.proto

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Limiter_RateLimitBackend int32

const (
	Limiter_netEaseLocalFlowControl Limiter_RateLimitBackend = 0
	Limiter_envoyLocalRateLimit     Limiter_RateLimitBackend = 1
)

var Limiter_RateLimitBackend_name = map[int32]string{
	0: "netEaseLocalFlowControl",
	1: "envoyLocalRateLimit",
}

var Limiter_RateLimitBackend_value = map[string]int32{
	"netEaseLocalFlowControl": 0,
	"envoyLocalRateLimit":     1,
}

func (x Limiter_RateLimitBackend) String() string {
	return proto.EnumName(Limiter_RateLimitBackend_name, int32(x))
}

func (Limiter_RateLimitBackend) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4827d40f7d98bcf0, []int{0, 0}
}

type Limiter struct {
	Backend Limiter_RateLimitBackend `protobuf:"varint,3,opt,name=backend,proto3,enum=slime.microservice.limiter.config.Limiter_RateLimitBackend" json:"backend,omitempty"`
	Refresh *time.Duration           `protobuf:"bytes,4,opt,name=refresh,proto3,stdduration" json:"refresh,omitempty"`
	// it will disable GlobalRateLimit
	DisableGlobalRateLimit bool `protobuf:"varint,5,opt,name=disableGlobalRateLimit,proto3" json:"disableGlobalRateLimit,omitempty"`
	// it will disable use promql
	DisableAdaptive    bool `protobuf:"varint,6,opt,name=disableAdaptive,proto3" json:"disableAdaptive,omitempty"`
	EnableServiceEntry bool `protobuf:"varint,7,opt,name=enableServiceEntry,proto3" json:"enableServiceEntry,omitempty"`
	// it will not generate envoy.filters.http.local_ratelimit insert before http.router when true
	DisableInsertLocalRateLimit bool `protobuf:"varint,8,opt,name=disableInsertLocalRateLimit,proto3" json:"disableInsertLocalRateLimit,omitempty"`
	// it will not generate envoy.filters.http.ratelimit insert before http.router when true
	DisableInsertGlobalRateLimit bool `protobuf:"varint,9,opt,name=disableInsertGlobalRateLimit,proto3" json:"disableInsertGlobalRateLimit,omitempty"`
	// specify the rls namespaces
	RlsConfigMap *RlsConfigMap `protobuf:"bytes,10,opt,name=rlsConfigMap,proto3" json:"rlsConfigMap,omitempty"`
	// specify domain, it is useful in global ratelimiter
	Domain string `protobuf:"bytes,11,opt,name=domain,proto3" json:"domain,omitempty"`
	// specify rls server address, if disableInsertGlobalRateLimit if false
	Rls                  *RateLimitService `protobuf:"bytes,12,opt,name=rls,proto3" json:"rls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Limiter) Reset()         { *m = Limiter{} }
func (m *Limiter) String() string { return proto.CompactTextString(m) }
func (*Limiter) ProtoMessage()    {}
func (*Limiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_4827d40f7d98bcf0, []int{0}
}
func (m *Limiter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Limiter.Unmarshal(m, b)
}
func (m *Limiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Limiter.Marshal(b, m, deterministic)
}
func (m *Limiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Limiter.Merge(m, src)
}
func (m *Limiter) XXX_Size() int {
	return xxx_messageInfo_Limiter.Size(m)
}
func (m *Limiter) XXX_DiscardUnknown() {
	xxx_messageInfo_Limiter.DiscardUnknown(m)
}

var xxx_messageInfo_Limiter proto.InternalMessageInfo

func (m *Limiter) GetBackend() Limiter_RateLimitBackend {
	if m != nil {
		return m.Backend
	}
	return Limiter_netEaseLocalFlowControl
}

func (m *Limiter) GetRefresh() *time.Duration {
	if m != nil {
		return m.Refresh
	}
	return nil
}

func (m *Limiter) GetDisableGlobalRateLimit() bool {
	if m != nil {
		return m.DisableGlobalRateLimit
	}
	return false
}

func (m *Limiter) GetDisableAdaptive() bool {
	if m != nil {
		return m.DisableAdaptive
	}
	return false
}

func (m *Limiter) GetEnableServiceEntry() bool {
	if m != nil {
		return m.EnableServiceEntry
	}
	return false
}

func (m *Limiter) GetDisableInsertLocalRateLimit() bool {
	if m != nil {
		return m.DisableInsertLocalRateLimit
	}
	return false
}

func (m *Limiter) GetDisableInsertGlobalRateLimit() bool {
	if m != nil {
		return m.DisableInsertGlobalRateLimit
	}
	return false
}

func (m *Limiter) GetRlsConfigMap() *RlsConfigMap {
	if m != nil {
		return m.RlsConfigMap
	}
	return nil
}

func (m *Limiter) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Limiter) GetRls() *RateLimitService {
	if m != nil {
		return m.Rls
	}
	return nil
}

// configmap will mount to RateLimitService '/data/ratelimit/config'
type RlsConfigMap struct {
	// configmap name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// configmap namespace
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RlsConfigMap) Reset()         { *m = RlsConfigMap{} }
func (m *RlsConfigMap) String() string { return proto.CompactTextString(m) }
func (*RlsConfigMap) ProtoMessage()    {}
func (*RlsConfigMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_4827d40f7d98bcf0, []int{1}
}
func (m *RlsConfigMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RlsConfigMap.Unmarshal(m, b)
}
func (m *RlsConfigMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RlsConfigMap.Marshal(b, m, deterministic)
}
func (m *RlsConfigMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RlsConfigMap.Merge(m, src)
}
func (m *RlsConfigMap) XXX_Size() int {
	return xxx_messageInfo_RlsConfigMap.Size(m)
}
func (m *RlsConfigMap) XXX_DiscardUnknown() {
	xxx_messageInfo_RlsConfigMap.DiscardUnknown(m)
}

var xxx_messageInfo_RlsConfigMap proto.InternalMessageInfo

func (m *RlsConfigMap) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RlsConfigMap) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type RateLimitService struct {
	// rate-limit.gateway-system.svc.cluster.local
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// service port
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitService) Reset()         { *m = RateLimitService{} }
func (m *RateLimitService) String() string { return proto.CompactTextString(m) }
func (*RateLimitService) ProtoMessage()    {}
func (*RateLimitService) Descriptor() ([]byte, []int) {
	return fileDescriptor_4827d40f7d98bcf0, []int{2}
}
func (m *RateLimitService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitService.Unmarshal(m, b)
}
func (m *RateLimitService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitService.Marshal(b, m, deterministic)
}
func (m *RateLimitService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitService.Merge(m, src)
}
func (m *RateLimitService) XXX_Size() int {
	return xxx_messageInfo_RateLimitService.Size(m)
}
func (m *RateLimitService) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitService.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitService proto.InternalMessageInfo

func (m *RateLimitService) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *RateLimitService) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterEnum("slime.microservice.limiter.config.Limiter_RateLimitBackend", Limiter_RateLimitBackend_name, Limiter_RateLimitBackend_value)
	proto.RegisterType((*Limiter)(nil), "slime.microservice.limiter.config.Limiter")
	proto.RegisterType((*RlsConfigMap)(nil), "slime.microservice.limiter.config.RlsConfigMap")
	proto.RegisterType((*RateLimitService)(nil), "slime.microservice.limiter.config.RateLimitService")
}

func init() { proto.RegisterFile("limiter_module.proto", fileDescriptor_4827d40f7d98bcf0) }

var fileDescriptor_4827d40f7d98bcf0 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x66, 0x69, 0x7e, 0x1a, 0x37, 0x82, 0xc8, 0xa0, 0xd6, 0xb4, 0x15, 0x84, 0x9c, 0x82, 0x10,
	0x5e, 0xa9, 0x95, 0x90, 0x10, 0x97, 0x90, 0x12, 0x7e, 0xa4, 0x72, 0xd9, 0x8a, 0x0b, 0x17, 0xe4,
	0xdd, 0x75, 0xb6, 0x16, 0x5e, 0xcf, 0xca, 0x76, 0x82, 0xfa, 0x26, 0xbc, 0x06, 0x4f, 0xc4, 0xab,
	0xa0, 0xcc, 0x7a, 0x69, 0x12, 0x41, 0xe9, 0x69, 0x67, 0x3c, 0xf3, 0x7d, 0xf3, 0xcd, 0x67, 0x2f,
	0x79, 0xa8, 0x55, 0xa9, 0xbc, 0xb4, 0x5f, 0x4b, 0xc8, 0x17, 0x5a, 0xf2, 0xca, 0x82, 0x07, 0xfa,
	0xd4, 0x69, 0x55, 0x4a, 0x5e, 0xaa, 0xcc, 0x82, 0x93, 0x76, 0xa9, 0x32, 0xc9, 0x43, 0x23, 0xcf,
	0xc0, 0xcc, 0x55, 0x71, 0xf8, 0xa2, 0x50, 0xfe, 0x72, 0x91, 0xf2, 0x0c, 0xca, 0xb8, 0x80, 0x02,
	0x62, 0x44, 0xa6, 0x8b, 0x39, 0x66, 0x98, 0x60, 0x54, 0x33, 0x1e, 0x3e, 0x2e, 0x00, 0x0a, 0x2d,
	0xaf, 0xbb, 0xf2, 0x85, 0x15, 0x5e, 0x81, 0xa9, 0xeb, 0xa3, 0x9f, 0x6d, 0xd2, 0x3d, 0xaf, 0x27,
	0xd0, 0xcf, 0xa4, 0x9b, 0x8a, 0xec, 0x9b, 0x34, 0x39, 0xdb, 0x19, 0x46, 0xe3, 0x7b, 0x27, 0xaf,
	0xf9, 0x7f, 0xf5, 0xf0, 0x00, 0xe6, 0x89, 0xf0, 0x12, 0xe3, 0x69, 0x4d, 0x91, 0x34, 0x5c, 0xf4,
	0x15, 0xe9, 0x5a, 0x39, 0xb7, 0xd2, 0x5d, 0xb2, 0xd6, 0x30, 0x1a, 0xef, 0x9d, 0x3c, 0xe2, 0xb5,
	0x28, 0xde, 0x88, 0xe2, 0x6f, 0x83, 0xa8, 0x69, 0xeb, 0xc7, 0xaf, 0x27, 0x51, 0xd2, 0xf4, 0xd3,
	0x97, 0x64, 0x3f, 0x57, 0x4e, 0xa4, 0x5a, 0xbe, 0xd7, 0x90, 0x0a, 0xfd, 0x67, 0x08, 0x6b, 0x0f,
	0xa3, 0xf1, 0x6e, 0xf2, 0x8f, 0x2a, 0x1d, 0x93, 0xfb, 0xa1, 0xf2, 0x26, 0x17, 0x95, 0x57, 0x4b,
	0xc9, 0x3a, 0x08, 0xd8, 0x3e, 0xa6, 0x9c, 0x50, 0x69, 0x56, 0x27, 0x17, 0xf5, 0x7a, 0x33, 0xe3,
	0xed, 0x15, 0xeb, 0x62, 0xf3, 0x5f, 0x2a, 0x74, 0x42, 0x8e, 0x02, 0xc5, 0x47, 0xe3, 0xa4, 0xf5,
	0xe7, 0x90, 0xad, 0xcb, 0xda, 0x45, 0xe0, 0x4d, 0x2d, 0x74, 0x4a, 0x8e, 0x37, 0xca, 0xdb, 0x9b,
	0xf5, 0x90, 0xe2, 0xc6, 0x1e, 0x7a, 0x41, 0xfa, 0x56, 0xbb, 0x33, 0xbc, 0x81, 0x4f, 0xa2, 0x62,
	0x04, 0x7d, 0x8d, 0x6f, 0x71, 0x5d, 0xc9, 0x1a, 0x2c, 0xd9, 0x20, 0xa1, 0xfb, 0xa4, 0x93, 0x43,
	0x29, 0x94, 0x61, 0x7b, 0xc3, 0x68, 0xdc, 0x4b, 0x42, 0x46, 0x67, 0x64, 0xc7, 0x6a, 0xc7, 0xfa,
	0x38, 0xe3, 0xf4, 0x36, 0x33, 0x1a, 0x9d, 0xc1, 0xb9, 0x64, 0x85, 0x1f, 0x7d, 0x20, 0x83, 0xed,
	0x37, 0x42, 0x8f, 0xc8, 0x81, 0x91, 0x7e, 0x26, 0x9c, 0x44, 0x93, 0xde, 0x69, 0xf8, 0x7e, 0x06,
	0xc6, 0x5b, 0xd0, 0x83, 0x3b, 0xf4, 0x80, 0x3c, 0x90, 0x66, 0x09, 0x57, 0x9b, 0xfe, 0x0d, 0xa2,
	0xd1, 0x84, 0xf4, 0xd7, 0xd7, 0xa0, 0x94, 0xb4, 0x8c, 0x28, 0x25, 0x8b, 0x50, 0x36, 0xc6, 0xf4,
	0x98, 0xf4, 0x56, 0x5f, 0x57, 0x89, 0x4c, 0xb2, 0xbb, 0x58, 0xb8, 0x3e, 0x18, 0x4d, 0xd6, 0xb4,
	0x04, 0x91, 0x94, 0x91, 0x6e, 0xd8, 0x27, 0x10, 0x35, 0xe9, 0x8a, 0xbf, 0x02, 0xeb, 0x91, 0xa6,
	0x9d, 0x60, 0x3c, 0x7d, 0xfe, 0xe5, 0x59, 0x6d, 0x84, 0x82, 0x18, 0x83, 0xb8, 0xfe, 0x91, 0x5d,
	0x1c, 0xcc, 0x88, 0x45, 0xa5, 0xe2, 0xda, 0x90, 0xb4, 0x83, 0x0f, 0xfd, 0xf4, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x41, 0xaa, 0xd3, 0x48, 0xf5, 0x03, 0x00, 0x00,
}
