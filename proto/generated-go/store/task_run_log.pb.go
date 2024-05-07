// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: store/task_run_log.proto

package store

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

type TaskRunLog_Type int32

const (
	TaskRunLog_TYPE_UNSPECIFIED  TaskRunLog_Type = 0
	TaskRunLog_SCHEMA_DUMP_START TaskRunLog_Type = 1
	TaskRunLog_SCHEMA_DUMP_END   TaskRunLog_Type = 2
	TaskRunLog_COMMAND_EXECUTE   TaskRunLog_Type = 3
	TaskRunLog_COMMAND_RESPONSE  TaskRunLog_Type = 4
)

// Enum value maps for TaskRunLog_Type.
var (
	TaskRunLog_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "SCHEMA_DUMP_START",
		2: "SCHEMA_DUMP_END",
		3: "COMMAND_EXECUTE",
		4: "COMMAND_RESPONSE",
	}
	TaskRunLog_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":  0,
		"SCHEMA_DUMP_START": 1,
		"SCHEMA_DUMP_END":   2,
		"COMMAND_EXECUTE":   3,
		"COMMAND_RESPONSE":  4,
	}
)

func (x TaskRunLog_Type) Enum() *TaskRunLog_Type {
	p := new(TaskRunLog_Type)
	*p = x
	return p
}

func (x TaskRunLog_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskRunLog_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_task_run_log_proto_enumTypes[0].Descriptor()
}

func (TaskRunLog_Type) Type() protoreflect.EnumType {
	return &file_store_task_run_log_proto_enumTypes[0]
}

func (x TaskRunLog_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskRunLog_Type.Descriptor instead.
func (TaskRunLog_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 0}
}

type TaskRunLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            TaskRunLog_Type             `protobuf:"varint,1,opt,name=type,proto3,enum=bytebase.store.TaskRunLog_Type" json:"type,omitempty"`
	SchemaDumpStart *TaskRunLog_SchemaDumpStart `protobuf:"bytes,2,opt,name=schema_dump_start,json=schemaDumpStart,proto3" json:"schema_dump_start,omitempty"`
	SchemaDumpEnd   *TaskRunLog_SchemaDumpEnd   `protobuf:"bytes,3,opt,name=schema_dump_end,json=schemaDumpEnd,proto3" json:"schema_dump_end,omitempty"`
	CommandExecute  *TaskRunLog_CommandExecute  `protobuf:"bytes,4,opt,name=command_execute,json=commandExecute,proto3" json:"command_execute,omitempty"`
	CommandResponse *TaskRunLog_CommandResponse `protobuf:"bytes,5,opt,name=command_response,json=commandResponse,proto3" json:"command_response,omitempty"`
}

func (x *TaskRunLog) Reset() {
	*x = TaskRunLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog) ProtoMessage() {}

func (x *TaskRunLog) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog.ProtoReflect.Descriptor instead.
func (*TaskRunLog) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0}
}

func (x *TaskRunLog) GetType() TaskRunLog_Type {
	if x != nil {
		return x.Type
	}
	return TaskRunLog_TYPE_UNSPECIFIED
}

func (x *TaskRunLog) GetSchemaDumpStart() *TaskRunLog_SchemaDumpStart {
	if x != nil {
		return x.SchemaDumpStart
	}
	return nil
}

func (x *TaskRunLog) GetSchemaDumpEnd() *TaskRunLog_SchemaDumpEnd {
	if x != nil {
		return x.SchemaDumpEnd
	}
	return nil
}

func (x *TaskRunLog) GetCommandExecute() *TaskRunLog_CommandExecute {
	if x != nil {
		return x.CommandExecute
	}
	return nil
}

func (x *TaskRunLog) GetCommandResponse() *TaskRunLog_CommandResponse {
	if x != nil {
		return x.CommandResponse
	}
	return nil
}

type TaskRunLog_SchemaDumpStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskRunLog_SchemaDumpStart) Reset() {
	*x = TaskRunLog_SchemaDumpStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_SchemaDumpStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_SchemaDumpStart) ProtoMessage() {}

func (x *TaskRunLog_SchemaDumpStart) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_SchemaDumpStart.ProtoReflect.Descriptor instead.
func (*TaskRunLog_SchemaDumpStart) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 0}
}

type TaskRunLog_SchemaDumpEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TaskRunLog_SchemaDumpEnd) Reset() {
	*x = TaskRunLog_SchemaDumpEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_SchemaDumpEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_SchemaDumpEnd) ProtoMessage() {}

func (x *TaskRunLog_SchemaDumpEnd) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_SchemaDumpEnd.ProtoReflect.Descriptor instead.
func (*TaskRunLog_SchemaDumpEnd) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TaskRunLog_SchemaDumpEnd) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type TaskRunLog_CommandExecute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Executed commands are in range [command_index, command_index + command_count).
	CommandIndex int32 `protobuf:"varint,1,opt,name=command_index,json=commandIndex,proto3" json:"command_index,omitempty"`
	CommandCount int32 `protobuf:"varint,2,opt,name=command_count,json=commandCount,proto3" json:"command_count,omitempty"`
}

func (x *TaskRunLog_CommandExecute) Reset() {
	*x = TaskRunLog_CommandExecute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_CommandExecute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_CommandExecute) ProtoMessage() {}

func (x *TaskRunLog_CommandExecute) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_CommandExecute.ProtoReflect.Descriptor instead.
func (*TaskRunLog_CommandExecute) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 2}
}

func (x *TaskRunLog_CommandExecute) GetCommandIndex() int32 {
	if x != nil {
		return x.CommandIndex
	}
	return 0
}

