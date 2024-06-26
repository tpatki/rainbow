// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: memory.proto

package service

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

type SatisfyResponse_ResultType int32

const (
	SatisfyResponse_RESULT_TYPE_UNSPECIFIED SatisfyResponse_ResultType = 0
	SatisfyResponse_RESULT_TYPE_SUCCESS     SatisfyResponse_ResultType = 1
	SatisfyResponse_RESULT_TYPE_ERROR       SatisfyResponse_ResultType = 2
)

// Enum value maps for SatisfyResponse_ResultType.
var (
	SatisfyResponse_ResultType_name = map[int32]string{
		0: "RESULT_TYPE_UNSPECIFIED",
		1: "RESULT_TYPE_SUCCESS",
		2: "RESULT_TYPE_ERROR",
	}
	SatisfyResponse_ResultType_value = map[string]int32{
		"RESULT_TYPE_UNSPECIFIED": 0,
		"RESULT_TYPE_SUCCESS":     1,
		"RESULT_TYPE_ERROR":       2,
	}
)

func (x SatisfyResponse_ResultType) Enum() *SatisfyResponse_ResultType {
	p := new(SatisfyResponse_ResultType)
	*p = x
	return p
}

func (x SatisfyResponse_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SatisfyResponse_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_memory_proto_enumTypes[0].Descriptor()
}

func (SatisfyResponse_ResultType) Type() protoreflect.EnumType {
	return &file_memory_proto_enumTypes[0]
}

