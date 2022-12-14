// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: proto/payment.proto

package proto

import (
	context "context"
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

type PaymentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerID        string  `protobuf:"bytes,1,opt,name=customerID,proto3" json:"customerID,omitempty"`
	TransactionAmount float64 `protobuf:"fixed64,2,opt,name=transactionAmount,proto3" json:"transactionAmount,omitempty"`
	Start             string  `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	Destination       string  `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Number            int32   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	RouteDetailID     int32   `protobuf:"varint,6,opt,name=routeDetailID,proto3" json:"routeDetailID,omitempty"`
	Date              string  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	Type              string  `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PaymentInfo) Reset() {
	*x = PaymentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentInfo) ProtoMessage() {}

func (x *PaymentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentInfo.ProtoReflect.Descriptor instead.
func (*PaymentInfo) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentInfo) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *PaymentInfo) GetTransactionAmount() float64 {
	if x != nil {
		return x.TransactionAmount
	}
	return 0
}

func (x *PaymentInfo) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *PaymentInfo) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *PaymentInfo) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *PaymentInfo) GetRouteDetailID() int32 {
	if x != nil {
		return x.RouteDetailID
	}
	return 0
}

func (x *PaymentInfo) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *PaymentInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ResPayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *ResPayInfo) Reset() {
	*x = ResPayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResPayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResPayInfo) ProtoMessage() {}

func (x *ResPayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResPayInfo.ProtoReflect.Descriptor instead.
func (*ResPayInfo) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{1}
}

func (x *ResPayInfo) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ResPayInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type QueryOrderUserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *QueryOrderUserID) Reset() {
	*x = QueryOrderUserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOrderUserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOrderUserID) ProtoMessage() {}

func (x *QueryOrderUserID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOrderUserID.ProtoReflect.Descriptor instead.
func (*QueryOrderUserID) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{2}
}

func (x *QueryOrderUserID) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type OrderInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderInfo `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *OrderInfoList) Reset() {
	*x = OrderInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoList) ProtoMessage() {}

func (x *OrderInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoList.ProtoReflect.Descriptor instead.
func (*OrderInfoList) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{3}
}

func (x *OrderInfoList) GetOrders() []*OrderInfo {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID       int32   `protobuf:"varint,1,opt,name=orderID,proto3" json:"orderID,omitempty"`
	Amount        float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Start         string  `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	Destination   string  `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Number        int32   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	RouteDetailID int32   `protobuf:"varint,6,opt,name=routeDetailID,proto3" json:"routeDetailID,omitempty"`
	Date          string  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	Type          string  `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Time          string  `protobuf:"bytes,9,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *OrderInfo) Reset() {
	*x = OrderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfo) ProtoMessage() {}

func (x *OrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfo.ProtoReflect.Descriptor instead.
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return file_proto_payment_proto_rawDescGZIP(), []int{4}
}

func (x *OrderInfo) GetOrderID() int32 {
	if x != nil {
		return x.OrderID
	}
	return 0
}

func (x *OrderInfo) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OrderInfo) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *OrderInfo) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *OrderInfo) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *OrderInfo) GetRouteDetailID() int32 {
	if x != nil {
		return x.RouteDetailID
	}
	return 0
}

func (x *OrderInfo) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *OrderInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *OrderInfo) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

var File_proto_payment_proto protoreflect.FileDescriptor

var file_proto_payment_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a,
	0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3c, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x50,
	0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x39, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0xef, 0x01,
	0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32,
	0x7f, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x30, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x61, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_payment_proto_rawDescOnce sync.Once
	file_proto_payment_proto_rawDescData = file_proto_payment_proto_rawDesc
)

func file_proto_payment_proto_rawDescGZIP() []byte {
	file_proto_payment_proto_rawDescOnce.Do(func() {
		file_proto_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_payment_proto_rawDescData)
	})
	return file_proto_payment_proto_rawDescData
}

var file_proto_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_payment_proto_goTypes = []interface{}{
	(*PaymentInfo)(nil),      // 0: proto.PaymentInfo
	(*ResPayInfo)(nil),       // 1: proto.ResPayInfo
	(*QueryOrderUserID)(nil), // 2: proto.QueryOrderUserID
	(*OrderInfoList)(nil),    // 3: proto.OrderInfoList
	(*OrderInfo)(nil),        // 4: proto.OrderInfo
}
var file_proto_payment_proto_depIdxs = []int32{
	4, // 0: proto.OrderInfoList.orders:type_name -> proto.OrderInfo
	0, // 1: proto.PaymentService.Payment:input_type -> proto.PaymentInfo
	2, // 2: proto.PaymentService.QueryOrder:input_type -> proto.QueryOrderUserID
	1, // 3: proto.PaymentService.Payment:output_type -> proto.ResPayInfo
	3, // 4: proto.PaymentService.QueryOrder:output_type -> proto.OrderInfoList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_payment_proto_init() }
func file_proto_payment_proto_init() {
	if File_proto_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentInfo); i {
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
		file_proto_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResPayInfo); i {
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
		file_proto_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOrderUserID); i {
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
		file_proto_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoList); i {
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
		file_proto_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfo); i {
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
			RawDescriptor: file_proto_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_payment_proto_goTypes,
		DependencyIndexes: file_proto_payment_proto_depIdxs,
		MessageInfos:      file_proto_payment_proto_msgTypes,
	}.Build()
	File_proto_payment_proto = out.File
	file_proto_payment_proto_rawDesc = nil
	file_proto_payment_proto_goTypes = nil
	file_proto_payment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentServiceClient interface {
	Payment(ctx context.Context, in *PaymentInfo, opts ...grpc.CallOption) (*ResPayInfo, error)
	QueryOrder(ctx context.Context, in *QueryOrderUserID, opts ...grpc.CallOption) (*OrderInfoList, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) Payment(ctx context.Context, in *PaymentInfo, opts ...grpc.CallOption) (*ResPayInfo, error) {
	out := new(ResPayInfo)
	err := c.cc.Invoke(ctx, "/proto.PaymentService/Payment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) QueryOrder(ctx context.Context, in *QueryOrderUserID, opts ...grpc.CallOption) (*OrderInfoList, error) {
	out := new(OrderInfoList)
	err := c.cc.Invoke(ctx, "/proto.PaymentService/QueryOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
type PaymentServiceServer interface {
	Payment(context.Context, *PaymentInfo) (*ResPayInfo, error)
	QueryOrder(context.Context, *QueryOrderUserID) (*OrderInfoList, error)
}

// UnimplementedPaymentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (*UnimplementedPaymentServiceServer) Payment(context.Context, *PaymentInfo) (*ResPayInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payment not implemented")
}
func (*UnimplementedPaymentServiceServer) QueryOrder(context.Context, *QueryOrderUserID) (*OrderInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrder not implemented")
}

func RegisterPaymentServiceServer(s *grpc.Server, srv PaymentServiceServer) {
	s.RegisterService(&_PaymentService_serviceDesc, srv)
}

func _PaymentService_Payment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Payment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PaymentService/Payment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Payment(ctx, req.(*PaymentInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_QueryOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrderUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).QueryOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PaymentService/QueryOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).QueryOrder(ctx, req.(*QueryOrderUserID))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Payment",
			Handler:    _PaymentService_Payment_Handler,
		},
		{
			MethodName: "QueryOrder",
			Handler:    _PaymentService_QueryOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/payment.proto",
}
