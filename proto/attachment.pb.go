// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: attachment.proto

package attachment

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type PrintSize int32

const (
	PrintSize_A4 PrintSize = 0
	PrintSize_A1 PrintSize = 1
	PrintSize_A2 PrintSize = 2
	PrintSize_A3 PrintSize = 3
	PrintSize_A0 PrintSize = 4
	PrintSize_A5 PrintSize = 5
	PrintSize_B1 PrintSize = 6
	PrintSize_B2 PrintSize = 7
	PrintSize_B3 PrintSize = 8
	PrintSize_B4 PrintSize = 9
	PrintSize_B5 PrintSize = 10
)

// Enum value maps for PrintSize.
var (
	PrintSize_name = map[int32]string{
		0:  "A4",
		1:  "A1",
		2:  "A2",
		3:  "A3",
		4:  "A0",
		5:  "A5",
		6:  "B1",
		7:  "B2",
		8:  "B3",
		9:  "B4",
		10: "B5",
	}
	PrintSize_value = map[string]int32{
		"A4": 0,
		"A1": 1,
		"A2": 2,
		"A3": 3,
		"A0": 4,
		"A5": 5,
		"B1": 6,
		"B2": 7,
		"B3": 8,
		"B4": 9,
		"B5": 10,
	}
)

func (x PrintSize) Enum() *PrintSize {
	p := new(PrintSize)
	*p = x
	return p
}

func (x PrintSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrintSize) Descriptor() protoreflect.EnumDescriptor {
	return file_attachment_proto_enumTypes[0].Descriptor()
}

func (PrintSize) Type() protoreflect.EnumType {
	return &file_attachment_proto_enumTypes[0]
}

func (x PrintSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrintSize.Descriptor instead.
func (PrintSize) EnumDescriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{0}
}

type Orientation int32

const (
	Orientation_Portrait  Orientation = 0
	Orientation_Landscape Orientation = 1
)

// Enum value maps for Orientation.
var (
	Orientation_name = map[int32]string{
		0: "Portrait",
		1: "Landscape",
	}
	Orientation_value = map[string]int32{
		"Portrait":  0,
		"Landscape": 1,
	}
)

func (x Orientation) Enum() *Orientation {
	p := new(Orientation)
	*p = x
	return p
}

func (x Orientation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Orientation) Descriptor() protoreflect.EnumDescriptor {
	return file_attachment_proto_enumTypes[1].Descriptor()
}

func (Orientation) Type() protoreflect.EnumType {
	return &file_attachment_proto_enumTypes[1]
}

func (x Orientation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Orientation.Descriptor instead.
func (Orientation) EnumDescriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{1}
}

type PDFMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       *string     `protobuf:"bytes,1,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Dpi         *uint32     `protobuf:"varint,2,opt,name=dpi,proto3,oneof" json:"dpi,omitempty"`
	Orientation Orientation `protobuf:"varint,3,opt,name=orientation,proto3,enum=attachment.v1.Orientation" json:"orientation,omitempty"`
	PrintSize   PrintSize   `protobuf:"varint,4,opt,name=printSize,proto3,enum=attachment.v1.PrintSize" json:"printSize,omitempty"`
}

func (x *PDFMeta) Reset() {
	*x = PDFMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDFMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDFMeta) ProtoMessage() {}

func (x *PDFMeta) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDFMeta.ProtoReflect.Descriptor instead.
func (*PDFMeta) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{0}
}

func (x *PDFMeta) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *PDFMeta) GetDpi() uint32 {
	if x != nil && x.Dpi != nil {
		return *x.Dpi
	}
	return 0
}

func (x *PDFMeta) GetOrientation() Orientation {
	if x != nil {
		return x.Orientation
	}
	return Orientation_Portrait
}

func (x *PDFMeta) GetPrintSize() PrintSize {
	if x != nil {
		return x.PrintSize
	}
	return PrintSize_A4
}

type PDFData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutFileUrl string `protobuf:"bytes,1,opt,name=outFileUrl,proto3" json:"outFileUrl,omitempty"`
}

func (x *PDFData) Reset() {
	*x = PDFData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDFData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDFData) ProtoMessage() {}

func (x *PDFData) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDFData.ProtoReflect.Descriptor instead.
func (*PDFData) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{1}
}

func (x *PDFData) GetOutFileUrl() string {
	if x != nil {
		return x.OutFileUrl
	}
	return ""
}

type CSVData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutFileUrl string `protobuf:"bytes,1,opt,name=outFileUrl,proto3" json:"outFileUrl,omitempty"`
}

func (x *CSVData) Reset() {
	*x = CSVData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSVData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSVData) ProtoMessage() {}

func (x *CSVData) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSVData.ProtoReflect.Descriptor instead.
func (*CSVData) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{2}
}

func (x *CSVData) GetOutFileUrl() string {
	if x != nil {
		return x.OutFileUrl
	}
	return ""
}

type PDFRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *PDFMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Url  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Html string   `protobuf:"bytes,3,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *PDFRequest) Reset() {
	*x = PDFRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDFRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDFRequest) ProtoMessage() {}

