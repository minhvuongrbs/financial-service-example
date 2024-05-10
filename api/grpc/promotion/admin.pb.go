// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: promotion/admin.proto

package promotion

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
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

type SetVoucherToCampaignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CampaignId int64                                  `protobuf:"varint,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	Vouchers   []*SetVoucherToCampaignRequest_Voucher `protobuf:"bytes,2,rep,name=vouchers,proto3" json:"vouchers,omitempty"`
}

func (x *SetVoucherToCampaignRequest) Reset() {
	*x = SetVoucherToCampaignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetVoucherToCampaignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetVoucherToCampaignRequest) ProtoMessage() {}

func (x *SetVoucherToCampaignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetVoucherToCampaignRequest.ProtoReflect.Descriptor instead.
func (*SetVoucherToCampaignRequest) Descriptor() ([]byte, []int) {
	return file_promotion_admin_proto_rawDescGZIP(), []int{0}
}

func (x *SetVoucherToCampaignRequest) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *SetVoucherToCampaignRequest) GetVouchers() []*SetVoucherToCampaignRequest_Voucher {
	if x != nil {
		return x.Vouchers
	}
	return nil
}

type SetVoucherToCampaignReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SetVoucherToCampaignReply) Reset() {
	*x = SetVoucherToCampaignReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetVoucherToCampaignReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetVoucherToCampaignReply) ProtoMessage() {}

func (x *SetVoucherToCampaignReply) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetVoucherToCampaignReply.ProtoReflect.Descriptor instead.
func (*SetVoucherToCampaignReply) Descriptor() ([]byte, []int) {
	return file_promotion_admin_proto_rawDescGZIP(), []int{1}
}

func (x *SetVoucherToCampaignReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SetVoucherToCampaignReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DefineVoucherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherType    string `protobuf:"bytes,1,opt,name=voucher_type,json=voucherType,proto3" json:"voucher_type,omitempty"`
	Amount         int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	ExpirationTime string `protobuf:"bytes,3,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"` // rfc3339 datetime
	AppId          string `protobuf:"bytes,4,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`                            // * will be all app_id
}

func (x *DefineVoucherRequest) Reset() {
	*x = DefineVoucherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefineVoucherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefineVoucherRequest) ProtoMessage() {}

