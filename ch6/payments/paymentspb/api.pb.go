// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: paymentspb/api.proto

package paymentspb

import (
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

type AuthorizePaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string  `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount     float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AuthorizePaymentRequest) Reset() {
	*x = AuthorizePaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizePaymentRequest) ProtoMessage() {}

func (x *AuthorizePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizePaymentRequest.ProtoReflect.Descriptor instead.
func (*AuthorizePaymentRequest) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizePaymentRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *AuthorizePaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AuthorizePaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorizePaymentResponse) Reset() {
	*x = AuthorizePaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizePaymentResponse) ProtoMessage() {}

func (x *AuthorizePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizePaymentResponse.ProtoReflect.Descriptor instead.
func (*AuthorizePaymentResponse) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizePaymentResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ConfirmPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConfirmPaymentRequest) Reset() {
	*x = ConfirmPaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmPaymentRequest) ProtoMessage() {}

func (x *ConfirmPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmPaymentRequest.ProtoReflect.Descriptor instead.
func (*ConfirmPaymentRequest) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmPaymentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ConfirmPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfirmPaymentResponse) Reset() {
	*x = ConfirmPaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmPaymentResponse) ProtoMessage() {}

func (x *ConfirmPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmPaymentResponse.ProtoReflect.Descriptor instead.
func (*ConfirmPaymentResponse) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{3}
}

type CreateInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string  `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PaymentId string  `protobuf:"bytes,2,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Amount    float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreateInvoiceRequest) Reset() {
	*x = CreateInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceRequest) ProtoMessage() {}

func (x *CreateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*CreateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateInvoiceRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateInvoiceRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *CreateInvoiceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateInvoiceResponse) Reset() {
	*x = CreateInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceResponse) ProtoMessage() {}

func (x *CreateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*CreateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{5}
}

func (x *CreateInvoiceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AdjustInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AdjustInvoiceRequest) Reset() {
	*x = AdjustInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustInvoiceRequest) ProtoMessage() {}

func (x *AdjustInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustInvoiceRequest.ProtoReflect.Descriptor instead.
func (*AdjustInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{6}
}

func (x *AdjustInvoiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdjustInvoiceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AdjustInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdjustInvoiceResponse) Reset() {
	*x = AdjustInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustInvoiceResponse) ProtoMessage() {}

func (x *AdjustInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustInvoiceResponse.ProtoReflect.Descriptor instead.
func (*AdjustInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{7}
}

type PayInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PayInvoiceRequest) Reset() {
	*x = PayInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayInvoiceRequest) ProtoMessage() {}

func (x *PayInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayInvoiceRequest.ProtoReflect.Descriptor instead.
func (*PayInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{8}
}

func (x *PayInvoiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PayInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PayInvoiceResponse) Reset() {
	*x = PayInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayInvoiceResponse) ProtoMessage() {}

func (x *PayInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayInvoiceResponse.ProtoReflect.Descriptor instead.
func (*PayInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{9}
}

type CancelInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelInvoiceRequest) Reset() {
	*x = CancelInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelInvoiceRequest) ProtoMessage() {}

func (x *CancelInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelInvoiceRequest.ProtoReflect.Descriptor instead.
func (*CancelInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{10}
}

func (x *CancelInvoiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelInvoiceResponse) Reset() {
	*x = CancelInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentspb_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelInvoiceResponse) ProtoMessage() {}

func (x *CancelInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentspb_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelInvoiceResponse.ProtoReflect.Descriptor instead.
func (*CancelInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_paymentspb_api_proto_rawDescGZIP(), []int{11}
}

var File_paymentspb_api_proto protoreflect.FileDescriptor

var file_paymentspb_api_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x70, 0x62, 0x1a, 0x19, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x17, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x2a, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a,
	0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x68, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x14, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11,
	0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x17, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa4, 0x04, 0x0a, 0x0f, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x10,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x23, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x21, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x0d, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62,
	0x2e, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0xa9, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x70, 0x62, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x75, 0x73, 0x2f, 0x65, 0x64, 0x61, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x63, 0x68, 0x36, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0xca, 0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x70, 0x62, 0xe2, 0x02, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_paymentspb_api_proto_rawDescOnce sync.Once
	file_paymentspb_api_proto_rawDescData = file_paymentspb_api_proto_rawDesc
)

func file_paymentspb_api_proto_rawDescGZIP() []byte {
	file_paymentspb_api_proto_rawDescOnce.Do(func() {
		file_paymentspb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_paymentspb_api_proto_rawDescData)
	})
	return file_paymentspb_api_proto_rawDescData
}

var file_paymentspb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_paymentspb_api_proto_goTypes = []interface{}{
	(*AuthorizePaymentRequest)(nil),  // 0: paymentspb.AuthorizePaymentRequest
	(*AuthorizePaymentResponse)(nil), // 1: paymentspb.AuthorizePaymentResponse
	(*ConfirmPaymentRequest)(nil),    // 2: paymentspb.ConfirmPaymentRequest
	(*ConfirmPaymentResponse)(nil),   // 3: paymentspb.ConfirmPaymentResponse
	(*CreateInvoiceRequest)(nil),     // 4: paymentspb.CreateInvoiceRequest
	(*CreateInvoiceResponse)(nil),    // 5: paymentspb.CreateInvoiceResponse
	(*AdjustInvoiceRequest)(nil),     // 6: paymentspb.AdjustInvoiceRequest
	(*AdjustInvoiceResponse)(nil),    // 7: paymentspb.AdjustInvoiceResponse
	(*PayInvoiceRequest)(nil),        // 8: paymentspb.PayInvoiceRequest
	(*PayInvoiceResponse)(nil),       // 9: paymentspb.PayInvoiceResponse
	(*CancelInvoiceRequest)(nil),     // 10: paymentspb.CancelInvoiceRequest
	(*CancelInvoiceResponse)(nil),    // 11: paymentspb.CancelInvoiceResponse
}
var file_paymentspb_api_proto_depIdxs = []int32{
	0,  // 0: paymentspb.PaymentsService.AuthorizePayment:input_type -> paymentspb.AuthorizePaymentRequest
	2,  // 1: paymentspb.PaymentsService.ConfirmPayment:input_type -> paymentspb.ConfirmPaymentRequest
	4,  // 2: paymentspb.PaymentsService.CreateInvoice:input_type -> paymentspb.CreateInvoiceRequest
	6,  // 3: paymentspb.PaymentsService.AdjustInvoice:input_type -> paymentspb.AdjustInvoiceRequest
	8,  // 4: paymentspb.PaymentsService.PayInvoice:input_type -> paymentspb.PayInvoiceRequest
	10, // 5: paymentspb.PaymentsService.CancelInvoice:input_type -> paymentspb.CancelInvoiceRequest
	1,  // 6: paymentspb.PaymentsService.AuthorizePayment:output_type -> paymentspb.AuthorizePaymentResponse
	3,  // 7: paymentspb.PaymentsService.ConfirmPayment:output_type -> paymentspb.ConfirmPaymentResponse
	5,  // 8: paymentspb.PaymentsService.CreateInvoice:output_type -> paymentspb.CreateInvoiceResponse
	7,  // 9: paymentspb.PaymentsService.AdjustInvoice:output_type -> paymentspb.AdjustInvoiceResponse
	9,  // 10: paymentspb.PaymentsService.PayInvoice:output_type -> paymentspb.PayInvoiceResponse
	11, // 11: paymentspb.PaymentsService.CancelInvoice:output_type -> paymentspb.CancelInvoiceResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_paymentspb_api_proto_init() }
func file_paymentspb_api_proto_init() {
	if File_paymentspb_api_proto != nil {
		return
	}
	file_paymentspb_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_paymentspb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizePaymentRequest); i {
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
		file_paymentspb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizePaymentResponse); i {
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
		file_paymentspb_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmPaymentRequest); i {
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
		file_paymentspb_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmPaymentResponse); i {
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
		file_paymentspb_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvoiceRequest); i {
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
		file_paymentspb_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvoiceResponse); i {
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
		file_paymentspb_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustInvoiceRequest); i {
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
		file_paymentspb_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustInvoiceResponse); i {
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
		file_paymentspb_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayInvoiceRequest); i {
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
		file_paymentspb_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayInvoiceResponse); i {
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
		file_paymentspb_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelInvoiceRequest); i {
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
		file_paymentspb_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelInvoiceResponse); i {
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
			RawDescriptor: file_paymentspb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paymentspb_api_proto_goTypes,
		DependencyIndexes: file_paymentspb_api_proto_depIdxs,
		MessageInfos:      file_paymentspb_api_proto_msgTypes,
	}.Build()
	File_paymentspb_api_proto = out.File
	file_paymentspb_api_proto_rawDesc = nil
	file_paymentspb_api_proto_goTypes = nil
	file_paymentspb_api_proto_depIdxs = nil
}
