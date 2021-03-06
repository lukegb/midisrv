// Code generated by protoc-gen-go. DO NOT EDIT.
// source: midisrv.proto

/*
Package midisrvpb is a generated protocol buffer package.

It is generated from these files:
	midisrv.proto

It has these top-level messages:
	MIDIWriteRequest
	SysExMIDIMessage
	MIDIEvents
	MIDIEvent
*/
package midisrvpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type MIDIWriteRequest struct {
	// Types that are valid to be assigned to Data:
	//	*MIDIWriteRequest_ImmediateEvent
	//	*MIDIWriteRequest_Events
	//	*MIDIWriteRequest_SysExMessage
	Data isMIDIWriteRequest_Data `protobuf_oneof:"data"`
}

func (m *MIDIWriteRequest) Reset()                    { *m = MIDIWriteRequest{} }
func (m *MIDIWriteRequest) String() string            { return proto.CompactTextString(m) }
func (*MIDIWriteRequest) ProtoMessage()               {}
func (*MIDIWriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isMIDIWriteRequest_Data interface {
	isMIDIWriteRequest_Data()
}

type MIDIWriteRequest_ImmediateEvent struct {
	ImmediateEvent *MIDIEvent `protobuf:"bytes,1,opt,name=immediateEvent,oneof"`
}
type MIDIWriteRequest_Events struct {
	Events *MIDIEvents `protobuf:"bytes,2,opt,name=events,oneof"`
}
type MIDIWriteRequest_SysExMessage struct {
	SysExMessage *SysExMIDIMessage `protobuf:"bytes,3,opt,name=sysExMessage,oneof"`
}

func (*MIDIWriteRequest_ImmediateEvent) isMIDIWriteRequest_Data() {}
func (*MIDIWriteRequest_Events) isMIDIWriteRequest_Data()         {}
func (*MIDIWriteRequest_SysExMessage) isMIDIWriteRequest_Data()   {}

func (m *MIDIWriteRequest) GetData() isMIDIWriteRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MIDIWriteRequest) GetImmediateEvent() *MIDIEvent {
	if x, ok := m.GetData().(*MIDIWriteRequest_ImmediateEvent); ok {
		return x.ImmediateEvent
	}
	return nil
}

func (m *MIDIWriteRequest) GetEvents() *MIDIEvents {
	if x, ok := m.GetData().(*MIDIWriteRequest_Events); ok {
		return x.Events
	}
	return nil
}

func (m *MIDIWriteRequest) GetSysExMessage() *SysExMIDIMessage {
	if x, ok := m.GetData().(*MIDIWriteRequest_SysExMessage); ok {
		return x.SysExMessage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MIDIWriteRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MIDIWriteRequest_OneofMarshaler, _MIDIWriteRequest_OneofUnmarshaler, _MIDIWriteRequest_OneofSizer, []interface{}{
		(*MIDIWriteRequest_ImmediateEvent)(nil),
		(*MIDIWriteRequest_Events)(nil),
		(*MIDIWriteRequest_SysExMessage)(nil),
	}
}

func _MIDIWriteRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MIDIWriteRequest)
	// data
	switch x := m.Data.(type) {
	case *MIDIWriteRequest_ImmediateEvent:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImmediateEvent); err != nil {
			return err
		}
	case *MIDIWriteRequest_Events:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Events); err != nil {
			return err
		}
	case *MIDIWriteRequest_SysExMessage:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SysExMessage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MIDIWriteRequest.Data has unexpected type %T", x)
	}
	return nil
}

