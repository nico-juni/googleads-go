// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/resources/smart_campaign_setting.proto

package resources

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

type SmartCampaignSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName            string                            `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	Campaign                string                            `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	PhoneNumber             *SmartCampaignSetting_PhoneNumber `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	FinalUrl                string                            `protobuf:"bytes,4,opt,name=final_url,json=finalUrl,proto3" json:"final_url,omitempty"`
	AdvertisingLanguageCode string                            `protobuf:"bytes,7,opt,name=advertising_language_code,json=advertisingLanguageCode,proto3" json:"advertising_language_code,omitempty"`
	// Types that are assignable to BusinessSetting:
	//	*SmartCampaignSetting_BusinessName
	//	*SmartCampaignSetting_BusinessLocationId
	BusinessSetting isSmartCampaignSetting_BusinessSetting `protobuf_oneof:"business_setting"`
}

func (x *SmartCampaignSetting) Reset() {
	*x = SmartCampaignSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartCampaignSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartCampaignSetting) ProtoMessage() {}

func (x *SmartCampaignSetting) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartCampaignSetting.ProtoReflect.Descriptor instead.
func (*SmartCampaignSetting) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescGZIP(), []int{0}
}

func (x *SmartCampaignSetting) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *SmartCampaignSetting) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *SmartCampaignSetting) GetPhoneNumber() *SmartCampaignSetting_PhoneNumber {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *SmartCampaignSetting) GetFinalUrl() string {
	if x != nil {
		return x.FinalUrl
	}
	return ""
}

func (x *SmartCampaignSetting) GetAdvertisingLanguageCode() string {
	if x != nil {
		return x.AdvertisingLanguageCode
	}
	return ""
}

func (m *SmartCampaignSetting) GetBusinessSetting() isSmartCampaignSetting_BusinessSetting {
	if m != nil {
		return m.BusinessSetting
	}
	return nil
}

func (x *SmartCampaignSetting) GetBusinessName() string {
	if x, ok := x.GetBusinessSetting().(*SmartCampaignSetting_BusinessName); ok {
		return x.BusinessName
	}
	return ""
}

func (x *SmartCampaignSetting) GetBusinessLocationId() int64 {
	if x, ok := x.GetBusinessSetting().(*SmartCampaignSetting_BusinessLocationId); ok {
		return x.BusinessLocationId
	}
	return 0
}

type isSmartCampaignSetting_BusinessSetting interface {
	isSmartCampaignSetting_BusinessSetting()
}

type SmartCampaignSetting_BusinessName struct {
	BusinessName string `protobuf:"bytes,5,opt,name=business_name,json=businessName,proto3,oneof"`
}

type SmartCampaignSetting_BusinessLocationId struct {
	BusinessLocationId int64 `protobuf:"varint,6,opt,name=business_location_id,json=businessLocationId,proto3,oneof"`
}

func (*SmartCampaignSetting_BusinessName) isSmartCampaignSetting_BusinessSetting() {}

func (*SmartCampaignSetting_BusinessLocationId) isSmartCampaignSetting_BusinessSetting() {}

type SmartCampaignSetting_PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber *string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	CountryCode *string `protobuf:"bytes,2,opt,name=country_code,json=countryCode,proto3,oneof" json:"country_code,omitempty"`
}

func (x *SmartCampaignSetting_PhoneNumber) Reset() {
	*x = SmartCampaignSetting_PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartCampaignSetting_PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartCampaignSetting_PhoneNumber) ProtoMessage() {}

func (x *SmartCampaignSetting_PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartCampaignSetting_PhoneNumber.ProtoReflect.Descriptor instead.
func (*SmartCampaignSetting_PhoneNumber) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SmartCampaignSetting_PhoneNumber) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *SmartCampaignSetting_PhoneNumber) GetCountryCode() string {
	if x != nil && x.CountryCode != nil {
		return *x.CountryCode
	}
	return ""
}

var File_google_ads_googleads_v10_resources_smart_campaign_setting_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdc, 0x05, 0x0a, 0x14, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x5a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x35, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2f, 0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x67, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72,
	0x6c, 0x12, 0x3a, 0x0a, 0x19, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e,
	0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a,
	0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x12, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x7f, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x6f, 0xea, 0x41, 0x6c, 0x0a, 0x2d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x12, 0x0a, 0x10, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x8b,
	0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x19, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x30, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescData = file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_goTypes = []interface{}{
	(*SmartCampaignSetting)(nil),             // 0: google.ads.googleads.v10.resources.SmartCampaignSetting
	(*SmartCampaignSetting_PhoneNumber)(nil), // 1: google.ads.googleads.v10.resources.SmartCampaignSetting.PhoneNumber
}
var file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.SmartCampaignSetting.phone_number:type_name -> google.ads.googleads.v10.resources.SmartCampaignSetting.PhoneNumber
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_init() }
func file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_init() {
	if File_google_ads_googleads_v10_resources_smart_campaign_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartCampaignSetting); i {
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
		file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartCampaignSetting_PhoneNumber); i {
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
	file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SmartCampaignSetting_BusinessName)(nil),
		(*SmartCampaignSetting_BusinessLocationId)(nil),
	}
	file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_smart_campaign_setting_proto = out.File
	file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_smart_campaign_setting_proto_depIdxs = nil
}
