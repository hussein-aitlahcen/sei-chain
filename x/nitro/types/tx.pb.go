// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nitro/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgRecordTransactionData struct {
	Sender    string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender"`
	Slot      uint64   `protobuf:"varint,2,opt,name=slot,proto3" json:"slot"`
	StateRoot string   `protobuf:"bytes,3,opt,name=stateRoot,proto3" json:"state_root"`
	Txs       []string `protobuf:"bytes,4,rep,name=txs,proto3" json:"txs"`
}

func (m *MsgRecordTransactionData) Reset()         { *m = MsgRecordTransactionData{} }
func (m *MsgRecordTransactionData) String() string { return proto.CompactTextString(m) }
func (*MsgRecordTransactionData) ProtoMessage()    {}
func (*MsgRecordTransactionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa387fcec7e73aca, []int{0}
}
func (m *MsgRecordTransactionData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRecordTransactionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRecordTransactionData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRecordTransactionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRecordTransactionData.Merge(m, src)
}
func (m *MsgRecordTransactionData) XXX_Size() int {
	return m.Size()
}
func (m *MsgRecordTransactionData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRecordTransactionData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRecordTransactionData proto.InternalMessageInfo

func (m *MsgRecordTransactionData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgRecordTransactionData) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *MsgRecordTransactionData) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func (m *MsgRecordTransactionData) GetTxs() []string {
	if m != nil {
		return m.Txs
	}
	return nil
}

type MsgSubmitFraudChallenge struct {
	Sender           string       `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender"`
	StartSlot        uint64       `protobuf:"varint,2,opt,name=startSlot,proto3" json:"start_slot"`
	EndSlot          uint64       `protobuf:"varint,3,opt,name=endSlot,proto3" json:"end_slot"`
	FraudStatePubKey string       `protobuf:"bytes,4,opt,name=fraudStatePubKey,proto3" json:"fraud_state_pub_key"`
	MerkleProof      *MerkleProof `protobuf:"bytes,5,opt,name=merkleProof,proto3" json:"merkle_proof"`
	AccountStates    []*Account   `protobuf:"bytes,6,rep,name=accountStates,proto3" json:"account_states"`
	Programs         []*Account   `protobuf:"bytes,7,rep,name=programs,proto3" json:"programs"`
	SysvarAccounts   []*Account   `protobuf:"bytes,8,rep,name=sysvarAccounts,proto3" json:"sysvar_accounts"`
}

func (m *MsgSubmitFraudChallenge) Reset()         { *m = MsgSubmitFraudChallenge{} }
func (m *MsgSubmitFraudChallenge) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitFraudChallenge) ProtoMessage()    {}
func (*MsgSubmitFraudChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa387fcec7e73aca, []int{1}
}
func (m *MsgSubmitFraudChallenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitFraudChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitFraudChallenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitFraudChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitFraudChallenge.Merge(m, src)
}
func (m *MsgSubmitFraudChallenge) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitFraudChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitFraudChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitFraudChallenge proto.InternalMessageInfo

func (m *MsgSubmitFraudChallenge) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgSubmitFraudChallenge) GetStartSlot() uint64 {
	if m != nil {
		return m.StartSlot
	}
	return 0
}

func (m *MsgSubmitFraudChallenge) GetEndSlot() uint64 {
	if m != nil {
		return m.EndSlot
	}
	return 0
}

func (m *MsgSubmitFraudChallenge) GetFraudStatePubKey() string {
	if m != nil {
		return m.FraudStatePubKey
	}
	return ""
}

func (m *MsgSubmitFraudChallenge) GetMerkleProof() *MerkleProof {
	if m != nil {
		return m.MerkleProof
	}
	return nil
}

func (m *MsgSubmitFraudChallenge) GetAccountStates() []*Account {
	if m != nil {
		return m.AccountStates
	}
	return nil
}

func (m *MsgSubmitFraudChallenge) GetPrograms() []*Account {
	if m != nil {
		return m.Programs
	}
	return nil
}

func (m *MsgSubmitFraudChallenge) GetSysvarAccounts() []*Account {
	if m != nil {
		return m.SysvarAccounts
	}
	return nil
}

type MsgRecordTransactionDataResponse struct {
}