func (x *TaskRunLog_CommandExecute) GetCommandCount() int32 {
	if x != nil {
		return x.CommandCount
	}
	return 0
}

type TaskRunLog_CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Executed commands are in range [command_index, command_index + command_count).
	CommandIndex int32  `protobuf:"varint,1,opt,name=command_index,json=commandIndex,proto3" json:"command_index,omitempty"`
	CommandCount int32  `protobuf:"varint,2,opt,name=command_count,json=commandCount,proto3" json:"command_count,omitempty"`
	Error        string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	AffectedRows int32  `protobuf:"varint,4,opt,name=affected_rows,json=affectedRows,proto3" json:"affected_rows,omitempty"`
}

func (x *TaskRunLog_CommandResponse) Reset() {
	*x = TaskRunLog_CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_task_run_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunLog_CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunLog_CommandResponse) ProtoMessage() {}

func (x *TaskRunLog_CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_task_run_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunLog_CommandResponse.ProtoReflect.Descriptor instead.
func (*TaskRunLog_CommandResponse) Descriptor() ([]byte, []int) {
	return file_store_task_run_log_proto_rawDescGZIP(), []int{0, 3}
}

func (x *TaskRunLog_CommandResponse) GetCommandIndex() int32 {
	if x != nil {
		return x.CommandIndex
	}
	return 0
}

func (x *TaskRunLog_CommandResponse) GetCommandCount() int32 {
	if x != nil {
		return x.CommandCount
	}
	return 0
}

func (x *TaskRunLog_CommandResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *TaskRunLog_CommandResponse) GetAffectedRows() int32 {
	if x != nil {
		return x.AffectedRows
	}
	return 0
}

var File_store_task_run_log_proto protoreflect.FileDescriptor

var file_store_task_run_log_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xba, 0x06, 0x0a, 0x0a, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e,
	0x4c, 0x6f, 0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x56,
	0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x75, 0x6d, 0x70,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x75, 0x6d,
	0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x50, 0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x5f, 0x64, 0x75, 0x6d, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x44, 0x75, 0x6d, 0x70, 0x45, 0x6e, 0x64, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x44, 0x75, 0x6d, 0x70, 0x45, 0x6e, 0x64, 0x12, 0x52, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x55, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x4c,
	0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x11, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x75, 0x6d,
	0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x1a, 0x25, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x44, 0x75, 0x6d, 0x70, 0x45, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x5a, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x96, 0x01, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f,
	0x77, 0x73, 0x22, 0x73, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x5f, 0x44, 0x55, 0x4d, 0x50, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x43, 0x48, 0x45, 0x4d,
	0x41, 0x5f, 0x44, 0x55, 0x4d, 0x50, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10,
	0x03, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x04, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_task_run_log_proto_rawDescOnce sync.Once
	file_store_task_run_log_proto_rawDescData = file_store_task_run_log_proto_rawDesc
)

func file_store_task_run_log_proto_rawDescGZIP() []byte {
	file_store_task_run_log_proto_rawDescOnce.Do(func() {
		file_store_task_run_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_task_run_log_proto_rawDescData)
	})
	return file_store_task_run_log_proto_rawDescData
}

var file_store_task_run_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_task_run_log_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_store_task_run_log_proto_goTypes = []interface{}{
	(TaskRunLog_Type)(0),               // 0: bytebase.store.TaskRunLog.Type
	(*TaskRunLog)(nil),                 // 1: bytebase.store.TaskRunLog
	(*TaskRunLog_SchemaDumpStart)(nil), // 2: bytebase.store.TaskRunLog.SchemaDumpStart
	(*TaskRunLog_SchemaDumpEnd)(nil),   // 3: bytebase.store.TaskRunLog.SchemaDumpEnd
	(*TaskRunLog_CommandExecute)(nil),  // 4: bytebase.store.TaskRunLog.CommandExecute
	(*TaskRunLog_CommandResponse)(nil), // 5: bytebase.store.TaskRunLog.CommandResponse
}
var file_store_task_run_log_proto_depIdxs = []int32{
	0, // 0: bytebase.store.TaskRunLog.type:type_name -> bytebase.store.TaskRunLog.Type
	2, // 1: bytebase.store.TaskRunLog.schema_dump_start:type_name -> bytebase.store.TaskRunLog.SchemaDumpStart
	3, // 2: bytebase.store.TaskRunLog.schema_dump_end:type_name -> bytebase.store.TaskRunLog.SchemaDumpEnd
	4, // 3: bytebase.store.TaskRunLog.command_execute:type_name -> bytebase.store.TaskRunLog.CommandExecute
	5, // 4: bytebase.store.TaskRunLog.command_response:type_name -> bytebase.store.TaskRunLog.CommandResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_store_task_run_log_proto_init() }
func file_store_task_run_log_proto_init() {
	if File_store_task_run_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_task_run_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRunLog); i {
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
		file_store_task_run_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRunLog_SchemaDumpStart); i {
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
		file_store_task_run_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRunLog_SchemaDumpEnd); i {
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
		file_store_task_run_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRunLog_CommandExecute); i {
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
		file_store_task_run_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRunLog_CommandResponse); i {
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
			RawDescriptor: file_store_task_run_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_task_run_log_proto_goTypes,
		DependencyIndexes: file_store_task_run_log_proto_depIdxs,
		EnumInfos:         file_store_task_run_log_proto_enumTypes,
		MessageInfos:      file_store_task_run_log_proto_msgTypes,
	}.Build()
	File_store_task_run_log_proto = out.File
	file_store_task_run_log_proto_rawDesc = nil
	file_store_task_run_log_proto_goTypes = nil
	file_store_task_run_log_proto_depIdxs = nil
}