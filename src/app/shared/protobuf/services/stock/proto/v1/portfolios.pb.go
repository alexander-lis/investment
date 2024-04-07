// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: services/stock/proto/v1/portfolios.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Create portfolio
type CreatePortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	From string `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	To   string `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`
}

func (x *CreatePortfolioRequest) Reset() {
	*x = CreatePortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortfolioRequest) ProtoMessage() {}

func (x *CreatePortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortfolioRequest.ProtoReflect.Descriptor instead.
func (*CreatePortfolioRequest) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePortfolioRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePortfolioRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreatePortfolioRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type CreatePortfolioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Portfolio *Portfolio `protobuf:"bytes,1,opt,name=Portfolio,proto3" json:"Portfolio,omitempty"`
}

func (x *CreatePortfolioResponse) Reset() {
	*x = CreatePortfolioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortfolioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortfolioResponse) ProtoMessage() {}

func (x *CreatePortfolioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortfolioResponse.ProtoReflect.Descriptor instead.
func (*CreatePortfolioResponse) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePortfolioResponse) GetPortfolio() *Portfolio {
	if x != nil {
		return x.Portfolio
	}
	return nil
}

// Get portfolios
type GetPortfoliosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPortfoliosRequest) Reset() {
	*x = GetPortfoliosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortfoliosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortfoliosRequest) ProtoMessage() {}

func (x *GetPortfoliosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortfoliosRequest.ProtoReflect.Descriptor instead.
func (*GetPortfoliosRequest) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{2}
}

type GetPortfoliosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Portfolios []*Portfolio `protobuf:"bytes,1,rep,name=Portfolios,proto3" json:"Portfolios,omitempty"`
}

func (x *GetPortfoliosResponse) Reset() {
	*x = GetPortfoliosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortfoliosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortfoliosResponse) ProtoMessage() {}

func (x *GetPortfoliosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortfoliosResponse.ProtoReflect.Descriptor instead.
func (*GetPortfoliosResponse) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{3}
}

func (x *GetPortfoliosResponse) GetPortfolios() []*Portfolio {
	if x != nil {
		return x.Portfolios
	}
	return nil
}

// Get portfolio
type GetPortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortfolioId string `protobuf:"bytes,1,opt,name=PortfolioId,proto3" json:"PortfolioId,omitempty"`
}

func (x *GetPortfolioRequest) Reset() {
	*x = GetPortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortfolioRequest) ProtoMessage() {}

func (x *GetPortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortfolioRequest.ProtoReflect.Descriptor instead.
func (*GetPortfolioRequest) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{4}
}

func (x *GetPortfolioRequest) GetPortfolioId() string {
	if x != nil {
		return x.PortfolioId
	}
	return ""
}

type GetPortfolioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Portfolio *Portfolio `protobuf:"bytes,1,opt,name=Portfolio,proto3" json:"Portfolio,omitempty"`
	Groups    []*Group   `protobuf:"bytes,2,rep,name=Groups,proto3" json:"Groups,omitempty"`
	Sectors   []*Sector  `protobuf:"bytes,3,rep,name=Sectors,proto3" json:"Sectors,omitempty"`
	Assets    []*Asset   `protobuf:"bytes,4,rep,name=Assets,proto3" json:"Assets,omitempty"`
}

func (x *GetPortfolioResponse) Reset() {
	*x = GetPortfolioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortfolioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortfolioResponse) ProtoMessage() {}

func (x *GetPortfolioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortfolioResponse.ProtoReflect.Descriptor instead.
func (*GetPortfolioResponse) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{5}
}

func (x *GetPortfolioResponse) GetPortfolio() *Portfolio {
	if x != nil {
		return x.Portfolio
	}
	return nil
}

func (x *GetPortfolioResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GetPortfolioResponse) GetSectors() []*Sector {
	if x != nil {
		return x.Sectors
	}
	return nil
}

func (x *GetPortfolioResponse) GetAssets() []*Asset {
	if x != nil {
		return x.Assets
	}
	return nil
}

// Update Portfolio
type UpdatePortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Portfolio *Portfolio `protobuf:"bytes,1,opt,name=Portfolio,proto3" json:"Portfolio,omitempty"`
}

func (x *UpdatePortfolioRequest) Reset() {
	*x = UpdatePortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePortfolioRequest) ProtoMessage() {}

func (x *UpdatePortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePortfolioRequest.ProtoReflect.Descriptor instead.
func (*UpdatePortfolioRequest) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePortfolioRequest) GetPortfolio() *Portfolio {
	if x != nil {
		return x.Portfolio
	}
	return nil
}

// Delete Portfolio
type DeletePortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortfolioId string `protobuf:"bytes,1,opt,name=PortfolioId,proto3" json:"PortfolioId,omitempty"`
}

func (x *DeletePortfolioRequest) Reset() {
	*x = DeletePortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePortfolioRequest) ProtoMessage() {}

func (x *DeletePortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePortfolioRequest.ProtoReflect.Descriptor instead.
func (*DeletePortfolioRequest) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePortfolioRequest) GetPortfolioId() string {
	if x != nil {
		return x.PortfolioId
	}
	return ""
}

// Types
type Portfolio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	From string `protobuf:"bytes,3,opt,name=From,proto3" json:"From,omitempty"`
	To   string `protobuf:"bytes,4,opt,name=To,proto3" json:"To,omitempty"`
}

func (x *Portfolio) Reset() {
	*x = Portfolio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Portfolio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Portfolio) ProtoMessage() {}

func (x *Portfolio) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Portfolio.ProtoReflect.Descriptor instead.
func (*Portfolio) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{8}
}

func (x *Portfolio) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Portfolio) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Portfolio) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Portfolio) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{9}
}

type Sector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Sector) Reset() {
	*x = Sector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sector) ProtoMessage() {}

func (x *Sector) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sector.ProtoReflect.Descriptor instead.
func (*Sector) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{10}
}

type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{11}
}

// Util types
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_services_stock_proto_v1_portfolios_proto_msgTypes[12]
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
	return file_services_stock_proto_v1_portfolios_proto_rawDescGZIP(), []int{12}
}

var File_services_stock_proto_v1_portfolios_proto protoreflect.FileDescriptor

var file_services_stock_proto_v1_portfolios_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f,
	0x6c, 0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x22, 0x5b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f,
	0x6c, 0x69, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f,
	0x6c, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x0a, 0x50, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49,
	0x64, 0x22, 0x83, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c,
	0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x50, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x52, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x36, 0x0a, 0x06,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x36, 0x0a, 0x06, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x06, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x22, 0x5a, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f,
	0x6c, 0x69, 0x6f, 0x22, 0x3a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x64, 0x22,
	0x53, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x54, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x08, 0x0a,
	0x06, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xb7, 0x04, 0x0a, 0x10, 0x50, 0x6f,
	0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x12, 0x2f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x12, 0x2d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x2f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x64, 0x0a,
	0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f,
	0x12, 0x2f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0xc7, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x42, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0xa2, 0x02, 0x03, 0x53, 0x53, 0x50, 0xaa, 0x02, 0x17, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x17, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x5c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_stock_proto_v1_portfolios_proto_rawDescOnce sync.Once
	file_services_stock_proto_v1_portfolios_proto_rawDescData = file_services_stock_proto_v1_portfolios_proto_rawDesc
)

func file_services_stock_proto_v1_portfolios_proto_rawDescGZIP() []byte {
	file_services_stock_proto_v1_portfolios_proto_rawDescOnce.Do(func() {
		file_services_stock_proto_v1_portfolios_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_stock_proto_v1_portfolios_proto_rawDescData)
	})
	return file_services_stock_proto_v1_portfolios_proto_rawDescData
}

var file_services_stock_proto_v1_portfolios_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_services_stock_proto_v1_portfolios_proto_goTypes = []interface{}{
	(*CreatePortfolioRequest)(nil),  // 0: services.stock.proto.v1.CreatePortfolioRequest
	(*CreatePortfolioResponse)(nil), // 1: services.stock.proto.v1.CreatePortfolioResponse
	(*GetPortfoliosRequest)(nil),    // 2: services.stock.proto.v1.GetPortfoliosRequest
	(*GetPortfoliosResponse)(nil),   // 3: services.stock.proto.v1.GetPortfoliosResponse
	(*GetPortfolioRequest)(nil),     // 4: services.stock.proto.v1.GetPortfolioRequest
	(*GetPortfolioResponse)(nil),    // 5: services.stock.proto.v1.GetPortfolioResponse
	(*UpdatePortfolioRequest)(nil),  // 6: services.stock.proto.v1.UpdatePortfolioRequest
	(*DeletePortfolioRequest)(nil),  // 7: services.stock.proto.v1.DeletePortfolioRequest
	(*Portfolio)(nil),               // 8: services.stock.proto.v1.Portfolio
	(*Group)(nil),                   // 9: services.stock.proto.v1.Group
	(*Sector)(nil),                  // 10: services.stock.proto.v1.Sector
	(*Asset)(nil),                   // 11: services.stock.proto.v1.Asset
	(*Empty)(nil),                   // 12: services.stock.proto.v1.Empty
}
var file_services_stock_proto_v1_portfolios_proto_depIdxs = []int32{
	8,  // 0: services.stock.proto.v1.CreatePortfolioResponse.Portfolio:type_name -> services.stock.proto.v1.Portfolio
	8,  // 1: services.stock.proto.v1.GetPortfoliosResponse.Portfolios:type_name -> services.stock.proto.v1.Portfolio
	8,  // 2: services.stock.proto.v1.GetPortfolioResponse.Portfolio:type_name -> services.stock.proto.v1.Portfolio
	9,  // 3: services.stock.proto.v1.GetPortfolioResponse.Groups:type_name -> services.stock.proto.v1.Group
	10, // 4: services.stock.proto.v1.GetPortfolioResponse.Sectors:type_name -> services.stock.proto.v1.Sector
	11, // 5: services.stock.proto.v1.GetPortfolioResponse.Assets:type_name -> services.stock.proto.v1.Asset
	8,  // 6: services.stock.proto.v1.UpdatePortfolioRequest.Portfolio:type_name -> services.stock.proto.v1.Portfolio
	0,  // 7: services.stock.proto.v1.PortfolioService.CreatePortfolio:input_type -> services.stock.proto.v1.CreatePortfolioRequest
	2,  // 8: services.stock.proto.v1.PortfolioService.GetPortfolios:input_type -> services.stock.proto.v1.GetPortfoliosRequest
	4,  // 9: services.stock.proto.v1.PortfolioService.GetPortfolio:input_type -> services.stock.proto.v1.GetPortfolioRequest
	6,  // 10: services.stock.proto.v1.PortfolioService.UpdatePortfolio:input_type -> services.stock.proto.v1.UpdatePortfolioRequest
	7,  // 11: services.stock.proto.v1.PortfolioService.DeletePortfolio:input_type -> services.stock.proto.v1.DeletePortfolioRequest
	1,  // 12: services.stock.proto.v1.PortfolioService.CreatePortfolio:output_type -> services.stock.proto.v1.CreatePortfolioResponse
	3,  // 13: services.stock.proto.v1.PortfolioService.GetPortfolios:output_type -> services.stock.proto.v1.GetPortfoliosResponse
	5,  // 14: services.stock.proto.v1.PortfolioService.GetPortfolio:output_type -> services.stock.proto.v1.GetPortfolioResponse
	12, // 15: services.stock.proto.v1.PortfolioService.UpdatePortfolio:output_type -> services.stock.proto.v1.Empty
	12, // 16: services.stock.proto.v1.PortfolioService.DeletePortfolio:output_type -> services.stock.proto.v1.Empty
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_services_stock_proto_v1_portfolios_proto_init() }
func file_services_stock_proto_v1_portfolios_proto_init() {
	if File_services_stock_proto_v1_portfolios_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_stock_proto_v1_portfolios_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePortfolioRequest); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePortfolioResponse); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortfoliosRequest); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortfoliosResponse); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortfolioRequest); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortfolioResponse); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePortfolioRequest); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePortfolioRequest); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Portfolio); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sector); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
		file_services_stock_proto_v1_portfolios_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_stock_proto_v1_portfolios_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_stock_proto_v1_portfolios_proto_goTypes,
		DependencyIndexes: file_services_stock_proto_v1_portfolios_proto_depIdxs,
		MessageInfos:      file_services_stock_proto_v1_portfolios_proto_msgTypes,
	}.Build()
	File_services_stock_proto_v1_portfolios_proto = out.File
	file_services_stock_proto_v1_portfolios_proto_rawDesc = nil
	file_services_stock_proto_v1_portfolios_proto_goTypes = nil
	file_services_stock_proto_v1_portfolios_proto_depIdxs = nil
}
