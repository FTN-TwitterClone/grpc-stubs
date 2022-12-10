// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: social_graph_service.proto

package social_graph

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type SocialGraphUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	FirstName   string `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Town        string `protobuf:"bytes,5,opt,name=Town,proto3" json:"Town,omitempty"`
	Gender      string `protobuf:"bytes,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
	YearOfBirth int32  `protobuf:"varint,7,opt,name=YearOfBirth,proto3" json:"YearOfBirth,omitempty"`
}

func (x *SocialGraphUser) Reset() {
	*x = SocialGraphUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGraphUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGraphUser) ProtoMessage() {}

func (x *SocialGraphUser) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGraphUser.ProtoReflect.Descriptor instead.
func (*SocialGraphUser) Descriptor() ([]byte, []int) {
	return file_social_graph_service_proto_rawDescGZIP(), []int{0}
}

func (x *SocialGraphUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SocialGraphUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SocialGraphUser) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SocialGraphUser) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SocialGraphUser) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

func (x *SocialGraphUser) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *SocialGraphUser) GetYearOfBirth() int32 {
	if x != nil {
		return x.YearOfBirth
	}
	return 0
}

type SocialGraphBusinessUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Website     string `protobuf:"bytes,3,opt,name=Website,proto3" json:"Website,omitempty"`
	CompanyName string `protobuf:"bytes,4,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
}

func (x *SocialGraphBusinessUser) Reset() {
	*x = SocialGraphBusinessUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGraphBusinessUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGraphBusinessUser) ProtoMessage() {}

func (x *SocialGraphBusinessUser) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGraphBusinessUser.ProtoReflect.Descriptor instead.
func (*SocialGraphBusinessUser) Descriptor() ([]byte, []int) {
	return file_social_graph_service_proto_rawDescGZIP(), []int{1}
}

func (x *SocialGraphBusinessUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SocialGraphBusinessUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SocialGraphBusinessUser) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *SocialGraphBusinessUser) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

type SocialGraphVisibilityUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visibility bool `protobuf:"varint,1,opt,name=Visibility,proto3" json:"Visibility,omitempty"`
}

func (x *SocialGraphVisibilityUserResponse) Reset() {
	*x = SocialGraphVisibilityUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGraphVisibilityUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGraphVisibilityUserResponse) ProtoMessage() {}

func (x *SocialGraphVisibilityUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGraphVisibilityUserResponse.ProtoReflect.Descriptor instead.
func (*SocialGraphVisibilityUserResponse) Descriptor() ([]byte, []int) {
	return file_social_graph_service_proto_rawDescGZIP(), []int{2}
}

func (x *SocialGraphVisibilityUserResponse) GetVisibility() bool {
	if x != nil {
		return x.Visibility
	}
	return false
}

type SocialGraphFollowers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usernames []*SocialGraphUsername `protobuf:"bytes,1,rep,name=usernames,proto3" json:"usernames,omitempty"`
}

func (x *SocialGraphFollowers) Reset() {
	*x = SocialGraphFollowers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGraphFollowers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGraphFollowers) ProtoMessage() {}

func (x *SocialGraphFollowers) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGraphFollowers.ProtoReflect.Descriptor instead.
func (*SocialGraphFollowers) Descriptor() ([]byte, []int) {
	return file_social_graph_service_proto_rawDescGZIP(), []int{3}
}

func (x *SocialGraphFollowers) GetUsernames() []*SocialGraphUsername {
	if x != nil {
		return x.Usernames
	}
	return nil
}

type SocialGraphTargetUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usernames []*SocialGraphUsername `protobuf:"bytes,1,rep,name=usernames,proto3" json:"usernames,omitempty"`
}

func (x *SocialGraphTargetUsers) Reset() {
	*x = SocialGraphTargetUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGraphTargetUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGraphTargetUsers) ProtoMessage() {}

func (x *SocialGraphTargetUsers) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGraphTargetUsers.ProtoReflect.Descriptor instead.
func (*SocialGraphTargetUsers) Descriptor() ([]byte, []int) {
	return file_social_graph_service_proto_rawDescGZIP(), []int{4}
}

func (x *SocialGraphTargetUsers) GetUsernames() []*SocialGraphUsername {
	if x != nil {
		return x.Usernames
	}
	return nil
}

type SocialGraphTargetUsersGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Town   string `protobuf:"bytes,1,opt,name=Town,proto3" json:"Town,omitempty"`
	Gender string `protobuf:"bytes,2,opt,name=Gender,proto3" json:"Gender,omitempty"`
	MinAge int32  `protobuf:"varint,3,opt,name=MinAge,proto3" json:"MinAge,omitempty"`
	MaxAge int32  `protobuf:"varint,4,opt,name=MaxAge,proto3" json:"MaxAge,omitempty"`
}

func (x *SocialGraphTargetUsersGroup) Reset() {
	*x = SocialGraphTargetUsersGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGraphTargetUsersGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGraphTargetUsersGroup) ProtoMessage() {}

func (x *SocialGraphTargetUsersGroup) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGraphTargetUsersGroup.ProtoReflect.Descriptor instead.
func (*SocialGraphTargetUsersGroup) Descriptor() ([]byte, []int) {
	return file_social_graph_service_proto_rawDescGZIP(), []int{5}
}

func (x *SocialGraphTargetUsersGroup) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

func (x *SocialGraphTargetUsersGroup) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *SocialGraphTargetUsersGroup) GetMinAge() int32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *SocialGraphTargetUsersGroup) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

type SocialGraphUsername struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *SocialGraphUsername) Reset() {
	*x = SocialGraphUsername{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGraphUsername) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGraphUsername) ProtoMessage() {}

func (x *SocialGraphUsername) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGraphUsername.ProtoReflect.Descriptor instead.
func (*SocialGraphUsername) Descriptor() ([]byte, []int) {
	return file_social_graph_service_proto_rawDescGZIP(), []int{6}
}

func (x *SocialGraphUsername) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SocialGraphUpdatedUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Private bool `protobuf:"varint,1,opt,name=Private,proto3" json:"Private,omitempty"`
}

func (x *SocialGraphUpdatedUser) Reset() {
	*x = SocialGraphUpdatedUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_graph_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialGraphUpdatedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialGraphUpdatedUser) ProtoMessage() {}

func (x *SocialGraphUpdatedUser) ProtoReflect() protoreflect.Message {
	mi := &file_social_graph_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialGraphUpdatedUser.ProtoReflect.Descriptor instead.
func (*SocialGraphUpdatedUser) Descriptor() ([]byte, []int) {
	return file_social_graph_service_proto_rawDescGZIP(), []int{7}
}

func (x *SocialGraphUpdatedUser) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

var File_social_graph_service_proto protoreflect.FileDescriptor

var file_social_graph_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0f, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x6f, 0x77, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x6f, 0x77, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x59, 0x65, 0x61, 0x72, 0x4f, 0x66, 0x42, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x59, 0x65, 0x61, 0x72, 0x4f, 0x66,
	0x42, 0x69, 0x72, 0x74, 0x68, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x43, 0x0a, 0x21, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x22, 0x57, 0x0a, 0x14, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x59, 0x0a,
	0x16, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x79, 0x0a, 0x1b, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x6f, 0x77, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x6f, 0x77, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x69, 0x6e, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x4d, 0x69, 0x6e, 0x41, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d,
	0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4d, 0x61, 0x78,
	0x41, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x32, 0xb1, 0x04, 0x0a, 0x12, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1d, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x14, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x25, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x21, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x2f, 0x2e, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x15,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x24, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x00, 0x42, 0x14,
	0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_social_graph_service_proto_rawDescOnce sync.Once
	file_social_graph_service_proto_rawDescData = file_social_graph_service_proto_rawDesc
)

func file_social_graph_service_proto_rawDescGZIP() []byte {
	file_social_graph_service_proto_rawDescOnce.Do(func() {
		file_social_graph_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_social_graph_service_proto_rawDescData)
	})
	return file_social_graph_service_proto_rawDescData
}

var file_social_graph_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_social_graph_service_proto_goTypes = []interface{}{
	(*SocialGraphUser)(nil),                   // 0: social_graph.SocialGraphUser
	(*SocialGraphBusinessUser)(nil),           // 1: social_graph.SocialGraphBusinessUser
	(*SocialGraphVisibilityUserResponse)(nil), // 2: social_graph.SocialGraphVisibilityUserResponse
	(*SocialGraphFollowers)(nil),              // 3: social_graph.SocialGraphFollowers
	(*SocialGraphTargetUsers)(nil),            // 4: social_graph.SocialGraphTargetUsers
	(*SocialGraphTargetUsersGroup)(nil),       // 5: social_graph.SocialGraphTargetUsersGroup
	(*SocialGraphUsername)(nil),               // 6: social_graph.SocialGraphUsername
	(*SocialGraphUpdatedUser)(nil),            // 7: social_graph.SocialGraphUpdatedUser
	(*empty.Empty)(nil),                       // 8: google.protobuf.Empty
}
var file_social_graph_service_proto_depIdxs = []int32{
	6, // 0: social_graph.SocialGraphFollowers.usernames:type_name -> social_graph.SocialGraphUsername
	6, // 1: social_graph.SocialGraphTargetUsers.usernames:type_name -> social_graph.SocialGraphUsername
	0, // 2: social_graph.SocialGraphService.RegisterUser:input_type -> social_graph.SocialGraphUser
	1, // 3: social_graph.SocialGraphService.RegisterBusinessUser:input_type -> social_graph.SocialGraphBusinessUser
	6, // 4: social_graph.SocialGraphService.CheckVisibility:input_type -> social_graph.SocialGraphUsername
	8, // 5: social_graph.SocialGraphService.GetMyFollowers:input_type -> google.protobuf.Empty
	7, // 6: social_graph.SocialGraphService.SocialGraphUpdateUser:input_type -> social_graph.SocialGraphUpdatedUser
	5, // 7: social_graph.SocialGraphService.GetTargetGroupUser:input_type -> social_graph.SocialGraphTargetUsersGroup
	8, // 8: social_graph.SocialGraphService.RegisterUser:output_type -> google.protobuf.Empty
	8, // 9: social_graph.SocialGraphService.RegisterBusinessUser:output_type -> google.protobuf.Empty
	2, // 10: social_graph.SocialGraphService.CheckVisibility:output_type -> social_graph.SocialGraphVisibilityUserResponse
	3, // 11: social_graph.SocialGraphService.GetMyFollowers:output_type -> social_graph.SocialGraphFollowers
	8, // 12: social_graph.SocialGraphService.SocialGraphUpdateUser:output_type -> google.protobuf.Empty
	4, // 13: social_graph.SocialGraphService.GetTargetGroupUser:output_type -> social_graph.SocialGraphTargetUsers
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_social_graph_service_proto_init() }
func file_social_graph_service_proto_init() {
	if File_social_graph_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_social_graph_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGraphUser); i {
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
		file_social_graph_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGraphBusinessUser); i {
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
		file_social_graph_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGraphVisibilityUserResponse); i {
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
		file_social_graph_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGraphFollowers); i {
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
		file_social_graph_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGraphTargetUsers); i {
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
		file_social_graph_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGraphTargetUsersGroup); i {
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
		file_social_graph_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGraphUsername); i {
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
		file_social_graph_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialGraphUpdatedUser); i {
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
			RawDescriptor: file_social_graph_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_social_graph_service_proto_goTypes,
		DependencyIndexes: file_social_graph_service_proto_depIdxs,
		MessageInfos:      file_social_graph_service_proto_msgTypes,
	}.Build()
	File_social_graph_service_proto = out.File
	file_social_graph_service_proto_rawDesc = nil
	file_social_graph_service_proto_goTypes = nil
	file_social_graph_service_proto_depIdxs = nil
}
