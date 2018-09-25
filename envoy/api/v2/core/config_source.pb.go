// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/core/config_source.proto

package core

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import time "time"

import bytes "bytes"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// APIs may be fetched via either REST or gRPC.
type ApiConfigSource_ApiType int32

const (
	// REST-JSON legacy corresponds to the v1 API.
	ApiConfigSource_REST_LEGACY ApiConfigSource_ApiType = 0
	// REST-JSON v2 API. The `canonical JSON encoding
	// <https://developers.google.com/protocol-buffers/docs/proto3#json>`_ for
	// the v2 protos is used.
	ApiConfigSource_REST ApiConfigSource_ApiType = 1
	// gRPC v2 API.
	ApiConfigSource_GRPC ApiConfigSource_ApiType = 2
)

var ApiConfigSource_ApiType_name = map[int32]string{
	0: "REST_LEGACY",
	1: "REST",
	2: "GRPC",
}
var ApiConfigSource_ApiType_value = map[string]int32{
	"REST_LEGACY": 0,
	"REST":        1,
	"GRPC":        2,
}

func (x ApiConfigSource_ApiType) String() string {
	return proto.EnumName(ApiConfigSource_ApiType_name, int32(x))
}
func (ApiConfigSource_ApiType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_source_97a6275449f3cb63, []int{0, 0}
}

// API configuration source. This identifies the API type and cluster that Envoy
// will use to fetch an xDS API.
type ApiConfigSource struct {
	ApiType ApiConfigSource_ApiType `protobuf:"varint,1,opt,name=api_type,json=apiType,proto3,enum=envoy.api.v2.core.ApiConfigSource_ApiType" json:"api_type,omitempty"`
	// Cluster names should be used only with REST_LEGACY/REST. If > 1
	// cluster is defined, clusters will be cycled through if any kind of failure
	// occurs.
	//
	// .. note::
	//
	//  The cluster with name ``cluster_name`` must be statically defined and its
	//  type must not be ``EDS``.
	ClusterNames []string `protobuf:"bytes,2,rep,name=cluster_names,json=clusterNames" json:"cluster_names,omitempty"`
	// Multiple gRPC services be provided for GRPC. If > 1 cluster is defined,
	// services will be cycled through if any kind of failure occurs.
	GrpcServices []*GrpcService `protobuf:"bytes,4,rep,name=grpc_services,json=grpcServices" json:"grpc_services,omitempty"`
	// For REST APIs, the delay between successive polls.
	RefreshDelay *time.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,stdduration" json:"refresh_delay,omitempty"`
	// For REST APIs, the request timeout. If not set, a default value of 1s will be used.
	RequestTimeout       *time.Duration `protobuf:"bytes,5,opt,name=request_timeout,json=requestTimeout,stdduration" json:"request_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApiConfigSource) Reset()         { *m = ApiConfigSource{} }
func (m *ApiConfigSource) String() string { return proto.CompactTextString(m) }
func (*ApiConfigSource) ProtoMessage()    {}
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_97a6275449f3cb63, []int{0}
}
func (m *ApiConfigSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApiConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApiConfigSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ApiConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiConfigSource.Merge(dst, src)
}
func (m *ApiConfigSource) XXX_Size() int {
	return m.Size()
}
func (m *ApiConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ApiConfigSource proto.InternalMessageInfo

func (m *ApiConfigSource) GetApiType() ApiConfigSource_ApiType {
	if m != nil {
		return m.ApiType
	}
	return ApiConfigSource_REST_LEGACY
}

func (m *ApiConfigSource) GetClusterNames() []string {
	if m != nil {
		return m.ClusterNames
	}
	return nil
}

func (m *ApiConfigSource) GetGrpcServices() []*GrpcService {
	if m != nil {
		return m.GrpcServices
	}
	return nil
}

func (m *ApiConfigSource) GetRefreshDelay() *time.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ApiConfigSource) GetRequestTimeout() *time.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

// Aggregated Discovery Service (ADS) options. This is currently empty, but when
// set in :ref:`ConfigSource <envoy_api_msg_core.ConfigSource>` can be used to
// specify that ADS is to be used.
type AggregatedConfigSource struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregatedConfigSource) Reset()         { *m = AggregatedConfigSource{} }
func (m *AggregatedConfigSource) String() string { return proto.CompactTextString(m) }
func (*AggregatedConfigSource) ProtoMessage()    {}
func (*AggregatedConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_97a6275449f3cb63, []int{1}
}
func (m *AggregatedConfigSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggregatedConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggregatedConfigSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *AggregatedConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatedConfigSource.Merge(dst, src)
}
func (m *AggregatedConfigSource) XXX_Size() int {
	return m.Size()
}
func (m *AggregatedConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatedConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatedConfigSource proto.InternalMessageInfo

// Configuration for :ref:`listeners <config_listeners>`, :ref:`clusters
// <config_cluster_manager>`, :ref:`routes
// <envoy_api_msg_RouteConfiguration>`, :ref:`endpoints
// <arch_overview_service_discovery>` etc. may either be sourced from the
// filesystem or from an xDS API source. Filesystem configs are watched with
// inotify for updates.
type ConfigSource struct {
	// Types that are valid to be assigned to ConfigSourceSpecifier:
	//	*ConfigSource_Path
	//	*ConfigSource_ApiConfigSource
	//	*ConfigSource_Ads
	ConfigSourceSpecifier isConfigSource_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	XXX_NoUnkeyedLiteral  struct{}                             `json:"-"`
	XXX_unrecognized      []byte                               `json:"-"`
	XXX_sizecache         int32                                `json:"-"`
}

func (m *ConfigSource) Reset()         { *m = ConfigSource{} }
func (m *ConfigSource) String() string { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()    {}
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_source_97a6275449f3cb63, []int{2}
}
func (m *ConfigSource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigSource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSource.Merge(dst, src)
}
func (m *ConfigSource) XXX_Size() int {
	return m.Size()
}
func (m *ConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSource proto.InternalMessageInfo

type isConfigSource_ConfigSourceSpecifier interface {
	isConfigSource_ConfigSourceSpecifier()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type ConfigSource_Path struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}
type ConfigSource_ApiConfigSource struct {
	ApiConfigSource *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,oneof"`
}
type ConfigSource_Ads struct {
	Ads *AggregatedConfigSource `protobuf:"bytes,3,opt,name=ads,oneof"`
}

func (*ConfigSource_Path) isConfigSource_ConfigSourceSpecifier()            {}
func (*ConfigSource_ApiConfigSource) isConfigSource_ConfigSourceSpecifier() {}
func (*ConfigSource_Ads) isConfigSource_ConfigSourceSpecifier()             {}

func (m *ConfigSource) GetConfigSourceSpecifier() isConfigSource_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (m *ConfigSource) GetPath() string {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Path); ok {
		return x.Path
	}
	return ""
}

func (m *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_ApiConfigSource); ok {
		return x.ApiConfigSource
	}
	return nil
}

