// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/errors/conversion_upload_error.proto

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

type ConversionUploadErrorEnum_ConversionUploadError int32

const (
	ConversionUploadErrorEnum_UNSPECIFIED                                                                   ConversionUploadErrorEnum_ConversionUploadError = 0
	ConversionUploadErrorEnum_UNKNOWN                                                                       ConversionUploadErrorEnum_ConversionUploadError = 1
	ConversionUploadErrorEnum_TOO_MANY_CONVERSIONS_IN_REQUEST                                               ConversionUploadErrorEnum_ConversionUploadError = 2
	ConversionUploadErrorEnum_UNPARSEABLE_GCLID                                                             ConversionUploadErrorEnum_ConversionUploadError = 3
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_EVENT                                                     ConversionUploadErrorEnum_ConversionUploadError = 42
	ConversionUploadErrorEnum_EXPIRED_EVENT                                                                 ConversionUploadErrorEnum_ConversionUploadError = 43
	ConversionUploadErrorEnum_TOO_RECENT_EVENT                                                              ConversionUploadErrorEnum_ConversionUploadError = 44
	ConversionUploadErrorEnum_EVENT_NOT_FOUND                                                               ConversionUploadErrorEnum_ConversionUploadError = 45
	ConversionUploadErrorEnum_UNAUTHORIZED_CUSTOMER                                                         ConversionUploadErrorEnum_ConversionUploadError = 8
	ConversionUploadErrorEnum_INVALID_CONVERSION_ACTION                                                     ConversionUploadErrorEnum_ConversionUploadError = 9
	ConversionUploadErrorEnum_TOO_RECENT_CONVERSION_ACTION                                                  ConversionUploadErrorEnum_ConversionUploadError = 10
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME                            ConversionUploadErrorEnum_ConversionUploadError = 11
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 12
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 13
	ConversionUploadErrorEnum_ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION            ConversionUploadErrorEnum_ConversionUploadError = 14
	ConversionUploadErrorEnum_ORDER_ID_ALREADY_IN_USE                                                       ConversionUploadErrorEnum_ConversionUploadError = 15
	ConversionUploadErrorEnum_DUPLICATE_ORDER_ID                                                            ConversionUploadErrorEnum_ConversionUploadError = 16
	ConversionUploadErrorEnum_TOO_RECENT_CALL                                                               ConversionUploadErrorEnum_ConversionUploadError = 17
	ConversionUploadErrorEnum_EXPIRED_CALL                                                                  ConversionUploadErrorEnum_ConversionUploadError = 18
	ConversionUploadErrorEnum_CALL_NOT_FOUND                                                                ConversionUploadErrorEnum_ConversionUploadError = 19
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_CALL                                                      ConversionUploadErrorEnum_ConversionUploadError = 20
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME                                  ConversionUploadErrorEnum_ConversionUploadError = 21
	ConversionUploadErrorEnum_UNPARSEABLE_CALLERS_PHONE_NUMBER                                              ConversionUploadErrorEnum_ConversionUploadError = 22
	ConversionUploadErrorEnum_CLICK_CONVERSION_ALREADY_EXISTS                                               ConversionUploadErrorEnum_ConversionUploadError = 23
	ConversionUploadErrorEnum_CALL_CONVERSION_ALREADY_EXISTS                                                ConversionUploadErrorEnum_ConversionUploadError = 24
	ConversionUploadErrorEnum_DUPLICATE_CLICK_CONVERSION_IN_REQUEST                                         ConversionUploadErrorEnum_ConversionUploadError = 25
	ConversionUploadErrorEnum_DUPLICATE_CALL_CONVERSION_IN_REQUEST                                          ConversionUploadErrorEnum_ConversionUploadError = 26
	ConversionUploadErrorEnum_CUSTOM_VARIABLE_NOT_ENABLED                                                   ConversionUploadErrorEnum_ConversionUploadError = 28
	ConversionUploadErrorEnum_CUSTOM_VARIABLE_VALUE_CONTAINS_PII                                            ConversionUploadErrorEnum_ConversionUploadError = 29
	ConversionUploadErrorEnum_INVALID_CUSTOMER_FOR_CLICK                                                    ConversionUploadErrorEnum_ConversionUploadError = 30
	ConversionUploadErrorEnum_INVALID_CUSTOMER_FOR_CALL                                                     ConversionUploadErrorEnum_ConversionUploadError = 31
	ConversionUploadErrorEnum_CONVERSION_NOT_COMPLIANT_WITH_ATT_POLICY                                      ConversionUploadErrorEnum_ConversionUploadError = 32
	ConversionUploadErrorEnum_CLICK_NOT_FOUND                                                               ConversionUploadErrorEnum_ConversionUploadError = 33
	ConversionUploadErrorEnum_INVALID_USER_IDENTIFIER                                                       ConversionUploadErrorEnum_ConversionUploadError = 34
	ConversionUploadErrorEnum_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION_NOT_PERMITTED_WITH_USER_IDENTIFIER    ConversionUploadErrorEnum_ConversionUploadError = 35
	ConversionUploadErrorEnum_UNSUPPORTED_USER_IDENTIFIER                                                   ConversionUploadErrorEnum_ConversionUploadError = 36
	ConversionUploadErrorEnum_GBRAID_WBRAID_BOTH_SET                                                        ConversionUploadErrorEnum_ConversionUploadError = 38
	ConversionUploadErrorEnum_UNPARSEABLE_WBRAID                                                            ConversionUploadErrorEnum_ConversionUploadError = 39
	ConversionUploadErrorEnum_UNPARSEABLE_GBRAID                                                            ConversionUploadErrorEnum_ConversionUploadError = 40
	ConversionUploadErrorEnum_EXTERNALLY_ATTRIBUTED_CONVERSION_TYPE_NOT_PERMITTED_WITH_BRAID                ConversionUploadErrorEnum_ConversionUploadError = 41
	ConversionUploadErrorEnum_ONE_PER_CLICK_CONVERSION_ACTION_NOT_PERMITTED_WITH_BRAID                      ConversionUploadErrorEnum_ConversionUploadError = 46
)

