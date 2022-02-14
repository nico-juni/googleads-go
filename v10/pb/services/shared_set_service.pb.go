// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/services/shared_set_service.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	enums "github.com/nico-juni/googleads-go/v10/pb/enums"
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

type MutateSharedSetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId          string                                            `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Operations          []*SharedSetOperation                             `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	PartialFailure      bool                                              `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	ValidateOnly        bool                                              `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v10.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateSharedSetsRequest) Reset() {
	*x = MutateSharedSetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedSetsRequest) ProtoMessage() {}

func (x *MutateSharedSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedSetsRequest.ProtoReflect.Descriptor instead.
func (*MutateSharedSetsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateSharedSetsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateSharedSetsRequest) GetOperations() []*SharedSetOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateSharedSetsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateSharedSetsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateSharedSetsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_UNSPECIFIED
}

type SharedSetOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Types that are assignable to Operation:
	//	*SharedSetOperation_Create
	//	*SharedSetOperation_Update
	//	*SharedSetOperation_Remove
	Operation isSharedSetOperation_Operation `protobuf_oneof:"operation"`
}

func (x *SharedSetOperation) Reset() {
	*x = SharedSetOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedSetOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedSetOperation) ProtoMessage() {}

func (x *SharedSetOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedSetOperation.ProtoReflect.Descriptor instead.
func (*SharedSetOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescGZIP(), []int{1}
}

func (x *SharedSetOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *SharedSetOperation) GetOperation() isSharedSetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *SharedSetOperation) GetCreate() *resources.SharedSet {
	if x, ok := x.GetOperation().(*SharedSetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *SharedSetOperation) GetUpdate() *resources.SharedSet {
	if x, ok := x.GetOperation().(*SharedSetOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *SharedSetOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*SharedSetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isSharedSetOperation_Operation interface {
	isSharedSetOperation_Operation()
}

type SharedSetOperation_Create struct {
	Create *resources.SharedSet `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type SharedSetOperation_Update struct {
	Update *resources.SharedSet `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type SharedSetOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*SharedSetOperation_Create) isSharedSetOperation_Operation() {}

func (*SharedSetOperation_Update) isSharedSetOperation_Operation() {}

func (*SharedSetOperation_Remove) isSharedSetOperation_Operation() {}

type MutateSharedSetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartialFailureError *status.Status           `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	Results             []*MutateSharedSetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateSharedSetsResponse) Reset() {
	*x = MutateSharedSetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedSetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedSetsResponse) ProtoMessage() {}

func (x *MutateSharedSetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedSetsResponse.ProtoReflect.Descriptor instead.
func (*MutateSharedSetsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateSharedSetsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateSharedSetsResponse) GetResults() []*MutateSharedSetResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type MutateSharedSetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName string               `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	SharedSet    *resources.SharedSet `protobuf:"bytes,2,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
}

func (x *MutateSharedSetResult) Reset() {
	*x = MutateSharedSetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedSetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedSetResult) ProtoMessage() {}

func (x *MutateSharedSetResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedSetResult.ProtoReflect.Descriptor instead.
func (*MutateSharedSetResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateSharedSetResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateSharedSetResult) GetSharedSet() *resources.SharedSet {
	if x != nil {
		return x.SharedSet
	}
	return nil
}

var File_google_ads_googleads_v10_services_shared_set_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_services_shared_set_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea,
	0x02, 0x0a, 0x17, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x5a, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x7f, 0x0a, 0x15, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb3, 0x02, 0x0a, 0x12,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12,
	0x47, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x27, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xb6, 0x01, 0x0a, 0x18, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x52, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x15, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x41, 0x24,
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74,
	0x32, 0xbd, 0x02, 0x0a, 0x10, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe1, 0x01, 0x0a, 0x10, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x12, 0x3a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x54, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x22, 0x30, 0x2f, 0x76, 0x31,
	0x30, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x42, 0x81, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x15, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x30, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescData = file_google_ads_googleads_v10_services_shared_set_service_proto_rawDesc
)

func file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_services_shared_set_service_proto_rawDescData
}

var file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v10_services_shared_set_service_proto_goTypes = []interface{}{
	(*MutateSharedSetsRequest)(nil),                        // 0: google.ads.googleads.v10.services.MutateSharedSetsRequest
	(*SharedSetOperation)(nil),                             // 1: google.ads.googleads.v10.services.SharedSetOperation
	(*MutateSharedSetsResponse)(nil),                       // 2: google.ads.googleads.v10.services.MutateSharedSetsResponse
	(*MutateSharedSetResult)(nil),                          // 3: google.ads.googleads.v10.services.MutateSharedSetResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v10.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.SharedSet)(nil),                            // 6: google.ads.googleads.v10.resources.SharedSet
	(*status.Status)(nil),                                  // 7: google.rpc.Status
}
var file_google_ads_googleads_v10_services_shared_set_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.services.MutateSharedSetsRequest.operations:type_name -> google.ads.googleads.v10.services.SharedSetOperation
	4, // 1: google.ads.googleads.v10.services.MutateSharedSetsRequest.response_content_type:type_name -> google.ads.googleads.v10.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v10.services.SharedSetOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v10.services.SharedSetOperation.create:type_name -> google.ads.googleads.v10.resources.SharedSet
	6, // 4: google.ads.googleads.v10.services.SharedSetOperation.update:type_name -> google.ads.googleads.v10.resources.SharedSet
	7, // 5: google.ads.googleads.v10.services.MutateSharedSetsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 6: google.ads.googleads.v10.services.MutateSharedSetsResponse.results:type_name -> google.ads.googleads.v10.services.MutateSharedSetResult
	6, // 7: google.ads.googleads.v10.services.MutateSharedSetResult.shared_set:type_name -> google.ads.googleads.v10.resources.SharedSet
	0, // 8: google.ads.googleads.v10.services.SharedSetService.MutateSharedSets:input_type -> google.ads.googleads.v10.services.MutateSharedSetsRequest
	2, // 9: google.ads.googleads.v10.services.SharedSetService.MutateSharedSets:output_type -> google.ads.googleads.v10.services.MutateSharedSetsResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_services_shared_set_service_proto_init() }
func file_google_ads_googleads_v10_services_shared_set_service_proto_init() {
	if File_google_ads_googleads_v10_services_shared_set_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedSetsRequest); i {
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
		file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedSetOperation); i {
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
		file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedSetsResponse); i {
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
		file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedSetResult); i {
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
	file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SharedSetOperation_Create)(nil),
		(*SharedSetOperation_Update)(nil),
		(*SharedSetOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_services_shared_set_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v10_services_shared_set_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_services_shared_set_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_services_shared_set_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_services_shared_set_service_proto = out.File
	file_google_ads_googleads_v10_services_shared_set_service_proto_rawDesc = nil
	file_google_ads_googleads_v10_services_shared_set_service_proto_goTypes = nil
	file_google_ads_googleads_v10_services_shared_set_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SharedSetServiceClient is the client API for SharedSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SharedSetServiceClient interface {
	MutateSharedSets(ctx context.Context, in *MutateSharedSetsRequest, opts ...grpc.CallOption) (*MutateSharedSetsResponse, error)
}

type sharedSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedSetServiceClient(cc grpc.ClientConnInterface) SharedSetServiceClient {
	return &sharedSetServiceClient{cc}
}

func (c *sharedSetServiceClient) MutateSharedSets(ctx context.Context, in *MutateSharedSetsRequest, opts ...grpc.CallOption) (*MutateSharedSetsResponse, error) {
	out := new(MutateSharedSetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.SharedSetService/MutateSharedSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedSetServiceServer is the server API for SharedSetService service.
type SharedSetServiceServer interface {
	MutateSharedSets(context.Context, *MutateSharedSetsRequest) (*MutateSharedSetsResponse, error)
}

// UnimplementedSharedSetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSharedSetServiceServer struct {
}

func (*UnimplementedSharedSetServiceServer) MutateSharedSets(context.Context, *MutateSharedSetsRequest) (*MutateSharedSetsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateSharedSets not implemented")
}

func RegisterSharedSetServiceServer(s *grpc.Server, srv SharedSetServiceServer) {
	s.RegisterService(&_SharedSetService_serviceDesc, srv)
}

func _SharedSetService_MutateSharedSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateSharedSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedSetServiceServer).MutateSharedSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.SharedSetService/MutateSharedSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedSetServiceServer).MutateSharedSets(ctx, req.(*MutateSharedSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SharedSetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v10.services.SharedSetService",
	HandlerType: (*SharedSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateSharedSets",
			Handler:    _SharedSetService_MutateSharedSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v10/services/shared_set_service.proto",
}
