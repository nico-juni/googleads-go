// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/errors/recommendation_error.proto

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

type RecommendationErrorEnum_RecommendationError int32

const (
	RecommendationErrorEnum_UNSPECIFIED                      RecommendationErrorEnum_RecommendationError = 0
	RecommendationErrorEnum_UNKNOWN                          RecommendationErrorEnum_RecommendationError = 1
	RecommendationErrorEnum_BUDGET_AMOUNT_TOO_SMALL          RecommendationErrorEnum_RecommendationError = 2
	RecommendationErrorEnum_BUDGET_AMOUNT_TOO_LARGE          RecommendationErrorEnum_RecommendationError = 3
	RecommendationErrorEnum_INVALID_BUDGET_AMOUNT            RecommendationErrorEnum_RecommendationError = 4
	RecommendationErrorEnum_POLICY_ERROR                     RecommendationErrorEnum_RecommendationError = 5
	RecommendationErrorEnum_INVALID_BID_AMOUNT               RecommendationErrorEnum_RecommendationError = 6
	RecommendationErrorEnum_ADGROUP_KEYWORD_LIMIT            RecommendationErrorEnum_RecommendationError = 7
	RecommendationErrorEnum_RECOMMENDATION_ALREADY_APPLIED   RecommendationErrorEnum_RecommendationError = 8
	RecommendationErrorEnum_RECOMMENDATION_INVALIDATED       RecommendationErrorEnum_RecommendationError = 9
	RecommendationErrorEnum_TOO_MANY_OPERATIONS              RecommendationErrorEnum_RecommendationError = 10
	RecommendationErrorEnum_NO_OPERATIONS                    RecommendationErrorEnum_RecommendationError = 11
	RecommendationErrorEnum_DIFFERENT_TYPES_NOT_SUPPORTED    RecommendationErrorEnum_RecommendationError = 12
	RecommendationErrorEnum_DUPLICATE_RESOURCE_NAME          RecommendationErrorEnum_RecommendationError = 13
	RecommendationErrorEnum_RECOMMENDATION_ALREADY_DISMISSED RecommendationErrorEnum_RecommendationError = 14
	RecommendationErrorEnum_INVALID_APPLY_REQUEST            RecommendationErrorEnum_RecommendationError = 15
)

// Enum value maps for RecommendationErrorEnum_RecommendationError.
var (
	RecommendationErrorEnum_RecommendationError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "BUDGET_AMOUNT_TOO_SMALL",
		3:  "BUDGET_AMOUNT_TOO_LARGE",
		4:  "INVALID_BUDGET_AMOUNT",
		5:  "POLICY_ERROR",
		6:  "INVALID_BID_AMOUNT",
		7:  "ADGROUP_KEYWORD_LIMIT",
		8:  "RECOMMENDATION_ALREADY_APPLIED",
		9:  "RECOMMENDATION_INVALIDATED",
		10: "TOO_MANY_OPERATIONS",
		11: "NO_OPERATIONS",
		12: "DIFFERENT_TYPES_NOT_SUPPORTED",
		13: "DUPLICATE_RESOURCE_NAME",
		14: "RECOMMENDATION_ALREADY_DISMISSED",
		15: "INVALID_APPLY_REQUEST",
	}
	RecommendationErrorEnum_RecommendationError_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"UNKNOWN":                          1,
		"BUDGET_AMOUNT_TOO_SMALL":          2,
		"BUDGET_AMOUNT_TOO_LARGE":          3,
		"INVALID_BUDGET_AMOUNT":            4,
		"POLICY_ERROR":                     5,
		"INVALID_BID_AMOUNT":               6,
		"ADGROUP_KEYWORD_LIMIT":            7,
		"RECOMMENDATION_ALREADY_APPLIED":   8,
		"RECOMMENDATION_INVALIDATED":       9,
		"TOO_MANY_OPERATIONS":              10,
		"NO_OPERATIONS":                    11,
		"DIFFERENT_TYPES_NOT_SUPPORTED":    12,
		"DUPLICATE_RESOURCE_NAME":          13,
		"RECOMMENDATION_ALREADY_DISMISSED": 14,
		"INVALID_APPLY_REQUEST":            15,
	}
)

func (x RecommendationErrorEnum_RecommendationError) Enum() *RecommendationErrorEnum_RecommendationError {
	p := new(RecommendationErrorEnum_RecommendationError)
	*p = x
	return p
}

func (x RecommendationErrorEnum_RecommendationError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecommendationErrorEnum_RecommendationError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_recommendation_error_proto_enumTypes[0].Descriptor()
}

func (RecommendationErrorEnum_RecommendationError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_recommendation_error_proto_enumTypes[0]
}

func (x RecommendationErrorEnum_RecommendationError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecommendationErrorEnum_RecommendationError.Descriptor instead.
func (RecommendationErrorEnum_RecommendationError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescGZIP(), []int{0, 0}
}

type RecommendationErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecommendationErrorEnum) Reset() {
	*x = RecommendationErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_recommendation_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendationErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendationErrorEnum) ProtoMessage() {}

func (x *RecommendationErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_recommendation_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendationErrorEnum.ProtoReflect.Descriptor instead.
func (*RecommendationErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_recommendation_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x03, 0x0a, 0x17,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xbe, 0x03, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1b, 0x0a,
	0x17, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x42, 0x55,
	0x44, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x54, 0x4f, 0x4f, 0x5f,
	0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54,
	0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x42, 0x49, 0x44, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15,
	0x41, 0x44, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x5f,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x07, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x45, 0x43, 0x4f, 0x4d,
	0x4d, 0x45, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x52,
	0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x54,
	0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x53, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x0b, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x49, 0x46, 0x46, 0x45,
	0x52, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x55,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0d, 0x12, 0x24, 0x0a, 0x20, 0x52, 0x45, 0x43, 0x4f, 0x4d,
	0x4d, 0x45, 0x4e, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x44, 0x49, 0x53, 0x4d, 0x49, 0x53, 0x53, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x19, 0x0a,
	0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x0f, 0x42, 0xf8, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x42, 0x18, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescData = file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_recommendation_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_recommendation_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_recommendation_error_proto_goTypes = []interface{}{
	(RecommendationErrorEnum_RecommendationError)(0), // 0: google.ads.googleads.v10.errors.RecommendationErrorEnum.RecommendationError
	(*RecommendationErrorEnum)(nil),                  // 1: google.ads.googleads.v10.errors.RecommendationErrorEnum
}
var file_google_ads_googleads_v10_errors_recommendation_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_recommendation_error_proto_init() }
func file_google_ads_googleads_v10_errors_recommendation_error_proto_init() {
	if File_google_ads_googleads_v10_errors_recommendation_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_recommendation_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendationErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_recommendation_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_recommendation_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_recommendation_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_recommendation_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_recommendation_error_proto = out.File
	file_google_ads_googleads_v10_errors_recommendation_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_recommendation_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_recommendation_error_proto_depIdxs = nil
}