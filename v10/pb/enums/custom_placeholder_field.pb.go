// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/enums/custom_placeholder_field.proto

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

type CustomPlaceholderFieldEnum_CustomPlaceholderField int32

const (
	CustomPlaceholderFieldEnum_UNSPECIFIED          CustomPlaceholderFieldEnum_CustomPlaceholderField = 0
	CustomPlaceholderFieldEnum_UNKNOWN              CustomPlaceholderFieldEnum_CustomPlaceholderField = 1
	CustomPlaceholderFieldEnum_ID                   CustomPlaceholderFieldEnum_CustomPlaceholderField = 2
	CustomPlaceholderFieldEnum_ID2                  CustomPlaceholderFieldEnum_CustomPlaceholderField = 3
	CustomPlaceholderFieldEnum_ITEM_TITLE           CustomPlaceholderFieldEnum_CustomPlaceholderField = 4
	CustomPlaceholderFieldEnum_ITEM_SUBTITLE        CustomPlaceholderFieldEnum_CustomPlaceholderField = 5
	CustomPlaceholderFieldEnum_ITEM_DESCRIPTION     CustomPlaceholderFieldEnum_CustomPlaceholderField = 6
	CustomPlaceholderFieldEnum_ITEM_ADDRESS         CustomPlaceholderFieldEnum_CustomPlaceholderField = 7
	CustomPlaceholderFieldEnum_PRICE                CustomPlaceholderFieldEnum_CustomPlaceholderField = 8
	CustomPlaceholderFieldEnum_FORMATTED_PRICE      CustomPlaceholderFieldEnum_CustomPlaceholderField = 9
	CustomPlaceholderFieldEnum_SALE_PRICE           CustomPlaceholderFieldEnum_CustomPlaceholderField = 10
	CustomPlaceholderFieldEnum_FORMATTED_SALE_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 11
	CustomPlaceholderFieldEnum_IMAGE_URL            CustomPlaceholderFieldEnum_CustomPlaceholderField = 12
	CustomPlaceholderFieldEnum_ITEM_CATEGORY        CustomPlaceholderFieldEnum_CustomPlaceholderField = 13
	CustomPlaceholderFieldEnum_FINAL_URLS           CustomPlaceholderFieldEnum_CustomPlaceholderField = 14
	CustomPlaceholderFieldEnum_FINAL_MOBILE_URLS    CustomPlaceholderFieldEnum_CustomPlaceholderField = 15
	CustomPlaceholderFieldEnum_TRACKING_URL         CustomPlaceholderFieldEnum_CustomPlaceholderField = 16
	CustomPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS  CustomPlaceholderFieldEnum_CustomPlaceholderField = 17
	CustomPlaceholderFieldEnum_ANDROID_APP_LINK     CustomPlaceholderFieldEnum_CustomPlaceholderField = 18
	CustomPlaceholderFieldEnum_SIMILAR_IDS          CustomPlaceholderFieldEnum_CustomPlaceholderField = 19
	CustomPlaceholderFieldEnum_IOS_APP_LINK         CustomPlaceholderFieldEnum_CustomPlaceholderField = 20
	CustomPlaceholderFieldEnum_IOS_APP_STORE_ID     CustomPlaceholderFieldEnum_CustomPlaceholderField = 21
)

// Enum value maps for CustomPlaceholderFieldEnum_CustomPlaceholderField.
var (
	CustomPlaceholderFieldEnum_CustomPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "ID",
		3:  "ID2",
		4:  "ITEM_TITLE",
		5:  "ITEM_SUBTITLE",
		6:  "ITEM_DESCRIPTION",
		7:  "ITEM_ADDRESS",
		8:  "PRICE",
		9:  "FORMATTED_PRICE",
		10: "SALE_PRICE",
		11: "FORMATTED_SALE_PRICE",
		12: "IMAGE_URL",
		13: "ITEM_CATEGORY",
		14: "FINAL_URLS",
		15: "FINAL_MOBILE_URLS",
		16: "TRACKING_URL",
		17: "CONTEXTUAL_KEYWORDS",
		18: "ANDROID_APP_LINK",
		19: "SIMILAR_IDS",
		20: "IOS_APP_LINK",
		21: "IOS_APP_STORE_ID",
	}
	CustomPlaceholderFieldEnum_CustomPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":          0,
		"UNKNOWN":              1,
		"ID":                   2,
		"ID2":                  3,
		"ITEM_TITLE":           4,
		"ITEM_SUBTITLE":        5,
		"ITEM_DESCRIPTION":     6,
		"ITEM_ADDRESS":         7,
		"PRICE":                8,
		"FORMATTED_PRICE":      9,
		"SALE_PRICE":           10,
		"FORMATTED_SALE_PRICE": 11,
		"IMAGE_URL":            12,
		"ITEM_CATEGORY":        13,
		"FINAL_URLS":           14,
		"FINAL_MOBILE_URLS":    15,
		"TRACKING_URL":         16,
		"CONTEXTUAL_KEYWORDS":  17,
		"ANDROID_APP_LINK":     18,
		"SIMILAR_IDS":          19,
		"IOS_APP_LINK":         20,
		"IOS_APP_STORE_ID":     21,
	}
)

func (x CustomPlaceholderFieldEnum_CustomPlaceholderField) Enum() *CustomPlaceholderFieldEnum_CustomPlaceholderField {
	p := new(CustomPlaceholderFieldEnum_CustomPlaceholderField)
	*p = x
	return p
}

func (x CustomPlaceholderFieldEnum_CustomPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomPlaceholderFieldEnum_CustomPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (CustomPlaceholderFieldEnum_CustomPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_enumTypes[0]
}

func (x CustomPlaceholderFieldEnum_CustomPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomPlaceholderFieldEnum_CustomPlaceholderField.Descriptor instead.
func (CustomPlaceholderFieldEnum_CustomPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

type CustomPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CustomPlaceholderFieldEnum) Reset() {
	*x = CustomPlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomPlaceholderFieldEnum) ProtoMessage() {}

func (x *CustomPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*CustomPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_custom_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x03,
	0x0a, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x9f, 0x03, 0x0a,
	0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x49, 0x44, 0x32, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x54,
	0x49, 0x54, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x53,
	0x55, 0x42, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x54, 0x45,
	0x4d, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x07, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f,
	0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10,
	0x09, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10,
	0x0a, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x53,
	0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10, 0x0d, 0x12, 0x0e, 0x0a,
	0x0a, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0e, 0x12, 0x15, 0x0a,
	0x11, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52,
	0x4c, 0x53, 0x10, 0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47,
	0x5f, 0x55, 0x52, 0x4c, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58,
	0x54, 0x55, 0x41, 0x4c, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x11, 0x12,
	0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c,
	0x49, 0x4e, 0x4b, 0x10, 0x12, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52,
	0x5f, 0x49, 0x44, 0x53, 0x10, 0x13, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x14, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4f, 0x53, 0x5f,
	0x41, 0x50, 0x50, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x15, 0x42, 0xf5,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescData = file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_goTypes = []interface{}{
	(CustomPlaceholderFieldEnum_CustomPlaceholderField)(0), // 0: google.ads.googleads.v10.enums.CustomPlaceholderFieldEnum.CustomPlaceholderField
	(*CustomPlaceholderFieldEnum)(nil),                     // 1: google.ads.googleads.v10.enums.CustomPlaceholderFieldEnum
}
var file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_init() }
func file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_init() {
	if File_google_ads_googleads_v10_enums_custom_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomPlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_custom_placeholder_field_proto = out.File
	file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_custom_placeholder_field_proto_depIdxs = nil
}