// Enum value maps for ConversionUploadErrorEnum_ConversionUploadError.
var (
	ConversionUploadErrorEnum_ConversionUploadError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "TOO_MANY_CONVERSIONS_IN_REQUEST",
		3:  "UNPARSEABLE_GCLID",
		42: "CONVERSION_PRECEDES_EVENT",
		43: "EXPIRED_EVENT",
		44: "TOO_RECENT_EVENT",
		45: "EVENT_NOT_FOUND",
		8:  "UNAUTHORIZED_CUSTOMER",
		9:  "INVALID_CONVERSION_ACTION",
		10: "TOO_RECENT_CONVERSION_ACTION",
		11: "CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME",
		12: "EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		13: "EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		14: "ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		15: "ORDER_ID_ALREADY_IN_USE",
		16: "DUPLICATE_ORDER_ID",
		17: "TOO_RECENT_CALL",
		18: "EXPIRED_CALL",
		19: "CALL_NOT_FOUND",
		20: "CONVERSION_PRECEDES_CALL",
		21: "CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME",
		22: "UNPARSEABLE_CALLERS_PHONE_NUMBER",
		23: "CLICK_CONVERSION_ALREADY_EXISTS",
		24: "CALL_CONVERSION_ALREADY_EXISTS",
		25: "DUPLICATE_CLICK_CONVERSION_IN_REQUEST",
		26: "DUPLICATE_CALL_CONVERSION_IN_REQUEST",
		28: "CUSTOM_VARIABLE_NOT_ENABLED",
		29: "CUSTOM_VARIABLE_VALUE_CONTAINS_PII",
		30: "INVALID_CUSTOMER_FOR_CLICK",
		31: "INVALID_CUSTOMER_FOR_CALL",
		32: "CONVERSION_NOT_COMPLIANT_WITH_ATT_POLICY",
		33: "CLICK_NOT_FOUND",
		34: "INVALID_USER_IDENTIFIER",
		35: "EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION_NOT_PERMITTED_WITH_USER_IDENTIFIER",
		36: "UNSUPPORTED_USER_IDENTIFIER",
		38: "GBRAID_WBRAID_BOTH_SET",
		39: "UNPARSEABLE_WBRAID",
		40: "UNPARSEABLE_GBRAID",
		41: "EXTERNALLY_ATTRIBUTED_CONVERSION_TYPE_NOT_PERMITTED_WITH_BRAID",
		46: "ONE_PER_CLICK_CONVERSION_ACTION_NOT_PERMITTED_WITH_BRAID",
	}
	ConversionUploadErrorEnum_ConversionUploadError_value = map[string]int32{
		"UNSPECIFIED":                     0,
		"UNKNOWN":                         1,
		"TOO_MANY_CONVERSIONS_IN_REQUEST": 2,
		"UNPARSEABLE_GCLID":               3,
		"CONVERSION_PRECEDES_EVENT":       42,
		"EXPIRED_EVENT":                   43,
		"TOO_RECENT_EVENT":                44,
		"EVENT_NOT_FOUND":                 45,
		"UNAUTHORIZED_CUSTOMER":           8,
		"INVALID_CONVERSION_ACTION":       9,
		"TOO_RECENT_CONVERSION_ACTION":    10,
		"CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME":                            11,
		"EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 12,
		"EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 13,
		"ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION":            14,
		"ORDER_ID_ALREADY_IN_USE":                      15,
		"DUPLICATE_ORDER_ID":                           16,
		"TOO_RECENT_CALL":                              17,
		"EXPIRED_CALL":                                 18,
		"CALL_NOT_FOUND":                               19,
		"CONVERSION_PRECEDES_CALL":                     20,
		"CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME": 21,
		"UNPARSEABLE_CALLERS_PHONE_NUMBER":             22,
		"CLICK_CONVERSION_ALREADY_EXISTS":              23,
		"CALL_CONVERSION_ALREADY_EXISTS":               24,
		"DUPLICATE_CLICK_CONVERSION_IN_REQUEST":        25,
		"DUPLICATE_CALL_CONVERSION_IN_REQUEST":         26,
		"CUSTOM_VARIABLE_NOT_ENABLED":                  28,
		"CUSTOM_VARIABLE_VALUE_CONTAINS_PII":           29,
		"INVALID_CUSTOMER_FOR_CLICK":                   30,
		"INVALID_CUSTOMER_FOR_CALL":                    31,
		"CONVERSION_NOT_COMPLIANT_WITH_ATT_POLICY":     32,
		"CLICK_NOT_FOUND":                              33,
		"INVALID_USER_IDENTIFIER":                      34,
		"EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION_NOT_PERMITTED_WITH_USER_IDENTIFIER": 35,
		"UNSUPPORTED_USER_IDENTIFIER": 36,
		"GBRAID_WBRAID_BOTH_SET":      38,
		"UNPARSEABLE_WBRAID":          39,
		"UNPARSEABLE_GBRAID":          40,
		"EXTERNALLY_ATTRIBUTED_CONVERSION_TYPE_NOT_PERMITTED_WITH_BRAID": 41,
		"ONE_PER_CLICK_CONVERSION_ACTION_NOT_PERMITTED_WITH_BRAID":       46,
	}
)

