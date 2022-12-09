// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: pb/usercenter.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Sex      int64  `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Avatar   string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Info     string `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetSex() int64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type UserAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	AuthType string `protobuf:"bytes,3,opt,name=authType,proto3" json:"authType,omitempty"`
	AuthKey  string `protobuf:"bytes,4,opt,name=authKey,proto3" json:"authKey,omitempty"`
}

func (x *UserAuth) Reset() {
	*x = UserAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuth) ProtoMessage() {}

func (x *UserAuth) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuth.ProtoReflect.Descriptor instead.
func (*UserAuth) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{1}
}

func (x *UserAuth) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserAuth) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserAuth) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *UserAuth) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile   string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	AuthKey  string `protobuf:"bytes,4,opt,name=authKey,proto3" json:"authKey,omitempty"`
	AuthType string `protobuf:"bytes,5,opt,name=authType,proto3" json:"authType,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *RegisterReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,2,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`
	RefreshAfter int64  `protobuf:"varint,3,opt,name=refreshAfter,proto3" json:"refreshAfter,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RegisterResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *RegisterResp) GetRefreshAfter() int64 {
	if x != nil {
		return x.RefreshAfter
	}
	return 0
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType string `protobuf:"bytes,1,opt,name=authType,proto3" json:"authType,omitempty"`
	AuthKey  string `protobuf:"bytes,2,opt,name=authKey,proto3" json:"authKey,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *LoginReq) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,2,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`
	RefreshAfter int64  `protobuf:"varint,3,opt,name=refreshAfter,proto3" json:"refreshAfter,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *LoginResp) GetRefreshAfter() int64 {
	if x != nil {
		return x.RefreshAfter
	}
	return 0
}

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserInfoReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserInfoResp) Reset() {
	*x = GetUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp) ProtoMessage() {}

func (x *GetUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserInfoResp) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserAuthByAuthKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthKey  string `protobuf:"bytes,1,opt,name=authKey,proto3" json:"authKey,omitempty"`
	AuthType string `protobuf:"bytes,2,opt,name=authType,proto3" json:"authType,omitempty"`
}

func (x *GetUserAuthByAuthKeyReq) Reset() {
	*x = GetUserAuthByAuthKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAuthByAuthKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAuthByAuthKeyReq) ProtoMessage() {}

func (x *GetUserAuthByAuthKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAuthByAuthKeyReq.ProtoReflect.Descriptor instead.
func (*GetUserAuthByAuthKeyReq) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserAuthByAuthKeyReq) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *GetUserAuthByAuthKeyReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

type GetUserAuthByAuthKeyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAuth *UserAuth `protobuf:"bytes,1,opt,name=userAuth,proto3" json:"userAuth,omitempty"`
}

func (x *GetUserAuthByAuthKeyResp) Reset() {
	*x = GetUserAuthByAuthKeyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAuthByAuthKeyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAuthByAuthKeyResp) ProtoMessage() {}

func (x *GetUserAuthByAuthKeyResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAuthByAuthKeyResp.ProtoReflect.Descriptor instead.
func (*GetUserAuthByAuthKeyResp) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserAuthByAuthKeyResp) GetUserAuth() *UserAuth {
	if x != nil {
		return x.UserAuth
	}
	return nil
}

type GetUserAuthByUserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserAuthByUserIdReq) Reset() {
	*x = GetUserAuthByUserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAuthByUserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAuthByUserIdReq) ProtoMessage() {}

func (x *GetUserAuthByUserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAuthByUserIdReq.ProtoReflect.Descriptor instead.
func (*GetUserAuthByUserIdReq) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserAuthByUserIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserAuthByUserIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAuth *UserAuth `protobuf:"bytes,2,opt,name=userAuth,proto3" json:"userAuth,omitempty"`
}

func (x *GetUserAuthByUserIdResp) Reset() {
	*x = GetUserAuthByUserIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAuthByUserIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAuthByUserIdResp) ProtoMessage() {}

func (x *GetUserAuthByUserIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAuthByUserIdResp.ProtoReflect.Descriptor instead.
func (*GetUserAuthByUserIdResp) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserAuthByUserIdResp) GetUserAuth() *UserAuth {
	if x != nil {
		return x.UserAuth
	}
	return nil
}

type GenerateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GenerateTokenReq) Reset() {
	*x = GenerateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenReq) ProtoMessage() {}

func (x *GenerateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenReq.ProtoReflect.Descriptor instead.
func (*GenerateTokenReq) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{12}
}

func (x *GenerateTokenReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GenerateTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,2,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`
	RefreshAfter int64  `protobuf:"varint,3,opt,name=refreshAfter,proto3" json:"refreshAfter,omitempty"`
}

func (x *GenerateTokenResp) Reset() {
	*x = GenerateTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_usercenter_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResp) ProtoMessage() {}

