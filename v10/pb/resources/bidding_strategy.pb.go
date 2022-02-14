// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/resources/bidding_strategy.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/nico-juni/googleads-go/v10/pb/common"
	enums "github.com/nico-juni/googleads-go/v10/pb/enums"
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

type BiddingStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName            string                                                `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	Id                      *int64                                                `protobuf:"varint,16,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name                    *string                                               `protobuf:"bytes,17,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Status                  enums.BiddingStrategyStatusEnum_BiddingStrategyStatus `protobuf:"varint,15,opt,name=status,proto3,enum=google.ads.googleads.v10.enums.BiddingStrategyStatusEnum_BiddingStrategyStatus" json:"status,omitempty"`
	Type                    enums.BiddingStrategyTypeEnum_BiddingStrategyType     `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v10.enums.BiddingStrategyTypeEnum_BiddingStrategyType" json:"type,omitempty"`
	CurrencyCode            string                                                `protobuf:"bytes,23,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	EffectiveCurrencyCode   *string                                               `protobuf:"bytes,20,opt,name=effective_currency_code,json=effectiveCurrencyCode,proto3,oneof" json:"effective_currency_code,omitempty"`
	CampaignCount           *int64                                                `protobuf:"varint,18,opt,name=campaign_count,json=campaignCount,proto3,oneof" json:"campaign_count,omitempty"`
	NonRemovedCampaignCount *int64                                                `protobuf:"varint,19,opt,name=non_removed_campaign_count,json=nonRemovedCampaignCount,proto3,oneof" json:"non_removed_campaign_count,omitempty"`
	// Types that are assignable to Scheme:
	//	*BiddingStrategy_EnhancedCpc
	//	*BiddingStrategy_MaximizeConversionValue
	//	*BiddingStrategy_MaximizeConversions
	//	*BiddingStrategy_TargetCpa
	//	*BiddingStrategy_TargetImpressionShare
	//	*BiddingStrategy_TargetRoas
	//	*BiddingStrategy_TargetSpend
	Scheme isBiddingStrategy_Scheme `protobuf_oneof:"scheme"`
}

func (x *BiddingStrategy) Reset() {
	*x = BiddingStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_bidding_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiddingStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingStrategy) ProtoMessage() {}

func (x *BiddingStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_bidding_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingStrategy.ProtoReflect.Descriptor instead.
func (*BiddingStrategy) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *BiddingStrategy) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *BiddingStrategy) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BiddingStrategy) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BiddingStrategy) GetStatus() enums.BiddingStrategyStatusEnum_BiddingStrategyStatus {
	if x != nil {
		return x.Status
	}
	return enums.BiddingStrategyStatusEnum_UNSPECIFIED
}

func (x *BiddingStrategy) GetType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	if x != nil {
		return x.Type
	}
	return enums.BiddingStrategyTypeEnum_UNSPECIFIED
}

func (x *BiddingStrategy) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *BiddingStrategy) GetEffectiveCurrencyCode() string {
	if x != nil && x.EffectiveCurrencyCode != nil {
		return *x.EffectiveCurrencyCode
	}
	return ""
}

func (x *BiddingStrategy) GetCampaignCount() int64 {
	if x != nil && x.CampaignCount != nil {
		return *x.CampaignCount
	}
	return 0
}

func (x *BiddingStrategy) GetNonRemovedCampaignCount() int64 {
	if x != nil && x.NonRemovedCampaignCount != nil {
		return *x.NonRemovedCampaignCount
	}
	return 0
}

func (m *BiddingStrategy) GetScheme() isBiddingStrategy_Scheme {
	if m != nil {
		return m.Scheme
	}
	return nil
}

func (x *BiddingStrategy) GetEnhancedCpc() *common.EnhancedCpc {
	if x, ok := x.GetScheme().(*BiddingStrategy_EnhancedCpc); ok {
		return x.EnhancedCpc
	}
	return nil
}

func (x *BiddingStrategy) GetMaximizeConversionValue() *common.MaximizeConversionValue {
	if x, ok := x.GetScheme().(*BiddingStrategy_MaximizeConversionValue); ok {
		return x.MaximizeConversionValue
	}
	return nil
}

func (x *BiddingStrategy) GetMaximizeConversions() *common.MaximizeConversions {
	if x, ok := x.GetScheme().(*BiddingStrategy_MaximizeConversions); ok {
		return x.MaximizeConversions
	}
	return nil
}

func (x *BiddingStrategy) GetTargetCpa() *common.TargetCpa {
	if x, ok := x.GetScheme().(*BiddingStrategy_TargetCpa); ok {
		return x.TargetCpa
	}
	return nil
}

func (x *BiddingStrategy) GetTargetImpressionShare() *common.TargetImpressionShare {
	if x, ok := x.GetScheme().(*BiddingStrategy_TargetImpressionShare); ok {
		return x.TargetImpressionShare
	}
	return nil
}

func (x *BiddingStrategy) GetTargetRoas() *common.TargetRoas {
	if x, ok := x.GetScheme().(*BiddingStrategy_TargetRoas); ok {
		return x.TargetRoas
	}
	return nil
}

func (x *BiddingStrategy) GetTargetSpend() *common.TargetSpend {
	if x, ok := x.GetScheme().(*BiddingStrategy_TargetSpend); ok {
		return x.TargetSpend
	}
	return nil
}

type isBiddingStrategy_Scheme interface {
	isBiddingStrategy_Scheme()
}

type BiddingStrategy_EnhancedCpc struct {
	EnhancedCpc *common.EnhancedCpc `protobuf:"bytes,7,opt,name=enhanced_cpc,json=enhancedCpc,proto3,oneof"`
}

type BiddingStrategy_MaximizeConversionValue struct {
	MaximizeConversionValue *common.MaximizeConversionValue `protobuf:"bytes,21,opt,name=maximize_conversion_value,json=maximizeConversionValue,proto3,oneof"`
}

type BiddingStrategy_MaximizeConversions struct {
	MaximizeConversions *common.MaximizeConversions `protobuf:"bytes,22,opt,name=maximize_conversions,json=maximizeConversions,proto3,oneof"`
}

type BiddingStrategy_TargetCpa struct {
	TargetCpa *common.TargetCpa `protobuf:"bytes,9,opt,name=target_cpa,json=targetCpa,proto3,oneof"`
}

type BiddingStrategy_TargetImpressionShare struct {
	TargetImpressionShare *common.TargetImpressionShare `protobuf:"bytes,48,opt,name=target_impression_share,json=targetImpressionShare,proto3,oneof"`
}

type BiddingStrategy_TargetRoas struct {
	TargetRoas *common.TargetRoas `protobuf:"bytes,11,opt,name=target_roas,json=targetRoas,proto3,oneof"`
}

type BiddingStrategy_TargetSpend struct {
	TargetSpend *common.TargetSpend `protobuf:"bytes,12,opt,name=target_spend,json=targetSpend,proto3,oneof"`
}

func (*BiddingStrategy_EnhancedCpc) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_MaximizeConversionValue) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_MaximizeConversions) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetCpa) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetImpressionShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetRoas) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetSpend) isBiddingStrategy_Scheme() {}

var File_google_ads_googleads_v10_resources_bidding_strategy_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x69, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc3, 0x0b, 0x0a, 0x0f, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x55, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0,
	0x41, 0x05, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x6c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x64,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x40,
	0x0a, 0x17, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52, 0x15, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x2f, 0x0a, 0x0e, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52,
	0x0d, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x45, 0x0a, 0x1a, 0x6e, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x05, 0x52, 0x17, 0x6e, 0x6f,
	0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x0c, 0x65, 0x6e, 0x68, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x43, 0x70, 0x63, 0x48, 0x00, 0x52, 0x0b,
	0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x43, 0x70, 0x63, 0x12, 0x76, 0x0a, 0x19, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x17, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x69, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x69, 0x7a, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4b,
	0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x70, 0x61, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x70, 0x61, 0x48, 0x00,
	0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x70, 0x61, 0x12, 0x70, 0x0a, 0x17, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x30, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x48, 0x00, 0x52, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x4e, 0x0a,
	0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x61, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x73, 0x48,
	0x00, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x73, 0x12, 0x51, 0x0a,
	0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x6e,
	0x64, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x70, 0x65, 0x6e, 0x64,
	0x3a, 0x6e, 0xea, 0x41, 0x6b, 0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x3f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x69, 0x64, 0x7d,
	0x42, 0x08, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x6e, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x86, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x14, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDescData = file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_bidding_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_resources_bidding_strategy_proto_goTypes = []interface{}{
	(*BiddingStrategy)(nil), // 0: google.ads.googleads.v10.resources.BiddingStrategy
	(enums.BiddingStrategyStatusEnum_BiddingStrategyStatus)(0), // 1: google.ads.googleads.v10.enums.BiddingStrategyStatusEnum.BiddingStrategyStatus
	(enums.BiddingStrategyTypeEnum_BiddingStrategyType)(0),     // 2: google.ads.googleads.v10.enums.BiddingStrategyTypeEnum.BiddingStrategyType
	(*common.EnhancedCpc)(nil),                                 // 3: google.ads.googleads.v10.common.EnhancedCpc
	(*common.MaximizeConversionValue)(nil),                     // 4: google.ads.googleads.v10.common.MaximizeConversionValue
	(*common.MaximizeConversions)(nil),                         // 5: google.ads.googleads.v10.common.MaximizeConversions
	(*common.TargetCpa)(nil),                                   // 6: google.ads.googleads.v10.common.TargetCpa
	(*common.TargetImpressionShare)(nil),                       // 7: google.ads.googleads.v10.common.TargetImpressionShare
	(*common.TargetRoas)(nil),                                  // 8: google.ads.googleads.v10.common.TargetRoas
	(*common.TargetSpend)(nil),                                 // 9: google.ads.googleads.v10.common.TargetSpend
}
var file_google_ads_googleads_v10_resources_bidding_strategy_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.BiddingStrategy.status:type_name -> google.ads.googleads.v10.enums.BiddingStrategyStatusEnum.BiddingStrategyStatus
	2, // 1: google.ads.googleads.v10.resources.BiddingStrategy.type:type_name -> google.ads.googleads.v10.enums.BiddingStrategyTypeEnum.BiddingStrategyType
	3, // 2: google.ads.googleads.v10.resources.BiddingStrategy.enhanced_cpc:type_name -> google.ads.googleads.v10.common.EnhancedCpc
	4, // 3: google.ads.googleads.v10.resources.BiddingStrategy.maximize_conversion_value:type_name -> google.ads.googleads.v10.common.MaximizeConversionValue
	5, // 4: google.ads.googleads.v10.resources.BiddingStrategy.maximize_conversions:type_name -> google.ads.googleads.v10.common.MaximizeConversions
	6, // 5: google.ads.googleads.v10.resources.BiddingStrategy.target_cpa:type_name -> google.ads.googleads.v10.common.TargetCpa
	7, // 6: google.ads.googleads.v10.resources.BiddingStrategy.target_impression_share:type_name -> google.ads.googleads.v10.common.TargetImpressionShare
	8, // 7: google.ads.googleads.v10.resources.BiddingStrategy.target_roas:type_name -> google.ads.googleads.v10.common.TargetRoas
	9, // 8: google.ads.googleads.v10.resources.BiddingStrategy.target_spend:type_name -> google.ads.googleads.v10.common.TargetSpend
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_bidding_strategy_proto_init() }
func file_google_ads_googleads_v10_resources_bidding_strategy_proto_init() {
	if File_google_ads_googleads_v10_resources_bidding_strategy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_bidding_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiddingStrategy); i {
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
	file_google_ads_googleads_v10_resources_bidding_strategy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BiddingStrategy_EnhancedCpc)(nil),
		(*BiddingStrategy_MaximizeConversionValue)(nil),
		(*BiddingStrategy_MaximizeConversions)(nil),
		(*BiddingStrategy_TargetCpa)(nil),
		(*BiddingStrategy_TargetImpressionShare)(nil),
		(*BiddingStrategy_TargetRoas)(nil),
		(*BiddingStrategy_TargetSpend)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_bidding_strategy_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_bidding_strategy_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_bidding_strategy_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_bidding_strategy_proto = out.File
	file_google_ads_googleads_v10_resources_bidding_strategy_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_bidding_strategy_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_bidding_strategy_proto_depIdxs = nil
}