func (m *ConfigSource) GetAds() *AggregatedConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Ads); ok {
		return x.Ads
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConfigSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConfigSource_OneofMarshaler, _ConfigSource_OneofUnmarshaler, _ConfigSource_OneofSizer, []interface{}{
		(*ConfigSource_Path)(nil),
		(*ConfigSource_ApiConfigSource)(nil),
		(*ConfigSource_Ads)(nil),
	}
}

func _ConfigSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConfigSource)
	// config_source_specifier
	switch x := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Path)
	case *ConfigSource_ApiConfigSource:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApiConfigSource); err != nil {
			return err
		}
	case *ConfigSource_Ads:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ads); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConfigSource.ConfigSourceSpecifier has unexpected type %T", x)
	}
	return nil
}

func _ConfigSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConfigSource)
	switch tag {
	case 1: // config_source_specifier.path
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ConfigSourceSpecifier = &ConfigSource_Path{x}
		return true, err
	case 2: // config_source_specifier.api_config_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApiConfigSource)
		err := b.DecodeMessage(msg)
		m.ConfigSourceSpecifier = &ConfigSource_ApiConfigSource{msg}
		return true, err
	case 3: // config_source_specifier.ads
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AggregatedConfigSource)
		err := b.DecodeMessage(msg)
		m.ConfigSourceSpecifier = &ConfigSource_Ads{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConfigSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConfigSource)
	// config_source_specifier
	switch x := m.ConfigSourceSpecifier.(type) {
	case *ConfigSource_Path:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Path)))
		n += len(x.Path)
	case *ConfigSource_ApiConfigSource:
		s := proto.Size(x.ApiConfigSource)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConfigSource_Ads:
		s := proto.Size(x.Ads)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ApiConfigSource)(nil), "envoy.api.v2.core.ApiConfigSource")
	proto.RegisterType((*AggregatedConfigSource)(nil), "envoy.api.v2.core.AggregatedConfigSource")
	proto.RegisterType((*ConfigSource)(nil), "envoy.api.v2.core.ConfigSource")
	proto.RegisterEnum("envoy.api.v2.core.ApiConfigSource_ApiType", ApiConfigSource_ApiType_name, ApiConfigSource_ApiType_value)
}
func (this *ApiConfigSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApiConfigSource)
	if !ok {
		that2, ok := that.(ApiConfigSource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ApiType != that1.ApiType {
		return false
	}
	if len(this.ClusterNames) != len(that1.ClusterNames) {
		return false
	}
	for i := range this.ClusterNames {
		if this.ClusterNames[i] != that1.ClusterNames[i] {
			return false
		}
	}
	if len(this.GrpcServices) != len(that1.GrpcServices) {
		return false
	}
	for i := range this.GrpcServices {
		if !this.GrpcServices[i].Equal(that1.GrpcServices[i]) {
			return false
		}
	}
	if this.RefreshDelay != nil && that1.RefreshDelay != nil {
		if *this.RefreshDelay != *that1.RefreshDelay {
			return false
		}
	} else if this.RefreshDelay != nil {
		return false
	} else if that1.RefreshDelay != nil {
		return false
	}
	if this.RequestTimeout != nil && that1.RequestTimeout != nil {
		if *this.RequestTimeout != *that1.RequestTimeout {
			return false
		}
	} else if this.RequestTimeout != nil {
		return false
	} else if that1.RequestTimeout != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AggregatedConfigSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AggregatedConfigSource)
	if !ok {
		that2, ok := that.(AggregatedConfigSource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ConfigSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfigSource)
	if !ok {
		that2, ok := that.(ConfigSource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.ConfigSourceSpecifier == nil {
		if this.ConfigSourceSpecifier != nil {
			return false
		}
	} else if this.ConfigSourceSpecifier == nil {
		return false
	} else if !this.ConfigSourceSpecifier.Equal(that1.ConfigSourceSpecifier) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ConfigSource_Path) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfigSource_Path)
	if !ok {
		that2, ok := that.(ConfigSource_Path)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	return true
}
func (this *ConfigSource_ApiConfigSource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfigSource_ApiConfigSource)
	if !ok {
		that2, ok := that.(ConfigSource_ApiConfigSource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ApiConfigSource.Equal(that1.ApiConfigSource) {
		return false
	}
	return true
}
func (this *ConfigSource_Ads) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfigSource_Ads)
	if !ok {
		that2, ok := that.(ConfigSource_Ads)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Ads.Equal(that1.Ads) {
		return false
	}
	return true
}
func (m *ApiConfigSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApiConfigSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ApiType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfigSource(dAtA, i, uint64(m.ApiType))
	}
	if len(m.ClusterNames) > 0 {
		for _, s := range m.ClusterNames {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.RefreshDelay != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfigSource(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.RefreshDelay)))
		n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.RefreshDelay, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.GrpcServices) > 0 {
		for _, msg := range m.GrpcServices {
			dAtA[i] = 0x22
			i++
			i = encodeVarintConfigSource(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.RequestTimeout != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintConfigSource(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*m.RequestTimeout)))
		n2, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.RequestTimeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *AggregatedConfigSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregatedConfigSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConfigSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ConfigSourceSpecifier != nil {
		nn3, err := m.ConfigSourceSpecifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ConfigSource_Path) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintConfigSource(dAtA, i, uint64(len(m.Path)))
	i += copy(dAtA[i:], m.Path)
	return i, nil
}
func (m *ConfigSource_ApiConfigSource) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ApiConfigSource != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfigSource(dAtA, i, uint64(m.ApiConfigSource.Size()))
		n4, err := m.ApiConfigSource.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *ConfigSource_Ads) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Ads != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConfigSource(dAtA, i, uint64(m.Ads.Size()))
		n5, err := m.Ads.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func encodeVarintConfigSource(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ApiConfigSource) Size() (n int) {
	var l int
	_ = l
	if m.ApiType != 0 {
		n += 1 + sovConfigSource(uint64(m.ApiType))
	}
	if len(m.ClusterNames) > 0 {
		for _, s := range m.ClusterNames {
			l = len(s)
			n += 1 + l + sovConfigSource(uint64(l))
		}
	}
	if m.RefreshDelay != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.RefreshDelay)
		n += 1 + l + sovConfigSource(uint64(l))
	}
	if len(m.GrpcServices) > 0 {
		for _, e := range m.GrpcServices {
			l = e.Size()
			n += 1 + l + sovConfigSource(uint64(l))
		}
	}
	if m.RequestTimeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.RequestTimeout)
		n += 1 + l + sovConfigSource(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AggregatedConfigSource) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigSource) Size() (n int) {
	var l int
	_ = l
	if m.ConfigSourceSpecifier != nil {
		n += m.ConfigSourceSpecifier.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ConfigSource_Path) Size() (n int) {
	var l int
	_ = l
	l = len(m.Path)
	n += 1 + l + sovConfigSource(uint64(l))
	return n
}
func (m *ConfigSource_ApiConfigSource) Size() (n int) {
	var l int
	_ = l
	if m.ApiConfigSource != nil {
		l = m.ApiConfigSource.Size()
		n += 1 + l + sovConfigSource(uint64(l))
	}
	return n
}
func (m *ConfigSource_Ads) Size() (n int) {
	var l int
	_ = l
	if m.Ads != nil {
		l = m.Ads.Size()
		n += 1 + l + sovConfigSource(uint64(l))
	}
	return n
}

func sovConfigSource(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfigSource(x uint64) (n int) {
	return sovConfigSource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ApiConfigSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfigSource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ApiConfigSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApiConfigSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiType", wireType)
			}
			m.ApiType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiType |= (ApiConfigSource_ApiType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfigSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterNames = append(m.ClusterNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefreshDelay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfigSource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RefreshDelay == nil {
				m.RefreshDelay = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.RefreshDelay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcServices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfigSource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GrpcServices = append(m.GrpcServices, &GrpcService{})
			if err := m.GrpcServices[len(m.GrpcServices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfigSource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestTimeout == nil {
				m.RequestTimeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.RequestTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfigSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfigSource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AggregatedConfigSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfigSource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AggregatedConfigSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregatedConfigSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipConfigSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfigSource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfigSource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConfigSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfigSourceSpecifier = &ConfigSource_Path{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiConfigSource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfigSource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ApiConfigSource{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigSourceSpecifier = &ConfigSource_ApiConfigSource{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ads", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfigSource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &AggregatedConfigSource{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigSourceSpecifier = &ConfigSource_Ads{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfigSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfigSource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfigSource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfigSource
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfigSource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConfigSource
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfigSource
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConfigSource(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConfigSource = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfigSource   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/api/v2/core/config_source.proto", fileDescriptor_config_source_97a6275449f3cb63)
}

var fileDescriptor_config_source_97a6275449f3cb63 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbd, 0x6e, 0x1a, 0x4d,
	0x14, 0x65, 0x58, 0xf8, 0xc0, 0x03, 0x98, 0xf5, 0xca, 0xfa, 0xbc, 0x71, 0xb1, 0x41, 0x24, 0x91,
	0x88, 0x8b, 0x5d, 0x89, 0xd4, 0x29, 0xf8, 0xb1, 0xa0, 0x88, 0x22, 0x67, 0xa1, 0x49, 0xb5, 0x1a,
	0x2f, 0x97, 0xf1, 0x48, 0x98, 0x99, 0xcc, 0xcc, 0x22, 0xd1, 0xe6, 0x29, 0xa2, 0x3c, 0x41, 0x94,
	0x27, 0x88, 0x52, 0xb9, 0x4c, 0x99, 0x2e, 0x65, 0x22, 0x3a, 0xbf, 0x45, 0x34, 0xbb, 0x8b, 0x62,
	0x6c, 0x24, 0x77, 0xf7, 0x9e, 0xb9, 0xe7, 0xea, 0x9c, 0x7b, 0x06, 0xbf, 0x80, 0xe5, 0x8a, 0xaf,
	0x03, 0x22, 0x58, 0xb0, 0xea, 0x06, 0x31, 0x97, 0x10, 0xc4, 0x7c, 0x39, 0x67, 0x34, 0x52, 0x3c,
	0x91, 0x31, 0xf8, 0x42, 0x72, 0xcd, 0x9d, 0xa3, 0x74, 0xcc, 0x27, 0x82, 0xf9, 0xab, 0xae, 0x6f,
	0xc6, 0x4e, 0x9f, 0x3f, 0x64, 0x52, 0x29, 0xe2, 0x48, 0x81, 0x5c, 0xb1, 0x2d, 0xf1, 0xd4, 0xa3,
	0x9c, 0xd3, 0x05, 0x04, 0x69, 0x77, 0x99, 0xcc, 0x83, 0x59, 0x22, 0x89, 0x66, 0x7c, 0x99, 0xbf,
	0x9f, 0xac, 0xc8, 0x82, 0xcd, 0x88, 0x86, 0x60, 0x5b, 0xe4, 0x0f, 0xc7, 0x94, 0x53, 0x9e, 0x96,
	0x81, 0xa9, 0x32, 0xb4, 0xfd, 0xd9, 0xc2, 0xcd, 0x9e, 0x60, 0x83, 0x54, 0xe2, 0x24, 0x55, 0xe8,
	0xbc, 0xc3, 0x55, 0x22, 0x58, 0xa4, 0xd7, 0x02, 0x5c, 0xd4, 0x42, 0x9d, 0xc3, 0xee, 0x99, 0xff,
	0x40, 0xae, 0x7f, 0x8f, 0x65, 0xfa, 0xe9, 0x5a, 0x40, 0x1f, 0x7f, 0xbf, 0xbd, 0xb1, 0xca, 0x1f,
	0x51, 0xd1, 0x46, 0x61, 0x85, 0x64, 0xa0, 0xf3, 0x0c, 0x37, 0xe2, 0x45, 0xa2, 0x34, 0xc8, 0x68,
	0x49, 0xae, 0x41, 0xb9, 0xc5, 0x96, 0xd5, 0x39, 0x08, 0xeb, 0x39, 0xf8, 0xd6, 0x60, 0xce, 0x10,
	0x37, 0x24, 0xcc, 0x25, 0xa8, 0xab, 0x68, 0x06, 0x0b, 0xb2, 0x76, 0xad, 0x16, 0xea, 0xd4, 0xba,
	0x4f, 0xfc, 0xcc, 0xb2, 0xbf, 0xb5, 0xec, 0x0f, 0x73, 0xcb, 0xfd, 0xd2, 0xa7, 0xdf, 0x4f, 0x51,
	0x58, 0xcf, 0x59, 0x43, 0x43, 0x72, 0x06, 0xb8, 0x71, 0xf7, 0x6c, 0xca, 0x2d, 0xb5, 0xac, 0x4e,
	0xad, 0xeb, 0xed, 0xb1, 0x30, 0x92, 0x22, 0x9e, 0x64, 0x63, 0x61, 0x9d, 0xfe, 0x6b, 0x94, 0x33,
	0xc5, 0x4d, 0x09, 0x1f, 0x12, 0x50, 0x3a, 0xd2, 0xec, 0x1a, 0x78, 0xa2, 0xdd, 0xf2, 0x63, 0x62,
	0x6c, 0x23, 0xc6, 0x98, 0xaf, 0x7c, 0x45, 0xa5, 0xb3, 0x62, 0xb5, 0x10, 0x1e, 0xe6, 0x3b, 0xa6,
	0xd9, 0x8a, 0xb6, 0x8f, 0x2b, 0xf9, 0x95, 0x9c, 0x26, 0xae, 0x85, 0xe7, 0x93, 0x69, 0xf4, 0xe6,
	0x7c, 0xd4, 0x1b, 0xbc, 0xb7, 0x0b, 0x4e, 0x15, 0x97, 0x0c, 0x60, 0x23, 0x53, 0x8d, 0xc2, 0x8b,
	0x81, 0x5d, 0x6c, 0xbb, 0xf8, 0xff, 0x1e, 0xa5, 0x12, 0x28, 0xd1, 0x30, 0xbb, 0x7b, 0xec, 0xf6,
	0x2f, 0x84, 0xeb, 0x3b, 0x99, 0x1d, 0xe3, 0x92, 0x20, 0xfa, 0x2a, 0xcd, 0xeb, 0x60, 0x5c, 0x08,
	0xd3, 0xce, 0xb9, 0xc0, 0x47, 0x26, 0xc9, 0x9d, 0x0f, 0xe8, 0x16, 0x53, 0x23, 0xed, 0xc7, 0x23,
	0x1d, 0x17, 0xc2, 0x26, 0xb9, 0xf7, 0x37, 0x5e, 0x63, 0x8b, 0xcc, 0x54, 0x9e, 0xcc, 0xcb, 0x7d,
	0x3b, 0xf6, 0x0a, 0x1e, 0x17, 0x42, 0xc3, 0xeb, 0xb7, 0xf0, 0xc9, 0x8e, 0x98, 0x48, 0x09, 0x88,
	0xd9, 0x9c, 0x81, 0x74, 0xca, 0xdf, 0x6e, 0x6f, 0x2c, 0xd4, 0xb7, 0xbf, 0x6c, 0x3c, 0xf4, 0x63,
	0xe3, 0xa1, 0x9f, 0x1b, 0x0f, 0xfd, 0xd9, 0x78, 0xe8, 0xf2, 0xbf, 0xf4, 0xd4, 0xaf, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xab, 0x0b, 0x66, 0x7e, 0x5a, 0x03, 0x00, 0x00,
}
