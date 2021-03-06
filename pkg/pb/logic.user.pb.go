// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logic.user.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Sex                  int32    `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Extra                string   `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	CreateTime           int64    `protobuf:"varint,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           int64    `protobuf:"varint,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Account              string   `protobuf:"bytes,8,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d96a09ab7af0adb, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *User) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *User) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *User) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type GetUserReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReq) Reset()         { *m = GetUserReq{} }
func (m *GetUserReq) String() string { return proto.CompactTextString(m) }
func (*GetUserReq) ProtoMessage()    {}
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d96a09ab7af0adb, []int{1}
}

func (m *GetUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReq.Unmarshal(m, b)
}
func (m *GetUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReq.Marshal(b, m, deterministic)
}
func (m *GetUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReq.Merge(m, src)
}
func (m *GetUserReq) XXX_Size() int {
	return xxx_messageInfo_GetUserReq.Size(m)
}
func (m *GetUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReq proto.InternalMessageInfo

func (m *GetUserReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type GetUserResp struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResp) Reset()         { *m = GetUserResp{} }
func (m *GetUserResp) String() string { return proto.CompactTextString(m) }
func (*GetUserResp) ProtoMessage()    {}
func (*GetUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d96a09ab7af0adb, []int{2}
}

func (m *GetUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResp.Unmarshal(m, b)
}
func (m *GetUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResp.Marshal(b, m, deterministic)
}
func (m *GetUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResp.Merge(m, src)
}
func (m *GetUserResp) XXX_Size() int {
	return xxx_messageInfo_GetUserResp.Size(m)
}
func (m *GetUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResp proto.InternalMessageInfo

func (m *GetUserResp) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*GetUserReq)(nil), "pb.GetUserReq")
	proto.RegisterType((*GetUserResp)(nil), "pb.GetUserResp")
}

func init() {
	proto.RegisterFile("logic.user.proto", fileDescriptor_8d96a09ab7af0adb)
}

var fileDescriptor_8d96a09ab7af0adb = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4d, 0x4b, 0xf4, 0x40,
	0x10, 0x84, 0x99, 0xdd, 0xcd, 0x57, 0x87, 0xf7, 0x75, 0x19, 0x04, 0x87, 0x45, 0x31, 0xec, 0x41,
	0x02, 0x42, 0x90, 0xf5, 0xec, 0x41, 0x2f, 0x22, 0x78, 0x0a, 0xee, 0x39, 0x4c, 0x92, 0x46, 0x82,
	0xf9, 0x18, 0x27, 0x13, 0x59, 0xff, 0xac, 0xbf, 0x45, 0xa6, 0x07, 0xa3, 0x7b, 0xeb, 0xaa, 0x7a,
	0xe8, 0xa6, 0x0b, 0xd6, 0xed, 0xf0, 0xda, 0x54, 0xd9, 0x34, 0xa2, 0xce, 0x94, 0x1e, 0xcc, 0xc0,
	0x17, 0xaa, 0xdc, 0x7e, 0x31, 0x58, 0xed, 0x47, 0xd4, 0xfc, 0x0c, 0x02, 0x1b, 0x15, 0x4d, 0x2d,
	0x58, 0xc2, 0xd2, 0x65, 0xee, 0x5b, 0xf9, 0x54, 0xf3, 0x0d, 0x84, 0x7d, 0x53, 0xbd, 0xf5, 0xb2,
	0x43, 0xb1, 0x48, 0x58, 0x1a, 0xe5, 0xb3, 0xe6, 0x6b, 0x58, 0x8e, 0x78, 0x10, 0xcb, 0x84, 0xa5,
	0x5e, 0x6e, 0x47, 0x7e, 0x01, 0x20, 0x3f, 0xa4, 0x91, 0xba, 0x98, 0x74, 0x2b, 0x56, 0xc4, 0x47,
	0xce, 0xd9, 0xeb, 0x96, 0x9f, 0x82, 0x87, 0x07, 0xa3, 0xa5, 0xf0, 0x28, 0x71, 0x82, 0x5f, 0x42,
	0x5c, 0x69, 0x94, 0x06, 0x0b, 0xd3, 0x74, 0x28, 0x7c, 0xba, 0x0f, 0xce, 0x7a, 0x69, 0x3a, 0xb4,
	0xc0, 0xa4, 0xea, 0x19, 0x08, 0x1c, 0xe0, 0x2c, 0x02, 0x04, 0x04, 0xb2, 0xaa, 0x86, 0xa9, 0x37,
	0x22, 0xa4, 0xcd, 0x3f, 0x72, 0x7b, 0x05, 0xf0, 0x88, 0xc6, 0xbe, 0x98, 0xe3, 0xfb, 0x5f, 0x8e,
	0x1d, 0x73, 0xd7, 0x10, 0xcf, 0xdc, 0xa8, 0xf8, 0x39, 0xac, 0xec, 0xff, 0x44, 0xc5, 0xbb, 0x30,
	0x53, 0x65, 0x46, 0x19, 0xb9, 0xbb, 0x3b, 0x88, 0x9e, 0x6d, 0x9b, 0xd4, 0xdc, 0x0d, 0xfc, 0xbb,
	0x77, 0x4b, 0x1e, 0x3e, 0xc9, 0xf8, 0x6f, 0xe9, 0xdf, 0xa3, 0x9b, 0x93, 0x23, 0x3d, 0xaa, 0xd2,
	0xa7, 0xfe, 0x6f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x09, 0x9a, 0x08, 0x54, 0x93, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogicUserClient is the client API for LogicUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicUserClient interface {
	//  登录
	AccountByUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
}

type logicUserClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicUserClient(cc grpc.ClientConnInterface) LogicUserClient {
	return &logicUserClient{cc}
}

func (c *logicUserClient) AccountByUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/pb.LogicUser/AccountByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicUserServer is the server API for LogicUser service.
type LogicUserServer interface {
	//  登录
	AccountByUser(context.Context, *GetUserReq) (*GetUserResp, error)
}

// UnimplementedLogicUserServer can be embedded to have forward compatible implementations.
type UnimplementedLogicUserServer struct {
}

func (*UnimplementedLogicUserServer) AccountByUser(ctx context.Context, req *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountByUser not implemented")
}

func RegisterLogicUserServer(s *grpc.Server, srv LogicUserServer) {
	s.RegisterService(&_LogicUser_serviceDesc, srv)
}

func _LogicUser_AccountByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicUserServer).AccountByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicUser/AccountByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicUserServer).AccountByUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogicUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogicUser",
	HandlerType: (*LogicUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountByUser",
			Handler:    _LogicUser_AccountByUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic.user.proto",
}