func (x *DefineVoucherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefineVoucherRequest.ProtoReflect.Descriptor instead.
func (*DefineVoucherRequest) Descriptor() ([]byte, []int) {
	return file_promotion_admin_proto_rawDescGZIP(), []int{2}
}

func (x *DefineVoucherRequest) GetVoucherType() string {
	if x != nil {
		return x.VoucherType
	}
	return ""
}

func (x *DefineVoucherRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DefineVoucherRequest) GetExpirationTime() string {
	if x != nil {
		return x.ExpirationTime
	}
	return ""
}

func (x *DefineVoucherRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type DefineVoucherReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefineVoucherReply) Reset() {
	*x = DefineVoucherReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefineVoucherReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefineVoucherReply) ProtoMessage() {}

func (x *DefineVoucherReply) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefineVoucherReply.ProtoReflect.Descriptor instead.
func (*DefineVoucherReply) Descriptor() ([]byte, []int) {
	return file_promotion_admin_proto_rawDescGZIP(), []int{3}
}

func (x *DefineVoucherReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DefineVoucherReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DefineCampaignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CampaignKey string `protobuf:"bytes,1,opt,name=campaign_key,json=campaignKey,proto3" json:"campaign_key,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DefineCampaignRequest) Reset() {
	*x = DefineCampaignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefineCampaignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefineCampaignRequest) ProtoMessage() {}

func (x *DefineCampaignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefineCampaignRequest.ProtoReflect.Descriptor instead.
func (*DefineCampaignRequest) Descriptor() ([]byte, []int) {
	return file_promotion_admin_proto_rawDescGZIP(), []int{4}
}

func (x *DefineCampaignRequest) GetCampaignKey() string {
	if x != nil {
		return x.CampaignKey
	}
	return ""
}

func (x *DefineCampaignRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DefineCampaignReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefineCampaignReply) Reset() {
	*x = DefineCampaignReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefineCampaignReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefineCampaignReply) ProtoMessage() {}

func (x *DefineCampaignReply) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefineCampaignReply.ProtoReflect.Descriptor instead.
func (*DefineCampaignReply) Descriptor() ([]byte, []int) {
	return file_promotion_admin_proto_rawDescGZIP(), []int{5}
}

func (x *DefineCampaignReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DefineCampaignReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SetVoucherToCampaignRequest_Voucher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherId int64 `protobuf:"varint,1,opt,name=voucher_id,json=voucherId,proto3" json:"voucher_id,omitempty"`
	Amount    int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SetVoucherToCampaignRequest_Voucher) Reset() {
	*x = SetVoucherToCampaignRequest_Voucher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promotion_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetVoucherToCampaignRequest_Voucher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetVoucherToCampaignRequest_Voucher) ProtoMessage() {}

func (x *SetVoucherToCampaignRequest_Voucher) ProtoReflect() protoreflect.Message {
	mi := &file_promotion_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetVoucherToCampaignRequest_Voucher.ProtoReflect.Descriptor instead.
func (*SetVoucherToCampaignRequest_Voucher) Descriptor() ([]byte, []int) {
	return file_promotion_admin_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SetVoucherToCampaignRequest_Voucher) GetVoucherId() int64 {
	if x != nil {
		return x.VoucherId
	}
	return 0
}

func (x *SetVoucherToCampaignRequest_Voucher) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_promotion_admin_proto protoreflect.FileDescriptor

var file_promotion_admin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01,
	0x0a, 0x1b, 0x53, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0x50,
	0x0a, 0x08, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x54, 0x6f, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x73,
	0x1a, 0x40, 0x0a, 0x07, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x49, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x54, 0x6f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x91, 0x01,
	0x0a, 0x14, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x22, 0x42, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4e, 0x0a, 0x15, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf0, 0x03, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x97, 0x01,
	0x0a, 0x0e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x37,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x3a, 0x01, 0x2a, 0x22, 0x2c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73,
	0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0x92, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a,
	0x22, 0x2a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12, 0xae, 0x01, 0x0a,
	0x14, 0x53, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68,
	0x65, 0x72, 0x54, 0x6f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x54, 0x6f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x3a, 0x01, 0x2a, 0x22, 0x31, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x73, 0x2f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x74, 0x42, 0xbf, 0x01,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x0a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x69, 0x6e, 0x68, 0x76, 0x75, 0x6f, 0x6e, 0x67, 0x72, 0x62, 0x73, 0x2f, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x50, 0x41, 0x58, 0xaa, 0x02, 0x0f, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0f, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02,
	0x1b, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x50,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_promotion_admin_proto_rawDescOnce sync.Once
	file_promotion_admin_proto_rawDescData = file_promotion_admin_proto_rawDesc
)

func file_promotion_admin_proto_rawDescGZIP() []byte {
	file_promotion_admin_proto_rawDescOnce.Do(func() {
		file_promotion_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_promotion_admin_proto_rawDescData)
	})
	return file_promotion_admin_proto_rawDescData
}

var file_promotion_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_promotion_admin_proto_goTypes = []interface{}{
	(*SetVoucherToCampaignRequest)(nil),         // 0: promotion.admin.SetVoucherToCampaignRequest
	(*SetVoucherToCampaignReply)(nil),           // 1: promotion.admin.SetVoucherToCampaignReply
	(*DefineVoucherRequest)(nil),                // 2: promotion.admin.DefineVoucherRequest
	(*DefineVoucherReply)(nil),                  // 3: promotion.admin.DefineVoucherReply
	(*DefineCampaignRequest)(nil),               // 4: promotion.admin.DefineCampaignRequest
	(*DefineCampaignReply)(nil),                 // 5: promotion.admin.DefineCampaignReply
	(*SetVoucherToCampaignRequest_Voucher)(nil), // 6: promotion.admin.SetVoucherToCampaignRequest.Voucher
}
var file_promotion_admin_proto_depIdxs = []int32{
	6, // 0: promotion.admin.SetVoucherToCampaignRequest.vouchers:type_name -> promotion.admin.SetVoucherToCampaignRequest.Voucher
	4, // 1: promotion.admin.PromotionAdmin.DefineCampaign:input_type -> promotion.admin.DefineCampaignRequest
	2, // 2: promotion.admin.PromotionAdmin.DefineVoucher:input_type -> promotion.admin.DefineVoucherRequest
	0, // 3: promotion.admin.PromotionAdmin.SetVoucherToCampaign:input_type -> promotion.admin.SetVoucherToCampaignRequest
	5, // 4: promotion.admin.PromotionAdmin.DefineCampaign:output_type -> promotion.admin.DefineCampaignReply
	3, // 5: promotion.admin.PromotionAdmin.DefineVoucher:output_type -> promotion.admin.DefineVoucherReply
	1, // 6: promotion.admin.PromotionAdmin.SetVoucherToCampaign:output_type -> promotion.admin.SetVoucherToCampaignReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_promotion_admin_proto_init() }
func file_promotion_admin_proto_init() {
	if File_promotion_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_promotion_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetVoucherToCampaignRequest); i {
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
		file_promotion_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetVoucherToCampaignReply); i {
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
		file_promotion_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefineVoucherRequest); i {
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
		file_promotion_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefineVoucherReply); i {
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
		file_promotion_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefineCampaignRequest); i {
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
		file_promotion_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefineCampaignReply); i {
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
		file_promotion_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetVoucherToCampaignRequest_Voucher); i {
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
			RawDescriptor: file_promotion_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_promotion_admin_proto_goTypes,
		DependencyIndexes: file_promotion_admin_proto_depIdxs,
		MessageInfos:      file_promotion_admin_proto_msgTypes,
	}.Build()
	File_promotion_admin_proto = out.File
	file_promotion_admin_proto_rawDesc = nil
	file_promotion_admin_proto_goTypes = nil
	file_promotion_admin_proto_depIdxs = nil
}