func _MIDIWriteRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MIDIWriteRequest)
	switch tag {
	case 1: // data.immediateEvent
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MIDIEvent)
		err := b.DecodeMessage(msg)
		m.Data = &MIDIWriteRequest_ImmediateEvent{msg}
		return true, err
	case 2: // data.events
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MIDIEvents)
		err := b.DecodeMessage(msg)
		m.Data = &MIDIWriteRequest_Events{msg}
		return true, err
	case 3: // data.sysExMessage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SysExMIDIMessage)
		err := b.DecodeMessage(msg)
		m.Data = &MIDIWriteRequest_SysExMessage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MIDIWriteRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MIDIWriteRequest)
	// data
	switch x := m.Data.(type) {
	case *MIDIWriteRequest_ImmediateEvent:
		s := proto.Size(x.ImmediateEvent)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MIDIWriteRequest_Events:
		s := proto.Size(x.Events)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MIDIWriteRequest_SysExMessage:
		s := proto.Size(x.SysExMessage)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SysExMIDIMessage struct {
	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SysExMIDIMessage) Reset()                    { *m = SysExMIDIMessage{} }
func (m *SysExMIDIMessage) String() string            { return proto.CompactTextString(m) }
func (*SysExMIDIMessage) ProtoMessage()               {}
func (*SysExMIDIMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SysExMIDIMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SysExMIDIMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type MIDIEvents struct {
	Events []*MIDIEvent `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *MIDIEvents) Reset()                    { *m = MIDIEvents{} }
func (m *MIDIEvents) String() string            { return proto.CompactTextString(m) }
func (*MIDIEvents) ProtoMessage()               {}
func (*MIDIEvents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MIDIEvents) GetEvents() []*MIDIEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type MIDIEvent struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Status    int64 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Data1     int64 `protobuf:"varint,3,opt,name=data1" json:"data1,omitempty"`
	Data2     int64 `protobuf:"varint,4,opt,name=data2" json:"data2,omitempty"`
}

func (m *MIDIEvent) Reset()                    { *m = MIDIEvent{} }
func (m *MIDIEvent) String() string            { return proto.CompactTextString(m) }
func (*MIDIEvent) ProtoMessage()               {}
func (*MIDIEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MIDIEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MIDIEvent) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MIDIEvent) GetData1() int64 {
	if m != nil {
		return m.Data1
	}
	return 0
}

func (m *MIDIEvent) GetData2() int64 {
	if m != nil {
		return m.Data2
	}
	return 0
}

func init() {
	proto.RegisterType((*MIDIWriteRequest)(nil), "MIDIWriteRequest")
	proto.RegisterType((*SysExMIDIMessage)(nil), "SysExMIDIMessage")
	proto.RegisterType((*MIDIEvents)(nil), "MIDIEvents")
	proto.RegisterType((*MIDIEvent)(nil), "MIDIEvent")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MIDIService service

type MIDIServiceClient interface {
	// Plays a stream of events.
	Output(ctx context.Context, opts ...grpc.CallOption) (MIDIService_OutputClient, error)
}

type mIDIServiceClient struct {
	cc *grpc.ClientConn
}

func NewMIDIServiceClient(cc *grpc.ClientConn) MIDIServiceClient {
	return &mIDIServiceClient{cc}
}

func (c *mIDIServiceClient) Output(ctx context.Context, opts ...grpc.CallOption) (MIDIService_OutputClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MIDIService_serviceDesc.Streams[0], c.cc, "/MIDIService/Output", opts...)
	if err != nil {
		return nil, err
	}
	x := &mIDIServiceOutputClient{stream}
	return x, nil
}

type MIDIService_OutputClient interface {
	Send(*MIDIWriteRequest) error
	CloseAndRecv() (*google_protobuf.Empty, error)
	grpc.ClientStream
}

type mIDIServiceOutputClient struct {
	grpc.ClientStream
}

func (x *mIDIServiceOutputClient) Send(m *MIDIWriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mIDIServiceOutputClient) CloseAndRecv() (*google_protobuf.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(google_protobuf.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MIDIService service

type MIDIServiceServer interface {
	// Plays a stream of events.
	Output(MIDIService_OutputServer) error
}

func RegisterMIDIServiceServer(s *grpc.Server, srv MIDIServiceServer) {
	s.RegisterService(&_MIDIService_serviceDesc, srv)
}

func _MIDIService_Output_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MIDIServiceServer).Output(&mIDIServiceOutputServer{stream})
}

type MIDIService_OutputServer interface {
	SendAndClose(*google_protobuf.Empty) error
	Recv() (*MIDIWriteRequest, error)
	grpc.ServerStream
}

type mIDIServiceOutputServer struct {
	grpc.ServerStream
}

func (x *mIDIServiceOutputServer) SendAndClose(m *google_protobuf.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mIDIServiceOutputServer) Recv() (*MIDIWriteRequest, error) {
	m := new(MIDIWriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MIDIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MIDIService",
	HandlerType: (*MIDIServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Output",
			Handler:       _MIDIService_Output_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "midisrv.proto",
}

func init() { proto.RegisterFile("midisrv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x1b, 0x53, 0x03, 0x9d, 0x54, 0x69, 0x17, 0x29, 0xa1, 0x7a, 0x28, 0x01, 0xa1, 0xa7,
	0xad, 0x46, 0xc5, 0x7b, 0x49, 0xa1, 0x39, 0x14, 0x61, 0x7b, 0x10, 0xbc, 0xa5, 0x76, 0x2c, 0x0b,
	0xae, 0x89, 0xd9, 0x49, 0xb1, 0xff, 0xca, 0x9f, 0x28, 0xd9, 0x4d, 0x1b, 0xda, 0x8b, 0xb7, 0xcc,
	0x7b, 0x5f, 0x32, 0x6f, 0xf3, 0x16, 0x2e, 0x94, 0x5c, 0x4b, 0x5d, 0x6c, 0x79, 0x5e, 0x64, 0x94,
	0x0d, 0xaf, 0x37, 0x59, 0xb6, 0xf9, 0xc4, 0x89, 0x99, 0x56, 0xe5, 0xc7, 0x04, 0x55, 0x4e, 0x3b,
	0x6b, 0x86, 0xbf, 0x0e, 0xf4, 0x16, 0x49, 0x9c, 0xbc, 0x16, 0x92, 0x50, 0xe0, 0x77, 0x89, 0x9a,
	0xd8, 0x23, 0x5c, 0x4a, 0xa5, 0x70, 0x2d, 0x53, 0xc2, 0xd9, 0x16, 0xbf, 0x28, 0x70, 0x46, 0xce,
	0xd8, 0x8f, 0x80, 0x57, 0xa8, 0x51, 0xe6, 0x2d, 0x71, 0xc2, 0xb0, 0x5b, 0xf0, 0xb0, 0x7a, 0xd0,
	0xc1, 0x99, 0xa1, 0xfd, 0x86, 0xd6, 0xf3, 0x96, 0xa8, 0x4d, 0xf6, 0x0c, 0x5d, 0xbd, 0xd3, 0xb3,
	0x9f, 0x05, 0x6a, 0x9d, 0x6e, 0x30, 0x70, 0x0d, 0xdc, 0xe7, 0x4b, 0x23, 0x26, 0x71, 0x52, 0x1b,
	0xf3, 0x96, 0x38, 0x02, 0xa7, 0x1e, 0xb4, 0xd7, 0x29, 0xa5, 0x61, 0x0c, 0xbd, 0x53, 0x96, 0xdd,
	0x40, 0x87, 0xa4, 0x42, 0x4d, 0xa9, 0xca, 0x4d, 0x58, 0x57, 0x34, 0x02, 0x63, 0xf6, 0x4d, 0x93,
	0xab, 0x2b, 0xec, 0x57, 0xee, 0x00, 0x9a, 0x78, 0x2c, 0x3c, 0x64, 0x77, 0x46, 0xee, 0xf1, 0x49,
	0xf7, 0xc1, 0x43, 0x05, 0x9d, 0x83, 0xf8, 0xcf, 0xc2, 0x01, 0x78, 0x9a, 0x52, 0x2a, 0xed, 0xaf,
	0x70, 0x45, 0x3d, 0xb1, 0x2b, 0x38, 0xaf, 0x96, 0xdf, 0x9b, 0x43, 0xbb, 0xc2, 0x0e, 0x7b, 0x35,
	0x0a, 0xda, 0x8d, 0x1a, 0x45, 0x31, 0xf8, 0xd5, 0xba, 0x25, 0x16, 0x5b, 0xf9, 0x8e, 0xec, 0x09,
	0xbc, 0x97, 0x92, 0xf2, 0x92, 0x58, 0x9f, 0x9f, 0x16, 0x36, 0x1c, 0x70, 0xdb, 0x31, 0xdf, 0x77,
	0xcc, 0x67, 0x55, 0xc7, 0x63, 0x67, 0xea, 0xbf, 0x75, 0xea, 0xdb, 0x90, 0xaf, 0x56, 0x9e, 0xb1,
	0x1f, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xcd, 0x20, 0x4b, 0x21, 0x02, 0x00, 0x00,
}
