// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/enums/budget_delivery_method.proto

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

type BudgetDeliveryMethodEnum_BudgetDeliveryMethod int32

const (
	BudgetDeliveryMethodEnum_UNSPECIFIED BudgetDeliveryMethodEnum_BudgetDeliveryMethod = 0
	BudgetDeliveryMethodEnum_UNKNOWN     BudgetDeliveryMethodEnum_BudgetDeliveryMethod = 1
	BudgetDeliveryMethodEnum_STANDARD    BudgetDeliveryMethodEnum_BudgetDeliveryMethod = 2
	BudgetDeliveryMethodEnum_ACCELERATED BudgetDeliveryMethodEnum_BudgetDeliveryMethod = 3
)

// Enum value maps for BudgetDeliveryMethodEnum_BudgetDeliveryMethod.
var (
	BudgetDeliveryMethodEnum_BudgetDeliveryMethod_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "STANDARD",
		3: "ACCELERATED",
	}
	BudgetDeliveryMethodEnum_BudgetDeliveryMethod_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"STANDARD":    2,
		"ACCELERATED": 3,
	}
)

func (x BudgetDeliveryMethodEnum_BudgetDeliveryMethod) Enum() *BudgetDeliveryMethodEnum_BudgetDeliveryMethod {
	p := new(BudgetDeliveryMethodEnum_BudgetDeliveryMethod)
	*p = x
	return p
}

func (x BudgetDeliveryMethodEnum_BudgetDeliveryMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BudgetDeliveryMethodEnum_BudgetDeliveryMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_budget_delivery_method_proto_enumTypes[0].Descriptor()
}

func (BudgetDeliveryMethodEnum_BudgetDeliveryMethod) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_budget_delivery_method_proto_enumTypes[0]
}

func (x BudgetDeliveryMethodEnum_BudgetDeliveryMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BudgetDeliveryMethodEnum_BudgetDeliveryMethod.Descriptor instead.
func (BudgetDeliveryMethodEnum_BudgetDeliveryMethod) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescGZIP(), []int{0, 0}
}

type BudgetDeliveryMethodEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BudgetDeliveryMethodEnum) Reset() {
	*x = BudgetDeliveryMethodEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_budget_delivery_method_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetDeliveryMethodEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetDeliveryMethodEnum) ProtoMessage() {}

func (x *BudgetDeliveryMethodEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_budget_delivery_method_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetDeliveryMethodEnum.ProtoReflect.Descriptor instead.
func (*BudgetDeliveryMethodEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_budget_delivery_method_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x18, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x53, 0x0a, 0x14, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x43, 0x43, 0x45, 0x4c, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0xf3, 0x01, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x42, 0x19, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescData = file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_budget_delivery_method_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_budget_delivery_method_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_budget_delivery_method_proto_goTypes = []interface{}{
	(BudgetDeliveryMethodEnum_BudgetDeliveryMethod)(0), // 0: google.ads.googleads.v10.enums.BudgetDeliveryMethodEnum.BudgetDeliveryMethod
	(*BudgetDeliveryMethodEnum)(nil),                   // 1: google.ads.googleads.v10.enums.BudgetDeliveryMethodEnum
}
var file_google_ads_googleads_v10_enums_budget_delivery_method_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_budget_delivery_method_proto_init() }
func file_google_ads_googleads_v10_enums_budget_delivery_method_proto_init() {
	if File_google_ads_googleads_v10_enums_budget_delivery_method_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_budget_delivery_method_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetDeliveryMethodEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_budget_delivery_method_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_budget_delivery_method_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_budget_delivery_method_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_budget_delivery_method_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_budget_delivery_method_proto = out.File
	file_google_ads_googleads_v10_enums_budget_delivery_method_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_budget_delivery_method_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_budget_delivery_method_proto_depIdxs = nil
}
