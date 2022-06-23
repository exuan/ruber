// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.1
// source: v1/recipe_backend.proto

package v1

import (
	v1 "github.com/exuan/ruber/api/recipe/service/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_v1_recipe_backend_proto_rawDescGZIP(), []int{0}
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

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_v1_recipe_backend_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_backend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_backend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_v1_recipe_backend_proto_rawDescGZIP(), []int{2}
}

type LogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutReply) Reset() {
	*x = LogoutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_backend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReply) ProtoMessage() {}

func (x *LogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_backend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReply.ProtoReflect.Descriptor instead.
func (*LogoutReply) Descriptor() ([]byte, []int) {
	return file_v1_recipe_backend_proto_rawDescGZIP(), []int{3}
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_backend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_backend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_v1_recipe_backend_proto_rawDescGZIP(), []int{4}
}

type UserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// @inject_tag: json:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username"`
	// @inject_tag: json:"phone"
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone"`
	// @inject_tag: json:"nickname"
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname"`
	// @inject_tag: json:"avatar"
	Avatar string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar"`
	// @inject_tag: json:"create_time"
	CreateTime int64 `protobuf:"varint,6,opt,name=create_time,json=createTime,proto3" json:"create_time"`
	// @inject_tag: json:"update_time"
	UpdateTime int64 `protobuf:"varint,7,opt,name=update_time,json=updateTime,proto3" json:"update_time"`
}

func (x *UserReply) Reset() {
	*x = UserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_backend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReply) ProtoMessage() {}

func (x *UserReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_backend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReply.ProtoReflect.Descriptor instead.
func (*UserReply) Descriptor() ([]byte, []int) {
	return file_v1_recipe_backend_proto_rawDescGZIP(), []int{5}
}

func (x *UserReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserReply) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserReply) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserReply) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserReply) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserReply) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *UserReply) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type UploadUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"ext" form:"ext"
	Ext string `protobuf:"bytes,1,opt,name=ext,proto3" json:"ext" form:"ext"`
}

func (x *UploadUrlRequest) Reset() {
	*x = UploadUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_backend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadUrlRequest) ProtoMessage() {}

func (x *UploadUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_backend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadUrlRequest.ProtoReflect.Descriptor instead.
func (*UploadUrlRequest) Descriptor() ([]byte, []int) {
	return file_v1_recipe_backend_proto_rawDescGZIP(), []int{6}
}

func (x *UploadUrlRequest) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type UploadUrlReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"url" form:"url"
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url" form:"url"`
	// @inject_tag: json:"presigned_url" form:"presigned_url"
	PresignedUrl string `protobuf:"bytes,2,opt,name=presignedUrl,proto3" json:"presigned_url" form:"presigned_url"`
}

func (x *UploadUrlReply) Reset() {
	*x = UploadUrlReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_recipe_backend_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadUrlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadUrlReply) ProtoMessage() {}

func (x *UploadUrlReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_recipe_backend_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadUrlReply.ProtoReflect.Descriptor instead.
func (*UploadUrlReply) Descriptor() ([]byte, []int) {
	return file_v1_recipe_backend_proto_rawDescGZIP(), []int{7}
}

func (x *UploadUrlReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UploadUrlReply) GetPresignedUrl() string {
	if x != nil {
		return x.PresignedUrl
	}
	return ""
}

var File_v1_recipe_backend_proto protoreflect.FileDescriptor

var file_v1_recipe_backend_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x22, 0x0a, 0x0a, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0f,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0d, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc3, 0x01,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x22, 0x46, 0x0a, 0x0e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x22, 0x0a,
	0x0c, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72,
	0x6c, 0x32, 0xba, 0x0c, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x5d, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x61, 0x0a, 0x06,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3a, 0x01, 0x2a, 0x12,
	0x56, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x67, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x7b, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a,
	0x07, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x73, 0x12, 0x5e, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x20, 0x2e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x12, 0x72, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12,
	0x24, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x73, 0x61,
	0x76, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x7b, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x73,
	0x0a, 0x0e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x22, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x87, 0x01, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x8f, 0x01,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x6d, 0x0a, 0x09, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x2e, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x3a, 0x01, 0x2a, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x75,
	0x61, 0x6e, 0x2f, 0x72, 0x75, 0x62, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_recipe_backend_proto_rawDescOnce sync.Once
	file_v1_recipe_backend_proto_rawDescData = file_v1_recipe_backend_proto_rawDesc
)

