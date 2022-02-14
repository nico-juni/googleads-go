// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/services/keyword_plan_ad_group_service.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/nico-juni/googleads-go/v10/pb/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

type MutateKeywordPlanAdGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId     string                         `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Operations     []*KeywordPlanAdGroupOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	PartialFailure bool                           `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	ValidateOnly   bool                           `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateKeywordPlanAdGroupsRequest) Reset() {
	*x = MutateKeywordPlanAdGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateKeywordPlanAdGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateKeywordPlanAdGroupsRequest) ProtoMessage() {}

func (x *MutateKeywordPlanAdGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateKeywordPlanAdGroupsRequest.ProtoReflect.Descriptor instead.
func (*MutateKeywordPlanAdGroupsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateKeywordPlanAdGroupsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateKeywordPlanAdGroupsRequest) GetOperations() []*KeywordPlanAdGroupOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateKeywordPlanAdGroupsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateKeywordPlanAdGroupsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

type KeywordPlanAdGroupOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Types that are assignable to Operation:
	//	*KeywordPlanAdGroupOperation_Create
	//	*KeywordPlanAdGroupOperation_Update
	//	*KeywordPlanAdGroupOperation_Remove
	Operation isKeywordPlanAdGroupOperation_Operation `protobuf_oneof:"operation"`
}

func (x *KeywordPlanAdGroupOperation) Reset() {
	*x = KeywordPlanAdGroupOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanAdGroupOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanAdGroupOperation) ProtoMessage() {}

func (x *KeywordPlanAdGroupOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanAdGroupOperation.ProtoReflect.Descriptor instead.
func (*KeywordPlanAdGroupOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescGZIP(), []int{1}
}

func (x *KeywordPlanAdGroupOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *KeywordPlanAdGroupOperation) GetOperation() isKeywordPlanAdGroupOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *KeywordPlanAdGroupOperation) GetCreate() *resources.KeywordPlanAdGroup {
	if x, ok := x.GetOperation().(*KeywordPlanAdGroupOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *KeywordPlanAdGroupOperation) GetUpdate() *resources.KeywordPlanAdGroup {
	if x, ok := x.GetOperation().(*KeywordPlanAdGroupOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *KeywordPlanAdGroupOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*KeywordPlanAdGroupOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isKeywordPlanAdGroupOperation_Operation interface {
	isKeywordPlanAdGroupOperation_Operation()
}

type KeywordPlanAdGroupOperation_Create struct {
	Create *resources.KeywordPlanAdGroup `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type KeywordPlanAdGroupOperation_Update struct {
	Update *resources.KeywordPlanAdGroup `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type KeywordPlanAdGroupOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*KeywordPlanAdGroupOperation_Create) isKeywordPlanAdGroupOperation_Operation() {}

func (*KeywordPlanAdGroupOperation_Update) isKeywordPlanAdGroupOperation_Operation() {}

func (*KeywordPlanAdGroupOperation_Remove) isKeywordPlanAdGroupOperation_Operation() {}

type MutateKeywordPlanAdGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartialFailureError *status.Status                    `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	Results             []*MutateKeywordPlanAdGroupResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateKeywordPlanAdGroupsResponse) Reset() {
	*x = MutateKeywordPlanAdGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateKeywordPlanAdGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateKeywordPlanAdGroupsResponse) ProtoMessage() {}

func (x *MutateKeywordPlanAdGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateKeywordPlanAdGroupsResponse.ProtoReflect.Descriptor instead.
func (*MutateKeywordPlanAdGroupsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateKeywordPlanAdGroupsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateKeywordPlanAdGroupsResponse) GetResults() []*MutateKeywordPlanAdGroupResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type MutateKeywordPlanAdGroupResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateKeywordPlanAdGroupResult) Reset() {
	*x = MutateKeywordPlanAdGroupResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateKeywordPlanAdGroupResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateKeywordPlanAdGroupResult) ProtoMessage() {}

func (x *MutateKeywordPlanAdGroupResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateKeywordPlanAdGroupResult.ProtoReflect.Descriptor instead.
func (*MutateKeywordPlanAdGroupResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateKeywordPlanAdGroupResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x64, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x20, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x63, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xd7, 0x02, 0x0a, 0x1b, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x73, 0x6b, 0x12, 0x50, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x06, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52,
	0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c,
	0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xc8, 0x01, 0x0a, 0x21, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5b,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x77, 0x0a, 0x1e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x55, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x32, 0xea, 0x02, 0x0a, 0x19, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x85, 0x02, 0x0a, 0x19, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x3e, 0x22, 0x39, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x3d, 0x2a, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x42, 0x8a, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1e, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescData = file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDesc
)

func file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDescData
}

