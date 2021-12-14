// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// AddCityClient is the client API for AddCity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddCityClient interface {
	AddCity(ctx context.Context, in *City, opts ...grpc.CallOption) (*Clock, error)
}

type addCityClient struct {
	cc grpc.ClientConnInterface
}

func NewAddCityClient(cc grpc.ClientConnInterface) AddCityClient {
	return &addCityClient{cc}
}

func (c *addCityClient) AddCity(ctx context.Context, in *City, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/AddCity/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddCityServer is the server API for AddCity service.
// All implementations must embed UnimplementedAddCityServer
// for forward compatibility
type AddCityServer interface {
	AddCity(context.Context, *City) (*Clock, error)
	mustEmbedUnimplementedAddCityServer()
}

// UnimplementedAddCityServer must be embedded to have forward compatible implementations.
type UnimplementedAddCityServer struct {
}

func (UnimplementedAddCityServer) AddCity(context.Context, *City) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedAddCityServer) mustEmbedUnimplementedAddCityServer() {}

// UnsafeAddCityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddCityServer will
// result in compilation errors.
type UnsafeAddCityServer interface {
	mustEmbedUnimplementedAddCityServer()
}

func RegisterAddCityServer(s grpc.ServiceRegistrar, srv AddCityServer) {
	s.RegisterService(&AddCity_ServiceDesc, srv)
}

func _AddCity_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(City)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddCityServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AddCity/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddCityServer).AddCity(ctx, req.(*City))
	}
	return interceptor(ctx, in, info, handler)
}

// AddCity_ServiceDesc is the grpc.ServiceDesc for AddCity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddCity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AddCity",
	HandlerType: (*AddCityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCity",
			Handler:    _AddCity_AddCity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Fulcrum-Informante.proto",
}

// DeleteCityClient is the client API for DeleteCity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeleteCityClient interface {
	DeleteCity(ctx context.Context, in *DelCt, opts ...grpc.CallOption) (*Clock, error)
}

type deleteCityClient struct {
	cc grpc.ClientConnInterface
}

func NewDeleteCityClient(cc grpc.ClientConnInterface) DeleteCityClient {
	return &deleteCityClient{cc}
}

func (c *deleteCityClient) DeleteCity(ctx context.Context, in *DelCt, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/DeleteCity/DeleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteCityServer is the server API for DeleteCity service.
// All implementations must embed UnimplementedDeleteCityServer
// for forward compatibility
type DeleteCityServer interface {
	DeleteCity(context.Context, *DelCt) (*Clock, error)
	mustEmbedUnimplementedDeleteCityServer()
}

// UnimplementedDeleteCityServer must be embedded to have forward compatible implementations.
type UnimplementedDeleteCityServer struct {
}

func (UnimplementedDeleteCityServer) DeleteCity(context.Context, *DelCt) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedDeleteCityServer) mustEmbedUnimplementedDeleteCityServer() {}

// UnsafeDeleteCityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeleteCityServer will
// result in compilation errors.
type UnsafeDeleteCityServer interface {
	mustEmbedUnimplementedDeleteCityServer()
}

func RegisterDeleteCityServer(s grpc.ServiceRegistrar, srv DeleteCityServer) {
	s.RegisterService(&DeleteCity_ServiceDesc, srv)
}

func _DeleteCity_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeleteCityServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DeleteCity/DeleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeleteCityServer).DeleteCity(ctx, req.(*DelCt))
	}
	return interceptor(ctx, in, info, handler)
}

// DeleteCity_ServiceDesc is the grpc.ServiceDesc for DeleteCity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeleteCity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DeleteCity",
	HandlerType: (*DeleteCityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteCity",
			Handler:    _DeleteCity_DeleteCity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Fulcrum-Informante.proto",
}

// GetNumberRebelsClient is the client API for GetNumberRebels service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetNumberRebelsClient interface {
	GetNumberRebels(ctx context.Context, in *DelCt, opts ...grpc.CallOption) (*NumClock, error)
}

type getNumberRebelsClient struct {
	cc grpc.ClientConnInterface
}

func NewGetNumberRebelsClient(cc grpc.ClientConnInterface) GetNumberRebelsClient {
	return &getNumberRebelsClient{cc}
}

func (c *getNumberRebelsClient) GetNumberRebels(ctx context.Context, in *DelCt, opts ...grpc.CallOption) (*NumClock, error) {
	out := new(NumClock)
	err := c.cc.Invoke(ctx, "/GetNumberRebels/GetNumberRebels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetNumberRebelsServer is the server API for GetNumberRebels service.
// All implementations must embed UnimplementedGetNumberRebelsServer
// for forward compatibility
type GetNumberRebelsServer interface {
	GetNumberRebels(context.Context, *DelCt) (*NumClock, error)
	mustEmbedUnimplementedGetNumberRebelsServer()
}

// UnimplementedGetNumberRebelsServer must be embedded to have forward compatible implementations.
type UnimplementedGetNumberRebelsServer struct {
}

func (UnimplementedGetNumberRebelsServer) GetNumberRebels(context.Context, *DelCt) (*NumClock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebels not implemented")
}
func (UnimplementedGetNumberRebelsServer) mustEmbedUnimplementedGetNumberRebelsServer() {}

// UnsafeGetNumberRebelsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetNumberRebelsServer will
// result in compilation errors.
type UnsafeGetNumberRebelsServer interface {
	mustEmbedUnimplementedGetNumberRebelsServer()
}

func RegisterGetNumberRebelsServer(s grpc.ServiceRegistrar, srv GetNumberRebelsServer) {
	s.RegisterService(&GetNumberRebels_ServiceDesc, srv)
}

func _GetNumberRebels_GetNumberRebels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetNumberRebelsServer).GetNumberRebels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GetNumberRebels/GetNumberRebels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetNumberRebelsServer).GetNumberRebels(ctx, req.(*DelCt))
	}
	return interceptor(ctx, in, info, handler)
}