func (m *MsgRecordTransactionDataResponse) Reset()         { *m = MsgRecordTransactionDataResponse{} }
func (m *MsgRecordTransactionDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRecordTransactionDataResponse) ProtoMessage()    {}
func (*MsgRecordTransactionDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa387fcec7e73aca, []int{2}
}
func (m *MsgRecordTransactionDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRecordTransactionDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRecordTransactionDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRecordTransactionDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRecordTransactionDataResponse.Merge(m, src)
}
func (m *MsgRecordTransactionDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRecordTransactionDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRecordTransactionDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRecordTransactionDataResponse proto.InternalMessageInfo

type MsgSubmitFraudChallengeResponse struct {
}

func (m *MsgSubmitFraudChallengeResponse) Reset()         { *m = MsgSubmitFraudChallengeResponse{} }
func (m *MsgSubmitFraudChallengeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitFraudChallengeResponse) ProtoMessage()    {}
func (*MsgSubmitFraudChallengeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa387fcec7e73aca, []int{3}
}
func (m *MsgSubmitFraudChallengeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitFraudChallengeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitFraudChallengeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitFraudChallengeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitFraudChallengeResponse.Merge(m, src)
}
func (m *MsgSubmitFraudChallengeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitFraudChallengeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitFraudChallengeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitFraudChallengeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRecordTransactionData)(nil), "seiprotocol.seichain.nitro.MsgRecordTransactionData")
	proto.RegisterType((*MsgSubmitFraudChallenge)(nil), "seiprotocol.seichain.nitro.MsgSubmitFraudChallenge")
	proto.RegisterType((*MsgRecordTransactionDataResponse)(nil), "seiprotocol.seichain.nitro.MsgRecordTransactionDataResponse")
	proto.RegisterType((*MsgSubmitFraudChallengeResponse)(nil), "seiprotocol.seichain.nitro.MsgSubmitFraudChallengeResponse")
}

func init() { proto.RegisterFile("nitro/tx.proto", fileDescriptor_fa387fcec7e73aca) }

