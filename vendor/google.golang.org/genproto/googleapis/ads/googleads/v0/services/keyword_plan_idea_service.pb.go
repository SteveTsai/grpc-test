// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/keyword_plan_idea_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v0/common"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeasRequest struct {
	// The ID of the customer with the recommendation.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The resource name of the language to target.
	// Required
	Language *wrappers.StringValue `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	// The resource names of the location to target.
	// Max 10
	GeoTargetConstants []*wrappers.StringValue `protobuf:"bytes,8,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	// Targeting network.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,9,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v0.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// The type of seed to generate keyword ideas.
	//
	// Types that are valid to be assigned to Seed:
	//	*GenerateKeywordIdeasRequest_KeywordAndUrlSeed
	//	*GenerateKeywordIdeasRequest_KeywordSeed
	//	*GenerateKeywordIdeasRequest_UrlSeed
	Seed                 isGenerateKeywordIdeasRequest_Seed `protobuf_oneof:"seed"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GenerateKeywordIdeasRequest) Reset()         { *m = GenerateKeywordIdeasRequest{} }
func (m *GenerateKeywordIdeasRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeasRequest) ProtoMessage()    {}
func (*GenerateKeywordIdeasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_8eed2c943de898df, []int{0}
}
func (m *GenerateKeywordIdeasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateKeywordIdeasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeasRequest.Merge(dst, src)
}
func (m *GenerateKeywordIdeasRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Size(m)
}
func (m *GenerateKeywordIdeasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeasRequest proto.InternalMessageInfo

func (m *GenerateKeywordIdeasRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *GenerateKeywordIdeasRequest) GetLanguage() *wrappers.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetGeoTargetConstants() []*wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstants
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

type isGenerateKeywordIdeasRequest_Seed interface {
	isGenerateKeywordIdeasRequest_Seed()
}

type GenerateKeywordIdeasRequest_KeywordAndUrlSeed struct {
	KeywordAndUrlSeed *KeywordAndUrlSeed `protobuf:"bytes,2,opt,name=keyword_and_url_seed,json=keywordAndUrlSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_KeywordSeed struct {
	KeywordSeed *KeywordSeed `protobuf:"bytes,3,opt,name=keyword_seed,json=keywordSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_UrlSeed struct {
	UrlSeed *UrlSeed `protobuf:"bytes,5,opt,name=url_seed,json=urlSeed,proto3,oneof"`
}

func (*GenerateKeywordIdeasRequest_KeywordAndUrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_KeywordSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_UrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (m *GenerateKeywordIdeasRequest) GetSeed() isGenerateKeywordIdeasRequest_Seed {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordAndUrlSeed() *KeywordAndUrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed); ok {
		return x.KeywordAndUrlSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordSeed() *KeywordSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordSeed); ok {
		return x.KeywordSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetUrlSeed() *UrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_UrlSeed); ok {
		return x.UrlSeed
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GenerateKeywordIdeasRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GenerateKeywordIdeasRequest_OneofMarshaler, _GenerateKeywordIdeasRequest_OneofUnmarshaler, _GenerateKeywordIdeasRequest_OneofSizer, []interface{}{
		(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed)(nil),
		(*GenerateKeywordIdeasRequest_KeywordSeed)(nil),
		(*GenerateKeywordIdeasRequest_UrlSeed)(nil),
	}
}

func _GenerateKeywordIdeasRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GenerateKeywordIdeasRequest)
	// seed
	switch x := m.Seed.(type) {
	case *GenerateKeywordIdeasRequest_KeywordAndUrlSeed:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KeywordAndUrlSeed); err != nil {
			return err
		}
	case *GenerateKeywordIdeasRequest_KeywordSeed:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KeywordSeed); err != nil {
			return err
		}
	case *GenerateKeywordIdeasRequest_UrlSeed:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UrlSeed); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GenerateKeywordIdeasRequest.Seed has unexpected type %T", x)
	}
	return nil
}

func _GenerateKeywordIdeasRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GenerateKeywordIdeasRequest)
	switch tag {
	case 2: // seed.keyword_and_url_seed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KeywordAndUrlSeed)
		err := b.DecodeMessage(msg)
		m.Seed = &GenerateKeywordIdeasRequest_KeywordAndUrlSeed{msg}
		return true, err
	case 3: // seed.keyword_seed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KeywordSeed)
		err := b.DecodeMessage(msg)
		m.Seed = &GenerateKeywordIdeasRequest_KeywordSeed{msg}
		return true, err
	case 5: // seed.url_seed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UrlSeed)
		err := b.DecodeMessage(msg)
		m.Seed = &GenerateKeywordIdeasRequest_UrlSeed{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GenerateKeywordIdeasRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GenerateKeywordIdeasRequest)
	// seed
	switch x := m.Seed.(type) {
	case *GenerateKeywordIdeasRequest_KeywordAndUrlSeed:
		s := proto.Size(x.KeywordAndUrlSeed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GenerateKeywordIdeasRequest_KeywordSeed:
		s := proto.Size(x.KeywordSeed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GenerateKeywordIdeasRequest_UrlSeed:
		s := proto.Size(x.UrlSeed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Keyword And Url Seed
type KeywordAndUrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordAndUrlSeed) Reset()         { *m = KeywordAndUrlSeed{} }
func (m *KeywordAndUrlSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordAndUrlSeed) ProtoMessage()    {}
func (*KeywordAndUrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_8eed2c943de898df, []int{1}
}
func (m *KeywordAndUrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordAndUrlSeed.Unmarshal(m, b)
}
func (m *KeywordAndUrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordAndUrlSeed.Marshal(b, m, deterministic)
}
func (dst *KeywordAndUrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordAndUrlSeed.Merge(dst, src)
}
func (m *KeywordAndUrlSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordAndUrlSeed.Size(m)
}
func (m *KeywordAndUrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordAndUrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordAndUrlSeed proto.InternalMessageInfo

func (m *KeywordAndUrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *KeywordAndUrlSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Keyword Seed
type KeywordSeed struct {
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordSeed) Reset()         { *m = KeywordSeed{} }
func (m *KeywordSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordSeed) ProtoMessage()    {}
func (*KeywordSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_8eed2c943de898df, []int{2}
}
func (m *KeywordSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordSeed.Unmarshal(m, b)
}
func (m *KeywordSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordSeed.Marshal(b, m, deterministic)
}
func (dst *KeywordSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordSeed.Merge(dst, src)
}
func (m *KeywordSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordSeed.Size(m)
}
func (m *KeywordSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordSeed proto.InternalMessageInfo

func (m *KeywordSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Url Seed
type UrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UrlSeed) Reset()         { *m = UrlSeed{} }
func (m *UrlSeed) String() string { return proto.CompactTextString(m) }
func (*UrlSeed) ProtoMessage()    {}
func (*UrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_8eed2c943de898df, []int{3}
}
func (m *UrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlSeed.Unmarshal(m, b)
}
func (m *UrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlSeed.Marshal(b, m, deterministic)
}
func (dst *UrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlSeed.Merge(dst, src)
}
func (m *UrlSeed) XXX_Size() int {
	return xxx_messageInfo_UrlSeed.Size(m)
}
func (m *UrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_UrlSeed proto.InternalMessageInfo

func (m *UrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

// Response message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeaResponse struct {
	// Results of generating keyword ideas.
	Results              []*GenerateKeywordIdeaResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GenerateKeywordIdeaResponse) Reset()         { *m = GenerateKeywordIdeaResponse{} }
func (m *GenerateKeywordIdeaResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResponse) ProtoMessage()    {}
func (*GenerateKeywordIdeaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_8eed2c943de898df, []int{4}
}
func (m *GenerateKeywordIdeaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Marshal(b, m, deterministic)
}
func (dst *GenerateKeywordIdeaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResponse.Merge(dst, src)
}
func (m *GenerateKeywordIdeaResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Size(m)
}
func (m *GenerateKeywordIdeaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResponse proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResponse) GetResults() []*GenerateKeywordIdeaResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result of generating keyword ideas.
type GenerateKeywordIdeaResult struct {
	// Text of the keyword idea.
	// As in Keyword Plan historical metrics, this text may not be an actual
	// keyword, but the canonical form of multiple keywords.
	// See KeywordPlanKeywordHistoricalMetrics message in KeywordPlanService.
	Text *wrappers.StringValue `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// The historical metrics for the keyword
	KeywordIdeaMetrics   *common.KeywordPlanHistoricalMetrics `protobuf:"bytes,3,opt,name=keyword_idea_metrics,json=keywordIdeaMetrics,proto3" json:"keyword_idea_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *GenerateKeywordIdeaResult) Reset()         { *m = GenerateKeywordIdeaResult{} }
func (m *GenerateKeywordIdeaResult) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResult) ProtoMessage()    {}
func (*GenerateKeywordIdeaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_8eed2c943de898df, []int{5}
}
func (m *GenerateKeywordIdeaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Marshal(b, m, deterministic)
}
func (dst *GenerateKeywordIdeaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResult.Merge(dst, src)
}
func (m *GenerateKeywordIdeaResult) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Size(m)
}
func (m *GenerateKeywordIdeaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResult.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResult proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResult) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *GenerateKeywordIdeaResult) GetKeywordIdeaMetrics() *common.KeywordPlanHistoricalMetrics {
	if m != nil {
		return m.KeywordIdeaMetrics
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateKeywordIdeasRequest)(nil), "google.ads.googleads.v0.services.GenerateKeywordIdeasRequest")
	proto.RegisterType((*KeywordAndUrlSeed)(nil), "google.ads.googleads.v0.services.KeywordAndUrlSeed")
	proto.RegisterType((*KeywordSeed)(nil), "google.ads.googleads.v0.services.KeywordSeed")
	proto.RegisterType((*UrlSeed)(nil), "google.ads.googleads.v0.services.UrlSeed")
	proto.RegisterType((*GenerateKeywordIdeaResponse)(nil), "google.ads.googleads.v0.services.GenerateKeywordIdeaResponse")
	proto.RegisterType((*GenerateKeywordIdeaResult)(nil), "google.ads.googleads.v0.services.GenerateKeywordIdeaResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordPlanIdeaServiceClient is the client API for KeywordPlanIdeaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanIdeaServiceClient interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error)
}

type keywordPlanIdeaServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordPlanIdeaServiceClient(cc *grpc.ClientConn) KeywordPlanIdeaServiceClient {
	return &keywordPlanIdeaServiceClient{cc}
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error) {
	out := new(GenerateKeywordIdeaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.KeywordPlanIdeaService/GenerateKeywordIdeas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanIdeaServiceServer is the server API for KeywordPlanIdeaService service.
type KeywordPlanIdeaServiceServer interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(context.Context, *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error)
}

func RegisterKeywordPlanIdeaServiceServer(s *grpc.Server, srv KeywordPlanIdeaServiceServer) {
	s.RegisterService(&_KeywordPlanIdeaService_serviceDesc, srv)
}

func _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordIdeasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.KeywordPlanIdeaService/GenerateKeywordIdeas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, req.(*GenerateKeywordIdeasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanIdeaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.KeywordPlanIdeaService",
	HandlerType: (*KeywordPlanIdeaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKeywordIdeas",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/keyword_plan_idea_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/keyword_plan_idea_service.proto", fileDescriptor_keyword_plan_idea_service_8eed2c943de898df)
}

var fileDescriptor_keyword_plan_idea_service_8eed2c943de898df = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x4e, 0x13, 0x4d,
	0x14, 0xff, 0xb6, 0xe5, 0xa3, 0x30, 0xfd, 0xf2, 0x25, 0x4c, 0x88, 0xa9, 0x40, 0xb4, 0x69, 0xbc,
	0x40, 0x12, 0x67, 0x9b, 0x72, 0x21, 0x82, 0x4d, 0x2c, 0x46, 0x5b, 0x62, 0x24, 0x64, 0x11, 0x2e,
	0x4c, 0x93, 0xcd, 0xd0, 0x3d, 0x6c, 0x36, 0xdd, 0xce, 0xd4, 0x99, 0x59, 0xf0, 0x4f, 0xb8, 0xf1,
	0x15, 0x7c, 0x03, 0x2f, 0x7d, 0x07, 0x5f, 0xc0, 0x2b, 0x13, 0x2e, 0xbd, 0xf5, 0xda, 0x67, 0x30,
	0x3b, 0x3b, 0xd3, 0x16, 0x68, 0x2d, 0x72, 0x77, 0xf6, 0xfc, 0xf9, 0x9d, 0xdf, 0xf9, 0x33, 0x67,
	0xd1, 0x93, 0x90, 0xf3, 0x30, 0x06, 0x97, 0x06, 0xd2, 0xcd, 0xc4, 0x54, 0x3a, 0xa9, 0xba, 0x12,
	0xc4, 0x49, 0xd4, 0x01, 0xe9, 0x76, 0xe1, 0xdd, 0x29, 0x17, 0x81, 0xdf, 0x8f, 0x29, 0xf3, 0xa3,
	0x00, 0xa8, 0x6f, 0x4c, 0xa4, 0x2f, 0xb8, 0xe2, 0xb8, 0x9c, 0x85, 0x11, 0x1a, 0x48, 0x32, 0x40,
	0x20, 0x27, 0x55, 0x62, 0x11, 0x96, 0x36, 0x26, 0xe5, 0xe8, 0xf0, 0x5e, 0x8f, 0xb3, 0x8b, 0x19,
	0x32, 0x5d, 0x86, 0x3d, 0x39, 0x12, 0x58, 0xd2, 0xbb, 0x44, 0x8d, 0x81, 0x3a, 0xe5, 0xa2, 0x6b,
	0x22, 0x57, 0x6c, 0x64, 0x3f, 0x72, 0x29, 0x63, 0x5c, 0x51, 0x15, 0x71, 0x26, 0x8d, 0xf5, 0x8e,
	0xb1, 0xea, 0xaf, 0xa3, 0xe4, 0xd8, 0x3d, 0x15, 0xb4, 0xdf, 0x07, 0x61, 0xec, 0x95, 0xef, 0x33,
	0x68, 0xb9, 0x09, 0x0c, 0x04, 0x55, 0xf0, 0x22, 0x4b, 0xb2, 0x13, 0x00, 0x95, 0x1e, 0xbc, 0x49,
	0x40, 0x2a, 0x7c, 0x17, 0x15, 0x3b, 0x89, 0x54, 0xbc, 0x07, 0xc2, 0x8f, 0x82, 0x92, 0x53, 0x76,
	0x56, 0xe7, 0x3d, 0x64, 0x55, 0x3b, 0x01, 0xde, 0x40, 0x73, 0x31, 0x65, 0x61, 0x42, 0x43, 0x28,
	0x15, 0xca, 0xce, 0x6a, 0xb1, 0xb6, 0x62, 0x9a, 0x43, 0x6c, 0x4e, 0xb2, 0xaf, 0x44, 0xc4, 0xc2,
	0x43, 0x1a, 0x27, 0xe0, 0x0d, 0xbc, 0xf1, 0x2e, 0x5a, 0x0c, 0x81, 0xfb, 0x8a, 0x8a, 0x10, 0x94,
	0xdf, 0xe1, 0x4c, 0x2a, 0xca, 0x94, 0x2c, 0xcd, 0x95, 0xf3, 0x53, 0x51, 0x70, 0x08, 0xfc, 0x95,
	0x0e, 0x7c, 0x6a, 0xe3, 0xf0, 0x7b, 0xb4, 0x38, 0xae, 0x4d, 0xa5, 0xf9, 0xb2, 0xb3, 0xfa, 0x7f,
	0xad, 0x45, 0x26, 0x4d, 0x4f, 0x77, 0x98, 0x98, 0xe2, 0xf7, 0x62, 0xca, 0x76, 0xb3, 0xc0, 0x67,
	0x2c, 0xe9, 0x8d, 0x51, 0x7b, 0xb8, 0x7b, 0x45, 0x87, 0x8f, 0x87, 0xb9, 0x29, 0x0b, 0xfc, 0x44,
	0xc4, 0xbe, 0x04, 0x08, 0x4a, 0x39, 0xdd, 0x91, 0x75, 0x32, 0x6d, 0x73, 0x6c, 0x9e, 0x06, 0x0b,
	0x0e, 0x44, 0xbc, 0x0f, 0x10, 0xb4, 0xfe, 0xf1, 0x16, 0xba, 0x97, 0x95, 0xd8, 0x43, 0xff, 0xd9,
	0x3c, 0x1a, 0x3f, 0xaf, 0xf1, 0x1f, 0x5c, 0x1b, 0xdf, 0x20, 0x17, 0xbb, 0xc3, 0x4f, 0xfc, 0x1c,
	0xcd, 0x0d, 0xf8, 0xfe, 0xab, 0xf1, 0xee, 0x4f, 0xc7, 0x1b, 0xb2, 0x2c, 0x24, 0x99, 0xb8, 0x3d,
	0x8b, 0x66, 0x52, 0x8c, 0xca, 0x19, 0x5a, 0xb8, 0x52, 0x0d, 0x26, 0x28, 0x9f, 0x88, 0x58, 0xef,
	0xcf, 0xb4, 0xd9, 0xa6, 0x8e, 0xe9, 0x5a, 0x19, 0x8e, 0xb2, 0x94, 0xbb, 0xc6, 0x42, 0x0c, 0xbc,
	0x2b, 0x4d, 0x54, 0x1c, 0x29, 0xf6, 0x02, 0x90, 0xf3, 0x57, 0x40, 0x8f, 0x50, 0xe1, 0x86, 0xec,
	0x2b, 0x6a, 0xec, 0xa3, 0xf2, 0x40, 0xf6, 0x39, 0x93, 0x80, 0x0f, 0x50, 0x41, 0x80, 0x4c, 0x62,
	0x65, 0x29, 0x6d, 0x4d, 0x6f, 0xf8, 0x78, 0xbc, 0x24, 0x56, 0x9e, 0xc5, 0xaa, 0x7c, 0x75, 0xd0,
	0xed, 0x89, 0x6e, 0xb8, 0x8a, 0x66, 0x14, 0xbc, 0x55, 0x66, 0x25, 0xff, 0x5c, 0x84, 0xf6, 0xc4,
	0x6c, 0xb8, 0xd4, 0xfa, 0x1a, 0xf6, 0x40, 0x89, 0xa8, 0x23, 0xcd, 0xd2, 0x3d, 0x9e, 0xc8, 0xd9,
	0x1c, 0xb6, 0x91, 0xa7, 0xd3, 0x8a, 0xa4, 0xe2, 0x22, 0xea, 0xd0, 0xf8, 0x65, 0x86, 0x31, 0x78,
	0x44, 0x29, 0x41, 0xa3, 0xab, 0xfd, 0x72, 0xd0, 0xad, 0x91, 0xa0, 0xd4, 0xb4, 0x9f, 0x95, 0x8f,
	0xcf, 0x1d, 0xb4, 0x38, 0xee, 0x4c, 0xe1, 0xfa, 0x8d, 0x3a, 0x67, 0xcf, 0xdb, 0x52, 0xfd, 0xa6,
	0x8d, 0xd7, 0x83, 0xac, 0xd4, 0x3f, 0x9e, 0xff, 0xfc, 0x94, 0x7b, 0x58, 0xa9, 0xe9, 0x1b, 0x6f,
	0x8e, 0xa2, 0x74, 0x3f, 0x8c, 0x9c, 0xcc, 0xfa, 0xda, 0xd9, 0x66, 0x38, 0x86, 0xc1, 0xa6, 0xb3,
	0xb6, 0xfd, 0xc3, 0x41, 0xf7, 0x3a, 0xbc, 0x37, 0x95, 0xc3, 0xf6, 0xf2, 0xf8, 0xb6, 0xec, 0xa5,
	0xb3, 0xdb, 0x73, 0x5e, 0xb7, 0x0c, 0x40, 0xc8, 0xd3, 0xe3, 0x4a, 0xb8, 0x08, 0xdd, 0x10, 0x98,
	0x9e, 0xac, 0xfd, 0x99, 0xf4, 0x23, 0x39, 0xf9, 0xcf, 0xb7, 0x65, 0x85, 0xcf, 0xb9, 0x7c, 0xb3,
	0xd1, 0xf8, 0x92, 0x2b, 0x37, 0x33, 0xc0, 0x46, 0x20, 0x49, 0x26, 0xa6, 0xd2, 0x61, 0x95, 0x98,
	0xc4, 0xf2, 0x9b, 0x75, 0x69, 0x37, 0x02, 0xd9, 0x1e, 0xb8, 0xb4, 0x0f, 0xab, 0x6d, 0xeb, 0x72,
	0x34, 0xab, 0x09, 0xac, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xdd, 0x93, 0x1f, 0x79, 0x07,
	0x00, 0x00,
}
