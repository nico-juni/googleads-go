// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/resources/campaign_simulation.proto

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

type CampaignSimulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName       string                                                              `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	CampaignId         int64                                                               `protobuf:"varint,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	Type               enums.SimulationTypeEnum_SimulationType                             `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v10.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,4,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v10.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	StartDate          string                                                              `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate            string                                                              `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// Types that are assignable to PointList:
	//	*CampaignSimulation_CpcBidPointList
	//	*CampaignSimulation_TargetCpaPointList
	//	*CampaignSimulation_TargetRoasPointList
	//	*CampaignSimulation_TargetImpressionSharePointList
	//	*CampaignSimulation_BudgetPointList
	PointList isCampaignSimulation_PointList `protobuf_oneof:"point_list"`
}

func (x *CampaignSimulation) Reset() {
	*x = CampaignSimulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_campaign_simulation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignSimulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignSimulation) ProtoMessage() {}

func (x *CampaignSimulation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_campaign_simulation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignSimulation.ProtoReflect.Descriptor instead.
func (*CampaignSimulation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDescGZIP(), []int{0}
}

func (x *CampaignSimulation) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CampaignSimulation) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *CampaignSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if x != nil {
		return x.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (x *CampaignSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if x != nil {
		return x.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (x *CampaignSimulation) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CampaignSimulation) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (m *CampaignSimulation) GetPointList() isCampaignSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (x *CampaignSimulation) GetCpcBidPointList() *common.CpcBidSimulationPointList {
	if x, ok := x.GetPointList().(*CampaignSimulation_CpcBidPointList); ok {
		return x.CpcBidPointList
	}
	return nil
}

func (x *CampaignSimulation) GetTargetCpaPointList() *common.TargetCpaSimulationPointList {
	if x, ok := x.GetPointList().(*CampaignSimulation_TargetCpaPointList); ok {
		return x.TargetCpaPointList
	}
	return nil
}

func (x *CampaignSimulation) GetTargetRoasPointList() *common.TargetRoasSimulationPointList {
	if x, ok := x.GetPointList().(*CampaignSimulation_TargetRoasPointList); ok {
		return x.TargetRoasPointList
	}
	return nil
}

func (x *CampaignSimulation) GetTargetImpressionSharePointList() *common.TargetImpressionShareSimulationPointList {
	if x, ok := x.GetPointList().(*CampaignSimulation_TargetImpressionSharePointList); ok {
		return x.TargetImpressionSharePointList
	}
	return nil
}

func (x *CampaignSimulation) GetBudgetPointList() *common.BudgetSimulationPointList {
	if x, ok := x.GetPointList().(*CampaignSimulation_BudgetPointList); ok {
		return x.BudgetPointList
	}
	return nil
}

type isCampaignSimulation_PointList interface {
	isCampaignSimulation_PointList()
}

type CampaignSimulation_CpcBidPointList struct {
	CpcBidPointList *common.CpcBidSimulationPointList `protobuf:"bytes,7,opt,name=cpc_bid_point_list,json=cpcBidPointList,proto3,oneof"`
}

type CampaignSimulation_TargetCpaPointList struct {
	TargetCpaPointList *common.TargetCpaSimulationPointList `protobuf:"bytes,8,opt,name=target_cpa_point_list,json=targetCpaPointList,proto3,oneof"`
}

type CampaignSimulation_TargetRoasPointList struct {
	TargetRoasPointList *common.TargetRoasSimulationPointList `protobuf:"bytes,9,opt,name=target_roas_point_list,json=targetRoasPointList,proto3,oneof"`
}

type CampaignSimulation_TargetImpressionSharePointList struct {
	TargetImpressionSharePointList *common.TargetImpressionShareSimulationPointList `protobuf:"bytes,10,opt,name=target_impression_share_point_list,json=targetImpressionSharePointList,proto3,oneof"`
}

type CampaignSimulation_BudgetPointList struct {
	BudgetPointList *common.BudgetSimulationPointList `protobuf:"bytes,11,opt,name=budget_point_list,json=budgetPointList,proto3,oneof"`
}

func (*CampaignSimulation_CpcBidPointList) isCampaignSimulation_PointList() {}

func (*CampaignSimulation_TargetCpaPointList) isCampaignSimulation_PointList() {}

func (*CampaignSimulation_TargetRoasPointList) isCampaignSimulation_PointList() {}

func (*CampaignSimulation_TargetImpressionSharePointList) isCampaignSimulation_PointList() {}

func (*CampaignSimulation_BudgetPointList) isCampaignSimulation_PointList() {}

var File_google_ads_googleads_v10_resources_campaign_simulation_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x73, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x09, 0x0a, 0x12, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x58, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2d, 0x0a,
	0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69,
	0x67, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64,
	0x12, 0x5a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x93, 0x01, 0x0a,
	0x13, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x6e, 0x0a, 0x12, 0x63, 0x70, 0x63, 0x5f, 0x62, 0x69,
	0x64, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x70, 0x63, 0x42, 0x69, 0x64, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x70, 0x63, 0x42, 0x69, 0x64, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x77, 0x0a, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x70, 0x61, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x70,
	0x61, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x12, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x43, 0x70, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x7a, 0x0a, 0x16, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x61, 0x73, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x61, 0x73, 0x53, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x6f,
	0x61, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x9c, 0x01, 0x0a, 0x22,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x1e, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x6d, 0x0a, 0x11, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0xa1, 0x01, 0xea, 0x41, 0x9d, 0x01,
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x53,
	0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x7e,
	0x7b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x7d, 0x7e, 0x7b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x7d, 0x7e, 0x7b, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x42, 0x0c, 0x0a,
	0x0a, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x89, 0x02, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x17, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x30, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDescData = file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_campaign_simulation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_resources_campaign_simulation_proto_goTypes = []interface{}{
	(*CampaignSimulation)(nil),                                               // 0: google.ads.googleads.v10.resources.CampaignSimulation
	(enums.SimulationTypeEnum_SimulationType)(0),                             // 1: google.ads.googleads.v10.enums.SimulationTypeEnum.SimulationType
	(enums.SimulationModificationMethodEnum_SimulationModificationMethod)(0), // 2: google.ads.googleads.v10.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	(*common.CpcBidSimulationPointList)(nil),                                 // 3: google.ads.googleads.v10.common.CpcBidSimulationPointList
	(*common.TargetCpaSimulationPointList)(nil),                              // 4: google.ads.googleads.v10.common.TargetCpaSimulationPointList
	(*common.TargetRoasSimulationPointList)(nil),                             // 5: google.ads.googleads.v10.common.TargetRoasSimulationPointList
	(*common.TargetImpressionShareSimulationPointList)(nil),                  // 6: google.ads.googleads.v10.common.TargetImpressionShareSimulationPointList
	(*common.BudgetSimulationPointList)(nil),                                 // 7: google.ads.googleads.v10.common.BudgetSimulationPointList
}
var file_google_ads_googleads_v10_resources_campaign_simulation_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.CampaignSimulation.type:type_name -> google.ads.googleads.v10.enums.SimulationTypeEnum.SimulationType
	2, // 1: google.ads.googleads.v10.resources.CampaignSimulation.modification_method:type_name -> google.ads.googleads.v10.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	3, // 2: google.ads.googleads.v10.resources.CampaignSimulation.cpc_bid_point_list:type_name -> google.ads.googleads.v10.common.CpcBidSimulationPointList
	4, // 3: google.ads.googleads.v10.resources.CampaignSimulation.target_cpa_point_list:type_name -> google.ads.googleads.v10.common.TargetCpaSimulationPointList
	5, // 4: google.ads.googleads.v10.resources.CampaignSimulation.target_roas_point_list:type_name -> google.ads.googleads.v10.common.TargetRoasSimulationPointList
	6, // 5: google.ads.googleads.v10.resources.CampaignSimulation.target_impression_share_point_list:type_name -> google.ads.googleads.v10.common.TargetImpressionShareSimulationPointList
	7, // 6: google.ads.googleads.v10.resources.CampaignSimulation.budget_point_list:type_name -> google.ads.googleads.v10.common.BudgetSimulationPointList
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_campaign_simulation_proto_init() }
func file_google_ads_googleads_v10_resources_campaign_simulation_proto_init() {
	if File_google_ads_googleads_v10_resources_campaign_simulation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_campaign_simulation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignSimulation); i {
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
	file_google_ads_googleads_v10_resources_campaign_simulation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CampaignSimulation_CpcBidPointList)(nil),
		(*CampaignSimulation_TargetCpaPointList)(nil),
		(*CampaignSimulation_TargetRoasPointList)(nil),
		(*CampaignSimulation_TargetImpressionSharePointList)(nil),
		(*CampaignSimulation_BudgetPointList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_campaign_simulation_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_campaign_simulation_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_campaign_simulation_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_campaign_simulation_proto = out.File
	file_google_ads_googleads_v10_resources_campaign_simulation_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_campaign_simulation_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_campaign_simulation_proto_depIdxs = nil
}
