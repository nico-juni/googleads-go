// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/errors/asset_set_asset_error.proto

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

type AssetSetAssetErrorEnum_AssetSetAssetError int32

const (
	AssetSetAssetErrorEnum_UNSPECIFIED            AssetSetAssetErrorEnum_AssetSetAssetError = 0
	AssetSetAssetErrorEnum_UNKNOWN                AssetSetAssetErrorEnum_AssetSetAssetError = 1
	AssetSetAssetErrorEnum_INVALID_ASSET_TYPE     AssetSetAssetErrorEnum_AssetSetAssetError = 2
	AssetSetAssetErrorEnum_INVALID_ASSET_SET_TYPE AssetSetAssetErrorEnum_AssetSetAssetError = 3
	AssetSetAssetErrorEnum_DUPLICATE_EXTERNAL_KEY AssetSetAssetErrorEnum_AssetSetAssetError = 4
)

// Enum value maps for AssetSetAssetErrorEnum_AssetSetAssetError.
var (
	AssetSetAssetErrorEnum_AssetSetAssetError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INVALID_ASSET_TYPE",
		3: "INVALID_ASSET_SET_TYPE",
		4: "DUPLICATE_EXTERNAL_KEY",
	}
	AssetSetAssetErrorEnum_AssetSetAssetError_value = map[string]int32{
		"UNSPECIFIED":            0,
		"UNKNOWN":                1,
		"INVALID_ASSET_TYPE":     2,
		"INVALID_ASSET_SET_TYPE": 3,
		"DUPLICATE_EXTERNAL_KEY": 4,
	}
)

func (x AssetSetAssetErrorEnum_AssetSetAssetError) Enum() *AssetSetAssetErrorEnum_AssetSetAssetError {
	p := new(AssetSetAssetErrorEnum_AssetSetAssetError)
	*p = x
	return p
}

func (x AssetSetAssetErrorEnum_AssetSetAssetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetSetAssetErrorEnum_AssetSetAssetError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_enumTypes[0].Descriptor()
}

func (AssetSetAssetErrorEnum_AssetSetAssetError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_enumTypes[0]
}

func (x AssetSetAssetErrorEnum_AssetSetAssetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetSetAssetErrorEnum_AssetSetAssetError.Descriptor instead.
func (AssetSetAssetErrorEnum_AssetSetAssetError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescGZIP(), []int{0, 0}
}

type AssetSetAssetErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetSetAssetErrorEnum) Reset() {
	*x = AssetSetAssetErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetSetAssetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSetAssetErrorEnum) ProtoMessage() {}

func (x *AssetSetAssetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSetAssetErrorEnum.ProtoReflect.Descriptor instead.
func (*AssetSetAssetErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_asset_set_asset_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a,
	0x16, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x03,
	0x12, 0x1a, 0x0a, 0x16, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x04, 0x42, 0xf7, 0x01, 0x0a,
	0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x42, 0x17, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescData = file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_goTypes = []interface{}{
	(AssetSetAssetErrorEnum_AssetSetAssetError)(0), // 0: google.ads.googleads.v10.errors.AssetSetAssetErrorEnum.AssetSetAssetError
	(*AssetSetAssetErrorEnum)(nil),                 // 1: google.ads.googleads.v10.errors.AssetSetAssetErrorEnum
}
var file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_init() }
func file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_init() {
	if File_google_ads_googleads_v10_errors_asset_set_asset_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetSetAssetErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_asset_set_asset_error_proto = out.File
	file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_asset_set_asset_error_proto_depIdxs = nil
}
