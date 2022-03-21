// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: position.proto

package positionpb

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

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	AccountID          string  `protobuf:"bytes,2,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	OrderID            string  `protobuf:"bytes,3,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	OpenPrice          float64 `protobuf:"fixed64,4,opt,name=OpenPrice,proto3" json:"OpenPrice,omitempty"`
	ClosePrice         float64 `protobuf:"fixed64,5,opt,name=ClosePrice,proto3" json:"ClosePrice,omitempty"`
	TakeProfit         float64 `protobuf:"fixed64,6,opt,name=TakeProfit,proto3" json:"TakeProfit,omitempty"`
	StopLoss           float64 `protobuf:"fixed64,7,opt,name=StopLoss,proto3" json:"StopLoss,omitempty"`
	Symbol             string  `protobuf:"bytes,8,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	GuaranteedStopLoss bool    `protobuf:"varint,9,opt,name=GuaranteedStopLoss,proto3" json:"GuaranteedStopLoss,omitempty"`
	State              string  `protobuf:"bytes,10,opt,name=State,proto3" json:"State,omitempty"`
	Quantity           float64 `protobuf:"fixed64,11,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Leverage           bool    `protobuf:"varint,12,opt,name=Leverage,proto3" json:"Leverage,omitempty"`
	Side               string  `protobuf:"bytes,13,opt,name=Side,proto3" json:"Side,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Position) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Position) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *Position) GetOpenPrice() float64 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *Position) GetClosePrice() float64 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *Position) GetTakeProfit() float64 {
	if x != nil {
		return x.TakeProfit
	}
	return 0
}

func (x *Position) GetStopLoss() float64 {
	if x != nil {
		return x.StopLoss
	}
	return 0
}

func (x *Position) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Position) GetGuaranteedStopLoss() bool {
	if x != nil {
		return x.GuaranteedStopLoss
	}
	return false
}

func (x *Position) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Position) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Position) GetLeverage() bool {
	if x != nil {
		return x.Leverage
	}
	return false
}

func (x *Position) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

type PositionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PositionID) Reset() {
	*x = PositionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionID) ProtoMessage() {}

func (x *PositionID) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionID.ProtoReflect.Descriptor instead.
func (*PositionID) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{1}
}

func (x *PositionID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ClosePrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ClosePrice) Reset() {
	*x = ClosePrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClosePrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosePrice) ProtoMessage() {}

func (x *ClosePrice) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosePrice.ProtoReflect.Descriptor instead.
func (*ClosePrice) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{2}
}

func (x *ClosePrice) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{3}
}

type AccountID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AccountID) Reset() {
	*x = AccountID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountID) ProtoMessage() {}

func (x *AccountID) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountID.ProtoReflect.Descriptor instead.
func (*AccountID) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{4}
}

func (x *AccountID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Positions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*Position `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *Positions) Reset() {
	*x = Positions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Positions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Positions) ProtoMessage() {}

func (x *Positions) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Positions.ProtoReflect.Descriptor instead.
func (*Positions) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{5}
}

func (x *Positions) GetValue() []*Position {
	if x != nil {
		return x.Value
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Ask    float64 `protobuf:"fixed64,2,opt,name=ask,proto3" json:"ask,omitempty"`
	Bid    float64 `protobuf:"fixed64,3,opt,name=bid,proto3" json:"bid,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{6}
}

func (x *Price) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Price) GetAsk() float64 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *Price) GetBid() float64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

type ProfitLoss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionID string  `protobuf:"bytes,1,opt,name=positionID,proto3" json:"positionID,omitempty"`
	Value      float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProfitLoss) Reset() {
	*x = ProfitLoss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_position_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfitLoss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfitLoss) ProtoMessage() {}

func (x *ProfitLoss) ProtoReflect() protoreflect.Message {
	mi := &file_position_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfitLoss.ProtoReflect.Descriptor instead.
func (*ProfitLoss) Descriptor() ([]byte, []int) {
	return file_position_proto_rawDescGZIP(), []int{7}
}

func (x *ProfitLoss) GetPositionID() string {
	if x != nil {
		return x.PositionID
	}
	return ""
}

func (x *ProfitLoss) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_position_proto protoreflect.FileDescriptor

var file_position_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf6, 0x02, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x54, 0x61, 0x6b, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x2e, 0x0a, 0x12, 0x47, 0x75, 0x61, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x12, 0x47, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x53,
	0x74, 0x6f, 0x70, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x65, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x4c, 0x65, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x64, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x69, 0x64, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x22, 0x0a,
	0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x0a, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3b, 0x0a,
	0x09, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x43, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x73, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x62, 0x69, 0x64, 0x22,
	0x42, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0xdd, 0x02, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x15, 0x2e, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x18, 0x2e,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x4c, 0x6f, 0x73, 0x73, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_position_proto_rawDescOnce sync.Once
	file_position_proto_rawDescData = file_position_proto_rawDesc
)

func file_position_proto_rawDescGZIP() []byte {
	file_position_proto_rawDescOnce.Do(func() {
		file_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_position_proto_rawDescData)
	})
	return file_position_proto_rawDescData
}

var file_position_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_position_proto_goTypes = []interface{}{
	(*Position)(nil),   // 0: position_proto.Position
	(*PositionID)(nil), // 1: position_proto.PositionID
	(*ClosePrice)(nil), // 2: position_proto.ClosePrice
	(*Empty)(nil),      // 3: position_proto.Empty
	(*AccountID)(nil),  // 4: position_proto.AccountID
	(*Positions)(nil),  // 5: position_proto.Positions
	(*Price)(nil),      // 6: position_proto.Price
	(*ProfitLoss)(nil), // 7: position_proto.ProfitLoss
}
var file_position_proto_depIdxs = []int32{
	0, // 0: position_proto.Positions.value:type_name -> position_proto.Position
	1, // 1: position_proto.PositionService.Close:input_type -> position_proto.PositionID
	0, // 2: position_proto.PositionService.Open:input_type -> position_proto.Position
	4, // 3: position_proto.PositionService.GetOpen:input_type -> position_proto.AccountID
	6, // 4: position_proto.PositionService.UpdatePrices:input_type -> position_proto.Price
	4, // 5: position_proto.PositionService.GetProfitLoss:input_type -> position_proto.AccountID
	3, // 6: position_proto.PositionService.Close:output_type -> position_proto.Empty
	3, // 7: position_proto.PositionService.Open:output_type -> position_proto.Empty
	5, // 8: position_proto.PositionService.GetOpen:output_type -> position_proto.Positions
	3, // 9: position_proto.PositionService.UpdatePrices:output_type -> position_proto.Empty
	7, // 10: position_proto.PositionService.GetProfitLoss:output_type -> position_proto.ProfitLoss
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_position_proto_init() }
func file_position_proto_init() {
	if File_position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_position_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionID); i {
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
		file_position_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClosePrice); i {
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
		file_position_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_position_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountID); i {
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
		file_position_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Positions); i {
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
		file_position_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_position_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfitLoss); i {
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
			RawDescriptor: file_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_position_proto_goTypes,
		DependencyIndexes: file_position_proto_depIdxs,
		MessageInfos:      file_position_proto_msgTypes,
	}.Build()
	File_position_proto = out.File
	file_position_proto_rawDesc = nil
	file_position_proto_goTypes = nil
	file_position_proto_depIdxs = nil
}