func (x SatisfyResponse_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SatisfyResponse_ResultType.Descriptor instead.
func (SatisfyResponse_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{2, 0}
}

// Enum to represent the result types of the operation.
type Response_ResultType int32

const (
	Response_RESULT_TYPE_UNSPECIFIED Response_ResultType = 0
	Response_RESULT_TYPE_SUCCESS     Response_ResultType = 1
	Response_RESULT_TYPE_ERROR       Response_ResultType = 2
)

// Enum value maps for Response_ResultType.
var (
	Response_ResultType_name = map[int32]string{
		0: "RESULT_TYPE_UNSPECIFIED",
		1: "RESULT_TYPE_SUCCESS",
		2: "RESULT_TYPE_ERROR",
	}
	Response_ResultType_value = map[string]int32{
		"RESULT_TYPE_UNSPECIFIED": 0,
		"RESULT_TYPE_SUCCESS":     1,
		"RESULT_TYPE_ERROR":       2,
	}
)

func (x Response_ResultType) Enum() *Response_ResultType {
	p := new(Response_ResultType)
	*p = x
	return p
}

func (x Response_ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response_ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_memory_proto_enumTypes[1].Descriptor()
}

func (Response_ResultType) Type() protoreflect.EnumType {
	return &file_memory_proto_enumTypes[1]
}

func (x Response_ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response_ResultType.Descriptor instead.
func (Response_ResultType) EnumDescriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{3, 0}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payload   string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Subsystem string `protobuf:"bytes,3,opt,name=subsystem,proto3" json:"subsystem,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *RegisterRequest) GetSubsystem() string {
	if x != nil {
		return x.Subsystem
	}
	return ""
}

type SatisfyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload string `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Matcher string `protobuf:"bytes,2,opt,name=matcher,proto3" json:"matcher,omitempty"`
}

func (x *SatisfyRequest) Reset() {
	*x = SatisfyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SatisfyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SatisfyRequest) ProtoMessage() {}

func (x *SatisfyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SatisfyRequest.ProtoReflect.Descriptor instead.
func (*SatisfyRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{1}
}

func (x *SatisfyRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *SatisfyRequest) GetMatcher() string {
	if x != nil {
		return x.Matcher
	}
	return ""
}

type SatisfyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters        []string                   `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	Status          SatisfyResponse_ResultType `protobuf:"varint,2,opt,name=status,proto3,enum=service.SatisfyResponse_ResultType" json:"status,omitempty"`
	TotalClusters   int32                      `protobuf:"varint,3,opt,name=total_clusters,json=totalClusters,proto3" json:"total_clusters,omitempty"`
	TotalMatches    int32                      `protobuf:"varint,4,opt,name=total_matches,json=totalMatches,proto3" json:"total_matches,omitempty"`
	TotalMismatches int32                      `protobuf:"varint,5,opt,name=total_mismatches,json=totalMismatches,proto3" json:"total_mismatches,omitempty"`
}

func (x *SatisfyResponse) Reset() {
	*x = SatisfyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SatisfyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SatisfyResponse) ProtoMessage() {}

func (x *SatisfyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SatisfyResponse.ProtoReflect.Descriptor instead.
func (*SatisfyResponse) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{2}
}

func (x *SatisfyResponse) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *SatisfyResponse) GetStatus() SatisfyResponse_ResultType {
	if x != nil {
		return x.Status
	}
	return SatisfyResponse_RESULT_TYPE_UNSPECIFIED
}

func (x *SatisfyResponse) GetTotalClusters() int32 {
	if x != nil {
		return x.TotalClusters
	}
	return 0
}

func (x *SatisfyResponse) GetTotalMatches() int32 {
	if x != nil {
		return x.TotalMatches
	}
	return 0
}

func (x *SatisfyResponse) GetTotalMismatches() int32 {
	if x != nil {
		return x.TotalMismatches
	}
	return 0
}

// Testing response - the server's response to a request.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Response_ResultType `protobuf:"varint,1,opt,name=status,proto3,enum=service.Response_ResultType" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetStatus() Response_ResultType {
	if x != nil {
		return x.Status
	}
	return Response_RESULT_TYPE_UNSPECIFIED
}

var File_memory_proto protoreflect.FileDescriptor

var file_memory_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5d, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x44, 0x0a, 0x0e, 0x53, 0x61, 0x74, 0x69, 0x73, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0xbc, 0x02, 0x0a,
	0x0f, 0x53, 0x61, 0x74, 0x69, 0x73, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x61, 0x74, 0x69, 0x73, 0x66, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d,
	0x69, 0x73, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x69, 0x73, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x22, 0x59, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x9b, 0x01, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x59,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32, 0x88, 0x01, 0x0a, 0x0b, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x3e, 0x0a, 0x07, 0x53, 0x61, 0x74,
	0x69, 0x73, 0x66, 0x79, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x61, 0x74, 0x69, 0x73, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x61, 0x74, 0x69, 0x73, 0x66, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x67, 0x65, 0x64, 0x2d, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x69, 0x6e, 0x62, 0x6f, 0x77, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memory_proto_rawDescOnce sync.Once
	file_memory_proto_rawDescData = file_memory_proto_rawDesc
)

func file_memory_proto_rawDescGZIP() []byte {
	file_memory_proto_rawDescOnce.Do(func() {
		file_memory_proto_rawDescData = protoimpl.X.CompressGZIP(file_memory_proto_rawDescData)
	})
	return file_memory_proto_rawDescData
}

var file_memory_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_memory_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_memory_proto_goTypes = []interface{}{
	(SatisfyResponse_ResultType)(0), // 0: service.SatisfyResponse.ResultType
	(Response_ResultType)(0),        // 1: service.Response.ResultType
	(*RegisterRequest)(nil),         // 2: service.RegisterRequest
	(*SatisfyRequest)(nil),          // 3: service.SatisfyRequest
	(*SatisfyResponse)(nil),         // 4: service.SatisfyResponse
	(*Response)(nil),                // 5: service.Response
}
var file_memory_proto_depIdxs = []int32{
	0, // 0: service.SatisfyResponse.status:type_name -> service.SatisfyResponse.ResultType
	1, // 1: service.Response.status:type_name -> service.Response.ResultType
	3, // 2: service.MemoryGraph.Satisfy:input_type -> service.SatisfyRequest
	2, // 3: service.MemoryGraph.Register:input_type -> service.RegisterRequest
	4, // 4: service.MemoryGraph.Satisfy:output_type -> service.SatisfyResponse
	5, // 5: service.MemoryGraph.Register:output_type -> service.Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_memory_proto_init() }
func file_memory_proto_init() {
	if File_memory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_memory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SatisfyRequest); i {
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
		file_memory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SatisfyResponse); i {
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
		file_memory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_memory_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memory_proto_goTypes,
		DependencyIndexes: file_memory_proto_depIdxs,
		EnumInfos:         file_memory_proto_enumTypes,
		MessageInfos:      file_memory_proto_msgTypes,
	}.Build()
	File_memory_proto = out.File
	file_memory_proto_rawDesc = nil
	file_memory_proto_goTypes = nil
	file_memory_proto_depIdxs = nil
}
