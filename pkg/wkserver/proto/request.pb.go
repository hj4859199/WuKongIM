// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: pkg/wkserver/proto/request.proto

package proto

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

type Status int32

const (
	Status_ERROR    Status = 0
	Status_OK       Status = 1
	Status_NotFound Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "ERROR",
		1: "OK",
		2: "NotFound",
	}
	Status_value = map[string]int32{
		"ERROR":    0,
		"OK":       1,
		"NotFound": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_wkserver_proto_request_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_wkserver_proto_request_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_wkserver_proto_request_proto_rawDescGZIP(), []int{0}
}

type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid   string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Body  []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wkserver_proto_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wkserver_proto_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_pkg_wkserver_proto_request_proto_rawDescGZIP(), []int{0}
}

func (x *Connect) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Connect) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Connect) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Connect) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Connack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Connack) Reset() {
	*x = Connack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wkserver_proto_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connack) ProtoMessage() {}

func (x *Connack) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wkserver_proto_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connack.ProtoReflect.Descriptor instead.
func (*Connack) Descriptor() ([]byte, []int) {
	return file_pkg_wkserver_proto_request_proto_rawDescGZIP(), []int{1}
}

func (x *Connack) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Connack) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *Connack) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Body []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wkserver_proto_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wkserver_proto_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_pkg_wkserver_proto_request_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Request) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status Status `protobuf:"varint,2,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_wkserver_proto_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_wkserver_proto_request_proto_msgTypes[3]
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
	return file_pkg_wkserver_proto_request_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Response) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_ERROR
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_pkg_wkserver_proto_request_proto protoreflect.FileDescriptor

var file_pkg_wkserver_proto_request_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x54, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x41, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x55, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x2a, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_wkserver_proto_request_proto_rawDescOnce sync.Once
	file_pkg_wkserver_proto_request_proto_rawDescData = file_pkg_wkserver_proto_request_proto_rawDesc
)

func file_pkg_wkserver_proto_request_proto_rawDescGZIP() []byte {
	file_pkg_wkserver_proto_request_proto_rawDescOnce.Do(func() {
		file_pkg_wkserver_proto_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_wkserver_proto_request_proto_rawDescData)
	})
	return file_pkg_wkserver_proto_request_proto_rawDescData
}

var file_pkg_wkserver_proto_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_wkserver_proto_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_wkserver_proto_request_proto_goTypes = []interface{}{
	(Status)(0),      // 0: proto.Status
	(*Connect)(nil),  // 1: proto.Connect
	(*Connack)(nil),  // 2: proto.Connack
	(*Request)(nil),  // 3: proto.Request
	(*Response)(nil), // 4: proto.Response
}
var file_pkg_wkserver_proto_request_proto_depIdxs = []int32{
	0, // 0: proto.Connack.status:type_name -> proto.Status
	0, // 1: proto.Response.status:type_name -> proto.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_wkserver_proto_request_proto_init() }
func file_pkg_wkserver_proto_request_proto_init() {
	if File_pkg_wkserver_proto_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_wkserver_proto_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
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
		file_pkg_wkserver_proto_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connack); i {
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
		file_pkg_wkserver_proto_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_pkg_wkserver_proto_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pkg_wkserver_proto_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_wkserver_proto_request_proto_goTypes,
		DependencyIndexes: file_pkg_wkserver_proto_request_proto_depIdxs,
		EnumInfos:         file_pkg_wkserver_proto_request_proto_enumTypes,
		MessageInfos:      file_pkg_wkserver_proto_request_proto_msgTypes,
	}.Build()
	File_pkg_wkserver_proto_request_proto = out.File
	file_pkg_wkserver_proto_request_proto_rawDesc = nil
	file_pkg_wkserver_proto_request_proto_goTypes = nil
	file_pkg_wkserver_proto_request_proto_depIdxs = nil
}
