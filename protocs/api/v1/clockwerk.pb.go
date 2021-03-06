// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: clockwerk.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Scheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EntryId        int32    `protobuf:"varint,2,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	ReferenceId    string   `protobuf:"bytes,3,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Name           string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Executor       string   `protobuf:"bytes,5,opt,name=executor,proto3" json:"executor,omitempty"`
	Command        string   `protobuf:"bytes,6,opt,name=command,proto3" json:"command,omitempty"`
	Url            string   `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Method         string   `protobuf:"bytes,8,opt,name=method,proto3" json:"method,omitempty"`
	Body           string   `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty"`
	Retry          int32    `protobuf:"varint,10,opt,name=retry,proto3" json:"retry,omitempty"`
	RetryThreshold int32    `protobuf:"varint,11,opt,name=retry_threshold,json=retryThreshold,proto3" json:"retry_threshold,omitempty"`
	Headers        []string `protobuf:"bytes,12,rep,name=headers,proto3" json:"headers,omitempty"`
	Spec           string   `protobuf:"bytes,13,opt,name=spec,proto3" json:"spec,omitempty"`
	Disabled       bool     `protobuf:"varint,14,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Persist        bool     `protobuf:"varint,15,opt,name=persist,proto3" json:"persist,omitempty"`
	CreatedAt      int64    `protobuf:"varint,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Username       string   `protobuf:"bytes,17,opt,name=username,proto3" json:"username,omitempty"`
	Password       string   `protobuf:"bytes,18,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Scheduler) Reset() {
	*x = Scheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clockwerk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scheduler) ProtoMessage() {}

func (x *Scheduler) ProtoReflect() protoreflect.Message {
	mi := &file_clockwerk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scheduler.ProtoReflect.Descriptor instead.
func (*Scheduler) Descriptor() ([]byte, []int) {
	return file_clockwerk_proto_rawDescGZIP(), []int{0}
}

func (x *Scheduler) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Scheduler) GetEntryId() int32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *Scheduler) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *Scheduler) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scheduler) GetExecutor() string {
	if x != nil {
		return x.Executor
	}
	return ""
}

func (x *Scheduler) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Scheduler) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Scheduler) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Scheduler) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Scheduler) GetRetry() int32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

func (x *Scheduler) GetRetryThreshold() int32 {
	if x != nil {
		return x.RetryThreshold
	}
	return 0
}

func (x *Scheduler) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Scheduler) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *Scheduler) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Scheduler) GetPersist() bool {
	if x != nil {
		return x.Persist
	}
	return false
}

func (x *Scheduler) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Scheduler) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Scheduler) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Schedulers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedulers []*Scheduler `protobuf:"bytes,1,rep,name=schedulers,proto3" json:"schedulers,omitempty"`
}

func (x *Schedulers) Reset() {
	*x = Schedulers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clockwerk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedulers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedulers) ProtoMessage() {}

func (x *Schedulers) ProtoReflect() protoreflect.Message {
	mi := &file_clockwerk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedulers.ProtoReflect.Descriptor instead.
func (*Schedulers) Descriptor() ([]byte, []int) {
	return file_clockwerk_proto_rawDescGZIP(), []int{1}
}

func (x *Schedulers) GetSchedulers() []*Scheduler {
	if x != nil {
		return x.Schedulers
	}
	return nil
}

type SelectScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReferenceId string `protobuf:"bytes,2,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Username    string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SelectScheduler) Reset() {
	*x = SelectScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clockwerk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectScheduler) ProtoMessage() {}

func (x *SelectScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_clockwerk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectScheduler.ProtoReflect.Descriptor instead.
func (*SelectScheduler) Descriptor() ([]byte, []int) {
	return file_clockwerk_proto_rawDescGZIP(), []int{2}
}

func (x *SelectScheduler) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SelectScheduler) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *SelectScheduler) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SelectScheduler) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SelectToggle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ReferenceId string `protobuf:"bytes,2,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	Disabled    bool   `protobuf:"varint,3,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Username    string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SelectToggle) Reset() {
	*x = SelectToggle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clockwerk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectToggle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectToggle) ProtoMessage() {}

func (x *SelectToggle) ProtoReflect() protoreflect.Message {
	mi := &file_clockwerk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectToggle.ProtoReflect.Descriptor instead.
func (*SelectToggle) Descriptor() ([]byte, []int) {
	return file_clockwerk_proto_rawDescGZIP(), []int{3}
}

func (x *SelectToggle) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SelectToggle) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *SelectToggle) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *SelectToggle) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SelectToggle) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_clockwerk_proto protoreflect.FileDescriptor