func (x ConversionUploadErrorEnum_ConversionUploadError) Enum() *ConversionUploadErrorEnum_ConversionUploadError {
	p := new(ConversionUploadErrorEnum_ConversionUploadError)
	*p = x
	return p
}

func (x ConversionUploadErrorEnum_ConversionUploadError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionUploadErrorEnum_ConversionUploadError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_conversion_upload_error_proto_enumTypes[0].Descriptor()
}

func (ConversionUploadErrorEnum_ConversionUploadError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_conversion_upload_error_proto_enumTypes[0]
}

func (x ConversionUploadErrorEnum_ConversionUploadError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionUploadErrorEnum_ConversionUploadError.Descriptor instead.
func (ConversionUploadErrorEnum_ConversionUploadError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescGZIP(), []int{0, 0}
}

type ConversionUploadErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionUploadErrorEnum) Reset() {
	*x = ConversionUploadErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_conversion_upload_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionUploadErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionUploadErrorEnum) ProtoMessage() {}

func (x *ConversionUploadErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_conversion_upload_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionUploadErrorEnum.ProtoReflect.Descriptor instead.
func (*ConversionUploadErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_conversion_upload_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99,
	0x0c, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xfb, 0x0b, 0x0a,
	0x15, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59,
	0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x49, 0x4e, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x50,
	0x41, 0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x47, 0x43, 0x4c, 0x49, 0x44, 0x10, 0x03,
	0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50,
	0x52, 0x45, 0x43, 0x45, 0x44, 0x45, 0x53, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x2a, 0x12,
	0x11, 0x0a, 0x0d, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x10, 0x2b, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x2c, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x2d, 0x12, 0x19, 0x0a,
	0x15, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4f, 0x5f, 0x52,
	0x45, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f,
	0x49, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10,
	0x0b, 0x12, 0x51, 0x0a, 0x4d, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x41, 0x54,
	0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53,
	0x45, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x44,
	0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x0c, 0x12, 0x51, 0x0a, 0x4d, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x45, 0x58,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55,
	0x54, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x46, 0x0a, 0x42, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45,
	0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59,
	0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0e, 0x12,
	0x1b, 0x0a, 0x17, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x4c, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x10, 0x0f, 0x12, 0x16, 0x0a, 0x12,
	0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x49, 0x44, 0x10, 0x10, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x45,
	0x4e, 0x54, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x11, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x43,
	0x41, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x13, 0x12,
	0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52,
	0x45, 0x43, 0x45, 0x44, 0x45, 0x53, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x14, 0x12, 0x30, 0x0a,
	0x2c, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43,
	0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x5f, 0x41, 0x54, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x15, 0x12,
	0x24, 0x0a, 0x20, 0x55, 0x4e, 0x50, 0x41, 0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x43,
	0x41, 0x4c, 0x4c, 0x45, 0x52, 0x53, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d,
	0x42, 0x45, 0x52, 0x10, 0x16, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x17, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x41,
	0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x18, 0x12, 0x29,
	0x0a, 0x25, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x43,
	0x4b, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x19, 0x12, 0x28, 0x0a, 0x24, 0x44, 0x55, 0x50,
	0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x1a, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x56, 0x41,
	0x52, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x1c, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x56,
	0x41, 0x52, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x43, 0x4f,
	0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f, 0x50, 0x49, 0x49, 0x10, 0x1d, 0x12, 0x1e, 0x0a, 0x1a,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x1e, 0x12, 0x1d, 0x0a, 0x19,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x1f, 0x12, 0x2c, 0x0a, 0x28, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x49, 0x41, 0x4e, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x54, 0x54,
	0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x20, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x49,
	0x43, 0x4b, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x21, 0x12, 0x1b,
	0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49,
	0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x22, 0x12, 0x4e, 0x0a, 0x4a, 0x45,
	0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42,
	0x55, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49,
	0x54, 0x54, 0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49,
	0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x23, 0x12, 0x1f, 0x0a, 0x1b, 0x55,
	0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x24, 0x12, 0x1a, 0x0a, 0x16,
	0x47, 0x42, 0x52, 0x41, 0x49, 0x44, 0x5f, 0x57, 0x42, 0x52, 0x41, 0x49, 0x44, 0x5f, 0x42, 0x4f,
	0x54, 0x48, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x26, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x50, 0x41,
	0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x42, 0x52, 0x41, 0x49, 0x44, 0x10, 0x27,
	0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x50, 0x41, 0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x47, 0x42, 0x52, 0x41, 0x49, 0x44, 0x10, 0x28, 0x12, 0x42, 0x0a, 0x3e, 0x45, 0x58, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f,
	0x57, 0x49, 0x54, 0x48, 0x5f, 0x42, 0x52, 0x41, 0x49, 0x44, 0x10, 0x29, 0x12, 0x3c, 0x0a, 0x38,
	0x4f, 0x4e, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x43, 0x4f,
	0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x5f, 0x42, 0x52, 0x41, 0x49, 0x44, 0x10, 0x2e, 0x42, 0xfa, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x42, 0x1a, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a,
	0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescData = file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_conversion_upload_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_conversion_upload_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_conversion_upload_error_proto_goTypes = []interface{}{
	(ConversionUploadErrorEnum_ConversionUploadError)(0), // 0: google.ads.googleads.v10.errors.ConversionUploadErrorEnum.ConversionUploadError
	(*ConversionUploadErrorEnum)(nil),                    // 1: google.ads.googleads.v10.errors.ConversionUploadErrorEnum
}
var file_google_ads_googleads_v10_errors_conversion_upload_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_conversion_upload_error_proto_init() }
func file_google_ads_googleads_v10_errors_conversion_upload_error_proto_init() {
	if File_google_ads_googleads_v10_errors_conversion_upload_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_conversion_upload_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionUploadErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_conversion_upload_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_conversion_upload_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_conversion_upload_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_conversion_upload_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_conversion_upload_error_proto = out.File
	file_google_ads_googleads_v10_errors_conversion_upload_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_conversion_upload_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_conversion_upload_error_proto_depIdxs = nil
}
