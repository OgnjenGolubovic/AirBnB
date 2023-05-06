// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: accomodation_service/accomodation_service.proto

package accomodation

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
	AccomodationService_Get_FullMethodName = "/accomodation.AccomodationService/Get"
)

// AccomodationServiceClient is the client API for AccomodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccomodationServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type accomodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccomodationServiceClient(cc grpc.ClientConnInterface) AccomodationServiceClient {
	return &accomodationServiceClient{cc}
}

func (c *accomodationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, AccomodationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccomodationServiceServer is the server API for AccomodationService service.
// All implementations must embed UnimplementedAccomodationServiceServer
// for forward compatibility
type AccomodationServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedAccomodationServiceServer()
}

// UnimplementedAccomodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccomodationServiceServer struct {
}

func (UnimplementedAccomodationServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccomodationServiceServer) mustEmbedUnimplementedAccomodationServiceServer() {}

// UnsafeAccomodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccomodationServiceServer will
// result in compilation errors.
type UnsafeAccomodationServiceServer interface {
	mustEmbedUnimplementedAccomodationServiceServer()
}

func RegisterAccomodationServiceServer(s grpc.ServiceRegistrar, srv AccomodationServiceServer) {
	s.RegisterService(&AccomodationService_ServiceDesc, srv)
}

func _AccomodationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccomodationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccomodationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccomodationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccomodationService_ServiceDesc is the grpc.ServiceDesc for AccomodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccomodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accomodation.AccomodationService",
	HandlerType: (*AccomodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AccomodationService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accomodation_service/accomodation_service.proto",
}
