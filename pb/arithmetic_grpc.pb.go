// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: proto/arithmetic.proto

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

const (
	ArithmeticService_Compute_FullMethodName = "/ArithmeticService/Compute"
)

// ArithmeticServiceClient is the client API for ArithmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithmeticServiceClient interface {
	Compute(ctx context.Context, in *ArithmeticParamsRequest, opts ...grpc.CallOption) (*ArithmeticParamsResponse, error)
}

type arithmeticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticServiceClient(cc grpc.ClientConnInterface) ArithmeticServiceClient {
	return &arithmeticServiceClient{cc}
}

func (c *arithmeticServiceClient) Compute(ctx context.Context, in *ArithmeticParamsRequest, opts ...grpc.CallOption) (*ArithmeticParamsResponse, error) {
	out := new(ArithmeticParamsResponse)
	err := c.cc.Invoke(ctx, ArithmeticService_Compute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServiceServer is the server API for ArithmeticService service.
// All implementations must embed UnimplementedArithmeticServiceServer
// for forward compatibility
type ArithmeticServiceServer interface {
	Compute(context.Context, *ArithmeticParamsRequest) (*ArithmeticParamsResponse, error)
	mustEmbedUnimplementedArithmeticServiceServer()
}

// UnimplementedArithmeticServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArithmeticServiceServer struct {
}

func (UnimplementedArithmeticServiceServer) Compute(context.Context, *ArithmeticParamsRequest) (*ArithmeticParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
}
func (UnimplementedArithmeticServiceServer) mustEmbedUnimplementedArithmeticServiceServer() {}

// UnsafeArithmeticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithmeticServiceServer will
// result in compilation errors.
type UnsafeArithmeticServiceServer interface {
	mustEmbedUnimplementedArithmeticServiceServer()
}

func RegisterArithmeticServiceServer(s grpc.ServiceRegistrar, srv ArithmeticServiceServer) {
	s.RegisterService(&ArithmeticService_ServiceDesc, srv)
}

func _ArithmeticService_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithmeticParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArithmeticService_Compute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).Compute(ctx, req.(*ArithmeticParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArithmeticService_ServiceDesc is the grpc.ServiceDesc for ArithmeticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArithmeticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ArithmeticService",
	HandlerType: (*ArithmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _ArithmeticService_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/arithmetic.proto",
}