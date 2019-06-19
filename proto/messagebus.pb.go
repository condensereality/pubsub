// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messagebus.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// PublishRequest is passed when publishing
type PublishRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f5cf36a55407aca, []int{0}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// SubscribeRequest is passed when subscribing
type SubscribeRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f5cf36a55407aca, []int{1}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

// SubscribeResponse object
type SubscribeResponse struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f5cf36a55407aca, []int{2}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func (m *SubscribeResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*PublishRequest)(nil), "proto.PublishRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "proto.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "proto.SubscribeResponse")
}

func init() { proto.RegisterFile("messagebus.proto", fileDescriptor_9f5cf36a55407aca) }

var fileDescriptor_9f5cf36a55407aca = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x89, 0x50, 0x4b, 0x07, 0xd1, 0x1a, 0xfc, 0xb3, 0xac, 0x97, 0x65, 0x4f, 0x7b, 0x31,
	0x15, 0x45, 0xf0, 0x58, 0x16, 0x3c, 0x0a, 0xcb, 0xfa, 0x04, 0xc9, 0x3a, 0xc6, 0x40, 0x6a, 0xe2,
	0x4e, 0x72, 0xe8, 0x1b, 0xf8, 0xd8, 0xd2, 0x4d, 0x6b, 0x5b, 0x85, 0x9e, 0xc2, 0x47, 0x26, 0xbf,
	0xfc, 0xbe, 0x81, 0xe9, 0x02, 0x89, 0xa4, 0x46, 0x15, 0x49, 0xf8, 0xde, 0x05, 0xc7, 0x47, 0xc3,
	0x91, 0xdf, 0x68, 0xe7, 0xb4, 0xc5, 0xd9, 0x90, 0x54, 0x7c, 0x9f, 0xe1, 0xc2, 0x87, 0x65, 0x9a,
	0x29, 0xe7, 0x70, 0xda, 0x44, 0x65, 0x0d, 0x7d, 0xb4, 0xf8, 0x15, 0x91, 0x02, 0xbf, 0x80, 0x51,
	0x70, 0xde, 0x74, 0x19, 0x2b, 0x58, 0x35, 0x69, 0x53, 0xe0, 0x19, 0x8c, 0xbd, 0x5c, 0x5a, 0x27,
	0xdf, 0xb2, 0xa3, 0x82, 0x55, 0x27, 0xed, 0x26, 0x96, 0x15, 0x4c, 0x5f, 0xa3, 0xa2, 0xae, 0x37,
	0x0a, 0x0f, 0x32, 0xca, 0x5b, 0x38, 0xdf, 0x99, 0x24, 0xef, 0x3e, 0x09, 0x77, 0xc1, 0x6c, 0x0f,
	0x7c, 0xff, 0xcd, 0x00, 0x5e, 0x52, 0xa7, 0x3a, 0x12, 0x7f, 0x82, 0xf1, 0xda, 0x94, 0x5f, 0x26,
	0x79, 0xb1, 0x6f, 0x9e, 0x5f, 0x89, 0xd4, 0x54, 0x6c, 0x9a, 0x8a, 0xe7, 0x55, 0x53, 0x3e, 0x87,
	0xc9, 0xef, 0xbf, 0xfc, 0x7a, 0xfd, 0xf6, 0xaf, 0x73, 0x9e, 0xfd, 0xbf, 0x48, 0x8a, 0x77, 0xac,
	0x7e, 0x84, 0x42, 0x3b, 0xe9, 0x8d, 0x72, 0xc6, 0x62, 0xef, 0xad, 0x0c, 0x28, 0x74, 0xef, 0x3b,
	0xb1, 0xdd, 0x79, 0x7d, 0xb6, 0x75, 0x6d, 0x56, 0xa0, 0x86, 0xa9, 0xe3, 0x81, 0xf8, 0xf0, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xff, 0x93, 0xcd, 0x57, 0x9b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageBusClient is the client API for MessageBus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageBusClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (MessageBus_SubscribeClient, error)
}

type messageBusClient struct {
	cc *grpc.ClientConn
}

func NewMessageBusClient(cc *grpc.ClientConn) MessageBusClient {
	return &messageBusClient{cc}
}

func (c *messageBusClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.MessageBus/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageBusClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (MessageBus_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageBus_serviceDesc.Streams[0], "/proto.MessageBus/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageBusSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageBus_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type messageBusSubscribeClient struct {
	grpc.ClientStream
}

func (x *messageBusSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageBusServer is the server API for MessageBus service.
type MessageBusServer interface {
	Publish(context.Context, *PublishRequest) (*empty.Empty, error)
	Subscribe(*SubscribeRequest, MessageBus_SubscribeServer) error
}

func RegisterMessageBusServer(s *grpc.Server, srv MessageBusServer) {
	s.RegisterService(&_MessageBus_serviceDesc, srv)
}

func _MessageBus_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageBusServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessageBus/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageBusServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageBus_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageBusServer).Subscribe(m, &messageBusSubscribeServer{stream})
}

type MessageBus_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type messageBusSubscribeServer struct {
	grpc.ServerStream
}

func (x *messageBusSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MessageBus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MessageBus",
	HandlerType: (*MessageBusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _MessageBus_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _MessageBus_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "messagebus.proto",
}