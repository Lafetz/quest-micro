// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: proto/knight/knight.proto

package protoknight

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
	KnightService_AddKnight_FullMethodName       = "/protoknight.KnightService/AddKnight"
	KnightService_GetKnightStatus_FullMethodName = "/protoknight.KnightService/GetKnightStatus"
	KnightService_UpdateStatus_FullMethodName    = "/protoknight.KnightService/UpdateStatus"
	KnightService_GetKnights_FullMethodName      = "/protoknight.KnightService/GetKnights"
)

// KnightServiceClient is the client API for KnightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KnightServiceClient interface {
	AddKnight(ctx context.Context, in *AddKnightReq, opts ...grpc.CallOption) (*AddKnightRes, error)
	GetKnightStatus(ctx context.Context, in *KnightStatusReq, opts ...grpc.CallOption) (*KnightStatusRes, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*UpdateStatusRes, error)
	GetKnights(ctx context.Context, in *GetKnightsReq, opts ...grpc.CallOption) (*UpdateStatusRes, error)
}

type knightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKnightServiceClient(cc grpc.ClientConnInterface) KnightServiceClient {
	return &knightServiceClient{cc}
}

func (c *knightServiceClient) AddKnight(ctx context.Context, in *AddKnightReq, opts ...grpc.CallOption) (*AddKnightRes, error) {
	out := new(AddKnightRes)
	err := c.cc.Invoke(ctx, KnightService_AddKnight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knightServiceClient) GetKnightStatus(ctx context.Context, in *KnightStatusReq, opts ...grpc.CallOption) (*KnightStatusRes, error) {
	out := new(KnightStatusRes)
	err := c.cc.Invoke(ctx, KnightService_GetKnightStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knightServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*UpdateStatusRes, error) {
	out := new(UpdateStatusRes)
	err := c.cc.Invoke(ctx, KnightService_UpdateStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knightServiceClient) GetKnights(ctx context.Context, in *GetKnightsReq, opts ...grpc.CallOption) (*UpdateStatusRes, error) {
	out := new(UpdateStatusRes)
	err := c.cc.Invoke(ctx, KnightService_GetKnights_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KnightServiceServer is the server API for KnightService service.
// All implementations must embed UnimplementedKnightServiceServer
// for forward compatibility
type KnightServiceServer interface {
	AddKnight(context.Context, *AddKnightReq) (*AddKnightRes, error)
	GetKnightStatus(context.Context, *KnightStatusReq) (*KnightStatusRes, error)
	UpdateStatus(context.Context, *UpdateStatusReq) (*UpdateStatusRes, error)
	GetKnights(context.Context, *GetKnightsReq) (*UpdateStatusRes, error)
	mustEmbedUnimplementedKnightServiceServer()
}

// UnimplementedKnightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKnightServiceServer struct {
}

func (UnimplementedKnightServiceServer) AddKnight(context.Context, *AddKnightReq) (*AddKnightRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKnight not implemented")
}
func (UnimplementedKnightServiceServer) GetKnightStatus(context.Context, *KnightStatusReq) (*KnightStatusRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKnightStatus not implemented")
}
func (UnimplementedKnightServiceServer) UpdateStatus(context.Context, *UpdateStatusReq) (*UpdateStatusRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedKnightServiceServer) GetKnights(context.Context, *GetKnightsReq) (*UpdateStatusRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKnights not implemented")
}
func (UnimplementedKnightServiceServer) mustEmbedUnimplementedKnightServiceServer() {}

// UnsafeKnightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KnightServiceServer will
// result in compilation errors.
type UnsafeKnightServiceServer interface {
	mustEmbedUnimplementedKnightServiceServer()
}

func RegisterKnightServiceServer(s grpc.ServiceRegistrar, srv KnightServiceServer) {
	s.RegisterService(&KnightService_ServiceDesc, srv)
}

func _KnightService_AddKnight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKnightReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnightServiceServer).AddKnight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnightService_AddKnight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnightServiceServer).AddKnight(ctx, req.(*AddKnightReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnightService_GetKnightStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KnightStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnightServiceServer).GetKnightStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnightService_GetKnightStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnightServiceServer).GetKnightStatus(ctx, req.(*KnightStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnightService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnightServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnightService_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnightServiceServer).UpdateStatus(ctx, req.(*UpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnightService_GetKnights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKnightsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnightServiceServer).GetKnights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KnightService_GetKnights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnightServiceServer).GetKnights(ctx, req.(*GetKnightsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// KnightService_ServiceDesc is the grpc.ServiceDesc for KnightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KnightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoknight.KnightService",
	HandlerType: (*KnightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddKnight",
			Handler:    _KnightService_AddKnight_Handler,
		},
		{
			MethodName: "GetKnightStatus",
			Handler:    _KnightService_GetKnightStatus_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _KnightService_UpdateStatus_Handler,
		},
		{
			MethodName: "GetKnights",
			Handler:    _KnightService_GetKnights_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/knight/knight.proto",
}
