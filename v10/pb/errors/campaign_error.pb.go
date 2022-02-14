// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/errors/campaign_error.proto

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

type CampaignErrorEnum_CampaignError int32

const (
	CampaignErrorEnum_UNSPECIFIED                                                             CampaignErrorEnum_CampaignError = 0
	CampaignErrorEnum_UNKNOWN                                                                 CampaignErrorEnum_CampaignError = 1
	CampaignErrorEnum_CANNOT_TARGET_CONTENT_NETWORK                                           CampaignErrorEnum_CampaignError = 3
	CampaignErrorEnum_CANNOT_TARGET_SEARCH_NETWORK                                            CampaignErrorEnum_CampaignError = 4
	CampaignErrorEnum_CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH                      CampaignErrorEnum_CampaignError = 5
	CampaignErrorEnum_CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN                            CampaignErrorEnum_CampaignError = 6
	CampaignErrorEnum_CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK                               CampaignErrorEnum_CampaignError = 7
	CampaignErrorEnum_CANNOT_TARGET_PARTNER_SEARCH_NETWORK                                    CampaignErrorEnum_CampaignError = 8
	CampaignErrorEnum_CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY CampaignErrorEnum_CampaignError = 9
	CampaignErrorEnum_CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS                      CampaignErrorEnum_CampaignError = 10
	CampaignErrorEnum_CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN                                        CampaignErrorEnum_CampaignError = 11
	CampaignErrorEnum_DUPLICATE_CAMPAIGN_NAME                                                 CampaignErrorEnum_CampaignError = 12
	CampaignErrorEnum_INCOMPATIBLE_CAMPAIGN_FIELD                                             CampaignErrorEnum_CampaignError = 13
	CampaignErrorEnum_INVALID_CAMPAIGN_NAME                                                   CampaignErrorEnum_CampaignError = 14
	CampaignErrorEnum_INVALID_AD_SERVING_OPTIMIZATION_STATUS                                  CampaignErrorEnum_CampaignError = 15
	CampaignErrorEnum_INVALID_TRACKING_URL                                                    CampaignErrorEnum_CampaignError = 16
	CampaignErrorEnum_CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING              CampaignErrorEnum_CampaignError = 17
	CampaignErrorEnum_MAX_IMPRESSIONS_NOT_IN_RANGE                                            CampaignErrorEnum_CampaignError = 18
	CampaignErrorEnum_TIME_UNIT_NOT_SUPPORTED                                                 CampaignErrorEnum_CampaignError = 19
	CampaignErrorEnum_INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED                           CampaignErrorEnum_CampaignError = 20
	CampaignErrorEnum_BUDGET_CANNOT_BE_SHARED                                                 CampaignErrorEnum_CampaignError = 21
	CampaignErrorEnum_CAMPAIGN_CANNOT_USE_SHARED_BUDGET                                       CampaignErrorEnum_CampaignError = 22
	CampaignErrorEnum_CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS                            CampaignErrorEnum_CampaignError = 23
	CampaignErrorEnum_CAMPAIGN_LABEL_DOES_NOT_EXIST                                           CampaignErrorEnum_CampaignError = 24
	CampaignErrorEnum_CAMPAIGN_LABEL_ALREADY_EXISTS                                           CampaignErrorEnum_CampaignError = 25
	CampaignErrorEnum_MISSING_SHOPPING_SETTING                                                CampaignErrorEnum_CampaignError = 26
	CampaignErrorEnum_INVALID_SHOPPING_SALES_COUNTRY                                          CampaignErrorEnum_CampaignError = 27
	CampaignErrorEnum_ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE                 CampaignErrorEnum_CampaignError = 31
	CampaignErrorEnum_INVALID_ADVERTISING_CHANNEL_SUB_TYPE                                    CampaignErrorEnum_CampaignError = 32
	CampaignErrorEnum_AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED                                CampaignErrorEnum_CampaignError = 33
	CampaignErrorEnum_CANNOT_SET_AD_ROTATION_MODE                                             CampaignErrorEnum_CampaignError = 34
	CampaignErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED                             CampaignErrorEnum_CampaignError = 35
	CampaignErrorEnum_CANNOT_SET_DATE_TO_PAST                                                 CampaignErrorEnum_CampaignError = 36
	CampaignErrorEnum_MISSING_HOTEL_CUSTOMER_LINK                                             CampaignErrorEnum_CampaignError = 37
	CampaignErrorEnum_INVALID_HOTEL_CUSTOMER_LINK                                             CampaignErrorEnum_CampaignError = 38
	CampaignErrorEnum_MISSING_HOTEL_SETTING                                                   CampaignErrorEnum_CampaignError = 39
	CampaignErrorEnum_CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP          CampaignErrorEnum_CampaignError = 40
	CampaignErrorEnum_APP_NOT_FOUND                                                           CampaignErrorEnum_CampaignError = 41
	CampaignErrorEnum_SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE                   CampaignErrorEnum_CampaignError = 42
	CampaignErrorEnum_MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS                         CampaignErrorEnum_CampaignError = 43
	CampaignErrorEnum_INSUFFICIENT_APP_INSTALLS_COUNT                                         CampaignErrorEnum_CampaignError = 44
	CampaignErrorEnum_SENSITIVE_CATEGORY_APP                                                  CampaignErrorEnum_CampaignError = 45
	CampaignErrorEnum_HEC_AGREEMENT_REQUIRED                                                  CampaignErrorEnum_CampaignError = 46
	CampaignErrorEnum_NOT_COMPATIBLE_WITH_VIEW_THROUGH_CONVERSION_OPTIMIZATION                CampaignErrorEnum_CampaignError = 49
	CampaignErrorEnum_INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE                                CampaignErrorEnum_CampaignError = 48
	CampaignErrorEnum_CANNOT_CREATE_APP_PRE_REGISTRATION_FOR_NON_ANDROID_APP                  CampaignErrorEnum_CampaignError = 50
	CampaignErrorEnum_APP_NOT_AVAILABLE_TO_CREATE_APP_PRE_REGISTRATION_CAMPAIGN               CampaignErrorEnum_CampaignError = 51
	CampaignErrorEnum_INCOMPATIBLE_BUDGET_TYPE                                                CampaignErrorEnum_CampaignError = 52
)

