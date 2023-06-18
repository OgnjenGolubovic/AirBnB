// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: accommodation_service/accommodation_service.proto

package accommodation_service

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
	AccommodationService_Get_FullMethodName             = "/accommodation.AccommodationService/Get"
	AccommodationService_GetAllFreeDates_FullMethodName = "/accommodation.AccommodationService/GetAllFreeDates"
	AccommodationService_GetAll_FullMethodName          = "/accommodation.AccommodationService/GetAll"
	AccommodationService_Create_FullMethodName          = "/accommodation.AccommodationService/Create"
	AccommodationService_AddFreeDates_FullMethodName    = "/accommodation.AccommodationService/AddFreeDates"
	AccommodationService_RemoveFreeDates_FullMethodName = "/accommodation.AccommodationService/RemoveFreeDates"
	AccommodationService_UpdatePrice_FullMethodName     = "/accommodation.AccommodationService/UpdatePrice"
)

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAllFreeDates(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DateResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	AddFreeDates(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*DateResponse, error)
	RemoveFreeDates(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*DateResponse, error)
	UpdatePrice(ctx context.Context, in *UpdatePriceRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, AccommodationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllFreeDates(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*DateResponse, error) {
	out := new(DateResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAllFreeDates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, AccommodationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) AddFreeDates(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*DateResponse, error) {
	out := new(DateResponse)
	err := c.cc.Invoke(ctx, AccommodationService_AddFreeDates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) RemoveFreeDates(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*DateResponse, error) {
	out := new(DateResponse)
	err := c.cc.Invoke(ctx, AccommodationService_RemoveFreeDates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) UpdatePrice(ctx context.Context, in *UpdatePriceRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, AccommodationService_UpdatePrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAllFreeDates(context.Context, *GetRequest) (*DateResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	AddFreeDates(context.Context, *DateRequest) (*DateResponse, error)
	RemoveFreeDates(context.Context, *DateRequest) (*DateResponse, error)
	UpdatePrice(context.Context, *UpdatePriceRequest) (*CreateResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAllFreeDates(context.Context, *GetRequest) (*DateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFreeDates not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAccommodationServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccommodationServiceServer) AddFreeDates(context.Context, *DateRequest) (*DateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFreeDates not implemented")
}
func (UnimplementedAccommodationServiceServer) RemoveFreeDates(context.Context, *DateRequest) (*DateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFreeDates not implemented")
}
func (UnimplementedAccommodationServiceServer) UpdatePrice(context.Context, *UpdatePriceRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrice not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllFreeDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllFreeDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAllFreeDates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllFreeDates(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_AddFreeDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).AddFreeDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_AddFreeDates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).AddFreeDates(ctx, req.(*DateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_RemoveFreeDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).RemoveFreeDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_RemoveFreeDates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).RemoveFreeDates(ctx, req.(*DateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_UpdatePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).UpdatePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_UpdatePrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).UpdatePrice(ctx, req.(*UpdatePriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AccommodationService_Get_Handler,
		},
		{
			MethodName: "GetAllFreeDates",
			Handler:    _AccommodationService_GetAllFreeDates_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AccommodationService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccommodationService_Create_Handler,
		},
		{
			MethodName: "AddFreeDates",
			Handler:    _AccommodationService_AddFreeDates_Handler,
		},
		{
			MethodName: "RemoveFreeDates",
			Handler:    _AccommodationService_RemoveFreeDates_Handler,
		},
		{
			MethodName: "UpdatePrice",
			Handler:    _AccommodationService_UpdatePrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_service/accommodation_service.proto",
}