// GetNumberRebels_ServiceDesc is the grpc.ServiceDesc for GetNumberRebels service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetNumberRebels_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GetNumberRebels",
	HandlerType: (*GetNumberRebelsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNumberRebels",
			Handler:    _GetNumberRebels_GetNumberRebels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Fulcrum-Informante.proto",
}

// UpdateNameClient is the client API for UpdateName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateNameClient interface {
	CambiarNombre(ctx context.Context, in *City, opts ...grpc.CallOption) (*Clock, error)
}

type updateNameClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateNameClient(cc grpc.ClientConnInterface) UpdateNameClient {
	return &updateNameClient{cc}
}

func (c *updateNameClient) CambiarNombre(ctx context.Context, in *City, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/UpdateName/CambiarNombre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateNameServer is the server API for UpdateName service.
// All implementations must embed UnimplementedUpdateNameServer
// for forward compatibility
type UpdateNameServer interface {
	CambiarNombre(context.Context, *City) (*Clock, error)
	mustEmbedUnimplementedUpdateNameServer()
}

// UnimplementedUpdateNameServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateNameServer struct {
}

func (UnimplementedUpdateNameServer) CambiarNombre(context.Context, *City) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CambiarNombre not implemented")
}
func (UnimplementedUpdateNameServer) mustEmbedUnimplementedUpdateNameServer() {}

// UnsafeUpdateNameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateNameServer will
// result in compilation errors.
type UnsafeUpdateNameServer interface {
	mustEmbedUnimplementedUpdateNameServer()
}

func RegisterUpdateNameServer(s grpc.ServiceRegistrar, srv UpdateNameServer) {
	s.RegisterService(&UpdateName_ServiceDesc, srv)
}

func _UpdateName_CambiarNombre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(City)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateNameServer).CambiarNombre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UpdateName/CambiarNombre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateNameServer).CambiarNombre(ctx, req.(*City))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdateName_ServiceDesc is the grpc.ServiceDesc for UpdateName service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateName_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UpdateName",
	HandlerType: (*UpdateNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CambiarNombre",
			Handler:    _UpdateName_CambiarNombre_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Fulcrum-Informante.proto",
}

// UpdateNumberClient is the client API for UpdateNumber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateNumberClient interface {
	CambiarValor(ctx context.Context, in *City, opts ...grpc.CallOption) (*Clock, error)
}

type updateNumberClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateNumberClient(cc grpc.ClientConnInterface) UpdateNumberClient {
	return &updateNumberClient{cc}
}

func (c *updateNumberClient) CambiarValor(ctx context.Context, in *City, opts ...grpc.CallOption) (*Clock, error) {
	out := new(Clock)
	err := c.cc.Invoke(ctx, "/UpdateNumber/CambiarValor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateNumberServer is the server API for UpdateNumber service.
// All implementations must embed UnimplementedUpdateNumberServer
// for forward compatibility
type UpdateNumberServer interface {
	CambiarValor(context.Context, *City) (*Clock, error)
	mustEmbedUnimplementedUpdateNumberServer()
}

// UnimplementedUpdateNumberServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateNumberServer struct {
}

func (UnimplementedUpdateNumberServer) CambiarValor(context.Context, *City) (*Clock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CambiarValor not implemented")
}
func (UnimplementedUpdateNumberServer) mustEmbedUnimplementedUpdateNumberServer() {}

// UnsafeUpdateNumberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateNumberServer will
// result in compilation errors.
type UnsafeUpdateNumberServer interface {
	mustEmbedUnimplementedUpdateNumberServer()
}

func RegisterUpdateNumberServer(s grpc.ServiceRegistrar, srv UpdateNumberServer) {
	s.RegisterService(&UpdateNumber_ServiceDesc, srv)
}

func _UpdateNumber_CambiarValor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(City)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateNumberServer).CambiarValor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UpdateNumber/CambiarValor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateNumberServer).CambiarValor(ctx, req.(*City))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdateNumber_ServiceDesc is the grpc.ServiceDesc for UpdateNumber service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateNumber_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UpdateNumber",
	HandlerType: (*UpdateNumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CambiarValor",
			Handler:    _UpdateNumber_CambiarValor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Fulcrum-Informante.proto",
}