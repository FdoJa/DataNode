// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: message.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DataNode_RegistrarNombre_FullMethodName         = "/grpc.DataNode/RegistrarNombre"
	DataNode_Solicitud_Info_DataNode_FullMethodName = "/grpc.DataNode/Solicitud_Info_DataNode"
)

// DataNodeClient is the client API for DataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodeClient interface {
	RegistrarNombre(ctx context.Context, in *Registro, opts ...grpc.CallOption) (*Recepcion, error)
	Solicitud_Info_DataNode(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Lista_Datos_DataNode, error)
}

type dataNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeClient(cc grpc.ClientConnInterface) DataNodeClient {
	return &dataNodeClient{cc}
}

func (c *dataNodeClient) RegistrarNombre(ctx context.Context, in *Registro, opts ...grpc.CallOption) (*Recepcion, error) {
	out := new(Recepcion)
	err := c.cc.Invoke(ctx, DataNode_RegistrarNombre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) Solicitud_Info_DataNode(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Lista_Datos_DataNode, error) {
	out := new(Lista_Datos_DataNode)
	err := c.cc.Invoke(ctx, DataNode_Solicitud_Info_DataNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServer is the server API for DataNode service.
// All implementations must embed UnimplementedDataNodeServer
// for forward compatibility
type DataNodeServer interface {
	RegistrarNombre(context.Context, *Registro) (*Recepcion, error)
	Solicitud_Info_DataNode(context.Context, *Id) (*Lista_Datos_DataNode, error)
	mustEmbedUnimplementedDataNodeServer()
}

// UnimplementedDataNodeServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodeServer struct {
}

func (UnimplementedDataNodeServer) RegistrarNombre(context.Context, *Registro) (*Recepcion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrarNombre not implemented")
}
func (UnimplementedDataNodeServer) Solicitud_Info_DataNode(context.Context, *Id) (*Lista_Datos_DataNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solicitud_Info_DataNode not implemented")
}
func (UnimplementedDataNodeServer) mustEmbedUnimplementedDataNodeServer() {}

// UnsafeDataNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServer will
// result in compilation errors.
type UnsafeDataNodeServer interface {
	mustEmbedUnimplementedDataNodeServer()
}

func RegisterDataNodeServer(s grpc.ServiceRegistrar, srv DataNodeServer) {
	s.RegisterService(&DataNode_ServiceDesc, srv)
}

func _DataNode_RegistrarNombre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Registro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).RegistrarNombre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataNode_RegistrarNombre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).RegistrarNombre(ctx, req.(*Registro))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNode_Solicitud_Info_DataNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).Solicitud_Info_DataNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataNode_Solicitud_Info_DataNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).Solicitud_Info_DataNode(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNode_ServiceDesc is the grpc.ServiceDesc for DataNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DataNode",
	HandlerType: (*DataNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistrarNombre",
			Handler:    _DataNode_RegistrarNombre_Handler,
		},
		{
			MethodName: "Solicitud_Info_DataNode",
			Handler:    _DataNode_Solicitud_Info_DataNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
