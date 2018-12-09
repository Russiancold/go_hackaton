// Code generated by protoc-gen-go. DO NOT EDIT.
// source: info.proto

package main

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OHLCV struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Time                 int32    `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	Interval             int32    `protobuf:"varint,3,opt,name=Interval,proto3" json:"Interval,omitempty"`
	Open                 float32  `protobuf:"fixed32,4,opt,name=Open,proto3" json:"Open,omitempty"`
	High                 float32  `protobuf:"fixed32,5,opt,name=High,proto3" json:"High,omitempty"`
	Low                  float32  `protobuf:"fixed32,6,opt,name=Low,proto3" json:"Low,omitempty"`
	Close                float32  `protobuf:"fixed32,7,opt,name=Close,proto3" json:"Close,omitempty"`
	Ticker               string   `protobuf:"bytes,9,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
	Volume               int64    `protobuf:"varint,10,opt,name=Volume,proto3" json:"Volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OHLCV) Reset()         { *m = OHLCV{} }
func (m *OHLCV) String() string { return proto.CompactTextString(m) }
func (*OHLCV) ProtoMessage()    {}
func (*OHLCV) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{0}
}

func (m *OHLCV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OHLCV.Unmarshal(m, b)
}
func (m *OHLCV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OHLCV.Marshal(b, m, deterministic)
}
func (m *OHLCV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OHLCV.Merge(m, src)
}
func (m *OHLCV) XXX_Size() int {
	return xxx_messageInfo_OHLCV.Size(m)
}
func (m *OHLCV) XXX_DiscardUnknown() {
	xxx_messageInfo_OHLCV.DiscardUnknown(m)
}

var xxx_messageInfo_OHLCV proto.InternalMessageInfo

func (m *OHLCV) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *OHLCV) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *OHLCV) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *OHLCV) GetOpen() float32 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *OHLCV) GetHigh() float32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *OHLCV) GetLow() float32 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *OHLCV) GetClose() float32 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *OHLCV) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *OHLCV) GetVolume() int64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type Deal struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BrokerID             int32    `protobuf:"varint,2,opt,name=BrokerID,proto3" json:"BrokerID,omitempty"`
	ClientID             int32    `protobuf:"varint,3,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	Ticker               string   `protobuf:"bytes,4,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
	Amount               int32    `protobuf:"varint,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Partial              bool     `protobuf:"varint,6,opt,name=Partial,proto3" json:"Partial,omitempty"`
	Time                 int32    `protobuf:"varint,7,opt,name=Time,proto3" json:"Time,omitempty"`
	Price                float32  `protobuf:"fixed32,8,opt,name=Price,proto3" json:"Price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deal) Reset()         { *m = Deal{} }
func (m *Deal) String() string { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()    {}
func (*Deal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{1}
}

func (m *Deal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deal.Unmarshal(m, b)
}
func (m *Deal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deal.Marshal(b, m, deterministic)
}
func (m *Deal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal.Merge(m, src)
}
func (m *Deal) XXX_Size() int {
	return xxx_messageInfo_Deal.Size(m)
}
func (m *Deal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal proto.InternalMessageInfo

func (m *Deal) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Deal) GetBrokerID() int32 {
	if m != nil {
		return m.BrokerID
	}
	return 0
}

func (m *Deal) GetClientID() int32 {
	if m != nil {
		return m.ClientID
	}
	return 0
}

func (m *Deal) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *Deal) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Deal) GetPartial() bool {
	if m != nil {
		return m.Partial
	}
	return false
}