func file_v1_recipe_backend_proto_rawDescGZIP() []byte {
	file_v1_recipe_backend_proto_rawDescOnce.Do(func() {
		file_v1_recipe_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_recipe_backend_proto_rawDescData)
	})
	return file_v1_recipe_backend_proto_rawDescData
}

var file_v1_recipe_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_recipe_backend_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),             // 0: recipe.backend.v1.LoginRequest
	(*LoginReply)(nil),               // 1: recipe.backend.v1.LoginReply
	(*LogoutRequest)(nil),            // 2: recipe.backend.v1.LogoutRequest
	(*LogoutReply)(nil),              // 3: recipe.backend.v1.LogoutReply
	(*UserRequest)(nil),              // 4: recipe.backend.v1.UserRequest
	(*UserReply)(nil),                // 5: recipe.backend.v1.UserReply
	(*UploadUrlRequest)(nil),         // 6: recipe.backend.v1.UploadUrlRequest
	(*UploadUrlReply)(nil),           // 7: recipe.backend.v1.UploadUrlReply
	(*v1.IndexRequest)(nil),          // 8: recipe.service.v1.IndexRequest
	(*v1.SaveIndexRequest)(nil),      // 9: recipe.service.v1.SaveIndexRequest
	(*v1.RecipesRequest)(nil),        // 10: recipe.service.v1.RecipesRequest
	(*v1.RecipeRequest)(nil),         // 11: recipe.service.v1.RecipeRequest
	(*v1.SaveRecipeRequest)(nil),     // 12: recipe.service.v1.SaveRecipeRequest
	(*v1.DeleteRecipeRequest)(nil),   // 13: recipe.service.v1.DeleteRecipeRequest
	(*v1.CategoriesRequest)(nil),     // 14: recipe.service.v1.CategoriesRequest
	(*v1.CategoryRequest)(nil),       // 15: recipe.service.v1.CategoryRequest
	(*v1.SaveCategoryRequest)(nil),   // 16: recipe.service.v1.SaveCategoryRequest
	(*v1.DeleteCategoryRequest)(nil), // 17: recipe.service.v1.DeleteCategoryRequest
	(*v1.IndexReply)(nil),            // 18: recipe.service.v1.IndexReply
	(*v1.SaveIndexReply)(nil),        // 19: recipe.service.v1.SaveIndexReply
	(*v1.RecipesReply)(nil),          // 20: recipe.service.v1.RecipesReply
	(*v1.RecipeReply)(nil),           // 21: recipe.service.v1.RecipeReply
	(*v1.SaveRecipeReply)(nil),       // 22: recipe.service.v1.SaveRecipeReply
	(*v1.DeleteRecipeReply)(nil),     // 23: recipe.service.v1.DeleteRecipeReply
	(*v1.CategoriesReply)(nil),       // 24: recipe.service.v1.CategoriesReply
	(*v1.CategoryReply)(nil),         // 25: recipe.service.v1.CategoryReply
	(*v1.SaveCategoryReply)(nil),     // 26: recipe.service.v1.SaveCategoryReply
	(*v1.DeleteCategoryReply)(nil),   // 27: recipe.service.v1.DeleteCategoryReply
}
var file_v1_recipe_backend_proto_depIdxs = []int32{
	0,  // 0: recipe.backend.v1.Backend.Login:input_type -> recipe.backend.v1.LoginRequest
	2,  // 1: recipe.backend.v1.Backend.Logout:input_type -> recipe.backend.v1.LogoutRequest
	4,  // 2: recipe.backend.v1.Backend.User:input_type -> recipe.backend.v1.UserRequest
	8,  // 3: recipe.backend.v1.Backend.RecipeIndex:input_type -> recipe.service.v1.IndexRequest
	9,  // 4: recipe.backend.v1.Backend.SaveRecipeIndex:input_type -> recipe.service.v1.SaveIndexRequest
	10, // 5: recipe.backend.v1.Backend.Recipes:input_type -> recipe.service.v1.RecipesRequest
	11, // 6: recipe.backend.v1.Backend.Recipe:input_type -> recipe.service.v1.RecipeRequest
	12, // 7: recipe.backend.v1.Backend.SaveRecipe:input_type -> recipe.service.v1.SaveRecipeRequest
	13, // 8: recipe.backend.v1.Backend.DeleteRecipe:input_type -> recipe.service.v1.DeleteRecipeRequest
	14, // 9: recipe.backend.v1.Backend.RecipeCategories:input_type -> recipe.service.v1.CategoriesRequest
	15, // 10: recipe.backend.v1.Backend.RecipeCategory:input_type -> recipe.service.v1.CategoryRequest
	16, // 11: recipe.backend.v1.Backend.SaveRecipeCategory:input_type -> recipe.service.v1.SaveCategoryRequest
	17, // 12: recipe.backend.v1.Backend.DeleteRecipeCategory:input_type -> recipe.service.v1.DeleteCategoryRequest
	6,  // 13: recipe.backend.v1.Backend.UploadUrl:input_type -> recipe.backend.v1.UploadUrlRequest
	1,  // 14: recipe.backend.v1.Backend.Login:output_type -> recipe.backend.v1.LoginReply
	3,  // 15: recipe.backend.v1.Backend.Logout:output_type -> recipe.backend.v1.LogoutReply
	5,  // 16: recipe.backend.v1.Backend.User:output_type -> recipe.backend.v1.UserReply
	18, // 17: recipe.backend.v1.Backend.RecipeIndex:output_type -> recipe.service.v1.IndexReply
	19, // 18: recipe.backend.v1.Backend.SaveRecipeIndex:output_type -> recipe.service.v1.SaveIndexReply
	20, // 19: recipe.backend.v1.Backend.Recipes:output_type -> recipe.service.v1.RecipesReply
	21, // 20: recipe.backend.v1.Backend.Recipe:output_type -> recipe.service.v1.RecipeReply
	22, // 21: recipe.backend.v1.Backend.SaveRecipe:output_type -> recipe.service.v1.SaveRecipeReply
	23, // 22: recipe.backend.v1.Backend.DeleteRecipe:output_type -> recipe.service.v1.DeleteRecipeReply
	24, // 23: recipe.backend.v1.Backend.RecipeCategories:output_type -> recipe.service.v1.CategoriesReply
	25, // 24: recipe.backend.v1.Backend.RecipeCategory:output_type -> recipe.service.v1.CategoryReply
	26, // 25: recipe.backend.v1.Backend.SaveRecipeCategory:output_type -> recipe.service.v1.SaveCategoryReply
	27, // 26: recipe.backend.v1.Backend.DeleteRecipeCategory:output_type -> recipe.service.v1.DeleteCategoryReply
	7,  // 27: recipe.backend.v1.Backend.UploadUrl:output_type -> recipe.backend.v1.UploadUrlReply
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_recipe_backend_proto_init() }
func file_v1_recipe_backend_proto_init() {
	if File_v1_recipe_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_recipe_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_v1_recipe_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_v1_recipe_backend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_v1_recipe_backend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReply); i {
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
		file_v1_recipe_backend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_v1_recipe_backend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReply); i {
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
		file_v1_recipe_backend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadUrlRequest); i {
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
		file_v1_recipe_backend_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadUrlReply); i {
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
			RawDescriptor: file_v1_recipe_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_recipe_backend_proto_goTypes,
		DependencyIndexes: file_v1_recipe_backend_proto_depIdxs,
		MessageInfos:      file_v1_recipe_backend_proto_msgTypes,
	}.Build()
	File_v1_recipe_backend_proto = out.File
	file_v1_recipe_backend_proto_rawDesc = nil
	file_v1_recipe_backend_proto_goTypes = nil
	file_v1_recipe_backend_proto_depIdxs = nil
}