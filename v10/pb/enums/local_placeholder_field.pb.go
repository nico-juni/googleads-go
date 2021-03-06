// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/enums/local_placeholder_field.proto

package enums

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

type LocalPlaceholderFieldEnum_LocalPlaceholderField int32

const (
	LocalPlaceholderFieldEnum_UNSPECIFIED          LocalPlaceholderFieldEnum_LocalPlaceholderField = 0
	LocalPlaceholderFieldEnum_UNKNOWN              LocalPlaceholderFieldEnum_LocalPlaceholderField = 1
	LocalPlaceholderFieldEnum_DEAL_ID              LocalPlaceholderFieldEnum_LocalPlaceholderField = 2
	LocalPlaceholderFieldEnum_DEAL_NAME            LocalPlaceholderFieldEnum_LocalPlaceholderField = 3
	LocalPlaceholderFieldEnum_SUBTITLE             LocalPlaceholderFieldEnum_LocalPlaceholderField = 4
	LocalPlaceholderFieldEnum_DESCRIPTION          LocalPlaceholderFieldEnum_LocalPlaceholderField = 5
	LocalPlaceholderFieldEnum_PRICE                LocalPlaceholderFieldEnum_LocalPlaceholderField = 6
	LocalPlaceholderFieldEnum_FORMATTED_PRICE      LocalPlaceholderFieldEnum_LocalPlaceholderField = 7
	LocalPlaceholderFieldEnum_SALE_PRICE           LocalPlaceholderFieldEnum_LocalPlaceholderField = 8
	LocalPlaceholderFieldEnum_FORMATTED_SALE_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 9
	LocalPlaceholderFieldEnum_IMAGE_URL            LocalPlaceholderFieldEnum_LocalPlaceholderField = 10
	LocalPlaceholderFieldEnum_ADDRESS              LocalPlaceholderFieldEnum_LocalPlaceholderField = 11
	LocalPlaceholderFieldEnum_CATEGORY             LocalPlaceholderFieldEnum_LocalPlaceholderField = 12
	LocalPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS  LocalPlaceholderFieldEnum_LocalPlaceholderField = 13
	LocalPlaceholderFieldEnum_FINAL_URLS           LocalPlaceholderFieldEnum_LocalPlaceholderField = 14
	LocalPlaceholderFieldEnum_FINAL_MOBILE_URLS    LocalPlaceholderFieldEnum_LocalPlaceholderField = 15
	LocalPlaceholderFieldEnum_TRACKING_URL         LocalPlaceholderFieldEnum_LocalPlaceholderField = 16
	LocalPlaceholderFieldEnum_ANDROID_APP_LINK     LocalPlaceholderFieldEnum_LocalPlaceholderField = 17
	LocalPlaceholderFieldEnum_SIMILAR_DEAL_IDS     LocalPlaceholderFieldEnum_LocalPlaceholderField = 18
	LocalPlaceholderFieldEnum_IOS_APP_LINK         LocalPlaceholderFieldEnum_LocalPlaceholderField = 19
	LocalPlaceholderFieldEnum_IOS_APP_STORE_ID     LocalPlaceholderFieldEnum_LocalPlaceholderField = 20
)

// Enum value maps for LocalPlaceholderFieldEnum_LocalPlaceholderField.
var (
	LocalPlaceholderFieldEnum_LocalPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DEAL_ID",
		3:  "DEAL_NAME",
		4:  "SUBTITLE",
		5:  "DESCRIPTION",
		6:  "PRICE",
		7:  "FORMATTED_PRICE",
		8:  "SALE_PRICE",
		9:  "FORMATTED_SALE_PRICE",
		10: "IMAGE_URL",
		11: "ADDRESS",
		12: "CATEGORY",
		13: "CONTEXTUAL_KEYWORDS",
		14: "FINAL_URLS",
		15: "FINAL_MOBILE_URLS",
		16: "TRACKING_URL",
		17: "ANDROID_APP_LINK",
		18: "SIMILAR_DEAL_IDS",
		19: "IOS_APP_LINK",
		20: "IOS_APP_STORE_ID",
	}
	LocalPlaceholderFieldEnum_LocalPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":          0,
		"UNKNOWN":              1,
		"DEAL_ID":              2,
		"DEAL_NAME":            3,
		"SUBTITLE":             4,
		"DESCRIPTION":          5,
		"PRICE":                6,
		"FORMATTED_PRICE":      7,
		"SALE_PRICE":           8,
		"FORMATTED_SALE_PRICE": 9,
		"IMAGE_URL":            10,
		"ADDRESS":              11,
		"CATEGORY":             12,
		"CONTEXTUAL_KEYWORDS":  13,
		"FINAL_URLS":           14,
		"FINAL_MOBILE_URLS":    15,
		"TRACKING_URL":         16,
		"ANDROID_APP_LINK":     17,
		"SIMILAR_DEAL_IDS":     18,
		"IOS_APP_LINK":         19,
		"IOS_APP_STORE_ID":     20,
	}
)

func (x LocalPlaceholderFieldEnum_LocalPlaceholderField) Enum() *LocalPlaceholderFieldEnum_LocalPlaceholderField {
	p := new(LocalPlaceholderFieldEnum_LocalPlaceholderField)
	*p = x
	return p
}

func (x LocalPlaceholderFieldEnum_LocalPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalPlaceholderFieldEnum_LocalPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_local_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (LocalPlaceholderFieldEnum_LocalPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_local_placeholder_field_proto_enumTypes[0]
}

func (x LocalPlaceholderFieldEnum_LocalPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocalPlaceholderFieldEnum_LocalPlaceholderField.Descriptor instead.
func (LocalPlaceholderFieldEnum_LocalPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

type LocalPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LocalPlaceholderFieldEnum) Reset() {
	*x = LocalPlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_local_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalPlaceholderFieldEnum) ProtoMessage() {}

func (x *LocalPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_local_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*LocalPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_local_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x03, 0x0a,
	0x19, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x8a, 0x03, 0x0a, 0x15, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x55, 0x42, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x09, 0x0a,
	0x05, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x07, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x08, 0x12, 0x18, 0x0a,
	0x14, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x5f,
	0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x41, 0x47, 0x45,
	0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10,
	0x0c, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x55, 0x41, 0x4c, 0x5f,
	0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x0d, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49,
	0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49,
	0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10,
	0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52,
	0x4c, 0x10, 0x10, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41,
	0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x49, 0x4d,
	0x49, 0x4c, 0x41, 0x52, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x12, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10,
	0x13, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x54, 0x4f,
	0x52, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x14, 0x42, 0xf4, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1a,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescData = file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_local_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_local_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_local_placeholder_field_proto_goTypes = []interface{}{
	(LocalPlaceholderFieldEnum_LocalPlaceholderField)(0), // 0: google.ads.googleads.v10.enums.LocalPlaceholderFieldEnum.LocalPlaceholderField
	(*LocalPlaceholderFieldEnum)(nil),                    // 1: google.ads.googleads.v10.enums.LocalPlaceholderFieldEnum
}
var file_google_ads_googleads_v10_enums_local_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_local_placeholder_field_proto_init() }
func file_google_ads_googleads_v10_enums_local_placeholder_field_proto_init() {
	if File_google_ads_googleads_v10_enums_local_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_local_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalPlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_local_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_local_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_local_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_local_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_local_placeholder_field_proto = out.File
	file_google_ads_googleads_v10_enums_local_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_local_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_local_placeholder_field_proto_depIdxs = nil
}
