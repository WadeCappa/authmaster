// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: authmaster/authmaster.proto

package authmaster

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestAuthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestAuthRequest) Reset() {
	*x = TestAuthRequest{}
	mi := &file_authmaster_authmaster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAuthRequest) ProtoMessage() {}

func (x *TestAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authmaster_authmaster_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAuthRequest.ProtoReflect.Descriptor instead.
func (*TestAuthRequest) Descriptor() ([]byte, []int) {
	return file_authmaster_authmaster_proto_rawDescGZIP(), []int{0}
}

type TestAuthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TestAuthResponse) Reset() {
	*x = TestAuthResponse{}
	mi := &file_authmaster_authmaster_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAuthResponse) ProtoMessage() {}

func (x *TestAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authmaster_authmaster_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAuthResponse.ProtoReflect.Descriptor instead.
func (*TestAuthResponse) Descriptor() ([]byte, []int) {
	return file_authmaster_authmaster_proto_rawDescGZIP(), []int{1}
}

func (x *TestAuthResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_authmaster_authmaster_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authmaster_authmaster_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_authmaster_authmaster_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_authmaster_authmaster_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authmaster_authmaster_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_authmaster_authmaster_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_authmaster_authmaster_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authmaster_authmaster_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_authmaster_authmaster_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_authmaster_authmaster_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authmaster_authmaster_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_authmaster_authmaster_proto_rawDescGZIP(), []int{5}
}

var File_authmaster_authmaster_proto protoreflect.FileDescriptor

const file_authmaster_authmaster_proto_rawDesc = "" +
	"\n" +
	"\x1bauthmaster/authmaster.proto\x12\n" +
	"authmaster\"\x11\n" +
	"\x0fTestAuthRequest\"*\n" +
	"\x10TestAuthResponse\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\x03R\x06userId\"F\n" +
	"\fLoginRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"%\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"K\n" +
	"\x11CreateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"\x14\n" +
	"\x12CreateUserResponse2\xe4\x01\n" +
	"\n" +
	"Authmaster\x12G\n" +
	"\bTestAuth\x12\x1b.authmaster.TestAuthRequest\x1a\x1c.authmaster.TestAuthResponse\"\x00\x12>\n" +
	"\x05Login\x12\x18.authmaster.LoginRequest\x1a\x19.authmaster.LoginResponse\"\x00\x12M\n" +
	"\n" +
	"CreateUser\x12\x1d.authmaster.CreateUserRequest\x1a\x1e.authmaster.CreateUserResponse\"\x00B,Z*github.com/WadeCappa/authmaster/authmasterb\x06proto3"

var (
	file_authmaster_authmaster_proto_rawDescOnce sync.Once
	file_authmaster_authmaster_proto_rawDescData []byte
)

func file_authmaster_authmaster_proto_rawDescGZIP() []byte {
	file_authmaster_authmaster_proto_rawDescOnce.Do(func() {
		file_authmaster_authmaster_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_authmaster_authmaster_proto_rawDesc), len(file_authmaster_authmaster_proto_rawDesc)))
	})
	return file_authmaster_authmaster_proto_rawDescData
}

var file_authmaster_authmaster_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_authmaster_authmaster_proto_goTypes = []any{
	(*TestAuthRequest)(nil),    // 0: authmaster.TestAuthRequest
	(*TestAuthResponse)(nil),   // 1: authmaster.TestAuthResponse
	(*LoginRequest)(nil),       // 2: authmaster.LoginRequest
	(*LoginResponse)(nil),      // 3: authmaster.LoginResponse
	(*CreateUserRequest)(nil),  // 4: authmaster.CreateUserRequest
	(*CreateUserResponse)(nil), // 5: authmaster.CreateUserResponse
}
var file_authmaster_authmaster_proto_depIdxs = []int32{
	0, // 0: authmaster.Authmaster.TestAuth:input_type -> authmaster.TestAuthRequest
	2, // 1: authmaster.Authmaster.Login:input_type -> authmaster.LoginRequest
	4, // 2: authmaster.Authmaster.CreateUser:input_type -> authmaster.CreateUserRequest
	1, // 3: authmaster.Authmaster.TestAuth:output_type -> authmaster.TestAuthResponse
	3, // 4: authmaster.Authmaster.Login:output_type -> authmaster.LoginResponse
	5, // 5: authmaster.Authmaster.CreateUser:output_type -> authmaster.CreateUserResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authmaster_authmaster_proto_init() }
func file_authmaster_authmaster_proto_init() {
	if File_authmaster_authmaster_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_authmaster_authmaster_proto_rawDesc), len(file_authmaster_authmaster_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authmaster_authmaster_proto_goTypes,
		DependencyIndexes: file_authmaster_authmaster_proto_depIdxs,
		MessageInfos:      file_authmaster_authmaster_proto_msgTypes,
	}.Build()
	File_authmaster_authmaster_proto = out.File
	file_authmaster_authmaster_proto_goTypes = nil
	file_authmaster_authmaster_proto_depIdxs = nil
}