func (x *PDFRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDFRequest.ProtoReflect.Descriptor instead.
func (*PDFRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{3}
}

func (x *PDFRequest) GetMeta() *PDFMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *PDFRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PDFRequest) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

type CSVRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CSVRequest) Reset() {
	*x = CSVRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSVRequest) ProtoMessage() {}

func (x *CSVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSVRequest.ProtoReflect.Descriptor instead.
func (*CSVRequest) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{4}
}

type PDFResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *PDFData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code    int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PDFResponse) Reset() {
	*x = PDFResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDFResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDFResponse) ProtoMessage() {}

func (x *PDFResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDFResponse.ProtoReflect.Descriptor instead.
func (*PDFResponse) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{5}
}

func (x *PDFResponse) GetData() *PDFData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PDFResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PDFResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CSVResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    *CSVData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code    int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CSVResponse) Reset() {
	*x = CSVResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSVResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSVResponse) ProtoMessage() {}

func (x *CSVResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSVResponse.ProtoReflect.Descriptor instead.
func (*CSVResponse) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{6}
}

func (x *CSVResponse) GetData() *CSVData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CSVResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CSVResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Err struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Err) Reset() {
	*x = Err{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Err) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Err) ProtoMessage() {}

func (x *Err) ProtoReflect() protoreflect.Message {
	mi := &file_attachment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Err.ProtoReflect.Descriptor instead.
func (*Err) Descriptor() ([]byte, []int) {
	return file_attachment_proto_rawDescGZIP(), []int{7}
}

func (x *Err) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Err) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_attachment_proto protoreflect.FileDescriptor

var file_attachment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9a, 0x02, 0x0a, 0x07, 0x50, 0x44, 0x46, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xe0, 0x41, 0x02,
	0x92, 0x41, 0x15, 0x32, 0x13, 0x54, 0x68, 0x65, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x20, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x03, 0x64, 0x70, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x22, 0x92, 0x41, 0x1f, 0x32, 0x18, 0x54, 0x68, 0x65, 0x20, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x20, 0x64, 0x6f, 0x74, 0x73, 0x20, 0x70, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x63, 0x68, 0x2e,
	0x3a, 0x03, 0x36, 0x30, 0x30, 0x48, 0x01, 0x52, 0x03, 0x64, 0x70, 0x69, 0x88, 0x01, 0x01, 0x12,
	0x3c, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x14, 0x92, 0x41, 0x11, 0x32,
	0x0f, 0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x70, 0x65, 0x72, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x2e,
	0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64, 0x70, 0x69, 0x22, 0x29, 0x0a,
	0x07, 0x50, 0x44, 0x46, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x29, 0x0a, 0x07, 0x43, 0x53, 0x56, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x72, 0x6c, 0x22, 0x90, 0x01, 0x0a, 0x0a, 0x50, 0x44, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x44, 0x46, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x26,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0x92, 0x41, 0x11,
	0x32, 0x0f, 0x54, 0x68, 0x65, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x75, 0x72, 0x6c,
	0x2e, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0x92, 0x41, 0x17, 0x32, 0x15, 0x54, 0x68, 0x65, 0x20, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x70, 0x61, 0x67, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x53, 0x56, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x0b, 0x50, 0x44, 0x46, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x44, 0x46, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x67, 0x0a,
	0x0b, 0x43, 0x53, 0x56, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x53, 0x56, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x63, 0x0a, 0x09, 0x50,
	0x72, 0x69, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x34, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x41, 0x31, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x32, 0x10, 0x02,
	0x12, 0x06, 0x0a, 0x02, 0x41, 0x33, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x30, 0x10, 0x04,
	0x12, 0x06, 0x0a, 0x02, 0x41, 0x35, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x31, 0x10, 0x06,
	0x12, 0x06, 0x0a, 0x02, 0x42, 0x32, 0x10, 0x07, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x33, 0x10, 0x08,
	0x12, 0x06, 0x0a, 0x02, 0x42, 0x34, 0x10, 0x09, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x35, 0x10, 0x0a,
	0x2a, 0x2a, 0x0a, 0x0b, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x10, 0x01, 0x32, 0xd7, 0x01, 0x0a,
	0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x60, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x44, 0x46, 0x12,
	0x19, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x44, 0x46, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x44, 0x46, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x44,
	0x46, 0x3a, 0x01, 0x2a, 0x12, 0x60, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x53,
	0x56, 0x12, 0x19, 0x2e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x53, 0x56, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x53, 0x56,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x53, 0x56, 0x3a, 0x01, 0x2a, 0x42, 0xc3, 0x02, 0x5a, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x2e, 0x63, 0x6e, 0x2d, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x2d, 0x31, 0x2e, 0x61, 0x6d, 0x61,
	0x7a, 0x6f, 0x6e, 0x61, 0x77, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x92, 0x41, 0xde, 0x01, 0x12, 0x3f,
	0x0a, 0x16, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x20, 0x41, 0x50, 0x49, 0x22, 0x20, 0x0a, 0x09, 0x6b, 0x65, 0x76, 0x69,
	0x6e, 0x20, 0x6c, 0x69, 0x75, 0x1a, 0x13, 0x6c, 0x63, 0x68, 0x65, 0x40, 0x6b, 0x6e, 0x69, 0x74,
	0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a,
	0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x74, 0x0a, 0x22, 0x67, 0x52, 0x50, 0x43, 0x20, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x4e, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x63, 0x6e, 0x2d, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x2d, 0x31, 0x2e,
	0x61, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x61, 0x77, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x2f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attachment_proto_rawDescOnce sync.Once
	file_attachment_proto_rawDescData = file_attachment_proto_rawDesc
)

