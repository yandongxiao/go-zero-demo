// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: transform.proto

package transform

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

// TransformerClient is the client API for Transformer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransformerClient interface {
	Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error)
	Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error)
}

type transformerClient struct {
	cc grpc.ClientConnInterface
}

func NewTransformerClient(cc grpc.ClientConnInterface) TransformerClient {
	return &transformerClient{cc}
}

func (c *transformerClient) Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error) {
	out := new(ExpandResp)
	err := c.cc.Invoke(ctx, "/transform.transformer/expand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transformerClient) Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error) {
	out := new(ShortenResp)
	err := c.cc.Invoke(ctx, "/transform.transformer/shorten", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransformerServer is the server API for Transformer service.
// All implementations must embed UnimplementedTransformerServer
// for forward compatibility
type TransformerServer interface {
	Expand(context.Context, *ExpandReq) (*ExpandResp, error)
	Shorten(context.Context, *ShortenReq) (*ShortenResp, error)
	mustEmbedUnimplementedTransformerServer()
}

// UnimplementedTransformerServer must be embedded to have forward compatible implementations.
type UnimplementedTransformerServer struct {
}

func (UnimplementedTransformerServer) Expand(context.Context, *ExpandReq) (*ExpandResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expand not implemented")
}
func (UnimplementedTransformerServer) Shorten(context.Context, *ShortenReq) (*ShortenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shorten not implemented")
}
func (UnimplementedTransformerServer) mustEmbedUnimplementedTransformerServer() {}

// UnsafeTransformerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransformerServer will
// result in compilation errors.
type UnsafeTransformerServer interface {
	mustEmbedUnimplementedTransformerServer()
}

func RegisterTransformerServer(s grpc.ServiceRegistrar, srv TransformerServer) {
	s.RegisterService(&Transformer_ServiceDesc, srv)
}

func _Transformer_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.transformer/expand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServer).Expand(ctx, req.(*ExpandReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transformer_Shorten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServer).Shorten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.transformer/shorten",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServer).Shorten(ctx, req.(*ShortenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Transformer_ServiceDesc is the grpc.ServiceDesc for Transformer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transformer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transform.transformer",
	HandlerType: (*TransformerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "expand",
			Handler:    _Transformer_Expand_Handler,
		},
		{
			MethodName: "shorten",
			Handler:    _Transformer_Shorten_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transform.proto",
}
