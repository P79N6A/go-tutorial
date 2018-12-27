// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter-config.proto

package proto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// 代表一个Dubbo，HSF的路由配置
type FilterConfig struct {
	StatPrefix           string         `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	ProtocolType         string         `protobuf:"bytes,2,opt,name=protocol_type,json=protocolType,proto3" json:"protocol_type,omitempty"`
	SerializationType    string         `protobuf:"bytes,3,opt,name=serialization_type,json=serializationType,proto3" json:"serialization_type,omitempty"`
	RouteConfig          []*RouteConfig `protobuf:"bytes,4,rep,name=route_config,json=routeConfig" json:"route_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{0}
}
func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (dst *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(dst, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *FilterConfig) GetProtocolType() string {
	if m != nil {
		return m.ProtocolType
	}
	return ""
}

func (m *FilterConfig) GetSerializationType() string {
	if m != nil {
		return m.SerializationType
	}
	return ""
}

func (m *FilterConfig) GetRouteConfig() []*RouteConfig {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

type RouteConfig struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interface            string              `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	Version              string              `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Group                string              `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Spec                 *RouteConfigSpec    `protobuf:"bytes,5,opt,name=spec" json:"spec,omitempty"`
	Routes               []*RouteConfigRoute `protobuf:"bytes,6,rep,name=routes" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RouteConfig) Reset()         { *m = RouteConfig{} }
func (m *RouteConfig) String() string { return proto.CompactTextString(m) }
func (*RouteConfig) ProtoMessage()    {}
func (*RouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{1}
}
func (m *RouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfig.Unmarshal(m, b)
}
func (m *RouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfig.Marshal(b, m, deterministic)
}
func (dst *RouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfig.Merge(dst, src)
}
func (m *RouteConfig) XXX_Size() int {
	return xxx_messageInfo_RouteConfig.Size(m)
}
func (m *RouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfig proto.InternalMessageInfo

func (m *RouteConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfig) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RouteConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RouteConfig) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RouteConfig) GetSpec() *RouteConfigSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *RouteConfig) GetRoutes() []*RouteConfigRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

