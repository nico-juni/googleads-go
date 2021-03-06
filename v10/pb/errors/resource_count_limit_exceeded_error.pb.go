// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/errors/resource_count_limit_exceeded_error.proto

package errors

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError int32

const (
	ResourceCountLimitExceededErrorEnum_UNSPECIFIED                 ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 0
	ResourceCountLimitExceededErrorEnum_UNKNOWN                     ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 1
	ResourceCountLimitExceededErrorEnum_ACCOUNT_LIMIT               ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 2
	ResourceCountLimitExceededErrorEnum_CAMPAIGN_LIMIT              ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 3
	ResourceCountLimitExceededErrorEnum_ADGROUP_LIMIT               ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 4
	ResourceCountLimitExceededErrorEnum_AD_GROUP_AD_LIMIT           ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 5
	ResourceCountLimitExceededErrorEnum_AD_GROUP_CRITERION_LIMIT    ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 6
	ResourceCountLimitExceededErrorEnum_SHARED_SET_LIMIT            ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 7
	ResourceCountLimitExceededErrorEnum_MATCHING_FUNCTION_LIMIT     ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 8
	ResourceCountLimitExceededErrorEnum_RESPONSE_ROW_LIMIT_EXCEEDED ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 9
	ResourceCountLimitExceededErrorEnum_RESOURCE_LIMIT              ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError = 10
)

// Enum value maps for ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError.
var (
	ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "ACCOUNT_LIMIT",
		3:  "CAMPAIGN_LIMIT",
		4:  "ADGROUP_LIMIT",
		5:  "AD_GROUP_AD_LIMIT",
		6:  "AD_GROUP_CRITERION_LIMIT",
		7:  "SHARED_SET_LIMIT",
		8:  "MATCHING_FUNCTION_LIMIT",
		9:  "RESPONSE_ROW_LIMIT_EXCEEDED",
		10: "RESOURCE_LIMIT",
	}
	ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError_value = map[string]int32{
		"UNSPECIFIED":                 0,
		"UNKNOWN":                     1,
		"ACCOUNT_LIMIT":               2,
		"CAMPAIGN_LIMIT":              3,
		"ADGROUP_LIMIT":               4,
		"AD_GROUP_AD_LIMIT":           5,
		"AD_GROUP_CRITERION_LIMIT":    6,
		"SHARED_SET_LIMIT":            7,
		"MATCHING_FUNCTION_LIMIT":     8,
		"RESPONSE_ROW_LIMIT_EXCEEDED": 9,
		"RESOURCE_LIMIT":              10,
	}
)

func (x ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError) Enum() *ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError {
	p := new(ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError)
	*p = x
	return p
}

func (x ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_enumTypes[0].Descriptor()
}

func (ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_enumTypes[0]
}

func (x ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError.Descriptor instead.
func (ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescGZIP(), []int{0, 0}
}

type ResourceCountLimitExceededErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResourceCountLimitExceededErrorEnum) Reset() {
	*x = ResourceCountLimitExceededErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceCountLimitExceededErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceCountLimitExceededErrorEnum) ProtoMessage() {}

func (x *ResourceCountLimitExceededErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceCountLimitExceededErrorEnum.ProtoReflect.Descriptor instead.
func (*ResourceCountLimitExceededErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x23, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0x96, 0x02, 0x0a, 0x1f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4d, 0x50, 0x41,
	0x49, 0x47, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41,
	0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x04, 0x12, 0x15,
	0x0a, 0x11, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x41, 0x44, 0x5f, 0x4c, 0x49,
	0x4d, 0x49, 0x54, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x53, 0x45,
	0x54, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x07, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c,
	0x49, 0x4d, 0x49, 0x54, 0x10, 0x08, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x5f, 0x52, 0x4f, 0x57, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43,
	0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x0a, 0x42, 0x84, 0x02, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x42, 0x24, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescData = file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_goTypes = []interface{}{
	(ResourceCountLimitExceededErrorEnum_ResourceCountLimitExceededError)(0), // 0: google.ads.googleads.v10.errors.ResourceCountLimitExceededErrorEnum.ResourceCountLimitExceededError
	(*ResourceCountLimitExceededErrorEnum)(nil),                              // 1: google.ads.googleads.v10.errors.ResourceCountLimitExceededErrorEnum
}
var file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_init() }
func file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_init() {
	if File_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceCountLimitExceededErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto = out.File
	file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_resource_count_limit_exceeded_error_proto_depIdxs = nil
}