func (m *Deal) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Deal) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type DealID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BrokerID             int64    `protobuf:"varint,2,opt,name=BrokerID,proto3" json:"BrokerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DealID) Reset()         { *m = DealID{} }
func (m *DealID) String() string { return proto.CompactTextString(m) }
func (*DealID) ProtoMessage()    {}
func (*DealID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{2}
}

func (m *DealID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealID.Unmarshal(m, b)
}
func (m *DealID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealID.Marshal(b, m, deterministic)
}
func (m *DealID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealID.Merge(m, src)
}
func (m *DealID) XXX_Size() int {
	return xxx_messageInfo_DealID.Size(m)
}
func (m *DealID) XXX_DiscardUnknown() {
	xxx_messageInfo_DealID.DiscardUnknown(m)
}

var xxx_messageInfo_DealID proto.InternalMessageInfo

func (m *DealID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *DealID) GetBrokerID() int64 {
	if m != nil {
		return m.BrokerID
	}
	return 0
}

type BrokerID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrokerID) Reset()         { *m = BrokerID{} }
func (m *BrokerID) String() string { return proto.CompactTextString(m) }
func (*BrokerID) ProtoMessage()    {}
func (*BrokerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{3}
}

func (m *BrokerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerID.Unmarshal(m, b)
}
func (m *BrokerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerID.Marshal(b, m, deterministic)
}
func (m *BrokerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerID.Merge(m, src)
}
func (m *BrokerID) XXX_Size() int {
	return xxx_messageInfo_BrokerID.Size(m)
}
func (m *BrokerID) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerID.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerID proto.InternalMessageInfo

func (m *BrokerID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type CancelResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelResult) Reset()         { *m = CancelResult{} }
func (m *CancelResult) String() string { return proto.CompactTextString(m) }
func (*CancelResult) ProtoMessage()    {}
func (*CancelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f140d5b28dddb141, []int{4}
}

func (m *CancelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelResult.Unmarshal(m, b)
}
func (m *CancelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelResult.Marshal(b, m, deterministic)
}
func (m *CancelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelResult.Merge(m, src)
}
func (m *CancelResult) XXX_Size() int {
	return xxx_messageInfo_CancelResult.Size(m)
}
func (m *CancelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelResult.DiscardUnknown(m)
}

var xxx_messageInfo_CancelResult proto.InternalMessageInfo

func (m *CancelResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*OHLCV)(nil), "OHLCV")
	proto.RegisterType((*Deal)(nil), "Deal")
	proto.RegisterType((*DealID)(nil), "DealID")
	proto.RegisterType((*BrokerID)(nil), "BrokerID")
	proto.RegisterType((*CancelResult)(nil), "CancelResult")
}

func init() { proto.RegisterFile("info.proto", fileDescriptor_f140d5b28dddb141) }

var fileDescriptor_f140d5b28dddb141 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x8a, 0xd5, 0x30,
	0x18, 0xc5, 0x6f, 0xee, 0xed, 0xdf, 0x0f, 0x15, 0x09, 0x83, 0x84, 0x6e, 0x2c, 0x59, 0x75, 0x55,
	0x44, 0x7d, 0x01, 0x6d, 0x85, 0x29, 0x0c, 0xcc, 0x10, 0x87, 0xd9, 0xc7, 0xf2, 0x39, 0x13, 0x26,
	0xb7, 0x19, 0xd2, 0x54, 0x7d, 0x0b, 0xdf, 0xc9, 0x8d, 0xaf, 0x25, 0x49, 0xdb, 0x6b, 0xb9, 0x2e,
	0xdc, 0x9d, 0xdf, 0xc9, 0x9f, 0xef, 0x9c, 0x10, 0x00, 0x35, 0x7c, 0x35, 0xf5, 0x93, 0x35, 0xce,
	0xf0, 0xdf, 0x04, 0xe2, 0xeb, 0xcb, 0xab, 0xe6, 0x8e, 0xbe, 0x80, 0x7d, 0xd7, 0x32, 0x52, 0x92,
	0xea, 0x20, 0xf6, 0x5d, 0x4b, 0x29, 0x44, 0xb7, 0xea, 0x88, 0x6c, 0x5f, 0x92, 0x2a, 0x16, 0x41,
	0xd3, 0x02, 0xb2, 0x6e, 0x70, 0x68, 0xbf, 0x49, 0xcd, 0x0e, 0xc1, 0x3f, 0xb1, 0xdf, 0x7f, 0xfd,
	0x84, 0x03, 0x8b, 0x4a, 0x52, 0xed, 0x45, 0xd0, 0xde, 0xbb, 0x54, 0xf7, 0x0f, 0x2c, 0x9e, 0x3d,
	0xaf, 0xe9, 0x4b, 0x38, 0x5c, 0x99, 0xef, 0x2c, 0x09, 0x96, 0x97, 0xf4, 0x02, 0xe2, 0x46, 0x9b,
	0x11, 0x59, 0x1a, 0xbc, 0x19, 0xe8, 0x2b, 0x48, 0x6e, 0x55, 0xff, 0x88, 0x96, 0xe5, 0x25, 0xa9,
	0x72, 0xb1, 0x90, 0xf7, 0xef, 0x8c, 0x9e, 0x8e, 0xc8, 0x20, 0x64, 0x5d, 0x88, 0xff, 0x22, 0x10,
	0xb5, 0x28, 0xf5, 0x3f, 0x45, 0x0a, 0xc8, 0x3e, 0x5a, 0xf3, 0x88, 0xb6, 0x6b, 0x97, 0x32, 0x27,
	0xf6, 0x6b, 0x8d, 0x56, 0x38, 0xb8, 0xae, 0x5d, 0x0b, 0xad, 0xbc, 0x09, 0x10, 0x9d, 0x07, 0xf8,
	0x70, 0x34, 0xd3, 0xe0, 0x42, 0xad, 0x58, 0x2c, 0x44, 0x19, 0xa4, 0x37, 0xd2, 0x3a, 0x25, 0x75,
	0x28, 0x97, 0x89, 0x15, 0x4f, 0x4f, 0x99, 0x6e, 0x9e, 0xf2, 0x02, 0xe2, 0x1b, 0xab, 0x7a, 0x64,
	0xd9, 0x5c, 0x3a, 0x00, 0x7f, 0x0f, 0x89, 0xef, 0xd0, 0xb5, 0xff, 0x6d, 0x71, 0xf8, 0xdb, 0x82,
	0x6f, 0xd6, 0xce, 0xcf, 0xf1, 0x0a, 0x9e, 0x35, 0x72, 0xe8, 0x51, 0x0b, 0x1c, 0x27, 0x1d, 0x52,
	0x8e, 0x53, 0xdf, 0xe3, 0x38, 0x86, 0x4d, 0x99, 0x58, 0xf1, 0xed, 0x4f, 0x02, 0xd9, 0xa7, 0x1f,
	0xfd, 0x83, 0x1c, 0xee, 0x91, 0x72, 0xc8, 0x3f, 0x3b, 0xe9, 0xd4, 0xe8, 0x54, 0x4f, 0xf3, 0x7a,
	0xbd, 0xbe, 0x48, 0xea, 0xf0, 0x5b, 0xf8, 0xee, 0x0d, 0xa1, 0x05, 0x24, 0x8d, 0x45, 0xe9, 0x90,
	0xc6, 0xb5, 0x4f, 0x5d, 0xa4, 0xf5, 0x1c, 0x9e, 0xef, 0x28, 0x87, 0x64, 0x1e, 0x4b, 0x57, 0xb3,
	0x78, 0x5e, 0x6f, 0x83, 0xf0, 0x1d, 0x7d, 0x0d, 0xe9, 0xac, 0xc7, 0xed, 0x84, 0xf9, 0x2e, 0x3f,
	0xe0, 0x4b, 0x12, 0xfe, 0xe8, 0xbb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xff, 0xb9, 0xe6,
	0xb1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeClient interface {
	// поток ценовых данных от биржи к брокеру
	// мы каждую секнуду будем получать отсюда событие с ценами, которые броке аггрегирует у себя в минуты и показывает клиентам
	// устанавливается 1 раз брокером
	Statistic(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (Exchange_StatisticClient, error)
	// отправка на биржу заявки от брокера
	Create(ctx context.Context, in *Deal, opts ...grpc.CallOption) (*DealID, error)
	// отмена заявки
	Cancel(ctx context.Context, in *DealID, opts ...grpc.CallOption) (*CancelResult, error)
	// исполнение заявок от биржи к брокеру
	// устанавливается 1 раз брокером и при исполнении какой-то заявки
	Results(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (Exchange_ResultsClient, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) Statistic(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (Exchange_StatisticClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Exchange_serviceDesc.Streams[0], "/Exchange/Statistic", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeStatisticClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Exchange_StatisticClient interface {
	Recv() (*OHLCV, error)
	grpc.ClientStream
}

type exchangeStatisticClient struct {
	grpc.ClientStream
}

func (x *exchangeStatisticClient) Recv() (*OHLCV, error) {
	m := new(OHLCV)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exchangeClient) Create(ctx context.Context, in *Deal, opts ...grpc.CallOption) (*DealID, error) {
	out := new(DealID)
	err := c.cc.Invoke(ctx, "/Exchange/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Cancel(ctx context.Context, in *DealID, opts ...grpc.CallOption) (*CancelResult, error) {
	out := new(CancelResult)
	err := c.cc.Invoke(ctx, "/Exchange/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Results(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (Exchange_ResultsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Exchange_serviceDesc.Streams[1], "/Exchange/Results", opts...)
	if err != nil {
		return nil, err
	}
	x := &exchangeResultsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Exchange_ResultsClient interface {
	Recv() (*Deal, error)
	grpc.ClientStream
}

type exchangeResultsClient struct {
	grpc.ClientStream
}

func (x *exchangeResultsClient) Recv() (*Deal, error) {
	m := new(Deal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExchangeServer is the server API for Exchange service.
type ExchangeServer interface {
	// поток ценовых данных от биржи к брокеру
	// мы каждую секнуду будем получать отсюда событие с ценами, которые броке аггрегирует у себя в минуты и показывает клиентам
	// устанавливается 1 раз брокером
	Statistic(*BrokerID, Exchange_StatisticServer) error
	// отправка на биржу заявки от брокера
	Create(context.Context, *Deal) (*DealID, error)
	// отмена заявки
	Cancel(context.Context, *DealID) (*CancelResult, error)
	// исполнение заявок от биржи к брокеру
	// устанавливается 1 раз брокером и при исполнении какой-то заявки
	Results(*BrokerID, Exchange_ResultsServer) error
}

func RegisterExchangeServer(s *grpc.Server, srv ExchangeServer) {
	s.RegisterService(&_Exchange_serviceDesc, srv)
}

func _Exchange_Statistic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BrokerID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExchangeServer).Statistic(m, &exchangeStatisticServer{stream})
}

type Exchange_StatisticServer interface {
	Send(*OHLCV) error
	grpc.ServerStream
}

type exchangeStatisticServer struct {
	grpc.ServerStream
}

func (x *exchangeStatisticServer) Send(m *OHLCV) error {
	return x.ServerStream.SendMsg(m)
}

func _Exchange_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Deal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Create(ctx, req.(*Deal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DealID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Exchange/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Cancel(ctx, req.(*DealID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Results_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BrokerID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExchangeServer).Results(m, &exchangeResultsServer{stream})
}

type Exchange_ResultsServer interface {
	Send(*Deal) error
	grpc.ServerStream
}

type exchangeResultsServer struct {
	grpc.ServerStream
}

func (x *exchangeResultsServer) Send(m *Deal) error {
	return x.ServerStream.SendMsg(m)
}

var _Exchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Exchange_Create_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Exchange_Cancel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Statistic",
			Handler:       _Exchange_Statistic_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Results",
			Handler:       _Exchange_Results_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "info.proto",
}
