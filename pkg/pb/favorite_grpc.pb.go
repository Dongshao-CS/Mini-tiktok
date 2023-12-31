// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: favorite.proto

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
	FavoriteService_FavoriteAction_FullMethodName         = "/FavoriteService/FavoriteAction"
	FavoriteService_GetFavoriteVideoIdList_FullMethodName = "/FavoriteService/GetFavoriteVideoIdList"
	FavoriteService_IsFavoriteVideoDict_FullMethodName    = "/FavoriteService/IsFavoriteVideoDict"
)

// FavoriteServiceClient is the client API for FavoriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteServiceClient interface {
	FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionRsp, error)
	GetFavoriteVideoIdList(ctx context.Context, in *GetFavoriteVideoIdListReq, opts ...grpc.CallOption) (*GetFavoriteVideoIdListRsp, error)
	IsFavoriteVideoDict(ctx context.Context, in *IsFavoriteVideoDictReq, opts ...grpc.CallOption) (*IsFavoriteVideoDictRsp, error)
}

type favoriteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteServiceClient(cc grpc.ClientConnInterface) FavoriteServiceClient {
	return &favoriteServiceClient{cc}
}

func (c *favoriteServiceClient) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionRsp, error) {
	out := new(FavoriteActionRsp)
	err := c.cc.Invoke(ctx, FavoriteService_FavoriteAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) GetFavoriteVideoIdList(ctx context.Context, in *GetFavoriteVideoIdListReq, opts ...grpc.CallOption) (*GetFavoriteVideoIdListRsp, error) {
	out := new(GetFavoriteVideoIdListRsp)
	err := c.cc.Invoke(ctx, FavoriteService_GetFavoriteVideoIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) IsFavoriteVideoDict(ctx context.Context, in *IsFavoriteVideoDictReq, opts ...grpc.CallOption) (*IsFavoriteVideoDictRsp, error) {
	out := new(IsFavoriteVideoDictRsp)
	err := c.cc.Invoke(ctx, FavoriteService_IsFavoriteVideoDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServiceServer is the server API for FavoriteService service.
// All implementations must embed UnimplementedFavoriteServiceServer
// for forward compatibility
type FavoriteServiceServer interface {
	FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionRsp, error)
	GetFavoriteVideoIdList(context.Context, *GetFavoriteVideoIdListReq) (*GetFavoriteVideoIdListRsp, error)
	IsFavoriteVideoDict(context.Context, *IsFavoriteVideoDictReq) (*IsFavoriteVideoDictRsp, error)
	mustEmbedUnimplementedFavoriteServiceServer()
}

// UnimplementedFavoriteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteServiceServer struct {
}

func (UnimplementedFavoriteServiceServer) FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedFavoriteServiceServer) GetFavoriteVideoIdList(context.Context, *GetFavoriteVideoIdListReq) (*GetFavoriteVideoIdListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteVideoIdList not implemented")
}
func (UnimplementedFavoriteServiceServer) IsFavoriteVideoDict(context.Context, *IsFavoriteVideoDictReq) (*IsFavoriteVideoDictRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFavoriteVideoDict not implemented")
}
func (UnimplementedFavoriteServiceServer) mustEmbedUnimplementedFavoriteServiceServer() {}

// UnsafeFavoriteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServiceServer will
// result in compilation errors.
type UnsafeFavoriteServiceServer interface {
	mustEmbedUnimplementedFavoriteServiceServer()
}

func RegisterFavoriteServiceServer(s grpc.ServiceRegistrar, srv FavoriteServiceServer) {
	s.RegisterService(&FavoriteService_ServiceDesc, srv)
}

func _FavoriteService_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteService_FavoriteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).FavoriteAction(ctx, req.(*FavoriteActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_GetFavoriteVideoIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoriteVideoIdListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).GetFavoriteVideoIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteService_GetFavoriteVideoIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).GetFavoriteVideoIdList(ctx, req.(*GetFavoriteVideoIdListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_IsFavoriteVideoDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFavoriteVideoDictReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).IsFavoriteVideoDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavoriteService_IsFavoriteVideoDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).IsFavoriteVideoDict(ctx, req.(*IsFavoriteVideoDictReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoriteService_ServiceDesc is the grpc.ServiceDesc for FavoriteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoriteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FavoriteService",
	HandlerType: (*FavoriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _FavoriteService_FavoriteAction_Handler,
		},
		{
			MethodName: "GetFavoriteVideoIdList",
			Handler:    _FavoriteService_GetFavoriteVideoIdList_Handler,
		},
		{
			MethodName: "IsFavoriteVideoDict",
			Handler:    _FavoriteService_IsFavoriteVideoDict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
