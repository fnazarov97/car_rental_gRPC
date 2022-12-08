// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: brand.proto

package brand_service

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

// BrandServiceClient is the client API for BrandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrandServiceClient interface {
	CreateBrand(ctx context.Context, in *CreateBrandRequest, opts ...grpc.CallOption) (*Brand, error)
	UpdateBrand(ctx context.Context, in *UpdateBrandRequest, opts ...grpc.CallOption) (*Brand, error)
	DeleteBrand(ctx context.Context, in *DeleteBrandRequest, opts ...grpc.CallOption) (*Brand, error)
	GetBrandByID(ctx context.Context, in *GetBrandByIDRequest, opts ...grpc.CallOption) (*Brand, error)
	GetBrandList(ctx context.Context, in *GetBrandListRequest, opts ...grpc.CallOption) (*GetBrandListResponse, error)
}

type brandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrandServiceClient(cc grpc.ClientConnInterface) BrandServiceClient {
	return &brandServiceClient{cc}
}

func (c *brandServiceClient) CreateBrand(ctx context.Context, in *CreateBrandRequest, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/genproto.BrandService/CreateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) UpdateBrand(ctx context.Context, in *UpdateBrandRequest, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/genproto.BrandService/UpdateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) DeleteBrand(ctx context.Context, in *DeleteBrandRequest, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/genproto.BrandService/DeleteBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) GetBrandByID(ctx context.Context, in *GetBrandByIDRequest, opts ...grpc.CallOption) (*Brand, error) {
	out := new(Brand)
	err := c.cc.Invoke(ctx, "/genproto.BrandService/GetBrandByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandServiceClient) GetBrandList(ctx context.Context, in *GetBrandListRequest, opts ...grpc.CallOption) (*GetBrandListResponse, error) {
	out := new(GetBrandListResponse)
	err := c.cc.Invoke(ctx, "/genproto.BrandService/GetBrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrandServiceServer is the server API for BrandService service.
// All implementations must embed UnimplementedBrandServiceServer
// for forward compatibility
type BrandServiceServer interface {
	CreateBrand(context.Context, *CreateBrandRequest) (*Brand, error)
	UpdateBrand(context.Context, *UpdateBrandRequest) (*Brand, error)
	DeleteBrand(context.Context, *DeleteBrandRequest) (*Brand, error)
	GetBrandByID(context.Context, *GetBrandByIDRequest) (*Brand, error)
	GetBrandList(context.Context, *GetBrandListRequest) (*GetBrandListResponse, error)
	mustEmbedUnimplementedBrandServiceServer()
}

// UnimplementedBrandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrandServiceServer struct {
}

func (UnimplementedBrandServiceServer) CreateBrand(context.Context, *CreateBrandRequest) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (UnimplementedBrandServiceServer) UpdateBrand(context.Context, *UpdateBrandRequest) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (UnimplementedBrandServiceServer) DeleteBrand(context.Context, *DeleteBrandRequest) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
func (UnimplementedBrandServiceServer) GetBrandByID(context.Context, *GetBrandByIDRequest) (*Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandByID not implemented")
}
func (UnimplementedBrandServiceServer) GetBrandList(context.Context, *GetBrandListRequest) (*GetBrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandList not implemented")
}
func (UnimplementedBrandServiceServer) mustEmbedUnimplementedBrandServiceServer() {}

// UnsafeBrandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrandServiceServer will
// result in compilation errors.
type UnsafeBrandServiceServer interface {
	mustEmbedUnimplementedBrandServiceServer()
}

func RegisterBrandServiceServer(s grpc.ServiceRegistrar, srv BrandServiceServer) {
	s.RegisterService(&BrandService_ServiceDesc, srv)
}

func _BrandService_CreateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).CreateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BrandService/CreateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).CreateBrand(ctx, req.(*CreateBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_UpdateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).UpdateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BrandService/UpdateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).UpdateBrand(ctx, req.(*UpdateBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BrandService/DeleteBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).DeleteBrand(ctx, req.(*DeleteBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_GetBrandByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrandByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).GetBrandByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BrandService/GetBrandByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).GetBrandByID(ctx, req.(*GetBrandByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrandService_GetBrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrandListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServiceServer).GetBrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.BrandService/GetBrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServiceServer).GetBrandList(ctx, req.(*GetBrandListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BrandService_ServiceDesc is the grpc.ServiceDesc for BrandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.BrandService",
	HandlerType: (*BrandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBrand",
			Handler:    _BrandService_CreateBrand_Handler,
		},
		{
			MethodName: "UpdateBrand",
			Handler:    _BrandService_UpdateBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _BrandService_DeleteBrand_Handler,
		},
		{
			MethodName: "GetBrandByID",
			Handler:    _BrandService_GetBrandByID_Handler,
		},
		{
			MethodName: "GetBrandList",
			Handler:    _BrandService_GetBrandList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "brand.proto",
}