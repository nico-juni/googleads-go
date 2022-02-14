// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: google/ads/googleads/v10/resources/account_budget_proposal.proto

package resources

import (
	proto "github.com/golang/protobuf/proto"
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

type AccountBudgetProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceName                string                                                            `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	Id                          *int64                                                            `protobuf:"varint,25,opt,name=id,proto3,oneof" json:"id,omitempty"`
	BillingSetup                *string                                                           `protobuf:"bytes,26,opt,name=billing_setup,json=billingSetup,proto3,oneof" json:"billing_setup,omitempty"`
	AccountBudget               *string                                                           `protobuf:"bytes,27,opt,name=account_budget,json=accountBudget,proto3,oneof" json:"account_budget,omitempty"`
	ProposalType                enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType     `protobuf:"varint,4,opt,name=proposal_type,json=proposalType,proto3,enum=google.ads.googleads.v10.enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType" json:"proposal_type,omitempty"`
	Status                      enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus `protobuf:"varint,15,opt,name=status,proto3,enum=google.ads.googleads.v10.enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus" json:"status,omitempty"`
	ProposedName                *string                                                           `protobuf:"bytes,28,opt,name=proposed_name,json=proposedName,proto3,oneof" json:"proposed_name,omitempty"`
	ApprovedStartDateTime       *string                                                           `protobuf:"bytes,30,opt,name=approved_start_date_time,json=approvedStartDateTime,proto3,oneof" json:"approved_start_date_time,omitempty"`
	ProposedPurchaseOrderNumber *string                                                           `protobuf:"bytes,35,opt,name=proposed_purchase_order_number,json=proposedPurchaseOrderNumber,proto3,oneof" json:"proposed_purchase_order_number,omitempty"`
	ProposedNotes               *string                                                           `protobuf:"bytes,36,opt,name=proposed_notes,json=proposedNotes,proto3,oneof" json:"proposed_notes,omitempty"`
	CreationDateTime            *string                                                           `protobuf:"bytes,37,opt,name=creation_date_time,json=creationDateTime,proto3,oneof" json:"creation_date_time,omitempty"`
	ApprovalDateTime            *string                                                           `protobuf:"bytes,38,opt,name=approval_date_time,json=approvalDateTime,proto3,oneof" json:"approval_date_time,omitempty"`
	// Types that are assignable to ProposedStartTime:
	//	*AccountBudgetProposal_ProposedStartDateTime
	//	*AccountBudgetProposal_ProposedStartTimeType
	ProposedStartTime isAccountBudgetProposal_ProposedStartTime `protobuf_oneof:"proposed_start_time"`
	// Types that are assignable to ProposedEndTime:
	//	*AccountBudgetProposal_ProposedEndDateTime
	//	*AccountBudgetProposal_ProposedEndTimeType
	ProposedEndTime isAccountBudgetProposal_ProposedEndTime `protobuf_oneof:"proposed_end_time"`
	// Types that are assignable to ApprovedEndTime:
	//	*AccountBudgetProposal_ApprovedEndDateTime
	//	*AccountBudgetProposal_ApprovedEndTimeType
	ApprovedEndTime isAccountBudgetProposal_ApprovedEndTime `protobuf_oneof:"approved_end_time"`
	// Types that are assignable to ProposedSpendingLimit:
	//	*AccountBudgetProposal_ProposedSpendingLimitMicros
	//	*AccountBudgetProposal_ProposedSpendingLimitType
	ProposedSpendingLimit isAccountBudgetProposal_ProposedSpendingLimit `protobuf_oneof:"proposed_spending_limit"`
	// Types that are assignable to ApprovedSpendingLimit:
	//	*AccountBudgetProposal_ApprovedSpendingLimitMicros
	//	*AccountBudgetProposal_ApprovedSpendingLimitType
	ApprovedSpendingLimit isAccountBudgetProposal_ApprovedSpendingLimit `protobuf_oneof:"approved_spending_limit"`
}

func (x *AccountBudgetProposal) Reset() {
	*x = AccountBudgetProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_account_budget_proposal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountBudgetProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountBudgetProposal) ProtoMessage() {}

func (x *AccountBudgetProposal) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_account_budget_proposal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountBudgetProposal.ProtoReflect.Descriptor instead.
func (*AccountBudgetProposal) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDescGZIP(), []int{0}
}

func (x *AccountBudgetProposal) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AccountBudgetProposal) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *AccountBudgetProposal) GetBillingSetup() string {
	if x != nil && x.BillingSetup != nil {
		return *x.BillingSetup
	}
	return ""
}

func (x *AccountBudgetProposal) GetAccountBudget() string {
	if x != nil && x.AccountBudget != nil {
		return *x.AccountBudget
	}
	return ""
}

func (x *AccountBudgetProposal) GetProposalType() enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType {
	if x != nil {
		return x.ProposalType
	}
	return enums.AccountBudgetProposalTypeEnum_UNSPECIFIED
}

func (x *AccountBudgetProposal) GetStatus() enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus {
	if x != nil {
		return x.Status
	}
	return enums.AccountBudgetProposalStatusEnum_UNSPECIFIED
}

func (x *AccountBudgetProposal) GetProposedName() string {
	if x != nil && x.ProposedName != nil {
		return *x.ProposedName
	}
	return ""
}

func (x *AccountBudgetProposal) GetApprovedStartDateTime() string {
	if x != nil && x.ApprovedStartDateTime != nil {
		return *x.ApprovedStartDateTime
	}
	return ""
}

func (x *AccountBudgetProposal) GetProposedPurchaseOrderNumber() string {
	if x != nil && x.ProposedPurchaseOrderNumber != nil {
		return *x.ProposedPurchaseOrderNumber
	}
	return ""
}

func (x *AccountBudgetProposal) GetProposedNotes() string {
	if x != nil && x.ProposedNotes != nil {
		return *x.ProposedNotes
	}
	return ""
}

func (x *AccountBudgetProposal) GetCreationDateTime() string {
	if x != nil && x.CreationDateTime != nil {
		return *x.CreationDateTime
	}
	return ""
}

func (x *AccountBudgetProposal) GetApprovalDateTime() string {
	if x != nil && x.ApprovalDateTime != nil {
		return *x.ApprovalDateTime
	}
	return ""
}

func (m *AccountBudgetProposal) GetProposedStartTime() isAccountBudgetProposal_ProposedStartTime {
	if m != nil {
		return m.ProposedStartTime
	}
	return nil
}

func (x *AccountBudgetProposal) GetProposedStartDateTime() string {
	if x, ok := x.GetProposedStartTime().(*AccountBudgetProposal_ProposedStartDateTime); ok {
		return x.ProposedStartDateTime
	}
	return ""
}

func (x *AccountBudgetProposal) GetProposedStartTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := x.GetProposedStartTime().(*AccountBudgetProposal_ProposedStartTimeType); ok {
		return x.ProposedStartTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

func (m *AccountBudgetProposal) GetProposedEndTime() isAccountBudgetProposal_ProposedEndTime {
	if m != nil {
		return m.ProposedEndTime
	}
	return nil
}

func (x *AccountBudgetProposal) GetProposedEndDateTime() string {
	if x, ok := x.GetProposedEndTime().(*AccountBudgetProposal_ProposedEndDateTime); ok {
		return x.ProposedEndDateTime
	}
	return ""
}

func (x *AccountBudgetProposal) GetProposedEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := x.GetProposedEndTime().(*AccountBudgetProposal_ProposedEndTimeType); ok {
		return x.ProposedEndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

func (m *AccountBudgetProposal) GetApprovedEndTime() isAccountBudgetProposal_ApprovedEndTime {
	if m != nil {
		return m.ApprovedEndTime
	}
	return nil
}

func (x *AccountBudgetProposal) GetApprovedEndDateTime() string {
	if x, ok := x.GetApprovedEndTime().(*AccountBudgetProposal_ApprovedEndDateTime); ok {
		return x.ApprovedEndDateTime
	}
	return ""
}

func (x *AccountBudgetProposal) GetApprovedEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := x.GetApprovedEndTime().(*AccountBudgetProposal_ApprovedEndTimeType); ok {
		return x.ApprovedEndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

func (m *AccountBudgetProposal) GetProposedSpendingLimit() isAccountBudgetProposal_ProposedSpendingLimit {
	if m != nil {
		return m.ProposedSpendingLimit
	}
	return nil
}

func (x *AccountBudgetProposal) GetProposedSpendingLimitMicros() int64 {
	if x, ok := x.GetProposedSpendingLimit().(*AccountBudgetProposal_ProposedSpendingLimitMicros); ok {
		return x.ProposedSpendingLimitMicros
	}
	return 0
}

func (x *AccountBudgetProposal) GetProposedSpendingLimitType() enums.SpendingLimitTypeEnum_SpendingLimitType {
	if x, ok := x.GetProposedSpendingLimit().(*AccountBudgetProposal_ProposedSpendingLimitType); ok {
		return x.ProposedSpendingLimitType
	}
	return enums.SpendingLimitTypeEnum_UNSPECIFIED
}

func (m *AccountBudgetProposal) GetApprovedSpendingLimit() isAccountBudgetProposal_ApprovedSpendingLimit {
	if m != nil {
		return m.ApprovedSpendingLimit
	}
	return nil
}

func (x *AccountBudgetProposal) GetApprovedSpendingLimitMicros() int64 {
	if x, ok := x.GetApprovedSpendingLimit().(*AccountBudgetProposal_ApprovedSpendingLimitMicros); ok {
		return x.ApprovedSpendingLimitMicros
	}
	return 0
}

func (x *AccountBudgetProposal) GetApprovedSpendingLimitType() enums.SpendingLimitTypeEnum_SpendingLimitType {
	if x, ok := x.GetApprovedSpendingLimit().(*AccountBudgetProposal_ApprovedSpendingLimitType); ok {
		return x.ApprovedSpendingLimitType
	}
	return enums.SpendingLimitTypeEnum_UNSPECIFIED
}

type isAccountBudgetProposal_ProposedStartTime interface {
	isAccountBudgetProposal_ProposedStartTime()
}

type AccountBudgetProposal_ProposedStartDateTime struct {
	ProposedStartDateTime string `protobuf:"bytes,29,opt,name=proposed_start_date_time,json=proposedStartDateTime,proto3,oneof"`
}

type AccountBudgetProposal_ProposedStartTimeType struct {
	ProposedStartTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,7,opt,name=proposed_start_time_type,json=proposedStartTimeType,proto3,enum=google.ads.googleads.v10.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudgetProposal_ProposedStartDateTime) isAccountBudgetProposal_ProposedStartTime() {}

func (*AccountBudgetProposal_ProposedStartTimeType) isAccountBudgetProposal_ProposedStartTime() {}

type isAccountBudgetProposal_ProposedEndTime interface {
	isAccountBudgetProposal_ProposedEndTime()
}

type AccountBudgetProposal_ProposedEndDateTime struct {
	ProposedEndDateTime string `protobuf:"bytes,31,opt,name=proposed_end_date_time,json=proposedEndDateTime,proto3,oneof"`
}

type AccountBudgetProposal_ProposedEndTimeType struct {
	ProposedEndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,9,opt,name=proposed_end_time_type,json=proposedEndTimeType,proto3,enum=google.ads.googleads.v10.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudgetProposal_ProposedEndDateTime) isAccountBudgetProposal_ProposedEndTime() {}

func (*AccountBudgetProposal_ProposedEndTimeType) isAccountBudgetProposal_ProposedEndTime() {}

type isAccountBudgetProposal_ApprovedEndTime interface {
	isAccountBudgetProposal_ApprovedEndTime()
}

type AccountBudgetProposal_ApprovedEndDateTime struct {
	ApprovedEndDateTime string `protobuf:"bytes,32,opt,name=approved_end_date_time,json=approvedEndDateTime,proto3,oneof"`
}

type AccountBudgetProposal_ApprovedEndTimeType struct {
	ApprovedEndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,22,opt,name=approved_end_time_type,json=approvedEndTimeType,proto3,enum=google.ads.googleads.v10.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*AccountBudgetProposal_ApprovedEndDateTime) isAccountBudgetProposal_ApprovedEndTime() {}

func (*AccountBudgetProposal_ApprovedEndTimeType) isAccountBudgetProposal_ApprovedEndTime() {}

type isAccountBudgetProposal_ProposedSpendingLimit interface {
	isAccountBudgetProposal_ProposedSpendingLimit()
}

type AccountBudgetProposal_ProposedSpendingLimitMicros struct {
	ProposedSpendingLimitMicros int64 `protobuf:"varint,33,opt,name=proposed_spending_limit_micros,json=proposedSpendingLimitMicros,proto3,oneof"`
}

type AccountBudgetProposal_ProposedSpendingLimitType struct {
	ProposedSpendingLimitType enums.SpendingLimitTypeEnum_SpendingLimitType `protobuf:"varint,11,opt,name=proposed_spending_limit_type,json=proposedSpendingLimitType,proto3,enum=google.ads.googleads.v10.enums.SpendingLimitTypeEnum_SpendingLimitType,oneof"`
}

func (*AccountBudgetProposal_ProposedSpendingLimitMicros) isAccountBudgetProposal_ProposedSpendingLimit() {
}

func (*AccountBudgetProposal_ProposedSpendingLimitType) isAccountBudgetProposal_ProposedSpendingLimit() {
}

type isAccountBudgetProposal_ApprovedSpendingLimit interface {
	isAccountBudgetProposal_ApprovedSpendingLimit()
}

type AccountBudgetProposal_ApprovedSpendingLimitMicros struct {
	ApprovedSpendingLimitMicros int64 `protobuf:"varint,34,opt,name=approved_spending_limit_micros,json=approvedSpendingLimitMicros,proto3,oneof"`
}

type AccountBudgetProposal_ApprovedSpendingLimitType struct {
	ApprovedSpendingLimitType enums.SpendingLimitTypeEnum_SpendingLimitType `protobuf:"varint,24,opt,name=approved_spending_limit_type,json=approvedSpendingLimitType,proto3,enum=google.ads.googleads.v10.enums.SpendingLimitTypeEnum_SpendingLimitType,oneof"`
}

func (*AccountBudgetProposal_ApprovedSpendingLimitMicros) isAccountBudgetProposal_ApprovedSpendingLimit() {
}

func (*AccountBudgetProposal_ApprovedSpendingLimitType) isAccountBudgetProposal_ApprovedSpendingLimit() {
}

var File_google_ads_googleads_v10_resources_account_budget_proposal_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xad, 0x12, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x5b, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x57, 0x0a, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x65, 0x74, 0x75, 0x70, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x48, 0x06, 0x52, 0x0c, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x48, 0x07, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x57, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x78, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x48, 0x08, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x18, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x09, 0x52, 0x15,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4d, 0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x0a, 0x52, 0x1b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x64, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x48, 0x0b, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x25,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x0c, 0x52, 0x10, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x36, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x48, 0x0d, 0x52, 0x10, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x18, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48,
	0x00, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x75, 0x0a, 0x18, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x3a, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x48, 0x01, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64,
	0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x71, 0x0a, 0x16, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x01, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a,
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x13, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x45,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x71, 0x0a, 0x16, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x13, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a,
	0x1e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18,
	0x21, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x03, 0x52, 0x1b, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x1c, 0x70, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x03,
	0x52, 0x19, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x1e, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x1b, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x1c, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x47,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x19,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x80, 0x01, 0xea, 0x41, 0x7d, 0x0a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12,
	0x4b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x2f,
	0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x15, 0x0a, 0x13,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x19, 0x0a,
	0x17, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x19, 0x0a, 0x17, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x21,
	0x0a, 0x1f, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x15, 0x0a, 0x13, 0x5f,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x42, 0x8c, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1a, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f,
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
	file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDescData = file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_account_budget_proposal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_resources_account_budget_proposal_proto_goTypes = []interface{}{
	(*AccountBudgetProposal)(nil),                                          // 0: google.ads.googleads.v10.resources.AccountBudgetProposal
	(enums.AccountBudgetProposalTypeEnum_AccountBudgetProposalType)(0),     // 1: google.ads.googleads.v10.enums.AccountBudgetProposalTypeEnum.AccountBudgetProposalType
	(enums.AccountBudgetProposalStatusEnum_AccountBudgetProposalStatus)(0), // 2: google.ads.googleads.v10.enums.AccountBudgetProposalStatusEnum.AccountBudgetProposalStatus
	(enums.TimeTypeEnum_TimeType)(0),                                       // 3: google.ads.googleads.v10.enums.TimeTypeEnum.TimeType
	(enums.SpendingLimitTypeEnum_SpendingLimitType)(0),                     // 4: google.ads.googleads.v10.enums.SpendingLimitTypeEnum.SpendingLimitType
}
var file_google_ads_googleads_v10_resources_account_budget_proposal_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.AccountBudgetProposal.proposal_type:type_name -> google.ads.googleads.v10.enums.AccountBudgetProposalTypeEnum.AccountBudgetProposalType
	2, // 1: google.ads.googleads.v10.resources.AccountBudgetProposal.status:type_name -> google.ads.googleads.v10.enums.AccountBudgetProposalStatusEnum.AccountBudgetProposalStatus
	3, // 2: google.ads.googleads.v10.resources.AccountBudgetProposal.proposed_start_time_type:type_name -> google.ads.googleads.v10.enums.TimeTypeEnum.TimeType
	3, // 3: google.ads.googleads.v10.resources.AccountBudgetProposal.proposed_end_time_type:type_name -> google.ads.googleads.v10.enums.TimeTypeEnum.TimeType
	3, // 4: google.ads.googleads.v10.resources.AccountBudgetProposal.approved_end_time_type:type_name -> google.ads.googleads.v10.enums.TimeTypeEnum.TimeType
	4, // 5: google.ads.googleads.v10.resources.AccountBudgetProposal.proposed_spending_limit_type:type_name -> google.ads.googleads.v10.enums.SpendingLimitTypeEnum.SpendingLimitType
	4, // 6: google.ads.googleads.v10.resources.AccountBudgetProposal.approved_spending_limit_type:type_name -> google.ads.googleads.v10.enums.SpendingLimitTypeEnum.SpendingLimitType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_account_budget_proposal_proto_init() }
func file_google_ads_googleads_v10_resources_account_budget_proposal_proto_init() {
	if File_google_ads_googleads_v10_resources_account_budget_proposal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_account_budget_proposal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountBudgetProposal); i {
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
	file_google_ads_googleads_v10_resources_account_budget_proposal_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AccountBudgetProposal_ProposedStartDateTime)(nil),
		(*AccountBudgetProposal_ProposedStartTimeType)(nil),
		(*AccountBudgetProposal_ProposedEndDateTime)(nil),
		(*AccountBudgetProposal_ProposedEndTimeType)(nil),
		(*AccountBudgetProposal_ApprovedEndDateTime)(nil),
		(*AccountBudgetProposal_ApprovedEndTimeType)(nil),
		(*AccountBudgetProposal_ProposedSpendingLimitMicros)(nil),
		(*AccountBudgetProposal_ProposedSpendingLimitType)(nil),
		(*AccountBudgetProposal_ApprovedSpendingLimitMicros)(nil),
		(*AccountBudgetProposal_ApprovedSpendingLimitType)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_account_budget_proposal_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_account_budget_proposal_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_account_budget_proposal_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_account_budget_proposal_proto = out.File
	file_google_ads_googleads_v10_resources_account_budget_proposal_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_account_budget_proposal_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_account_budget_proposal_proto_depIdxs = nil
}
