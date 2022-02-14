// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/services/customer_user_access_invitation_service.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/nico-juni/googleads-go/v10/pb/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MutateCustomerUserAccessInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string                                 `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Operation  *CustomerUserAccessInvitationOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *MutateCustomerUserAccessInvitationRequest) Reset() {
	*x = MutateCustomerUserAccessInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessInvitationRequest) ProtoMessage() {}

func (x *MutateCustomerUserAccessInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessInvitationRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessInvitationRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateCustomerUserAccessInvitationRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerUserAccessInvitationRequest) GetOperation() *CustomerUserAccessInvitationOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type CustomerUserAccessInvitationOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Operation:
	//	*CustomerUserAccessInvitationOperation_Create
	//	*CustomerUserAccessInvitationOperation_Remove
	Operation isCustomerUserAccessInvitationOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CustomerUserAccessInvitationOperation) Reset() {
	*x = CustomerUserAccessInvitationOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerUserAccessInvitationOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerUserAccessInvitationOperation) ProtoMessage() {}

func (x *CustomerUserAccessInvitationOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerUserAccessInvitationOperation.ProtoReflect.Descriptor instead.
func (*CustomerUserAccessInvitationOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{1}
}

func (m *CustomerUserAccessInvitationOperation) GetOperation() isCustomerUserAccessInvitationOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CustomerUserAccessInvitationOperation) GetCreate() *resources.CustomerUserAccessInvitation {
	if x, ok := x.GetOperation().(*CustomerUserAccessInvitationOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *CustomerUserAccessInvitationOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*CustomerUserAccessInvitationOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isCustomerUserAccessInvitationOperation_Operation interface {
	isCustomerUserAccessInvitationOperation_Operation()
}

type CustomerUserAccessInvitationOperation_Create struct {
	Create *resources.CustomerUserAccessInvitation `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerUserAccessInvitationOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CustomerUserAccessInvitationOperation_Create) isCustomerUserAccessInvitationOperation_Operation() {
}

func (*CustomerUserAccessInvitationOperation_Remove) isCustomerUserAccessInvitationOperation_Operation() {
}

type MutateCustomerUserAccessInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *MutateCustomerUserAccessInvitationResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MutateCustomerUserAccessInvitationResponse) Reset() {
	*x = MutateCustomerUserAccessInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessInvitationResponse) ProtoMessage() {}

func (x *MutateCustomerUserAccessInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessInvitationResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessInvitationResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateCustomerUserAccessInvitationResponse) GetResult() *MutateCustomerUserAccessInvitationResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type MutateCustomerUserAccessInvitationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCustomerUserAccessInvitationResult) Reset() {
	*x = MutateCustomerUserAccessInvitationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerUserAccessInvitationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerUserAccessInvitationResult) ProtoMessage() {}

func (x *MutateCustomerUserAccessInvitationResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerUserAccessInvitationResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerUserAccessInvitationResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerUserAccessInvitationResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x29, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x6b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xe6, 0x01, 0x0a, 0x25, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x06,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xfa, 0x41, 0x37, 0x0a, 0x35, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x2a,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x8b, 0x01, 0x0a, 0x28, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5f, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x3a, 0xfa, 0x41, 0x37, 0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x98, 0x03,
	0x0a, 0x23, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa9, 0x02, 0x0a, 0x22, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x48, 0x22, 0x43, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a,
	0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x15, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41,
	0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x94, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x42, 0x28, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescData = file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDesc
)

func file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDescData
}

var file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_goTypes = []interface{}{
	(*MutateCustomerUserAccessInvitationRequest)(nil),  // 0: google.ads.googleads.v10.services.MutateCustomerUserAccessInvitationRequest
	(*CustomerUserAccessInvitationOperation)(nil),      // 1: google.ads.googleads.v10.services.CustomerUserAccessInvitationOperation
	(*MutateCustomerUserAccessInvitationResponse)(nil), // 2: google.ads.googleads.v10.services.MutateCustomerUserAccessInvitationResponse
	(*MutateCustomerUserAccessInvitationResult)(nil),   // 3: google.ads.googleads.v10.services.MutateCustomerUserAccessInvitationResult
	(*resources.CustomerUserAccessInvitation)(nil),     // 4: google.ads.googleads.v10.resources.CustomerUserAccessInvitation
}
var file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.services.MutateCustomerUserAccessInvitationRequest.operation:type_name -> google.ads.googleads.v10.services.CustomerUserAccessInvitationOperation
	4, // 1: google.ads.googleads.v10.services.CustomerUserAccessInvitationOperation.create:type_name -> google.ads.googleads.v10.resources.CustomerUserAccessInvitation
	3, // 2: google.ads.googleads.v10.services.MutateCustomerUserAccessInvitationResponse.result:type_name -> google.ads.googleads.v10.services.MutateCustomerUserAccessInvitationResult
	0, // 3: google.ads.googleads.v10.services.CustomerUserAccessInvitationService.MutateCustomerUserAccessInvitation:input_type -> google.ads.googleads.v10.services.MutateCustomerUserAccessInvitationRequest
	2, // 4: google.ads.googleads.v10.services.CustomerUserAccessInvitationService.MutateCustomerUserAccessInvitation:output_type -> google.ads.googleads.v10.services.MutateCustomerUserAccessInvitationResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_init()
}
func file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_init() {
	if File_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessInvitationRequest); i {
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
		file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerUserAccessInvitationOperation); i {
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
		file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessInvitationResponse); i {
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
		file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerUserAccessInvitationResult); i {
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
	file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CustomerUserAccessInvitationOperation_Create)(nil),
		(*CustomerUserAccessInvitationOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto = out.File
	file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_rawDesc = nil
	file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_goTypes = nil
	file_google_ads_googleads_v10_services_customer_user_access_invitation_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerUserAccessInvitationServiceClient is the client API for CustomerUserAccessInvitationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerUserAccessInvitationServiceClient interface {
	MutateCustomerUserAccessInvitation(ctx context.Context, in *MutateCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessInvitationResponse, error)
}

type customerUserAccessInvitationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerUserAccessInvitationServiceClient(cc grpc.ClientConnInterface) CustomerUserAccessInvitationServiceClient {
	return &customerUserAccessInvitationServiceClient{cc}
}

func (c *customerUserAccessInvitationServiceClient) MutateCustomerUserAccessInvitation(ctx context.Context, in *MutateCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessInvitationResponse, error) {
	out := new(MutateCustomerUserAccessInvitationResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v10.services.CustomerUserAccessInvitationService/MutateCustomerUserAccessInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerUserAccessInvitationServiceServer is the server API for CustomerUserAccessInvitationService service.
type CustomerUserAccessInvitationServiceServer interface {
	MutateCustomerUserAccessInvitation(context.Context, *MutateCustomerUserAccessInvitationRequest) (*MutateCustomerUserAccessInvitationResponse, error)
}

// UnimplementedCustomerUserAccessInvitationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerUserAccessInvitationServiceServer struct {
}

func (*UnimplementedCustomerUserAccessInvitationServiceServer) MutateCustomerUserAccessInvitation(context.Context, *MutateCustomerUserAccessInvitationRequest) (*MutateCustomerUserAccessInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerUserAccessInvitation not implemented")
}

func RegisterCustomerUserAccessInvitationServiceServer(s *grpc.Server, srv CustomerUserAccessInvitationServiceServer) {
	s.RegisterService(&_CustomerUserAccessInvitationService_serviceDesc, srv)
}

func _CustomerUserAccessInvitationService_MutateCustomerUserAccessInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerUserAccessInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserAccessInvitationServiceServer).MutateCustomerUserAccessInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v10.services.CustomerUserAccessInvitationService/MutateCustomerUserAccessInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserAccessInvitationServiceServer).MutateCustomerUserAccessInvitation(ctx, req.(*MutateCustomerUserAccessInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerUserAccessInvitationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v10.services.CustomerUserAccessInvitationService",
	HandlerType: (*CustomerUserAccessInvitationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerUserAccessInvitation",
			Handler:    _CustomerUserAccessInvitationService_MutateCustomerUserAccessInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v10/services/customer_user_access_invitation_service.proto",
}
