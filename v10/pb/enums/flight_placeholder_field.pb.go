// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/enums/flight_placeholder_field.proto

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

type FlightPlaceholderFieldEnum_FlightPlaceholderField int32

const (
	FlightPlaceholderFieldEnum_UNSPECIFIED             FlightPlaceholderFieldEnum_FlightPlaceholderField = 0
	FlightPlaceholderFieldEnum_UNKNOWN                 FlightPlaceholderFieldEnum_FlightPlaceholderField = 1
	FlightPlaceholderFieldEnum_DESTINATION_ID          FlightPlaceholderFieldEnum_FlightPlaceholderField = 2
	FlightPlaceholderFieldEnum_ORIGIN_ID               FlightPlaceholderFieldEnum_FlightPlaceholderField = 3
	FlightPlaceholderFieldEnum_FLIGHT_DESCRIPTION      FlightPlaceholderFieldEnum_FlightPlaceholderField = 4
	FlightPlaceholderFieldEnum_ORIGIN_NAME             FlightPlaceholderFieldEnum_FlightPlaceholderField = 5
	FlightPlaceholderFieldEnum_DESTINATION_NAME        FlightPlaceholderFieldEnum_FlightPlaceholderField = 6
	FlightPlaceholderFieldEnum_FLIGHT_PRICE            FlightPlaceholderFieldEnum_FlightPlaceholderField = 7
	FlightPlaceholderFieldEnum_FORMATTED_PRICE         FlightPlaceholderFieldEnum_FlightPlaceholderField = 8
	FlightPlaceholderFieldEnum_FLIGHT_SALE_PRICE       FlightPlaceholderFieldEnum_FlightPlaceholderField = 9
	FlightPlaceholderFieldEnum_FORMATTED_SALE_PRICE    FlightPlaceholderFieldEnum_FlightPlaceholderField = 10
	FlightPlaceholderFieldEnum_IMAGE_URL               FlightPlaceholderFieldEnum_FlightPlaceholderField = 11
	FlightPlaceholderFieldEnum_FINAL_URLS              FlightPlaceholderFieldEnum_FlightPlaceholderField = 12
	FlightPlaceholderFieldEnum_FINAL_MOBILE_URLS       FlightPlaceholderFieldEnum_FlightPlaceholderField = 13
	FlightPlaceholderFieldEnum_TRACKING_URL            FlightPlaceholderFieldEnum_FlightPlaceholderField = 14
	FlightPlaceholderFieldEnum_ANDROID_APP_LINK        FlightPlaceholderFieldEnum_FlightPlaceholderField = 15
	FlightPlaceholderFieldEnum_SIMILAR_DESTINATION_IDS FlightPlaceholderFieldEnum_FlightPlaceholderField = 16
	FlightPlaceholderFieldEnum_IOS_APP_LINK            FlightPlaceholderFieldEnum_FlightPlaceholderField = 17
	FlightPlaceholderFieldEnum_IOS_APP_STORE_ID        FlightPlaceholderFieldEnum_FlightPlaceholderField = 18
)

// Enum value maps for FlightPlaceholderFieldEnum_FlightPlaceholderField.
var (
	FlightPlaceholderFieldEnum_FlightPlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DESTINATION_ID",
		3:  "ORIGIN_ID",
		4:  "FLIGHT_DESCRIPTION",
		5:  "ORIGIN_NAME",
		6:  "DESTINATION_NAME",
		7:  "FLIGHT_PRICE",
		8:  "FORMATTED_PRICE",
		9:  "FLIGHT_SALE_PRICE",
		10: "FORMATTED_SALE_PRICE",
		11: "IMAGE_URL",
		12: "FINAL_URLS",
		13: "FINAL_MOBILE_URLS",
		14: "TRACKING_URL",
		15: "ANDROID_APP_LINK",
		16: "SIMILAR_DESTINATION_IDS",
		17: "IOS_APP_LINK",
		18: "IOS_APP_STORE_ID",
	}
	FlightPlaceholderFieldEnum_FlightPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":             0,
		"UNKNOWN":                 1,
		"DESTINATION_ID":          2,
		"ORIGIN_ID":               3,
		"FLIGHT_DESCRIPTION":      4,
		"ORIGIN_NAME":             5,
		"DESTINATION_NAME":        6,
		"FLIGHT_PRICE":            7,
		"FORMATTED_PRICE":         8,
		"FLIGHT_SALE_PRICE":       9,
		"FORMATTED_SALE_PRICE":    10,
		"IMAGE_URL":               11,
		"FINAL_URLS":              12,
		"FINAL_MOBILE_URLS":       13,
		"TRACKING_URL":            14,
		"ANDROID_APP_LINK":        15,
		"SIMILAR_DESTINATION_IDS": 16,
		"IOS_APP_LINK":            17,
		"IOS_APP_STORE_ID":        18,
	}
)

func (x FlightPlaceholderFieldEnum_FlightPlaceholderField) Enum() *FlightPlaceholderFieldEnum_FlightPlaceholderField {
	p := new(FlightPlaceholderFieldEnum_FlightPlaceholderField)
	*p = x
	return p
}

func (x FlightPlaceholderFieldEnum_FlightPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlightPlaceholderFieldEnum_FlightPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (FlightPlaceholderFieldEnum_FlightPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_enumTypes[0]
}

func (x FlightPlaceholderFieldEnum_FlightPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlightPlaceholderFieldEnum_FlightPlaceholderField.Descriptor instead.
func (FlightPlaceholderFieldEnum_FlightPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

type FlightPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlightPlaceholderFieldEnum) Reset() {
	*x = FlightPlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlightPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlightPlaceholderFieldEnum) ProtoMessage() {}

func (x *FlightPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlightPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*FlightPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_flight_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x03,
	0x0a, 0x1a, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x93, 0x03, 0x0a,
	0x16, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x52, 0x49,
	0x47, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x4c, 0x49, 0x47,
	0x48, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04,
	0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10,
	0x05, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x4c, 0x49, 0x47, 0x48,
	0x54, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4f, 0x52,
	0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x08, 0x12, 0x15,
	0x0a, 0x11, 0x46, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52,
	0x49, 0x43, 0x45, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54,
	0x45, 0x44, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x0a, 0x12,
	0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0b, 0x12, 0x0e,
	0x0a, 0x0a, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0c, 0x12, 0x15,
	0x0a, 0x11, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55,
	0x52, 0x4c, 0x53, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e,
	0x47, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44, 0x52, 0x4f,
	0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x0f, 0x12, 0x1b, 0x0a,
	0x17, 0x53, 0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x49, 0x4e, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f,
	0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x11, 0x12, 0x14, 0x0a, 0x10,
	0x49, 0x4f, 0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44,
	0x10, 0x12, 0x42, 0xf6, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1c, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30,
	0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescData = file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_goTypes = []interface{}{
	(FlightPlaceholderFieldEnum_FlightPlaceholderField)(0), // 0: google.ads.googleads.v10.enums.FlightPlaceholderFieldEnum.FlightPlaceholderField
	(*FlightPlaceholderFieldEnum)(nil),                     // 1: google.ads.googleads.v10.enums.FlightPlaceholderFieldEnum
}
var file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_init() }
func file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_init() {
	if File_google_ads_googleads_v10_enums_flight_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlightPlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_flight_placeholder_field_proto = out.File
	file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_flight_placeholder_field_proto_depIdxs = nil
}