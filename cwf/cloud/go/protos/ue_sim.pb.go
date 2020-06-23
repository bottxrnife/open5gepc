// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cwf/protos/ue_sim.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
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

// --------------------------------------------------------------------------
// UE config
// --------------------------------------------------------------------------
type UEConfig struct {
	// Unique identifier for the UE.
	Imsi string `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	// Authentication key (k).
	AuthKey []byte `protobuf:"bytes,2,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
	// Operator configuration field (Op) signed with authentication key (k).
	AuthOpc []byte `protobuf:"bytes,3,opt,name=auth_opc,json=authOpc,proto3" json:"auth_opc,omitempty"`
	// Sequence Number (SEQ).
	Seq                  uint64   `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UEConfig) Reset()         { *m = UEConfig{} }
func (m *UEConfig) String() string { return proto.CompactTextString(m) }
func (*UEConfig) ProtoMessage()    {}
func (*UEConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{0}
}

func (m *UEConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UEConfig.Unmarshal(m, b)
}
func (m *UEConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UEConfig.Marshal(b, m, deterministic)
}
func (m *UEConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UEConfig.Merge(m, src)
}
func (m *UEConfig) XXX_Size() int {
	return xxx_messageInfo_UEConfig.Size(m)
}
func (m *UEConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UEConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UEConfig proto.InternalMessageInfo

func (m *UEConfig) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *UEConfig) GetAuthKey() []byte {
	if m != nil {
		return m.AuthKey
	}
	return nil
}

func (m *UEConfig) GetAuthOpc() []byte {
	if m != nil {
		return m.AuthOpc
	}
	return nil
}

func (m *UEConfig) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type AuthenticateRequest struct {
	Imsi                 string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	CalledStationID      string   `protobuf:"bytes,2,opt,name=calledStationID,proto3" json:"calledStationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{1}
}

func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(m, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *AuthenticateRequest) GetCalledStationID() string {
	if m != nil {
		return m.CalledStationID
	}
	return ""
}

type AuthenticateResponse struct {
	RadiusPacket         []byte   `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{2}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetRadiusPacket() []byte {
	if m != nil {
		return m.RadiusPacket
	}
	return nil
}

type DisconnectRequest struct {
	Imsi                 string   `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	CalledStationID      string   `protobuf:"bytes,2,opt,name=calledStationID,proto3" json:"calledStationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{3}
}

func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (m *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(m, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *DisconnectRequest) GetCalledStationID() string {
	if m != nil {
		return m.CalledStationID
	}
	return ""
}

type DisconnectResponse struct {
	RadiusPacket         []byte   `protobuf:"bytes,1,opt,name=radiusPacket,proto3" json:"radiusPacket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectResponse) Reset()         { *m = DisconnectResponse{} }
func (m *DisconnectResponse) String() string { return proto.CompactTextString(m) }
func (*DisconnectResponse) ProtoMessage()    {}
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{4}
}