var fileDescriptor_fa387fcec7e73aca = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x5d, 0x96, 0xb2, 0x75, 0xde, 0x28, 0x93, 0x37, 0xb4, 0x50, 0xa1, 0x24, 0x04, 0x09, 0x7a,
	0x80, 0x44, 0xea, 0xb8, 0xc1, 0x85, 0x0c, 0x21, 0x21, 0x54, 0x69, 0x72, 0x39, 0x81, 0x50, 0x70,
	0x13, 0x2f, 0x8d, 0xd6, 0xc6, 0x91, 0xed, 0xa0, 0xf6, 0x1b, 0x20, 0x71, 0xe1, 0x13, 0x70, 0xe3,
	0xa3, 0x20, 0x71, 0xdc, 0x91, 0x53, 0x84, 0xda, 0x5b, 0x3e, 0x05, 0x8a, 0x93, 0x50, 0xfe, 0xb4,
	0xa3, 0xda, 0xc5, 0xb5, 0x9f, 0xdf, 0x7b, 0x7e, 0xbf, 0xfe, 0xec, 0x80, 0x56, 0x1c, 0x09, 0x46,
	0x1d, 0x31, 0xb1, 0x13, 0x46, 0x05, 0x85, 0x6d, 0x4e, 0x22, 0x39, 0xf3, 0xe9, 0xc8, 0xe6, 0x24,
	0xf2, 0x87, 0x38, 0x8a, 0x6d, 0x49, 0x6a, 0x1f, 0x86, 0x34, 0xa4, 0x72, 0xd3, 0x29, 0x66, 0xa5,
	0xa2, 0xbd, 0x5f, 0x3a, 0x04, 0x58, 0xe0, 0x0a, 0x39, 0x28, 0x11, 0xec, 0xfb, 0x34, 0x8d, 0x45,
	0x09, 0x5a, 0x5f, 0x14, 0xa0, 0xf5, 0x78, 0x88, 0x88, 0x4f, 0x59, 0xf0, 0x8a, 0xe1, 0x98, 0x63,
	0x5f, 0x44, 0x34, 0x7e, 0x86, 0x05, 0x86, 0x16, 0xd8, 0xe2, 0x24, 0x0e, 0x08, 0xd3, 0x14, 0x53,
	0xe9, 0xec, 0xb8, 0x20, 0xcf, 0x8c, 0x0a, 0x41, 0xd5, 0x2f, 0xbc, 0x0d, 0x1a, 0x7c, 0x44, 0x85,
	0xb6, 0x69, 0x2a, 0x9d, 0x86, 0xdb, 0xcc, 0x33, 0x43, 0xae, 0x91, 0x1c, 0xe1, 0x03, 0xb0, 0xc3,
	0x05, 0x16, 0x04, 0x51, 0x2a, 0x34, 0x55, 0x9a, 0xb4, 0xf2, 0xcc, 0x00, 0x12, 0xf4, 0x18, 0xa5,
	0x02, 0x2d, 0x08, 0xf0, 0x16, 0x50, 0xc5, 0x84, 0x6b, 0x0d, 0x53, 0xed, 0xec, 0xb8, 0xdb, 0x79,
	0x66, 0x14, 0x4b, 0x54, 0x0c, 0xd6, 0xd7, 0x06, 0x38, 0xea, 0xf1, 0xb0, 0x9f, 0x0e, 0xc6, 0x91,
	0x78, 0xce, 0x70, 0x1a, 0x9c, 0x0c, 0xf1, 0x68, 0x44, 0xe2, 0x90, 0xac, 0x15, 0xb3, 0x0c, 0xc2,
	0x44, 0x7f, 0x91, 0xb5, 0x0e, 0xc2, 0x84, 0x27, 0x13, 0x2f, 0x08, 0xf0, 0x1e, 0xd8, 0x26, 0x71,
	0x20, 0xb9, 0xaa, 0xe4, 0xee, 0xe5, 0x99, 0xd1, 0x24, 0x71, 0x50, 0x32, 0xeb, 0x4d, 0x78, 0x02,
	0xf6, 0xcf, 0x8a, 0x2c, 0xfd, 0xa2, 0x84, 0xd3, 0x74, 0xf0, 0x92, 0x4c, 0xb5, 0x86, 0xcc, 0x70,
	0x94, 0x67, 0xc6, 0x81, 0xdc, 0xf3, 0xca, 0x5a, 0x93, 0x74, 0xe0, 0x9d, 0x93, 0x29, 0xfa, 0x47,
	0x00, 0xdf, 0x80, 0xdd, 0x31, 0x61, 0xe7, 0x23, 0x72, 0xca, 0x28, 0x3d, 0xd3, 0xae, 0x99, 0x4a,
	0x67, 0xb7, 0x7b, 0xdf, 0x5e, 0xdd, 0x71, 0xbb, 0xb7, 0xa0, 0xbb, 0xfb, 0x79, 0x66, 0xec, 0x95,
	0x7a, 0x2f, 0x29, 0x10, 0xf4, 0xbb, 0x1b, 0x7c, 0x0b, 0xae, 0x57, 0x0d, 0x97, 0x47, 0x72, 0x6d,
	0xcb, 0x54, 0x3b, 0xbb, 0xdd, 0xbb, 0x97, 0xd9, 0x3f, 0x2d, 0x05, 0x2e, 0xcc, 0x33, 0xa3, 0x55,
	0xa9, 0xcb, 0x2a, 0x38, 0xfa, 0xd3, 0x0d, 0xf6, 0x40, 0x33, 0x61, 0x34, 0x64, 0x78, 0xcc, 0xb5,
	0xed, 0xf5, 0x9d, 0xe5, 0xdf, 0x59, 0x0b, 0xd1, 0xaf, 0x19, 0x7c, 0x07, 0x5a, 0x7c, 0xca, 0xdf,
	0x63, 0x56, 0x11, 0xb9, 0xd6, 0x5c, 0xdf, 0xf4, 0x20, 0xcf, 0x8c, 0x1b, 0xa5, 0xdc, 0xab, 0x52,
	0x72, 0xf4, 0x97, 0x9f, 0x65, 0x01, 0x73, 0xd5, 0x75, 0x47, 0x84, 0x27, 0x34, 0xe6, 0xc4, 0xba,
	0x03, 0x8c, 0x15, 0x57, 0xad, 0xa6, 0x74, 0x3f, 0x6f, 0x02, 0xb5, 0xc7, 0x43, 0xf8, 0x51, 0x01,
	0x37, 0x97, 0xbf, 0x9d, 0x47, 0x97, 0x36, 0x70, 0x45, 0x84, 0xf6, 0x93, 0xab, 0xa8, 0xea, 0x54,
	0xf0, 0x83, 0x02, 0x0e, 0x97, 0xbe, 0x90, 0xe3, 0xff, 0xd8, 0x2e, 0x13, 0xb5, 0x1f, 0x5f, 0x41,
	0x54, 0x47, 0x71, 0x5f, 0x7c, 0x9b, 0xe9, 0xca, 0xc5, 0x4c, 0x57, 0x7e, 0xcc, 0x74, 0xe5, 0xd3,
	0x5c, 0xdf, 0xb8, 0x98, 0xeb, 0x1b, 0xdf, 0xe7, 0xfa, 0xc6, 0x6b, 0x27, 0x8c, 0xc4, 0x30, 0x1d,
	0xd8, 0x3e, 0x1d, 0x3b, 0x9c, 0x44, 0x0f, 0xeb, 0x13, 0xe4, 0x42, 0x1e, 0xe1, 0x4c, 0x9c, 0xea,
	0xf3, 0x37, 0x4d, 0x08, 0x1f, 0x6c, 0x49, 0xc6, 0xf1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0x52, 0xe3, 0x68, 0x14, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RecordTransactionData(ctx context.Context, in *MsgRecordTransactionData, opts ...grpc.CallOption) (*MsgRecordTransactionDataResponse, error)
	SubmitFraudChallenge(ctx context.Context, in *MsgSubmitFraudChallenge, opts ...grpc.CallOption) (*MsgSubmitFraudChallengeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RecordTransactionData(ctx context.Context, in *MsgRecordTransactionData, opts ...grpc.CallOption) (*MsgRecordTransactionDataResponse, error) {
	out := new(MsgRecordTransactionDataResponse)
	err := c.cc.Invoke(ctx, "/seiprotocol.seichain.nitro.Msg/RecordTransactionData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitFraudChallenge(ctx context.Context, in *MsgSubmitFraudChallenge, opts ...grpc.CallOption) (*MsgSubmitFraudChallengeResponse, error) {
	out := new(MsgSubmitFraudChallengeResponse)
	err := c.cc.Invoke(ctx, "/seiprotocol.seichain.nitro.Msg/SubmitFraudChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RecordTransactionData(context.Context, *MsgRecordTransactionData) (*MsgRecordTransactionDataResponse, error)
	SubmitFraudChallenge(context.Context, *MsgSubmitFraudChallenge) (*MsgSubmitFraudChallengeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RecordTransactionData(ctx context.Context, req *MsgRecordTransactionData) (*MsgRecordTransactionDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordTransactionData not implemented")
}
func (*UnimplementedMsgServer) SubmitFraudChallenge(ctx context.Context, req *MsgSubmitFraudChallenge) (*MsgSubmitFraudChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitFraudChallenge not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RecordTransactionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRecordTransactionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RecordTransactionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seiprotocol.seichain.nitro.Msg/RecordTransactionData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RecordTransactionData(ctx, req.(*MsgRecordTransactionData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitFraudChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitFraudChallenge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitFraudChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seiprotocol.seichain.nitro.Msg/SubmitFraudChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitFraudChallenge(ctx, req.(*MsgSubmitFraudChallenge))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seiprotocol.seichain.nitro.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordTransactionData",
			Handler:    _Msg_RecordTransactionData_Handler,
		},
		{
			MethodName: "SubmitFraudChallenge",
			Handler:    _Msg_SubmitFraudChallenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nitro/tx.proto",
}

func (m *MsgRecordTransactionData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRecordTransactionData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRecordTransactionData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Txs[iNdEx])
			copy(dAtA[i:], m.Txs[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Txs[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StateRoot) > 0 {
		i -= len(m.StateRoot)
		copy(dAtA[i:], m.StateRoot)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StateRoot)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Slot != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitFraudChallenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitFraudChallenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitFraudChallenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SysvarAccounts) > 0 {
		for iNdEx := len(m.SysvarAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SysvarAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Programs) > 0 {
		for iNdEx := len(m.Programs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Programs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AccountStates) > 0 {
		for iNdEx := len(m.AccountStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MerkleProof != nil {
		{
			size, err := m.MerkleProof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FraudStatePubKey) > 0 {
		i -= len(m.FraudStatePubKey)
		copy(dAtA[i:], m.FraudStatePubKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FraudStatePubKey)))
		i--
		dAtA[i] = 0x22
	}
	if m.EndSlot != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.EndSlot))
		i--
		dAtA[i] = 0x18
	}
	if m.StartSlot != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StartSlot))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRecordTransactionDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRecordTransactionDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRecordTransactionDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitFraudChallengeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitFraudChallengeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitFraudChallengeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRecordTransactionData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Slot != 0 {
		n += 1 + sovTx(uint64(m.Slot))
	}
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Txs) > 0 {
		for _, s := range m.Txs {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgSubmitFraudChallenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StartSlot != 0 {
		n += 1 + sovTx(uint64(m.StartSlot))
	}
	if m.EndSlot != 0 {
		n += 1 + sovTx(uint64(m.EndSlot))
	}
	l = len(m.FraudStatePubKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MerkleProof != nil {
		l = m.MerkleProof.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.AccountStates) > 0 {
		for _, e := range m.AccountStates {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Programs) > 0 {
		for _, e := range m.Programs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.SysvarAccounts) > 0 {
		for _, e := range m.SysvarAccounts {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgRecordTransactionDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitFraudChallengeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRecordTransactionData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRecordTransactionData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRecordTransactionData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitFraudChallenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitFraudChallenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitFraudChallenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartSlot", wireType)
			}
			m.StartSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartSlot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndSlot", wireType)
			}
			m.EndSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndSlot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FraudStatePubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FraudStatePubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MerkleProof == nil {
				m.MerkleProof = &MerkleProof{}
			}
			if err := m.MerkleProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountStates = append(m.AccountStates, &Account{})
			if err := m.AccountStates[len(m.AccountStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Programs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Programs = append(m.Programs, &Account{})
			if err := m.Programs[len(m.Programs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SysvarAccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SysvarAccounts = append(m.SysvarAccounts, &Account{})
			if err := m.SysvarAccounts[len(m.SysvarAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRecordTransactionDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRecordTransactionDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRecordTransactionDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitFraudChallengeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitFraudChallengeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitFraudChallengeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
