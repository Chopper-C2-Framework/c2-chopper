// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: listener.proto

package proto

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

type ListenerListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool        `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Data    []*Listener `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ListenerListResponse) Reset() {
	*x = ListenerListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerListResponse) ProtoMessage() {}

func (x *ListenerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerListResponse.ProtoReflect.Descriptor instead.
func (*ListenerListResponse) Descriptor() ([]byte, []int) {
	return file_listener_proto_rawDescGZIP(), []int{0}
}

func (x *ListenerListResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ListenerListResponse) GetData() []*Listener {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetListenerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetListenerInfoRequest) Reset() {
	*x = GetListenerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListenerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListenerInfoRequest) ProtoMessage() {}

func (x *GetListenerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListenerInfoRequest.ProtoReflect.Descriptor instead.
func (*GetListenerInfoRequest) Descriptor() ([]byte, []int) {
	return file_listener_proto_rawDescGZIP(), []int{1}
}

func (x *GetListenerInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListenerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool      `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Data    *Listener `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetListenerInfoResponse) Reset() {
	*x = GetListenerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListenerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListenerInfoResponse) ProtoMessage() {}

func (x *GetListenerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListenerInfoResponse.ProtoReflect.Descriptor instead.
func (*GetListenerInfoResponse) Descriptor() ([]byte, []int) {
	return file_listener_proto_rawDescGZIP(), []int{2}
}

func (x *GetListenerInfoResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetListenerInfoResponse) GetData() *Listener {
	if x != nil {
		return x.Data
	}
	return nil
}

type ExecuteCmdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenerId string `protobuf:"bytes,1,opt,name=ListenerId,proto3" json:"ListenerId,omitempty"`
	Cmd        string `protobuf:"bytes,2,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
}

func (x *ExecuteCmdRequest) Reset() {
	*x = ExecuteCmdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCmdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCmdRequest) ProtoMessage() {}

func (x *ExecuteCmdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCmdRequest.ProtoReflect.Descriptor instead.
func (*ExecuteCmdRequest) Descriptor() ([]byte, []int) {
	return file_listener_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteCmdRequest) GetListenerId() string {
	if x != nil {
		return x.ListenerId
	}
	return ""
}

func (x *ExecuteCmdRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type ExecuteCmdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Result  string `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ExecuteCmdResponse) Reset() {
	*x = ExecuteCmdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteCmdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteCmdResponse) ProtoMessage() {}

func (x *ExecuteCmdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteCmdResponse.ProtoReflect.Descriptor instead.
func (*ExecuteCmdResponse) Descriptor() ([]byte, []int) {
	return file_listener_proto_rawDescGZIP(), []int{4}
}

func (x *ExecuteCmdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ExecuteCmdResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_listener_proto protoreflect.FileDescriptor

var file_listener_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a, 0x11, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x6d, 0x64,
	0x22, 0x46, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6d, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xd4, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x43, 0x6d, 0x64, 0x12, 0x12, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6d,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_listener_proto_rawDescOnce sync.Once
	file_listener_proto_rawDescData = file_listener_proto_rawDesc
)

func file_listener_proto_rawDescGZIP() []byte {
	file_listener_proto_rawDescOnce.Do(func() {
		file_listener_proto_rawDescData = protoimpl.X.CompressGZIP(file_listener_proto_rawDescData)
	})
	return file_listener_proto_rawDescData
}

var file_listener_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_listener_proto_goTypes = []interface{}{
	(*ListenerListResponse)(nil),    // 0: ListenerListResponse
	(*GetListenerInfoRequest)(nil),  // 1: GetListenerInfoRequest
	(*GetListenerInfoResponse)(nil), // 2: GetListenerInfoResponse
	(*ExecuteCmdRequest)(nil),       // 3: ExecuteCmdRequest
	(*ExecuteCmdResponse)(nil),      // 4: ExecuteCmdResponse
	(*Listener)(nil),                // 5: Listener
	(*emptypb.Empty)(nil),           // 6: google.protobuf.Empty
}
var file_listener_proto_depIdxs = []int32{
	5, // 0: ListenerListResponse.Data:type_name -> Listener
	5, // 1: GetListenerInfoResponse.Data:type_name -> Listener
	6, // 2: ListenerService.ListListeners:input_type -> google.protobuf.Empty
	1, // 3: ListenerService.GetListenerInfo:input_type -> GetListenerInfoRequest
	3, // 4: ListenerService.ExecuteCmd:input_type -> ExecuteCmdRequest
	0, // 5: ListenerService.ListListeners:output_type -> ListenerListResponse
	2, // 6: ListenerService.GetListenerInfo:output_type -> GetListenerInfoResponse
	4, // 7: ListenerService.ExecuteCmd:output_type -> ExecuteCmdResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_listener_proto_init() }
func file_listener_proto_init() {
	if File_listener_proto != nil {
		return
	}
	file_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_listener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerListResponse); i {
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
		file_listener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListenerInfoRequest); i {
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
		file_listener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListenerInfoResponse); i {
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
		file_listener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCmdRequest); i {
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
		file_listener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteCmdResponse); i {
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
			RawDescriptor: file_listener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listener_proto_goTypes,
		DependencyIndexes: file_listener_proto_depIdxs,
		MessageInfos:      file_listener_proto_msgTypes,
	}.Build()
	File_listener_proto = out.File
	file_listener_proto_rawDesc = nil
	file_listener_proto_goTypes = nil
	file_listener_proto_depIdxs = nil
}