var file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_goTypes = []interface{}{
	(*MutateKeywordPlanAdGroupsRequest)(nil),  // 0: google.ads.googleads.v10.services.MutateKeywordPlanAdGroupsRequest
	(*KeywordPlanAdGroupOperation)(nil),       // 1: google.ads.googleads.v10.services.KeywordPlanAdGroupOperation
	(*MutateKeywordPlanAdGroupsResponse)(nil), // 2: google.ads.googleads.v10.services.MutateKeywordPlanAdGroupsResponse
	(*MutateKeywordPlanAdGroupResult)(nil),    // 3: google.ads.googleads.v10.services.MutateKeywordPlanAdGroupResult
	(*fieldmaskpb.FieldMask)(nil),             // 4: google.protobuf.FieldMask
	(*resources.KeywordPlanAdGroup)(nil),      // 5: google.ads.googleads.v10.resources.KeywordPlanAdGroup
	(*status.Status)(nil),                     // 6: google.rpc.Status
}
var file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.services.MutateKeywordPlanAdGroupsRequest.operations:type_name -> google.ads.googleads.v10.services.KeywordPlanAdGroupOperation
	4, // 1: google.ads.googleads.v10.services.KeywordPlanAdGroupOperation.update_mask:type_name -> google.protobuf.FieldMask
	5, // 2: google.ads.googleads.v10.services.KeywordPlanAdGroupOperation.create:type_name -> google.ads.googleads.v10.resources.KeywordPlanAdGroup
	5, // 3: google.ads.googleads.v10.services.KeywordPlanAdGroupOperation.update:type_name -> google.ads.googleads.v10.resources.KeywordPlanAdGroup
	6, // 4: google.ads.googleads.v10.services.MutateKeywordPlanAdGroupsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 5: google.ads.googleads.v10.services.MutateKeywordPlanAdGroupsResponse.results:type_name -> google.ads.googleads.v10.services.MutateKeywordPlanAdGroupResult
	0, // 6: google.ads.googleads.v10.services.KeywordPlanAdGroupService.MutateKeywordPlanAdGroups:input_type -> google.ads.googleads.v10.services.MutateKeywordPlanAdGroupsRequest
	2, // 7: google.ads.googleads.v10.services.KeywordPlanAdGroupService.MutateKeywordPlanAdGroups:output_type -> google.ads.googleads.v10.services.MutateKeywordPlanAdGroupsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_init() }
func file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_init() {
	if File_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateKeywordPlanAdGroupsRequest); i {
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
		file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPlanAdGroupOperation); i {
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
		file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateKeywordPlanAdGroupsResponse); i {
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
		file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateKeywordPlanAdGroupResult); i {
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
	file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*KeywordPlanAdGroupOperation_Create)(nil),
		(*KeywordPlanAdGroupOperation_Update)(nil),
		(*KeywordPlanAdGroupOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto = out.File
	file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_rawDesc = nil
	file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_goTypes = nil
	file_google_ads_googleads_v10_services_keyword_plan_ad_group_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordPlanAdGroupServiceClient is the client API for KeywordPlanAdGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanAdGroupServiceClient interface {
	MutateKeywordPlanAdGroups(ctx context.Context, in *MutateKeywordPlanAdGroupsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupsResponse, error)
}

type keywordPlanAdGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanAdGroupServiceClient(cc grpc.ClientConnInterface) KeywordPlanAdGroupServiceClient {
	return &keywordPlanAdGroupServiceClient{cc}
}

func (c *keywordPlanAdGroupServiceClient) MutateKeywordPlanAdGroups(ctx context.Context, in *MutateKeywordPlanAdGroupsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupsResponse, error) {
	out := new(MutateKeywordPlanAdGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.KeywordPlanAdGroupService/MutateKeywordPlanAdGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanAdGroupServiceServer is the server API for KeywordPlanAdGroupService service.
type KeywordPlanAdGroupServiceServer interface {
	MutateKeywordPlanAdGroups(context.Context, *MutateKeywordPlanAdGroupsRequest) (*MutateKeywordPlanAdGroupsResponse, error)
}

// UnimplementedKeywordPlanAdGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanAdGroupServiceServer struct {
}

func (*UnimplementedKeywordPlanAdGroupServiceServer) MutateKeywordPlanAdGroups(context.Context, *MutateKeywordPlanAdGroupsRequest) (*MutateKeywordPlanAdGroupsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateKeywordPlanAdGroups not implemented")
}

func RegisterKeywordPlanAdGroupServiceServer(s *grpc.Server, srv KeywordPlanAdGroupServiceServer) {
	s.RegisterService(&_KeywordPlanAdGroupService_serviceDesc, srv)
}

func _KeywordPlanAdGroupService_MutateKeywordPlanAdGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanAdGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanAdGroupServiceServer).MutateKeywordPlanAdGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.KeywordPlanAdGroupService/MutateKeywordPlanAdGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanAdGroupServiceServer).MutateKeywordPlanAdGroups(ctx, req.(*MutateKeywordPlanAdGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanAdGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v10.services.KeywordPlanAdGroupService",
	HandlerType: (*KeywordPlanAdGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateKeywordPlanAdGroups",
			Handler:    _KeywordPlanAdGroupService_MutateKeywordPlanAdGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v10/services/keyword_plan_ad_group_service.proto",
}