// Enum value maps for CampaignErrorEnum_CampaignError.
var (
	CampaignErrorEnum_CampaignError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "CANNOT_TARGET_CONTENT_NETWORK",
		4:  "CANNOT_TARGET_SEARCH_NETWORK",
		5:  "CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH",
		6:  "CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN",
		7:  "CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK",
		8:  "CANNOT_TARGET_PARTNER_SEARCH_NETWORK",
		9:  "CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY",
		10: "CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS",
		11: "CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN",
		12: "DUPLICATE_CAMPAIGN_NAME",
		13: "INCOMPATIBLE_CAMPAIGN_FIELD",
		14: "INVALID_CAMPAIGN_NAME",
		15: "INVALID_AD_SERVING_OPTIMIZATION_STATUS",
		16: "INVALID_TRACKING_URL",
		17: "CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING",
		18: "MAX_IMPRESSIONS_NOT_IN_RANGE",
		19: "TIME_UNIT_NOT_SUPPORTED",
		20: "INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED",
		21: "BUDGET_CANNOT_BE_SHARED",
		22: "CAMPAIGN_CANNOT_USE_SHARED_BUDGET",
		23: "CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS",
		24: "CAMPAIGN_LABEL_DOES_NOT_EXIST",
		25: "CAMPAIGN_LABEL_ALREADY_EXISTS",
		26: "MISSING_SHOPPING_SETTING",
		27: "INVALID_SHOPPING_SALES_COUNTRY",
		31: "ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
		32: "INVALID_ADVERTISING_CHANNEL_SUB_TYPE",
		33: "AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED",
		34: "CANNOT_SET_AD_ROTATION_MODE",
		35: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
		36: "CANNOT_SET_DATE_TO_PAST",
		37: "MISSING_HOTEL_CUSTOMER_LINK",
		38: "INVALID_HOTEL_CUSTOMER_LINK",
		39: "MISSING_HOTEL_SETTING",
		40: "CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP",
		41: "APP_NOT_FOUND",
		42: "SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE",
		43: "MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS",
		44: "INSUFFICIENT_APP_INSTALLS_COUNT",
		45: "SENSITIVE_CATEGORY_APP",
		46: "HEC_AGREEMENT_REQUIRED",
		49: "NOT_COMPATIBLE_WITH_VIEW_THROUGH_CONVERSION_OPTIMIZATION",
		48: "INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE",
		50: "CANNOT_CREATE_APP_PRE_REGISTRATION_FOR_NON_ANDROID_APP",
		51: "APP_NOT_AVAILABLE_TO_CREATE_APP_PRE_REGISTRATION_CAMPAIGN",
		52: "INCOMPATIBLE_BUDGET_TYPE",
	}
	CampaignErrorEnum_CampaignError_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"CANNOT_TARGET_CONTENT_NETWORK": 3,
		"CANNOT_TARGET_SEARCH_NETWORK":  4,
		"CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH":                      5,
		"CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN":                            6,
		"CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK":                               7,
		"CANNOT_TARGET_PARTNER_SEARCH_NETWORK":                                    8,
		"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY": 9,
		"CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS":                      10,
		"CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN":                                        11,
		"DUPLICATE_CAMPAIGN_NAME":                                                 12,
		"INCOMPATIBLE_CAMPAIGN_FIELD":                                             13,
		"INVALID_CAMPAIGN_NAME":                                                   14,
		"INVALID_AD_SERVING_OPTIMIZATION_STATUS":                                  15,
		"INVALID_TRACKING_URL":                                                    16,
		"CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING":              17,
		"MAX_IMPRESSIONS_NOT_IN_RANGE":                                            18,
		"TIME_UNIT_NOT_SUPPORTED":                                                 19,
		"INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED":                           20,
		"BUDGET_CANNOT_BE_SHARED":                                                 21,
		"CAMPAIGN_CANNOT_USE_SHARED_BUDGET":                                       22,
		"CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS":                            23,
		"CAMPAIGN_LABEL_DOES_NOT_EXIST":                                           24,
		"CAMPAIGN_LABEL_ALREADY_EXISTS":                                           25,
		"MISSING_SHOPPING_SETTING":                                                26,
		"INVALID_SHOPPING_SALES_COUNTRY":                                          27,
		"ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE":                 31,
		"INVALID_ADVERTISING_CHANNEL_SUB_TYPE":                                    32,
		"AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED":                                33,
		"CANNOT_SET_AD_ROTATION_MODE":                                             34,
		"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED":                             35,
		"CANNOT_SET_DATE_TO_PAST":                                                 36,
		"MISSING_HOTEL_CUSTOMER_LINK":                                             37,
		"INVALID_HOTEL_CUSTOMER_LINK":                                             38,
		"MISSING_HOTEL_SETTING":                                                   39,
		"CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP":          40,
		"APP_NOT_FOUND":                                                           41,
		"SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE":                   42,
		"MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS":                         43,
		"INSUFFICIENT_APP_INSTALLS_COUNT":                                         44,
		"SENSITIVE_CATEGORY_APP":                                                  45,
		"HEC_AGREEMENT_REQUIRED":                                                  46,
		"NOT_COMPATIBLE_WITH_VIEW_THROUGH_CONVERSION_OPTIMIZATION":                49,
		"INVALID_EXCLUDED_PARENT_ASSET_FIELD_TYPE":                                48,
		"CANNOT_CREATE_APP_PRE_REGISTRATION_FOR_NON_ANDROID_APP":                  50,
		"APP_NOT_AVAILABLE_TO_CREATE_APP_PRE_REGISTRATION_CAMPAIGN":               51,
		"INCOMPATIBLE_BUDGET_TYPE":                                                52,
	}
)

