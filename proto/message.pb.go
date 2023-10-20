// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: message.proto

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

type Registro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre   string `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,3,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *Registro) Reset() {
	*x = Registro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registro) ProtoMessage() {}

func (x *Registro) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registro.ProtoReflect.Descriptor instead.
func (*Registro) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Registro) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Registro) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Registro) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type Datos_DataNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *Datos_DataNode) Reset() {
	*x = Datos_DataNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datos_DataNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datos_DataNode) ProtoMessage() {}

func (x *Datos_DataNode) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datos_DataNode.ProtoReflect.Descriptor instead.
func (*Datos_DataNode) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Datos_DataNode) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Datos_DataNode) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type Lista_Datos_DataNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListaDatos_DataNode []*Datos_DataNode `protobuf:"bytes,1,rep,name=lista_datos_DataNode,json=listaDatosDataNode,proto3" json:"lista_datos_DataNode,omitempty"`
}

func (x *Lista_Datos_DataNode) Reset() {
	*x = Lista_Datos_DataNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lista_Datos_DataNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lista_Datos_DataNode) ProtoMessage() {}

func (x *Lista_Datos_DataNode) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lista_Datos_DataNode.ProtoReflect.Descriptor instead.
func (*Lista_Datos_DataNode) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Lista_Datos_DataNode) GetListaDatos_DataNode() []*Datos_DataNode {
	if x != nil {
		return x.ListaDatos_DataNode
	}
	return nil
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListaId []string `protobuf:"bytes,1,rep,name=lista_id,json=listaId,proto3" json:"lista_id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetListaId() []string {
	if x != nil {
		return x.ListaId
	}
	return nil
}

type Recepcion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok string `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Recepcion) Reset() {
	*x = Recepcion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recepcion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recepcion) ProtoMessage() {}

func (x *Recepcion) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recepcion.ProtoReflect.Descriptor instead.
func (*Recepcion) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *Recepcion) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x4e, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65,
	0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65,
	0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x22, 0x44, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x22, 0x5e, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x61, 0x5f, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x46, 0x0a, 0x14, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74,
	0x6f, 0x73, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x44,
	0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x6f, 0x73, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x1f, 0x0a, 0x02, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x09,
	0x52, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x81, 0x01, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x72, 0x4e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x17, 0x53, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x5f, 0x44, 0x61, 0x74,
	0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x1a,
	0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x5f, 0x44, 0x61, 0x74,
	0x6f, 0x73, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x64, 0x6f, 0x4a,
	0x61, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(*Registro)(nil),             // 0: grpc.Registro
	(*Datos_DataNode)(nil),       // 1: grpc.Datos_DataNode
	(*Lista_Datos_DataNode)(nil), // 2: grpc.Lista_Datos_DataNode
	(*Id)(nil),                   // 3: grpc.Id
	(*Recepcion)(nil),            // 4: grpc.Recepcion
}
var file_message_proto_depIdxs = []int32{
	1, // 0: grpc.Lista_Datos_DataNode.lista_datos_DataNode:type_name -> grpc.Datos_DataNode
	0, // 1: grpc.DataNode.RegistrarNombre:input_type -> grpc.Registro
	3, // 2: grpc.DataNode.Solicitud_Info_DataNode:input_type -> grpc.Id
	4, // 3: grpc.DataNode.RegistrarNombre:output_type -> grpc.Recepcion
	2, // 4: grpc.DataNode.Solicitud_Info_DataNode:output_type -> grpc.Lista_Datos_DataNode
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registro); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datos_DataNode); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lista_Datos_DataNode); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recepcion); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
