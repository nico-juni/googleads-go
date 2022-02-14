// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/enums/offline_user_data_job_match_rate_range.proto

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

type OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange int32

const (
	OfflineUserDataJobMatchRateRangeEnum_UNSPECIFIED              OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 0
	OfflineUserDataJobMatchRateRangeEnum_UNKNOWN                  OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 1
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_LESS_THAN_20 OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 2
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_20_TO_30     OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 3
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_31_TO_40     OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 4
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_41_TO_50     OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 5
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_51_TO_60     OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 6
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_61_TO_70     OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 7
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_71_TO_80     OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 8
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_81_TO_90     OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 9
	OfflineUserDataJobMatchRateRangeEnum_MATCH_RANGE_91_TO_100    OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange = 10
)

// Enum value maps for OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange.
var (
	OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "MATCH_RANGE_LESS_THAN_20",
		3:  "MATCH_RANGE_20_TO_30",
		4:  "MATCH_RANGE_31_TO_40",
		5:  "MATCH_RANGE_41_TO_50",
		6:  "MATCH_RANGE_51_TO_60",
		7:  "MATCH_RANGE_61_TO_70",
		8:  "MATCH_RANGE_71_TO_80",
		9:  "MATCH_RANGE_81_TO_90",
		10: "MATCH_RANGE_91_TO_100",
	}
	OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange_value = map[string]int32{
		"UNSPECIFIED":              0,
		"UNKNOWN":                  1,
		"MATCH_RANGE_LESS_THAN_20": 2,
		"MATCH_RANGE_20_TO_30":     3,
		"MATCH_RANGE_31_TO_40":     4,
		"MATCH_RANGE_41_TO_50":     5,
		"MATCH_RANGE_51_TO_60":     6,
		"MATCH_RANGE_61_TO_70":     7,
		"MATCH_RANGE_71_TO_80":     8,
		"MATCH_RANGE_81_TO_90":     9,
		"MATCH_RANGE_91_TO_100":    10,
	}
)

func (x OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange) Enum() *OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange {
	p := new(OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange)
	*p = x
	return p
}

func (x OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_enumTypes[0].Descriptor()
}

func (OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_enumTypes[0]
}

func (x OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange.Descriptor instead.
func (OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescGZIP(), []int{0, 0}
}

type OfflineUserDataJobMatchRateRangeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OfflineUserDataJobMatchRateRangeEnum) Reset() {
	*x = OfflineUserDataJobMatchRateRangeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineUserDataJobMatchRateRangeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineUserDataJobMatchRateRangeEnum) ProtoMessage() {}

func (x *OfflineUserDataJobMatchRateRangeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineUserDataJobMatchRateRangeEnum.ProtoReflect.Descriptor instead.
func (*OfflineUserDataJobMatchRateRangeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x24,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4a,
	0x6f, 0x62, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0xaf, 0x02, 0x0a, 0x20, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x6f, 0x62, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e,
	0x5f, 0x32, 0x30, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x32, 0x30, 0x5f, 0x54, 0x4f, 0x5f, 0x33, 0x30, 0x10, 0x03, 0x12,
	0x18, 0x0a, 0x14, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x33,
	0x31, 0x5f, 0x54, 0x4f, 0x5f, 0x34, 0x30, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x34, 0x31, 0x5f, 0x54, 0x4f, 0x5f, 0x35,
	0x30, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x35, 0x31, 0x5f, 0x54, 0x4f, 0x5f, 0x36, 0x30, 0x10, 0x06, 0x12, 0x18, 0x0a,
	0x14, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x36, 0x31, 0x5f,
	0x54, 0x4f, 0x5f, 0x37, 0x30, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x37, 0x31, 0x5f, 0x54, 0x4f, 0x5f, 0x38, 0x30, 0x10,
	0x08, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x38, 0x31, 0x5f, 0x54, 0x4f, 0x5f, 0x39, 0x30, 0x10, 0x09, 0x12, 0x19, 0x0a, 0x15, 0x4d,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x39, 0x31, 0x5f, 0x54, 0x4f,
	0x5f, 0x31, 0x30, 0x30, 0x10, 0x0a, 0x42, 0xff, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x25, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4a, 0x6f,
	0x62, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x30, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescData = file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_goTypes = []interface{}{
	(OfflineUserDataJobMatchRateRangeEnum_OfflineUserDataJobMatchRateRange)(0), // 0: google.ads.googleads.v10.enums.OfflineUserDataJobMatchRateRangeEnum.OfflineUserDataJobMatchRateRange
	(*OfflineUserDataJobMatchRateRangeEnum)(nil),                               // 1: google.ads.googleads.v10.enums.OfflineUserDataJobMatchRateRangeEnum
}
var file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_init() }
func file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_init() {
	if File_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineUserDataJobMatchRateRangeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto = out.File
	file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_offline_user_data_job_match_rate_range_proto_depIdxs = nil
}