func (x *GenerateTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_usercenter_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResp.ProtoReflect.Descriptor instead.
func (*GenerateTokenResp) Descriptor() ([]byte, []int) {
	return file_pb_usercenter_proto_rawDescGZIP(), []int{13}
}

func (x *GenerateTokenResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GenerateTokenResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *GenerateTokenResp) GetRefreshAfter() int64 {
	if x != nil {
		return x.RefreshAfter
	}
	return 0
}

var File_pb_usercenter_proto protoreflect.FileDescriptor

var file_pb_usercenter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x88, 0x01, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x68, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x93,
	0x01, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x78, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0x5c,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x75, 0x0a, 0x09,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x22, 0x30, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x43, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x7d, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65, 0x72, 0x32,
	0xfa, 0x02, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x14, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4e,
	0x0a, 0x13, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c,
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_usercenter_proto_rawDescOnce sync.Once
	file_pb_usercenter_proto_rawDescData = file_pb_usercenter_proto_rawDesc
)

func file_pb_usercenter_proto_rawDescGZIP() []byte {
	file_pb_usercenter_proto_rawDescOnce.Do(func() {
		file_pb_usercenter_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_usercenter_proto_rawDescData)
	})
	return file_pb_usercenter_proto_rawDescData
}

var file_pb_usercenter_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pb_usercenter_proto_goTypes = []interface{}{
	(*User)(nil),                     // 0: pb.User
	(*UserAuth)(nil),                 // 1: pb.UserAuth
	(*RegisterReq)(nil),              // 2: pb.RegisterReq
	(*RegisterResp)(nil),             // 3: pb.RegisterResp
	(*LoginReq)(nil),                 // 4: pb.LoginReq
	(*LoginResp)(nil),                // 5: pb.LoginResp
	(*GetUserInfoReq)(nil),           // 6: pb.GetUserInfoReq
	(*GetUserInfoResp)(nil),          // 7: pb.GetUserInfoResp
	(*GetUserAuthByAuthKeyReq)(nil),  // 8: pb.GetUserAuthByAuthKeyReq
	(*GetUserAuthByAuthKeyResp)(nil), // 9: pb.GetUserAuthByAuthKeyResp
	(*GetUserAuthByUserIdReq)(nil),   // 10: pb.GetUserAuthByUserIdReq
	(*GetUserAuthByUserIdResp)(nil),  // 11: pb.GetUserAuthByUserIdResp
	(*GenerateTokenReq)(nil),         // 12: pb.GenerateTokenReq
	(*GenerateTokenResp)(nil),        // 13: pb.GenerateTokenResp
}
var file_pb_usercenter_proto_depIdxs = []int32{
	0,  // 0: pb.GetUserInfoResp.user:type_name -> pb.User
	1,  // 1: pb.GetUserAuthByAuthKeyResp.userAuth:type_name -> pb.UserAuth
	1,  // 2: pb.GetUserAuthByUserIdResp.userAuth:type_name -> pb.UserAuth
	4,  // 3: pb.usercenter.login:input_type -> pb.LoginReq
	2,  // 4: pb.usercenter.register:input_type -> pb.RegisterReq
	6,  // 5: pb.usercenter.getUserInfo:input_type -> pb.GetUserInfoReq
	8,  // 6: pb.usercenter.getUserAuthByAuthKey:input_type -> pb.GetUserAuthByAuthKeyReq
	10, // 7: pb.usercenter.getUserAuthByUserId:input_type -> pb.GetUserAuthByUserIdReq
	12, // 8: pb.usercenter.generateToken:input_type -> pb.GenerateTokenReq
	5,  // 9: pb.usercenter.login:output_type -> pb.LoginResp
	3,  // 10: pb.usercenter.register:output_type -> pb.RegisterResp
	7,  // 11: pb.usercenter.getUserInfo:output_type -> pb.GetUserInfoResp
	9,  // 12: pb.usercenter.getUserAuthByAuthKey:output_type -> pb.GetUserAuthByAuthKeyResp
	11, // 13: pb.usercenter.getUserAuthByUserId:output_type -> pb.GetUserAuthByUserIdResp
	13, // 14: pb.usercenter.generateToken:output_type -> pb.GenerateTokenResp
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pb_usercenter_proto_init() }
func file_pb_usercenter_proto_init() {
	if File_pb_usercenter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_usercenter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuth); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAuthByAuthKeyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAuthByAuthKeyResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAuthByUserIdReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAuthByUserIdResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_usercenter_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_usercenter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_usercenter_proto_goTypes,
		DependencyIndexes: file_pb_usercenter_proto_depIdxs,
		MessageInfos:      file_pb_usercenter_proto_msgTypes,
	}.Build()
	File_pb_usercenter_proto = out.File
	file_pb_usercenter_proto_rawDesc = nil
	file_pb_usercenter_proto_goTypes = nil
	file_pb_usercenter_proto_depIdxs = nil
}
