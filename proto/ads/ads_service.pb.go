// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: ads_service.proto

package ads

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

type AdInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetId  string `protobuf:"bytes,1,opt,name=TweetId,proto3" json:"TweetId,omitempty"`
	PostedBy string `protobuf:"bytes,2,opt,name=PostedBy,proto3" json:"PostedBy,omitempty"`
	Town     string `protobuf:"bytes,3,opt,name=Town,proto3" json:"Town,omitempty"`
	MinAge   int32  `protobuf:"varint,4,opt,name=MinAge,proto3" json:"MinAge,omitempty"`
	MaxAge   int32  `protobuf:"varint,5,opt,name=MaxAge,proto3" json:"MaxAge,omitempty"`
	Gender   string `protobuf:"bytes,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
}

func (x *AdInfo) Reset() {
	*x = AdInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ads_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdInfo) ProtoMessage() {}

func (x *AdInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ads_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdInfo.ProtoReflect.Descriptor instead.
func (*AdInfo) Descriptor() ([]byte, []int) {
	return file_ads_service_proto_rawDescGZIP(), []int{0}
}

func (x *AdInfo) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

func (x *AdInfo) GetPostedBy() string {
	if x != nil {
		return x.PostedBy
	}
	return ""
}

func (x *AdInfo) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

func (x *AdInfo) GetMinAge() int32 {
	if x != nil {
		return x.MinAge
	}
	return 0
}

func (x *AdInfo) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *AdInfo) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type LikeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetId  string `protobuf:"bytes,1,opt,name=TweetId,proto3" json:"TweetId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *LikeEvent) Reset() {
	*x = LikeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ads_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeEvent) ProtoMessage() {}

func (x *LikeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_ads_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeEvent.ProtoReflect.Descriptor instead.
func (*LikeEvent) Descriptor() ([]byte, []int) {
	return file_ads_service_proto_rawDescGZIP(), []int{1}
}

func (x *LikeEvent) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

func (x *LikeEvent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UnlikeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetId  string `protobuf:"bytes,1,opt,name=TweetId,proto3" json:"TweetId,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *UnlikeEvent) Reset() {
	*x = UnlikeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ads_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlikeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlikeEvent) ProtoMessage() {}

func (x *UnlikeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_ads_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlikeEvent.ProtoReflect.Descriptor instead.
func (*UnlikeEvent) Descriptor() ([]byte, []int) {
	return file_ads_service_proto_rawDescGZIP(), []int{2}
}

func (x *UnlikeEvent) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

func (x *UnlikeEvent) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_ads_service_proto protoreflect.FileDescriptor

var file_ads_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x64, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x64, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x06, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x6f, 0x77, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x6f, 0x77, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x69,
	0x6e, 0x41, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4d, 0x69, 0x6e, 0x41,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x22, 0x41, 0x0a, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x0b, 0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xbb, 0x01, 0x0a, 0x0a, 0x41,
	0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x53, 0x61, 0x76,
	0x65, 0x41, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x41, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0f, 0x53, 0x61, 0x76,
	0x65, 0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x55, 0x6e, 0x6c, 0x69, 0x6b, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ads_service_proto_rawDescOnce sync.Once
	file_ads_service_proto_rawDescData = file_ads_service_proto_rawDesc
)

func file_ads_service_proto_rawDescGZIP() []byte {
	file_ads_service_proto_rawDescOnce.Do(func() {
		file_ads_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ads_service_proto_rawDescData)
	})
	return file_ads_service_proto_rawDescData
}

var file_ads_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ads_service_proto_goTypes = []interface{}{
	(*AdInfo)(nil),      // 0: ads.AdInfo
	(*LikeEvent)(nil),   // 1: ads.LikeEvent
	(*UnlikeEvent)(nil), // 2: ads.UnlikeEvent
	(*empty.Empty)(nil), // 3: google.protobuf.Empty
}
var file_ads_service_proto_depIdxs = []int32{
	0, // 0: ads.AdsService.SaveAdInfo:input_type -> ads.AdInfo
	1, // 1: ads.AdsService.SaveLikeEvent:input_type -> ads.LikeEvent
	2, // 2: ads.AdsService.SaveUnlikeEvent:input_type -> ads.UnlikeEvent
	3, // 3: ads.AdsService.SaveAdInfo:output_type -> google.protobuf.Empty
	3, // 4: ads.AdsService.SaveLikeEvent:output_type -> google.protobuf.Empty
	3, // 5: ads.AdsService.SaveUnlikeEvent:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ads_service_proto_init() }
func file_ads_service_proto_init() {
	if File_ads_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ads_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdInfo); i {
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
		file_ads_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeEvent); i {
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
		file_ads_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlikeEvent); i {
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
			RawDescriptor: file_ads_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ads_service_proto_goTypes,
		DependencyIndexes: file_ads_service_proto_depIdxs,
		MessageInfos:      file_ads_service_proto_msgTypes,
	}.Build()
	File_ads_service_proto = out.File
	file_ads_service_proto_rawDesc = nil
	file_ads_service_proto_goTypes = nil
	file_ads_service_proto_depIdxs = nil
}