type RouteConfigSpec struct {
	AppName              string   `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	WriteMode            string   `protobuf:"bytes,2,opt,name=write_mode,json=writeMode,proto3" json:"write_mode,omitempty"`
	RouteIdx             int32    `protobuf:"varint,3,opt,name=route_idx,json=routeIdx,proto3" json:"route_idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfigSpec) Reset()         { *m = RouteConfigSpec{} }
func (m *RouteConfigSpec) String() string { return proto.CompactTextString(m) }
func (*RouteConfigSpec) ProtoMessage()    {}
func (*RouteConfigSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{2}
}
func (m *RouteConfigSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfigSpec.Unmarshal(m, b)
}
func (m *RouteConfigSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfigSpec.Marshal(b, m, deterministic)
}
func (dst *RouteConfigSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfigSpec.Merge(dst, src)
}
func (m *RouteConfigSpec) XXX_Size() int {
	return xxx_messageInfo_RouteConfigSpec.Size(m)
}
func (m *RouteConfigSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfigSpec.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfigSpec proto.InternalMessageInfo

func (m *RouteConfigSpec) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *RouteConfigSpec) GetWriteMode() string {
	if m != nil {
		return m.WriteMode
	}
	return ""
}

func (m *RouteConfigSpec) GetRouteIdx() int32 {
	if m != nil {
		return m.RouteIdx
	}
	return 0
}

type RouteConfigRoute struct {
	Match                *RouteConfigRouteMatch `protobuf:"bytes,1,opt,name=match" json:"match,omitempty"`
	Route                *RouteConfigRouteRoute `protobuf:"bytes,2,opt,name=route" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteConfigRoute) Reset()         { *m = RouteConfigRoute{} }
func (m *RouteConfigRoute) String() string { return proto.CompactTextString(m) }
func (*RouteConfigRoute) ProtoMessage()    {}
func (*RouteConfigRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{3}
}
func (m *RouteConfigRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfigRoute.Unmarshal(m, b)
}
func (m *RouteConfigRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfigRoute.Marshal(b, m, deterministic)
}
func (dst *RouteConfigRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfigRoute.Merge(dst, src)
}
func (m *RouteConfigRoute) XXX_Size() int {
	return xxx_messageInfo_RouteConfigRoute.Size(m)
}
func (m *RouteConfigRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfigRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfigRoute proto.InternalMessageInfo

func (m *RouteConfigRoute) GetMatch() *RouteConfigRouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *RouteConfigRoute) GetRoute() *RouteConfigRouteRoute {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteConfigRouteRoute struct {
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfigRouteRoute) Reset()         { *m = RouteConfigRouteRoute{} }
func (m *RouteConfigRouteRoute) String() string { return proto.CompactTextString(m) }
func (*RouteConfigRouteRoute) ProtoMessage()    {}
func (*RouteConfigRouteRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{4}
}
func (m *RouteConfigRouteRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfigRouteRoute.Unmarshal(m, b)
}
func (m *RouteConfigRouteRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfigRouteRoute.Marshal(b, m, deterministic)
}
func (dst *RouteConfigRouteRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfigRouteRoute.Merge(dst, src)
}
func (m *RouteConfigRouteRoute) XXX_Size() int {
	return xxx_messageInfo_RouteConfigRouteRoute.Size(m)
}
func (m *RouteConfigRouteRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfigRouteRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfigRouteRoute proto.InternalMessageInfo

func (m *RouteConfigRouteRoute) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type RouteConfigRouteMatch struct {
	Method               *RouteConfigRouteMatchMethod `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Headers              []*HeaderMatch               `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RouteConfigRouteMatch) Reset()         { *m = RouteConfigRouteMatch{} }
func (m *RouteConfigRouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteConfigRouteMatch) ProtoMessage()    {}
func (*RouteConfigRouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{5}
}
func (m *RouteConfigRouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfigRouteMatch.Unmarshal(m, b)
}
func (m *RouteConfigRouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfigRouteMatch.Marshal(b, m, deterministic)
}
func (dst *RouteConfigRouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfigRouteMatch.Merge(dst, src)
}
func (m *RouteConfigRouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteConfigRouteMatch.Size(m)
}
func (m *RouteConfigRouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfigRouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfigRouteMatch proto.InternalMessageInfo

func (m *RouteConfigRouteMatch) GetMethod() *RouteConfigRouteMatchMethod {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RouteConfigRouteMatch) GetHeaders() []*HeaderMatch {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteConfigRouteMatchMethod struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParamsMatch          []*ParameterMatch `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch" json:"params_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RouteConfigRouteMatchMethod) Reset()         { *m = RouteConfigRouteMatchMethod{} }
func (m *RouteConfigRouteMatchMethod) String() string { return proto.CompactTextString(m) }
func (*RouteConfigRouteMatchMethod) ProtoMessage()    {}
func (*RouteConfigRouteMatchMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{6}
}
func (m *RouteConfigRouteMatchMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfigRouteMatchMethod.Unmarshal(m, b)
}
func (m *RouteConfigRouteMatchMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfigRouteMatchMethod.Marshal(b, m, deterministic)
}
func (dst *RouteConfigRouteMatchMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfigRouteMatchMethod.Merge(dst, src)
}
func (m *RouteConfigRouteMatchMethod) XXX_Size() int {
	return xxx_messageInfo_RouteConfigRouteMatchMethod.Size(m)
}
func (m *RouteConfigRouteMatchMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfigRouteMatchMethod.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfigRouteMatchMethod proto.InternalMessageInfo

func (m *RouteConfigRouteMatchMethod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfigRouteMatchMethod) GetParamsMatch() []*ParameterMatch {
	if m != nil {
		return m.ParamsMatch
	}
	return nil
}

type ParameterMatch struct {
	Index                int32       `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Type                 string      `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	RangeMatch           *RangeMatch `protobuf:"bytes,3,opt,name=range_match,json=rangeMatch" json:"range_match,omitempty"`
	PrefixMatch          string      `protobuf:"bytes,4,opt,name=prefix_match,json=prefixMatch,proto3" json:"prefix_match,omitempty"`
	RegexMatch           string      `protobuf:"bytes,5,opt,name=regex_match,json=regexMatch,proto3" json:"regex_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ParameterMatch) Reset()         { *m = ParameterMatch{} }
func (m *ParameterMatch) String() string { return proto.CompactTextString(m) }
func (*ParameterMatch) ProtoMessage()    {}
func (*ParameterMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{7}
}
func (m *ParameterMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParameterMatch.Unmarshal(m, b)
}
func (m *ParameterMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParameterMatch.Marshal(b, m, deterministic)
}
func (dst *ParameterMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterMatch.Merge(dst, src)
}
func (m *ParameterMatch) XXX_Size() int {
	return xxx_messageInfo_ParameterMatch.Size(m)
}
func (m *ParameterMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterMatch proto.InternalMessageInfo

func (m *ParameterMatch) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ParameterMatch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ParameterMatch) GetRangeMatch() *RangeMatch {
	if m != nil {
		return m.RangeMatch
	}
	return nil
}

func (m *ParameterMatch) GetPrefixMatch() string {
	if m != nil {
		return m.PrefixMatch
	}
	return ""
}

func (m *ParameterMatch) GetRegexMatch() string {
	if m != nil {
		return m.RegexMatch
	}
	return ""
}

type HeaderMatch struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExactMatch           string   `protobuf:"bytes,2,opt,name=exact_match,json=exactMatch,proto3" json:"exact_match,omitempty"`
	InvertMatch          bool     `protobuf:"varint,3,opt,name=invert_match,json=invertMatch,proto3" json:"invert_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderMatch) Reset()         { *m = HeaderMatch{} }
func (m *HeaderMatch) String() string { return proto.CompactTextString(m) }
func (*HeaderMatch) ProtoMessage()    {}
func (*HeaderMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{8}
}
func (m *HeaderMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderMatch.Unmarshal(m, b)
}
func (m *HeaderMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderMatch.Marshal(b, m, deterministic)
}
func (dst *HeaderMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderMatch.Merge(dst, src)
}
func (m *HeaderMatch) XXX_Size() int {
	return xxx_messageInfo_HeaderMatch.Size(m)
}
func (m *HeaderMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderMatch.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderMatch proto.InternalMessageInfo

func (m *HeaderMatch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HeaderMatch) GetExactMatch() string {
	if m != nil {
		return m.ExactMatch
	}
	return ""
}

func (m *HeaderMatch) GetInvertMatch() bool {
	if m != nil {
		return m.InvertMatch
	}
	return false
}

type RangeMatch struct {
	Start                int32    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int32    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RangeMatch) Reset()         { *m = RangeMatch{} }
func (m *RangeMatch) String() string { return proto.CompactTextString(m) }
func (*RangeMatch) ProtoMessage()    {}
func (*RangeMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_config_939c919159233461, []int{9}
}
func (m *RangeMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RangeMatch.Unmarshal(m, b)
}
func (m *RangeMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RangeMatch.Marshal(b, m, deterministic)
}
func (dst *RangeMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeMatch.Merge(dst, src)
}
func (m *RangeMatch) XXX_Size() int {
	return xxx_messageInfo_RangeMatch.Size(m)
}
func (m *RangeMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RangeMatch proto.InternalMessageInfo

func (m *RangeMatch) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *RangeMatch) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "proto.FilterConfig")
	proto.RegisterType((*RouteConfig)(nil), "proto.RouteConfig")
	proto.RegisterType((*RouteConfigSpec)(nil), "proto.RouteConfigSpec")
	proto.RegisterType((*RouteConfigRoute)(nil), "proto.RouteConfigRoute")
	proto.RegisterType((*RouteConfigRouteRoute)(nil), "proto.RouteConfigRouteRoute")
	proto.RegisterType((*RouteConfigRouteMatch)(nil), "proto.RouteConfigRouteMatch")
	proto.RegisterType((*RouteConfigRouteMatchMethod)(nil), "proto.RouteConfigRouteMatchMethod")
	proto.RegisterType((*ParameterMatch)(nil), "proto.ParameterMatch")
	proto.RegisterType((*HeaderMatch)(nil), "proto.HeaderMatch")
	proto.RegisterType((*RangeMatch)(nil), "proto.RangeMatch")
}

func init() { proto.RegisterFile("filter-config.proto", fileDescriptor_filter_config_939c919159233461) }

var fileDescriptor_filter_config_939c919159233461 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0x9b, 0x38, 0x6d, 0x67, 0xf3, 0xfb, 0xd1, 0x2e, 0x14, 0x8c, 0x5a, 0xd4, 0x62, 0x2e,
	0x11, 0xa2, 0x45, 0x04, 0x90, 0x10, 0x57, 0x24, 0x04, 0x87, 0xa0, 0xca, 0x70, 0xb7, 0x16, 0x7b,
	0x92, 0xac, 0x48, 0xec, 0xd5, 0x7a, 0x53, 0xd2, 0xde, 0xf8, 0x4a, 0x1c, 0xf9, 0x0c, 0x7c, 0x28,
	0xb4, 0x33, 0x76, 0xfe, 0xb4, 0x69, 0x2f, 0xc9, 0xce, 0x9b, 0x37, 0x33, 0xef, 0x4d, 0x76, 0x03,
	0xf7, 0x87, 0x7a, 0xe2, 0xd0, 0x9e, 0x66, 0x65, 0x31, 0xd4, 0xa3, 0x33, 0x63, 0x4b, 0x57, 0xca,
	0x90, 0xbe, 0xe2, 0x3f, 0x01, 0x74, 0x3f, 0x52, 0xfa, 0x03, 0x65, 0xe5, 0x31, 0x88, 0xca, 0x29,
	0x97, 0x1a, 0x8b, 0x43, 0x3d, 0x8f, 0x82, 0x93, 0xa0, 0xb7, 0x9b, 0x80, 0x87, 0xce, 0x09, 0x91,
	0xcf, 0xe0, 0x3f, 0x2a, 0xcd, 0xca, 0x49, 0xea, 0x2e, 0x0d, 0x46, 0x5b, 0x44, 0xe9, 0x36, 0xe0,
	0xb7, 0x4b, 0x83, 0xf2, 0x14, 0x64, 0x85, 0x56, 0xab, 0x89, 0xbe, 0x52, 0x4e, 0x97, 0x05, 0x33,
	0x5b, 0xc4, 0xdc, 0x5f, 0xcb, 0x10, 0xfd, 0x2d, 0x74, 0x6d, 0x39, 0x73, 0x98, 0xb2, 0xc4, 0xa8,
	0x7d, 0xd2, 0xea, 0x89, 0xbe, 0x64, 0xa9, 0x67, 0x89, 0x4f, 0xb1, 0xbc, 0x44, 0xd8, 0x65, 0x10,
	0xff, 0x0d, 0x40, 0xac, 0x24, 0xa5, 0x84, 0x76, 0xa1, 0xa6, 0x58, 0x8b, 0xa6, 0xb3, 0x3c, 0x82,
	0x5d, 0x5d, 0x38, 0xb4, 0x43, 0x95, 0x35, 0x52, 0x97, 0x80, 0x8c, 0x60, 0xfb, 0x02, 0x6d, 0xa5,
	0xcb, 0xa2, 0x16, 0xd7, 0x84, 0xf2, 0x01, 0x84, 0x23, 0x5b, 0xce, 0x4c, 0xd4, 0x26, 0x9c, 0x03,
	0xf9, 0x1c, 0xda, 0x95, 0xc1, 0x2c, 0x0a, 0x4f, 0x82, 0x9e, 0xe8, 0x3f, 0xbc, 0x29, 0xf0, 0xab,
	0xc1, 0x2c, 0x21, 0x8e, 0x7c, 0x09, 0x1d, 0x12, 0x5b, 0x45, 0x1d, 0xb2, 0xf3, 0x68, 0x83, 0x1d,
	0x7f, 0x4c, 0x6a, 0x5a, 0x3c, 0x86, 0x7b, 0xd7, 0x3a, 0xc9, 0xc7, 0xb0, 0xa3, 0x8c, 0x49, 0x57,
	0x5c, 0x6d, 0x2b, 0x63, 0xbe, 0x78, 0x63, 0x4f, 0x00, 0x7e, 0x5a, 0xed, 0x30, 0x9d, 0x96, 0xf9,
	0xc2, 0x19, 0x21, 0x83, 0x32, 0x47, 0x79, 0x08, 0xbb, 0xbc, 0x52, 0x9d, 0xcf, 0xc9, 0x5b, 0x98,
	0xec, 0x10, 0xf0, 0x39, 0x9f, 0xc7, 0x57, 0xb0, 0x77, 0x5d, 0x85, 0xec, 0x43, 0x38, 0x55, 0x2e,
	0x1b, 0xd3, 0x1c, 0xd1, 0x3f, 0xba, 0x45, 0xed, 0xc0, 0x73, 0x12, 0xa6, 0xfa, 0x1a, 0xea, 0x49,
	0xe3, 0x6f, 0xaf, 0xe1, 0x0f, 0xa6, 0xc6, 0xaf, 0xe0, 0x60, 0x63, 0xde, 0xff, 0x16, 0xd9, 0x64,
	0x56, 0x39, 0xb4, 0x8d, 0xd5, 0x3a, 0x8c, 0x7f, 0x05, 0x37, 0x6b, 0x48, 0x87, 0x7c, 0x0f, 0x9d,
	0x29, 0xba, 0x71, 0x99, 0xd7, 0xaa, 0xe3, 0xbb, 0x54, 0x0f, 0x88, 0x99, 0xd4, 0x15, 0xf2, 0x05,
	0x6c, 0x8f, 0x51, 0xe5, 0x68, 0xab, 0x68, 0x6b, 0xed, 0xbe, 0x7d, 0x22, 0x94, 0x8d, 0x36, 0x94,
	0xf8, 0x07, 0x1c, 0xde, 0xd1, 0x74, 0xe3, 0xd5, 0x7b, 0x07, 0x5d, 0xa3, 0xac, 0x9a, 0x56, 0x29,
	0x2f, 0x96, 0xa7, 0x1c, 0xd4, 0x53, 0xce, 0x7d, 0x0a, 0x5d, 0x33, 0x48, 0x30, 0x95, 0x82, 0xf8,
	0x77, 0x00, 0xff, 0xaf, 0xe7, 0xfd, 0x7d, 0xd4, 0x45, 0x8e, 0xfc, 0x22, 0xc3, 0x84, 0x03, 0x3f,
	0x76, 0xe5, 0x0d, 0xd2, 0x59, 0xf6, 0x41, 0x58, 0x55, 0x8c, 0xb0, 0x9e, 0xda, 0xa2, 0xc5, 0xec,
	0x37, 0x8b, 0xf1, 0x19, 0x9e, 0x08, 0x76, 0x71, 0x96, 0x4f, 0xa1, 0xcb, 0x0f, 0xbe, 0x2e, 0xe2,
	0x4b, 0x2f, 0x18, 0x63, 0xca, 0x31, 0x08, 0x8b, 0x23, 0x6c, 0x18, 0x21, 0xff, 0x31, 0x10, 0xc4,
	0xa2, 0x11, 0xc4, 0xca, 0xe6, 0x36, 0x6e, 0xe4, 0x18, 0x04, 0xce, 0x55, 0xe6, 0x16, 0x0b, 0xa1,
	0x1e, 0x04, 0x2d, 0x74, 0xe8, 0xe2, 0x02, 0xad, 0x5b, 0x11, 0xbf, 0x93, 0x08, 0xc6, 0x78, 0xcc,
	0x1b, 0x80, 0xa5, 0x09, 0xbf, 0x96, 0xca, 0x29, 0xeb, 0x9a, 0xb5, 0x50, 0x20, 0xf7, 0xa0, 0x85,
	0x45, 0x4e, 0xfd, 0xc3, 0xc4, 0x1f, 0xbf, 0x77, 0xc8, 0xfe, 0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x27, 0x56, 0x68, 0x8b, 0x0c, 0x05, 0x00, 0x00,
}