func (m *DisconnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectResponse.Unmarshal(m, b)
}
func (m *DisconnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectResponse.Marshal(b, m, deterministic)
}
func (m *DisconnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectResponse.Merge(m, src)
}
func (m *DisconnectResponse) XXX_Size() int {
	return xxx_messageInfo_DisconnectResponse.Size(m)
}
func (m *DisconnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectResponse proto.InternalMessageInfo

func (m *DisconnectResponse) GetRadiusPacket() []byte {
	if m != nil {
		return m.RadiusPacket
	}
	return nil
}

type GenTrafficRequest struct {
	Imsi                    string                `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Volume                  *wrappers.StringValue `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Bitrate                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	TimeInSecs              uint64                `protobuf:"varint,4,opt,name=timeInSecs,proto3" json:"timeInSecs,omitempty"`
	ReportingIntervalInSecs uint64                `protobuf:"varint,5,opt,name=reportingIntervalInSecs,proto3" json:"reportingIntervalInSecs,omitempty"`
	ReverseMode             bool                  `protobuf:"varint,6,opt,name=reverseMode,proto3" json:"reverseMode,omitempty"`
	Timeout                 uint32                `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *GenTrafficRequest) Reset()         { *m = GenTrafficRequest{} }
func (m *GenTrafficRequest) String() string { return proto.CompactTextString(m) }
func (*GenTrafficRequest) ProtoMessage()    {}
func (*GenTrafficRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{5}
}

func (m *GenTrafficRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenTrafficRequest.Unmarshal(m, b)
}
func (m *GenTrafficRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenTrafficRequest.Marshal(b, m, deterministic)
}
func (m *GenTrafficRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenTrafficRequest.Merge(m, src)
}
func (m *GenTrafficRequest) XXX_Size() int {
	return xxx_messageInfo_GenTrafficRequest.Size(m)
}
func (m *GenTrafficRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenTrafficRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenTrafficRequest proto.InternalMessageInfo

func (m *GenTrafficRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *GenTrafficRequest) GetVolume() *wrappers.StringValue {
	if m != nil {
		return m.Volume
	}
	return nil
}

func (m *GenTrafficRequest) GetBitrate() *wrappers.StringValue {
	if m != nil {
		return m.Bitrate
	}
	return nil
}

func (m *GenTrafficRequest) GetTimeInSecs() uint64 {
	if m != nil {
		return m.TimeInSecs
	}
	return 0
}

func (m *GenTrafficRequest) GetReportingIntervalInSecs() uint64 {
	if m != nil {
		return m.ReportingIntervalInSecs
	}
	return 0
}

func (m *GenTrafficRequest) GetReverseMode() bool {
	if m != nil {
		return m.ReverseMode
	}
	return false
}

func (m *GenTrafficRequest) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type GenTrafficResponse struct {
	Output               []byte   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenTrafficResponse) Reset()         { *m = GenTrafficResponse{} }
func (m *GenTrafficResponse) String() string { return proto.CompactTextString(m) }
func (*GenTrafficResponse) ProtoMessage()    {}
func (*GenTrafficResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bc05ea16f96cbc, []int{6}
}

func (m *GenTrafficResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenTrafficResponse.Unmarshal(m, b)
}
func (m *GenTrafficResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenTrafficResponse.Marshal(b, m, deterministic)
}
func (m *GenTrafficResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenTrafficResponse.Merge(m, src)
}
func (m *GenTrafficResponse) XXX_Size() int {
	return xxx_messageInfo_GenTrafficResponse.Size(m)
}
func (m *GenTrafficResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenTrafficResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenTrafficResponse proto.InternalMessageInfo

func (m *GenTrafficResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*UEConfig)(nil), "magma.cwf.UEConfig")
	proto.RegisterType((*AuthenticateRequest)(nil), "magma.cwf.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "magma.cwf.AuthenticateResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "magma.cwf.DisconnectRequest")
	proto.RegisterType((*DisconnectResponse)(nil), "magma.cwf.DisconnectResponse")
	proto.RegisterType((*GenTrafficRequest)(nil), "magma.cwf.GenTrafficRequest")
	proto.RegisterType((*GenTrafficResponse)(nil), "magma.cwf.GenTrafficResponse")
}

func init() { proto.RegisterFile("cwf/protos/ue_sim.proto", fileDescriptor_01bc05ea16f96cbc) }

var fileDescriptor_01bc05ea16f96cbc = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xfb, 0x91, 0xb4, 0xd3, 0x20, 0xe8, 0x16, 0x51, 0x37, 0x94, 0x62, 0xf9, 0xe4, 0x03,
	0x72, 0x44, 0x41, 0xa8, 0xe2, 0x56, 0x68, 0x85, 0xaa, 0x0a, 0x41, 0x1d, 0xda, 0x03, 0x97, 0x6a,
	0xb3, 0x1e, 0xbb, 0xab, 0xda, 0xbb, 0xee, 0x7e, 0x24, 0xea, 0x2f, 0xe2, 0xc0, 0x9f, 0x44, 0x71,
	0xec, 0xc6, 0x21, 0x0d, 0x08, 0x89, 0x93, 0x3d, 0xf3, 0xde, 0xbe, 0x19, 0xcd, 0xbc, 0x5d, 0xd8,
	0x61, 0xa3, 0xa4, 0x57, 0x28, 0x69, 0xa4, 0xee, 0x59, 0xbc, 0xd2, 0x3c, 0x0f, 0xcb, 0x88, 0x6c,
	0xe4, 0x34, 0xcd, 0x69, 0xc8, 0x46, 0x49, 0x77, 0x57, 0x2a, 0x76, 0xa8, 0x6a, 0x16, 0x93, 0x79,
	0x2e, 0xc5, 0x84, 0xd5, 0xdd, 0x4f, 0xa5, 0x4c, 0x33, 0x9c, 0x60, 0x03, 0x9b, 0xf4, 0x46, 0x8a,
	0x16, 0x05, 0x2a, 0x3d, 0xc1, 0xfd, 0x04, 0xd6, 0x2f, 0x4e, 0x3e, 0x4a, 0x91, 0xf0, 0x94, 0x10,
	0x58, 0xe5, 0xb9, 0xe6, 0xae, 0xe3, 0x39, 0xc1, 0x46, 0x54, 0xfe, 0x93, 0x5d, 0x58, 0xa7, 0xd6,
	0x5c, 0x5f, 0xdd, 0xe0, 0x9d, 0xbb, 0xec, 0x39, 0x41, 0x27, 0x6a, 0x8f, 0xe3, 0x33, 0xbc, 0xbb,
	0x87, 0x64, 0xc1, 0xdc, 0x95, 0x29, 0xf4, 0xa5, 0x60, 0xe4, 0x09, 0xac, 0x68, 0xbc, 0x75, 0x57,
	0x3d, 0x27, 0x58, 0x8d, 0xc6, 0xbf, 0x7e, 0x1f, 0xb6, 0x8f, 0xac, 0xb9, 0x46, 0x61, 0x38, 0xa3,
	0x06, 0x23, 0xbc, 0xb5, 0xa8, 0xcd, 0x83, 0x25, 0x03, 0x78, 0xcc, 0x68, 0x96, 0x61, 0xdc, 0x37,
	0xd4, 0x70, 0x29, 0x4e, 0x8f, 0xcb, 0xca, 0x1b, 0xd1, 0xef, 0x69, 0xff, 0x3d, 0x3c, 0x9d, 0x15,
	0xd5, 0x85, 0x14, 0x1a, 0x89, 0x0f, 0x1d, 0x45, 0x63, 0x6e, 0xf5, 0x57, 0xca, 0x6e, 0xd0, 0x94,
	0xea, 0x9d, 0x68, 0x26, 0xe7, 0x9f, 0xc3, 0xd6, 0x31, 0xd7, 0x4c, 0x0a, 0x81, 0xcc, 0xfc, 0x9f,
	0x76, 0x0e, 0x81, 0x34, 0x25, 0xff, 0xa1, 0x99, 0x9f, 0xcb, 0xb0, 0xf5, 0x09, 0xc5, 0x37, 0x45,
	0x93, 0x84, 0xb3, 0x3f, 0x75, 0xf3, 0x16, 0x5a, 0x43, 0x99, 0xd9, 0x1c, 0xcb, 0x26, 0x36, 0x0f,
	0xf6, 0xc2, 0xc9, 0x82, 0xc3, 0x7a, 0xc1, 0x61, 0xdf, 0x28, 0x2e, 0xd2, 0x4b, 0x9a, 0x59, 0x8c,
	0x2a, 0x2e, 0x79, 0x07, 0xed, 0x01, 0x37, 0x8a, 0x1a, 0x2c, 0x37, 0xf5, 0xb7, 0x63, 0x35, 0x99,
	0xec, 0x03, 0x18, 0x9e, 0xe3, 0xa9, 0xe8, 0x23, 0xd3, 0xd5, 0x3a, 0x1b, 0x19, 0x72, 0x08, 0x3b,
	0x0a, 0x0b, 0xa9, 0x0c, 0x17, 0xe9, 0xa9, 0x30, 0xa8, 0x86, 0x34, 0xab, 0xc8, 0x6b, 0x25, 0x79,
	0x11, 0x4c, 0x3c, 0xd8, 0x54, 0x38, 0x44, 0xa5, 0xf1, 0xb3, 0x8c, 0xd1, 0x6d, 0x79, 0x4e, 0xb0,
	0x1e, 0x35, 0x53, 0xc4, 0x85, 0xf6, 0xb8, 0x92, 0xb4, 0xc6, 0x6d, 0x7b, 0x4e, 0xf0, 0x28, 0xaa,
	0x43, 0xff, 0x15, 0x90, 0xe6, 0xb0, 0xaa, 0x39, 0x3f, 0x83, 0x96, 0xb4, 0xa6, 0xb0, 0xf5, 0x84,
	0xab, 0xe8, 0xe0, 0xc7, 0x32, 0xac, 0x5d, 0x9c, 0xf4, 0x79, 0x4e, 0x5e, 0xc3, 0xda, 0x51, 0x1c,
	0x5f, 0x9c, 0x90, 0xed, 0xf0, 0xfe, 0xee, 0x84, 0xb5, 0xfb, 0xbb, 0x5b, 0x55, 0xb2, 0xbc, 0x4b,
	0xe1, 0xa5, 0xe4, 0xb1, 0xbf, 0x44, 0xce, 0xa1, 0xd3, 0x74, 0x18, 0xd9, 0x6f, 0x9c, 0x7c, 0xc0,
	0xcf, 0xdd, 0x97, 0x0b, 0xf1, 0x49, 0x97, 0xfe, 0x12, 0x39, 0x03, 0x98, 0xba, 0x84, 0xec, 0x35,
	0x0e, 0xcc, 0xf9, 0xb1, 0xfb, 0x62, 0x01, 0xda, 0x14, 0x9b, 0x8e, 0x62, 0x46, 0x6c, 0xce, 0x4e,
	0x33, 0x62, 0xf3, 0xf3, 0xf3, 0x97, 0x3e, 0x3c, 0xff, 0xbe, 0x5b, 0x32, 0x7a, 0xe3, 0x27, 0x87,
	0x65, 0xd2, 0xc6, 0xbd, 0x54, 0x56, 0xaf, 0xca, 0xa0, 0x55, 0x7e, 0xdf, 0xfc, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x06, 0x6e, 0x6e, 0x90, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UESimClient is the client API for UESim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UESimClient interface {
	// Adds a new UE to the store.
	//
	AddUE(ctx context.Context, in *UEConfig, opts ...grpc.CallOption) (*protos.Void, error)
	// Triggers an authentication for the UE with the specified imsi.
	//
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	// Triggers iperf traffic towards the CWAG
	GenTraffic(ctx context.Context, in *GenTrafficRequest, opts ...grpc.CallOption) (*GenTrafficResponse, error)
}

type uESimClient struct {
	cc grpc.ClientConnInterface
}

func NewUESimClient(cc grpc.ClientConnInterface) UESimClient {
	return &uESimClient{cc}
}

func (c *uESimClient) AddUE(ctx context.Context, in *UEConfig, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/AddUE", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uESimClient) GenTraffic(ctx context.Context, in *GenTrafficRequest, opts ...grpc.CallOption) (*GenTrafficResponse, error) {
	out := new(GenTrafficResponse)
	err := c.cc.Invoke(ctx, "/magma.cwf.UESim/GenTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UESimServer is the server API for UESim service.
type UESimServer interface {
	// Adds a new UE to the store.
	//
	AddUE(context.Context, *UEConfig) (*protos.Void, error)
	// Triggers an authentication for the UE with the specified imsi.
	//
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	// Triggers iperf traffic towards the CWAG
	GenTraffic(context.Context, *GenTrafficRequest) (*GenTrafficResponse, error)
}

// UnimplementedUESimServer can be embedded to have forward compatible implementations.
type UnimplementedUESimServer struct {
}

func (*UnimplementedUESimServer) AddUE(ctx context.Context, req *UEConfig) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUE not implemented")
}
func (*UnimplementedUESimServer) Authenticate(ctx context.Context, req *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (*UnimplementedUESimServer) Disconnect(ctx context.Context, req *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (*UnimplementedUESimServer) GenTraffic(ctx context.Context, req *GenTrafficRequest) (*GenTrafficResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenTraffic not implemented")
}

func RegisterUESimServer(s *grpc.Server, srv UESimServer) {
	s.RegisterService(&_UESim_serviceDesc, srv)
}

func _UESim_AddUE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UEConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).AddUE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/AddUE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).AddUE(ctx, req.(*UEConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UESim_GenTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UESimServer).GenTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.cwf.UESim/GenTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UESimServer).GenTraffic(ctx, req.(*GenTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UESim_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.cwf.UESim",
	HandlerType: (*UESimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUE",
			Handler:    _UESim_AddUE_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UESim_Authenticate_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _UESim_Disconnect_Handler,
		},
		{
			MethodName: "GenTraffic",
			Handler:    _UESim_GenTraffic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cwf/protos/ue_sim.proto",
}