func file_attachment_proto_rawDescGZIP() []byte {
	file_attachment_proto_rawDescOnce.Do(func() {
		file_attachment_proto_rawDescData = protoimpl.X.CompressGZIP(file_attachment_proto_rawDescData)
	})
	return file_attachment_proto_rawDescData
}

var file_attachment_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_attachment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_attachment_proto_goTypes = []interface{}{
	(PrintSize)(0),      // 0: attachment.v1.PrintSize
	(Orientation)(0),    // 1: attachment.v1.Orientation
	(*PDFMeta)(nil),     // 2: attachment.v1.PDFMeta
	(*PDFData)(nil),     // 3: attachment.v1.PDFData
	(*CSVData)(nil),     // 4: attachment.v1.CSVData
	(*PDFRequest)(nil),  // 5: attachment.v1.PDFRequest
	(*CSVRequest)(nil),  // 6: attachment.v1.CSVRequest
	(*PDFResponse)(nil), // 7: attachment.v1.PDFResponse
	(*CSVResponse)(nil), // 8: attachment.v1.CSVResponse
	(*Err)(nil),         // 9: attachment.v1.Err
}
var file_attachment_proto_depIdxs = []int32{
	1, // 0: attachment.v1.PDFMeta.orientation:type_name -> attachment.v1.Orientation
	0, // 1: attachment.v1.PDFMeta.printSize:type_name -> attachment.v1.PrintSize
	2, // 2: attachment.v1.PDFRequest.meta:type_name -> attachment.v1.PDFMeta
	3, // 3: attachment.v1.PDFResponse.data:type_name -> attachment.v1.PDFData
	4, // 4: attachment.v1.CSVResponse.data:type_name -> attachment.v1.CSVData
	5, // 5: attachment.v1.AttachmentService.CreatePDF:input_type -> attachment.v1.PDFRequest
	6, // 6: attachment.v1.AttachmentService.CreateCSV:input_type -> attachment.v1.CSVRequest
	7, // 7: attachment.v1.AttachmentService.CreatePDF:output_type -> attachment.v1.PDFResponse
	8, // 8: attachment.v1.AttachmentService.CreateCSV:output_type -> attachment.v1.CSVResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_attachment_proto_init() }
func file_attachment_proto_init() {
	if File_attachment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attachment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDFMeta); i {
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
		file_attachment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDFData); i {
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
		file_attachment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSVData); i {
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
		file_attachment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDFRequest); i {
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
		file_attachment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSVRequest); i {
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
		file_attachment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDFResponse); i {
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
		file_attachment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSVResponse); i {
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
		file_attachment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Err); i {
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
	file_attachment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_attachment_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attachment_proto_goTypes,
		DependencyIndexes: file_attachment_proto_depIdxs,
		EnumInfos:         file_attachment_proto_enumTypes,
		MessageInfos:      file_attachment_proto_msgTypes,
	}.Build()
	File_attachment_proto = out.File
	file_attachment_proto_rawDesc = nil
	file_attachment_proto_goTypes = nil
	file_attachment_proto_depIdxs = nil
}