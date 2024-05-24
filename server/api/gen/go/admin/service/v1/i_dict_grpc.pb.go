// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: admin/service/v1/i_dict.proto

package servicev1

import (
	context "context"
	v11 "github.com/bugsmo/chunyu/api/gen/go/system/service/v1"
	v1 "github.com/bugsmo/cy/contrib/kratos/bootstrap/gen/api/go/pagination/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DictService_ListDict_FullMethodName   = "/admin.service.v1.DictService/ListDict"
	DictService_GetDict_FullMethodName    = "/admin.service.v1.DictService/GetDict"
	DictService_CreateDict_FullMethodName = "/admin.service.v1.DictService/CreateDict"
	DictService_UpdateDict_FullMethodName = "/admin.service.v1.DictService/UpdateDict"
	DictService_DeleteDict_FullMethodName = "/admin.service.v1.DictService/DeleteDict"
)

// DictServiceClient is the client API for DictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictServiceClient interface {
	// 查询字典列表
	ListDict(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListDictResponse, error)
	// 查询字典
	GetDict(ctx context.Context, in *v11.GetDictRequest, opts ...grpc.CallOption) (*v11.Dict, error)
	// 创建字典
	CreateDict(ctx context.Context, in *v11.CreateDictRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 更新字典
	UpdateDict(ctx context.Context, in *v11.UpdateDictRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 删除字典
	DeleteDict(ctx context.Context, in *v11.DeleteDictRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dictServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDictServiceClient(cc grpc.ClientConnInterface) DictServiceClient {
	return &dictServiceClient{cc}
}

func (c *dictServiceClient) ListDict(ctx context.Context, in *v1.PagingRequest, opts ...grpc.CallOption) (*v11.ListDictResponse, error) {
	out := new(v11.ListDictResponse)
	err := c.cc.Invoke(ctx, DictService_ListDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) GetDict(ctx context.Context, in *v11.GetDictRequest, opts ...grpc.CallOption) (*v11.Dict, error) {
	out := new(v11.Dict)
	err := c.cc.Invoke(ctx, DictService_GetDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) CreateDict(ctx context.Context, in *v11.CreateDictRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictService_CreateDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) UpdateDict(ctx context.Context, in *v11.UpdateDictRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictService_UpdateDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictServiceClient) DeleteDict(ctx context.Context, in *v11.DeleteDictRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DictService_DeleteDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictServiceServer is the server API for DictService service.
// All implementations must embed UnimplementedDictServiceServer
// for forward compatibility
type DictServiceServer interface {
	// 查询字典列表
	ListDict(context.Context, *v1.PagingRequest) (*v11.ListDictResponse, error)
	// 查询字典
	GetDict(context.Context, *v11.GetDictRequest) (*v11.Dict, error)
	// 创建字典
	CreateDict(context.Context, *v11.CreateDictRequest) (*emptypb.Empty, error)
	// 更新字典
	UpdateDict(context.Context, *v11.UpdateDictRequest) (*emptypb.Empty, error)
	// 删除字典
	DeleteDict(context.Context, *v11.DeleteDictRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDictServiceServer()
}

// UnimplementedDictServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDictServiceServer struct {
}

func (UnimplementedDictServiceServer) ListDict(context.Context, *v1.PagingRequest) (*v11.ListDictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDict not implemented")
}
func (UnimplementedDictServiceServer) GetDict(context.Context, *v11.GetDictRequest) (*v11.Dict, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDict not implemented")
}
func (UnimplementedDictServiceServer) CreateDict(context.Context, *v11.CreateDictRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDict not implemented")
}
func (UnimplementedDictServiceServer) UpdateDict(context.Context, *v11.UpdateDictRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDict not implemented")
}
func (UnimplementedDictServiceServer) DeleteDict(context.Context, *v11.DeleteDictRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDict not implemented")
}
func (UnimplementedDictServiceServer) mustEmbedUnimplementedDictServiceServer() {}

// UnsafeDictServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictServiceServer will
// result in compilation errors.
type UnsafeDictServiceServer interface {
	mustEmbedUnimplementedDictServiceServer()
}

func RegisterDictServiceServer(s grpc.ServiceRegistrar, srv DictServiceServer) {
	s.RegisterService(&DictService_ServiceDesc, srv)
}

func _DictService_ListDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).ListDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_ListDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).ListDict(ctx, req.(*v1.PagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_GetDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.GetDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).GetDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_GetDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).GetDict(ctx, req.(*v11.GetDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_CreateDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.CreateDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).CreateDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_CreateDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).CreateDict(ctx, req.(*v11.CreateDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_UpdateDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.UpdateDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).UpdateDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_UpdateDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).UpdateDict(ctx, req.(*v11.UpdateDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictService_DeleteDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.DeleteDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictServiceServer).DeleteDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictService_DeleteDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictServiceServer).DeleteDict(ctx, req.(*v11.DeleteDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DictService_ServiceDesc is the grpc.ServiceDesc for DictService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DictService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.service.v1.DictService",
	HandlerType: (*DictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDict",
			Handler:    _DictService_ListDict_Handler,
		},
		{
			MethodName: "GetDict",
			Handler:    _DictService_GetDict_Handler,
		},
		{
			MethodName: "CreateDict",
			Handler:    _DictService_CreateDict_Handler,
		},
		{
			MethodName: "UpdateDict",
			Handler:    _DictService_UpdateDict_Handler,
		},
		{
			MethodName: "DeleteDict",
			Handler:    _DictService_DeleteDict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/service/v1/i_dict.proto",
}