func (x CampaignErrorEnum_CampaignError) Enum() *CampaignErrorEnum_CampaignError {
	p := new(CampaignErrorEnum_CampaignError)
	*p = x
	return p
}

func (x CampaignErrorEnum_CampaignError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignErrorEnum_CampaignError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_campaign_error_proto_enumTypes[0].Descriptor()
}

func (CampaignErrorEnum_CampaignError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_campaign_error_proto_enumTypes[0]
}

func (x CampaignErrorEnum_CampaignError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignErrorEnum_CampaignError.Descriptor instead.
func (CampaignErrorEnum_CampaignError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescGZIP(), []int{0, 0}
}

type CampaignErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignErrorEnum) Reset() {
	*x = CampaignErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_campaign_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignErrorEnum) ProtoMessage() {}

func (x *CampaignErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_campaign_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignErrorEnum.ProtoReflect.Descriptor instead.
func (*CampaignErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_campaign_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_campaign_error_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x0f, 0x0a, 0x11, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xb9, 0x0f, 0x0a, 0x0d,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x43,
	0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x03, 0x12, 0x20,
	0x0a, 0x1c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f,
	0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x04,
	0x12, 0x36, 0x0a, 0x32, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45,
	0x54, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f,
	0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x05, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45,
	0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x50, 0x4d, 0x5f,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x06, 0x12, 0x2d, 0x0a, 0x29, 0x43, 0x41,
	0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47,
	0x45, 0x54, 0x5f, 0x41, 0x54, 0x5f, 0x4c, 0x45, 0x41, 0x53, 0x54, 0x5f, 0x4f, 0x4e, 0x45, 0x5f,
	0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x07, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x41, 0x4e,
	0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x4e,
	0x45, 0x52, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x10, 0x08, 0x12, 0x4b, 0x0a, 0x47, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x41,
	0x52, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x54,
	0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43,
	0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x41, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x42, 0x49,
	0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x10, 0x09,
	0x12, 0x36, 0x0a, 0x32, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x44, 0x55, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41,
	0x49, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x54, 0x52, 0x49, 0x41, 0x4c, 0x53, 0x10, 0x0a, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x54, 0x52,
	0x49, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x0b, 0x12, 0x1b,
	0x0a, 0x17, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0c, 0x12, 0x1f, 0x0a, 0x1b, 0x49,
	0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x0d, 0x12, 0x19, 0x0a, 0x15,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0e, 0x12, 0x2a, 0x0a, 0x26, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x41, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x50,
	0x54, 0x49, 0x4d, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x10, 0x0f, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54,
	0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x10, 0x12, 0x3e, 0x0a,
	0x3a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x54, 0x48,
	0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x45,
	0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x11, 0x12, 0x20, 0x0a,
	0x1c, 0x4d, 0x41, 0x58, 0x5f, 0x49, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x12, 0x12,
	0x1b, 0x0a, 0x17, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x13, 0x12, 0x31, 0x0a, 0x2d,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x49, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x14, 0x12,
	0x1b, 0x0a, 0x17, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x42, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x15, 0x12, 0x25, 0x0a, 0x21,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x55, 0x53, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45,
	0x54, 0x10, 0x16, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x54, 0x52, 0x49,
	0x41, 0x4c, 0x53, 0x10, 0x17, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x18, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x19, 0x12, 0x1c, 0x0a, 0x18, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x1a, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x41,
	0x4c, 0x45, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x1b, 0x12, 0x3b, 0x0a,
	0x37, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x1f, 0x12, 0x28, 0x0a, 0x24, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x49, 0x4e,
	0x47, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x20, 0x12, 0x2c, 0x0a, 0x28, 0x41, 0x54, 0x5f, 0x4c, 0x45, 0x41, 0x53, 0x54,
	0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x45, 0x44,
	0x10, 0x21, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54,
	0x5f, 0x41, 0x44, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x10, 0x22, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f,
	0x44, 0x49, 0x46, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x49, 0x46, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x23, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x41, 0x53, 0x54, 0x10,
	0x24, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x48, 0x4f, 0x54,
	0x45, 0x4c, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4e, 0x4b,
	0x10, 0x25, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x48, 0x4f,
	0x54, 0x45, 0x4c, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4e,
	0x4b, 0x10, 0x26, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x48,
	0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x27, 0x12, 0x42,
	0x0a, 0x3e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x53, 0x48, 0x41,
	0x52, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x42, 0x55, 0x44,
	0x47, 0x45, 0x54, 0x5f, 0x57, 0x48, 0x49, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x5f, 0x4f,
	0x46, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x10, 0x28, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x29, 0x12, 0x39, 0x0a, 0x35, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e,
	0x47, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x2a,
	0x12, 0x33, 0x0a, 0x2f, 0x4d, 0x45, 0x52, 0x43, 0x48, 0x41, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x41, 0x44, 0x53, 0x10, 0x2b, 0x12, 0x23, 0x0a, 0x1f, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49,
	0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x2c, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45,
	0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x41, 0x50, 0x50, 0x10, 0x2d, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x45, 0x43, 0x5f, 0x41, 0x47,
	0x52, 0x45, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x2e, 0x12, 0x3c, 0x0a, 0x38, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54,
	0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x54,
	0x48, 0x52, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x31,
	0x12, 0x2c, 0x0a, 0x28, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45, 0x58, 0x43, 0x4c,
	0x55, 0x44, 0x45, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x53, 0x53, 0x45,
	0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x30, 0x12, 0x3a,
	0x0a, 0x36, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f,
	0x41, 0x50, 0x50, 0x5f, 0x50, 0x52, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x41, 0x4e, 0x44,
	0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x32, 0x12, 0x3d, 0x0a, 0x39, 0x41, 0x50,
	0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x54, 0x4f, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x52,
	0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x33, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x43,
	0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x34, 0x42, 0xf2, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42,
	0x12, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescData = file_google_ads_googleads_v10_errors_campaign_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_campaign_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_campaign_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_campaign_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_campaign_error_proto_goTypes = []interface{}{
	(CampaignErrorEnum_CampaignError)(0), // 0: google.ads.googleads.v10.errors.CampaignErrorEnum.CampaignError
	(*CampaignErrorEnum)(nil),            // 1: google.ads.googleads.v10.errors.CampaignErrorEnum
}
var file_google_ads_googleads_v10_errors_campaign_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_campaign_error_proto_init() }
func file_google_ads_googleads_v10_errors_campaign_error_proto_init() {
	if File_google_ads_googleads_v10_errors_campaign_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_campaign_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_campaign_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_campaign_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_campaign_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_campaign_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_campaign_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_campaign_error_proto = out.File
	file_google_ads_googleads_v10_errors_campaign_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_campaign_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_campaign_error_proto_depIdxs = nil
}
