// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: echo/v1/echo.proto

package v1

import (
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

// SayReq is the request message for the Echo.Say method.
type SayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 提交内容
	// v:  required
	// eg: hello world
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// only comment,not rule
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Sex      string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"` // tail comment
	// 结构体调用
	Data *SayRes `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// map 调用
	MapData map[string]*SayRes `protobuf:"bytes,5,rep,name=map_data,json=mapData,proto3" json:"map_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 数组调用
	ArrayData []*SayRes `protobuf:"bytes,6,rep,name=array_data,json=arrayData,proto3" json:"array_data,omitempty"`
	// int
	IntData []int32 `protobuf:"varint,7,rep,packed,name=int_data,json=intData,proto3" json:"int_data,omitempty"`
	// uint32
	Uint32Data []uint32 `protobuf:"varint,8,rep,packed,name=uint32_data,json=uint32Data,proto3" json:"uint32_data,omitempty"`
	// uint64
	// eg: 0
	// v: required
	// json: uint64_data
	Uint64Data int64 `protobuf:"varint,9,opt,name=uint64_data,json=uint64Data,proto3" json:"uint64_data,omitempty"`
}

func (x *SayReq) Reset() {
	*x = SayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayReq) ProtoMessage() {}

func (x *SayReq) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayReq.ProtoReflect.Descriptor instead.
func (*SayReq) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{0}
}

func (x *SayReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SayReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SayReq) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *SayReq) GetData() *SayRes {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SayReq) GetMapData() map[string]*SayRes {
	if x != nil {
		return x.MapData
	}
	return nil
}

func (x *SayReq) GetArrayData() []*SayRes {
	if x != nil {
		return x.ArrayData
	}
	return nil
}

func (x *SayReq) GetIntData() []int32 {
	if x != nil {
		return x.IntData
	}
	return nil
}

func (x *SayReq) GetUint32Data() []uint32 {
	if x != nil {
		return x.Uint32Data
	}
	return nil
}

func (x *SayReq) GetUint64Data() int64 {
	if x != nil {
		return x.Uint64Data
	}
	return 0
}

// SayRes is the response message for the Echo.Say method.
type SayRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SayRes) Reset() {
	*x = SayRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_v1_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayRes) ProtoMessage() {}

func (x *SayRes) ProtoReflect() protoreflect.Message {
	mi := &file_echo_v1_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayRes.ProtoReflect.Descriptor instead.
func (*SayRes) Descriptor() ([]byte, []int) {
	return file_echo_v1_echo_proto_rawDescGZIP(), []int{1}
}

func (x *SayRes) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_echo_v1_echo_proto protoreflect.FileDescriptor

var file_echo_v1_echo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x03, 0x0a,
	0x06, 0x53, 0x61, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12,
	0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52, 0x65, 0x73, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x4d, 0x61, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2f, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x52, 0x09, 0x61, 0x72, 0x72, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x75,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x4c, 0x0a,
	0x0c, 0x4d, 0x61, 0x70, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x06, 0x53,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32,
	0x4a, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x42, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x63, 0x68, 0x6f, 0x2f, 0x73, 0x61, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0x0a, 0x5a, 0x08, 0x2f,
	0x65, 0x63, 0x68, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_v1_echo_proto_rawDescOnce sync.Once
	file_echo_v1_echo_proto_rawDescData = file_echo_v1_echo_proto_rawDesc
)

func file_echo_v1_echo_proto_rawDescGZIP() []byte {
	file_echo_v1_echo_proto_rawDescOnce.Do(func() {
		file_echo_v1_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_v1_echo_proto_rawDescData)
	})
	return file_echo_v1_echo_proto_rawDescData
}

var file_echo_v1_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_echo_v1_echo_proto_goTypes = []interface{}{
	(*SayReq)(nil), // 0: proto.v1.SayReq
	(*SayRes)(nil), // 1: proto.v1.SayRes
	nil,            // 2: proto.v1.SayReq.MapDataEntry
}
var file_echo_v1_echo_proto_depIdxs = []int32{
	1, // 0: proto.v1.SayReq.data:type_name -> proto.v1.SayRes
	2, // 1: proto.v1.SayReq.map_data:type_name -> proto.v1.SayReq.MapDataEntry
	1, // 2: proto.v1.SayReq.array_data:type_name -> proto.v1.SayRes
	1, // 3: proto.v1.SayReq.MapDataEntry.value:type_name -> proto.v1.SayRes
	0, // 4: proto.v1.Echo.Say:input_type -> proto.v1.SayReq
	1, // 5: proto.v1.Echo.Say:output_type -> proto.v1.SayRes
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_echo_v1_echo_proto_init() }
func file_echo_v1_echo_proto_init() {
	if File_echo_v1_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_v1_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayReq); i {
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
		file_echo_v1_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayRes); i {
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
			RawDescriptor: file_echo_v1_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_v1_echo_proto_goTypes,
		DependencyIndexes: file_echo_v1_echo_proto_depIdxs,
		MessageInfos:      file_echo_v1_echo_proto_msgTypes,
	}.Build()
	File_echo_v1_echo_proto = out.File
	file_echo_v1_echo_proto_rawDesc = nil
	file_echo_v1_echo_proto_goTypes = nil
	file_echo_v1_echo_proto_depIdxs = nil
}
