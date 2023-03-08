// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: proto/exercise.proto

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

// ExerciseStoreClient is the client API for ExerciseStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExerciseStoreClient interface {
	AddExercises(ctx context.Context, in *AddExercisesRequest, opts ...grpc.CallOption) (*Empty, error)
	GetExercises(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetExercisesResponse, error)
	GetExerciseByTags(ctx context.Context, in *GetExerciseByTagsRequest, opts ...grpc.CallOption) (*GetExercisesResponse, error)
	GetExerciseByCategory(ctx context.Context, in *GetExerciseByCategoryRequest, opts ...grpc.CallOption) (*GetExercisesResponse, error)
	AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*Empty, error)
	GetCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	GetCategoriesByName(ctx context.Context, in *GetCategoriesByNameReq, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
}

type exerciseStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewExerciseStoreClient(cc grpc.ClientConnInterface) ExerciseStoreClient {
	return &exerciseStoreClient{cc}
}

func (c *exerciseStoreClient) AddExercises(ctx context.Context, in *AddExercisesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/exercise.ExerciseStore/AddExercises", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseStoreClient) GetExercises(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetExercisesResponse, error) {
	out := new(GetExercisesResponse)
	err := c.cc.Invoke(ctx, "/exercise.ExerciseStore/GetExercises", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseStoreClient) GetExerciseByTags(ctx context.Context, in *GetExerciseByTagsRequest, opts ...grpc.CallOption) (*GetExercisesResponse, error) {
	out := new(GetExercisesResponse)
	err := c.cc.Invoke(ctx, "/exercise.ExerciseStore/GetExerciseByTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseStoreClient) GetExerciseByCategory(ctx context.Context, in *GetExerciseByCategoryRequest, opts ...grpc.CallOption) (*GetExercisesResponse, error) {
	out := new(GetExercisesResponse)
	err := c.cc.Invoke(ctx, "/exercise.ExerciseStore/GetExerciseByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseStoreClient) AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/exercise.ExerciseStore/AddCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseStoreClient) GetCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, "/exercise.ExerciseStore/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exerciseStoreClient) GetCategoriesByName(ctx context.Context, in *GetCategoriesByNameReq, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, "/exercise.ExerciseStore/GetCategoriesByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExerciseStoreServer is the server API for ExerciseStore service.
// All implementations must embed UnimplementedExerciseStoreServer
// for forward compatibility
type ExerciseStoreServer interface {
	AddExercises(context.Context, *AddExercisesRequest) (*Empty, error)
	GetExercises(context.Context, *Empty) (*GetExercisesResponse, error)
	GetExerciseByTags(context.Context, *GetExerciseByTagsRequest) (*GetExercisesResponse, error)
	GetExerciseByCategory(context.Context, *GetExerciseByCategoryRequest) (*GetExercisesResponse, error)
	AddCategory(context.Context, *AddCategoryRequest) (*Empty, error)
	GetCategories(context.Context, *Empty) (*GetCategoriesResponse, error)
	GetCategoriesByName(context.Context, *GetCategoriesByNameReq) (*GetCategoriesResponse, error)
	mustEmbedUnimplementedExerciseStoreServer()
}

// UnimplementedExerciseStoreServer must be embedded to have forward compatible implementations.
type UnimplementedExerciseStoreServer struct {
}

func (UnimplementedExerciseStoreServer) AddExercises(context.Context, *AddExercisesRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExercises not implemented")
}
func (UnimplementedExerciseStoreServer) GetExercises(context.Context, *Empty) (*GetExercisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExercises not implemented")
}
func (UnimplementedExerciseStoreServer) GetExerciseByTags(context.Context, *GetExerciseByTagsRequest) (*GetExercisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExerciseByTags not implemented")
}
func (UnimplementedExerciseStoreServer) GetExerciseByCategory(context.Context, *GetExerciseByCategoryRequest) (*GetExercisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExerciseByCategory not implemented")
}
func (UnimplementedExerciseStoreServer) AddCategory(context.Context, *AddCategoryRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedExerciseStoreServer) GetCategories(context.Context, *Empty) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedExerciseStoreServer) GetCategoriesByName(context.Context, *GetCategoriesByNameReq) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoriesByName not implemented")
}
func (UnimplementedExerciseStoreServer) mustEmbedUnimplementedExerciseStoreServer() {}

// UnsafeExerciseStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExerciseStoreServer will
// result in compilation errors.
type UnsafeExerciseStoreServer interface {
	mustEmbedUnimplementedExerciseStoreServer()
}

func RegisterExerciseStoreServer(s grpc.ServiceRegistrar, srv ExerciseStoreServer) {
	s.RegisterService(&ExerciseStore_ServiceDesc, srv)
}

func _ExerciseStore_AddExercises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExercisesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseStoreServer).AddExercises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exercise.ExerciseStore/AddExercises",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseStoreServer).AddExercises(ctx, req.(*AddExercisesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseStore_GetExercises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseStoreServer).GetExercises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exercise.ExerciseStore/GetExercises",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseStoreServer).GetExercises(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseStore_GetExerciseByTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExerciseByTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseStoreServer).GetExerciseByTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exercise.ExerciseStore/GetExerciseByTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseStoreServer).GetExerciseByTags(ctx, req.(*GetExerciseByTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseStore_GetExerciseByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExerciseByCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseStoreServer).GetExerciseByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exercise.ExerciseStore/GetExerciseByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseStoreServer).GetExerciseByCategory(ctx, req.(*GetExerciseByCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseStore_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseStoreServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exercise.ExerciseStore/AddCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseStoreServer).AddCategory(ctx, req.(*AddCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseStore_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseStoreServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exercise.ExerciseStore/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseStoreServer).GetCategories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExerciseStore_GetCategoriesByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesByNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseStoreServer).GetCategoriesByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exercise.ExerciseStore/GetCategoriesByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseStoreServer).GetCategoriesByName(ctx, req.(*GetCategoriesByNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ExerciseStore_ServiceDesc is the grpc.ServiceDesc for ExerciseStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExerciseStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exercise.ExerciseStore",
	HandlerType: (*ExerciseStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddExercises",
			Handler:    _ExerciseStore_AddExercises_Handler,
		},
		{
			MethodName: "GetExercises",
			Handler:    _ExerciseStore_GetExercises_Handler,
		},
		{
			MethodName: "GetExerciseByTags",
			Handler:    _ExerciseStore_GetExerciseByTags_Handler,
		},
		{
			MethodName: "GetExerciseByCategory",
			Handler:    _ExerciseStore_GetExerciseByCategory_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _ExerciseStore_AddCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _ExerciseStore_GetCategories_Handler,
		},
		{
			MethodName: "GetCategoriesByName",
			Handler:    _ExerciseStore_GetCategoriesByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/exercise.proto",
}