var file_clockwerk_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x77, 0x65, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x03, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x3f, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x73, 0x22, 0x7c, 0x0a, 0x0f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x95, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x6f,
	0x67, 0x67, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xb1, 0x04, 0x0a, 0x09,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x77, 0x65, 0x72, 0x6b, 0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x75, 0x6d,
	0x6d, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3a, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x0c, 0x41, 0x64,
	0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x12, 0x42, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0f, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x67, 0x67, 0x6c, 0x65, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x69, 0x6c, 0x76, 0x65, 0x72, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x63, 0x6c,
	0x6f, 0x63, 0x6b, 0x77, 0x65, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clockwerk_proto_rawDescOnce sync.Once
	file_clockwerk_proto_rawDescData = file_clockwerk_proto_rawDesc
)

func file_clockwerk_proto_rawDescGZIP() []byte {
	file_clockwerk_proto_rawDescOnce.Do(func() {
		file_clockwerk_proto_rawDescData = protoimpl.X.CompressGZIP(file_clockwerk_proto_rawDescData)
	})
	return file_clockwerk_proto_rawDescData
}

var file_clockwerk_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_clockwerk_proto_goTypes = []interface{}{
	(*Scheduler)(nil),       // 0: api.v1.Scheduler
	(*Schedulers)(nil),      // 1: api.v1.Schedulers
	(*SelectScheduler)(nil), // 2: api.v1.SelectScheduler
	(*SelectToggle)(nil),    // 3: api.v1.SelectToggle
	(*emptypb.Empty)(nil),   // 4: google.protobuf.Empty
}
var file_clockwerk_proto_depIdxs = []int32{
	0,  // 0: api.v1.Schedulers.schedulers:type_name -> api.v1.Scheduler
	4,  // 1: api.v1.Clockwerk.GetDummy:input_type -> google.protobuf.Empty
	4,  // 2: api.v1.Clockwerk.PostDummy:input_type -> google.protobuf.Empty
	4,  // 3: api.v1.Clockwerk.DeleteDummy:input_type -> google.protobuf.Empty
	4,  // 4: api.v1.Clockwerk.PutDummy:input_type -> google.protobuf.Empty
	4,  // 5: api.v1.Clockwerk.GetSchedulers:input_type -> google.protobuf.Empty
	0,  // 6: api.v1.Clockwerk.AddScheduler:input_type -> api.v1.Scheduler
	2,  // 7: api.v1.Clockwerk.DeleteScheduler:input_type -> api.v1.SelectScheduler
	3,  // 8: api.v1.Clockwerk.ToggleScheduler:input_type -> api.v1.SelectToggle
	4,  // 9: api.v1.Clockwerk.Backup:input_type -> google.protobuf.Empty
	4,  // 10: api.v1.Clockwerk.GetDummy:output_type -> google.protobuf.Empty
	4,  // 11: api.v1.Clockwerk.PostDummy:output_type -> google.protobuf.Empty
	4,  // 12: api.v1.Clockwerk.DeleteDummy:output_type -> google.protobuf.Empty
	4,  // 13: api.v1.Clockwerk.PutDummy:output_type -> google.protobuf.Empty
	1,  // 14: api.v1.Clockwerk.GetSchedulers:output_type -> api.v1.Schedulers
	0,  // 15: api.v1.Clockwerk.AddScheduler:output_type -> api.v1.Scheduler
	4,  // 16: api.v1.Clockwerk.DeleteScheduler:output_type -> google.protobuf.Empty
	4,  // 17: api.v1.Clockwerk.ToggleScheduler:output_type -> google.protobuf.Empty
	4,  // 18: api.v1.Clockwerk.Backup:output_type -> google.protobuf.Empty
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_clockwerk_proto_init() }
func file_clockwerk_proto_init() {
	if File_clockwerk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clockwerk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scheduler); i {
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
		file_clockwerk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedulers); i {
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
		file_clockwerk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectScheduler); i {
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
		file_clockwerk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectToggle); i {
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
			RawDescriptor: file_clockwerk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clockwerk_proto_goTypes,
		DependencyIndexes: file_clockwerk_proto_depIdxs,
		MessageInfos:      file_clockwerk_proto_msgTypes,
	}.Build()
	File_clockwerk_proto = out.File
	file_clockwerk_proto_rawDesc = nil
	file_clockwerk_proto_goTypes = nil
	file_clockwerk_proto_depIdxs = nil
}
