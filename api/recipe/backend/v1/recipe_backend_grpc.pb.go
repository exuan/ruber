// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	v1 "github.com/exuan/ruber/api/recipe/service/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackendClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
	User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error)
	RecipeIndex(ctx context.Context, in *v1.IndexRequest, opts ...grpc.CallOption) (*v1.IndexReply, error)
	SaveRecipeIndex(ctx context.Context, in *v1.SaveIndexRequest, opts ...grpc.CallOption) (*v1.SaveIndexReply, error)
	Recipes(ctx context.Context, in *v1.RecipesRequest, opts ...grpc.CallOption) (*v1.RecipesReply, error)
	Recipe(ctx context.Context, in *v1.RecipeRequest, opts ...grpc.CallOption) (*v1.RecipeReply, error)
	SaveRecipe(ctx context.Context, in *v1.SaveRecipeRequest, opts ...grpc.CallOption) (*v1.SaveRecipeReply, error)
	DeleteRecipe(ctx context.Context, in *v1.DeleteRecipeRequest, opts ...grpc.CallOption) (*v1.DeleteRecipeReply, error)
	RecipeCategories(ctx context.Context, in *v1.CategoriesRequest, opts ...grpc.CallOption) (*v1.CategoriesReply, error)
	RecipeCategory(ctx context.Context, in *v1.CategoryRequest, opts ...grpc.CallOption) (*v1.CategoryReply, error)
	SaveRecipeCategory(ctx context.Context, in *v1.SaveCategoryRequest, opts ...grpc.CallOption) (*v1.SaveCategoryReply, error)
	DeleteRecipeCategory(ctx context.Context, in *v1.DeleteCategoryRequest, opts ...grpc.CallOption) (*v1.DeleteCategoryReply, error)
	UploadUrl(ctx context.Context, in *UploadUrlRequest, opts ...grpc.CallOption) (*UploadUrlReply, error)
}

type backendClient struct {
	cc grpc.ClientConnInterface
}

func NewBackendClient(cc grpc.ClientConnInterface) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) RecipeIndex(ctx context.Context, in *v1.IndexRequest, opts ...grpc.CallOption) (*v1.IndexReply, error) {
	out := new(v1.IndexReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/RecipeIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) SaveRecipeIndex(ctx context.Context, in *v1.SaveIndexRequest, opts ...grpc.CallOption) (*v1.SaveIndexReply, error) {
	out := new(v1.SaveIndexReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/SaveRecipeIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) Recipes(ctx context.Context, in *v1.RecipesRequest, opts ...grpc.CallOption) (*v1.RecipesReply, error) {
	out := new(v1.RecipesReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/Recipes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) Recipe(ctx context.Context, in *v1.RecipeRequest, opts ...grpc.CallOption) (*v1.RecipeReply, error) {
	out := new(v1.RecipeReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/Recipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) SaveRecipe(ctx context.Context, in *v1.SaveRecipeRequest, opts ...grpc.CallOption) (*v1.SaveRecipeReply, error) {
	out := new(v1.SaveRecipeReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/SaveRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) DeleteRecipe(ctx context.Context, in *v1.DeleteRecipeRequest, opts ...grpc.CallOption) (*v1.DeleteRecipeReply, error) {
	out := new(v1.DeleteRecipeReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/DeleteRecipe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) RecipeCategories(ctx context.Context, in *v1.CategoriesRequest, opts ...grpc.CallOption) (*v1.CategoriesReply, error) {
	out := new(v1.CategoriesReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/RecipeCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) RecipeCategory(ctx context.Context, in *v1.CategoryRequest, opts ...grpc.CallOption) (*v1.CategoryReply, error) {
	out := new(v1.CategoryReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/RecipeCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) SaveRecipeCategory(ctx context.Context, in *v1.SaveCategoryRequest, opts ...grpc.CallOption) (*v1.SaveCategoryReply, error) {
	out := new(v1.SaveCategoryReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/SaveRecipeCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) DeleteRecipeCategory(ctx context.Context, in *v1.DeleteCategoryRequest, opts ...grpc.CallOption) (*v1.DeleteCategoryReply, error) {
	out := new(v1.DeleteCategoryReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/DeleteRecipeCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) UploadUrl(ctx context.Context, in *UploadUrlRequest, opts ...grpc.CallOption) (*UploadUrlReply, error) {
	out := new(UploadUrlReply)
	err := c.cc.Invoke(ctx, "/recipe.backend.v1.Backend/UploadUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
// All implementations must embed UnimplementedBackendServer
// for forward compatibility
type BackendServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	User(context.Context, *UserRequest) (*UserReply, error)
	RecipeIndex(context.Context, *v1.IndexRequest) (*v1.IndexReply, error)
	SaveRecipeIndex(context.Context, *v1.SaveIndexRequest) (*v1.SaveIndexReply, error)
	Recipes(context.Context, *v1.RecipesRequest) (*v1.RecipesReply, error)
	Recipe(context.Context, *v1.RecipeRequest) (*v1.RecipeReply, error)
	SaveRecipe(context.Context, *v1.SaveRecipeRequest) (*v1.SaveRecipeReply, error)
	DeleteRecipe(context.Context, *v1.DeleteRecipeRequest) (*v1.DeleteRecipeReply, error)
	RecipeCategories(context.Context, *v1.CategoriesRequest) (*v1.CategoriesReply, error)
	RecipeCategory(context.Context, *v1.CategoryRequest) (*v1.CategoryReply, error)
	SaveRecipeCategory(context.Context, *v1.SaveCategoryRequest) (*v1.SaveCategoryReply, error)
	DeleteRecipeCategory(context.Context, *v1.DeleteCategoryRequest) (*v1.DeleteCategoryReply, error)
	UploadUrl(context.Context, *UploadUrlRequest) (*UploadUrlReply, error)
	mustEmbedUnimplementedBackendServer()
}

// UnimplementedBackendServer must be embedded to have forward compatible implementations.
type UnimplementedBackendServer struct {
}

func (UnimplementedBackendServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBackendServer) Logout(context.Context, *LogoutRequest) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedBackendServer) User(context.Context, *UserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedBackendServer) RecipeIndex(context.Context, *v1.IndexRequest) (*v1.IndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecipeIndex not implemented")
}
func (UnimplementedBackendServer) SaveRecipeIndex(context.Context, *v1.SaveIndexRequest) (*v1.SaveIndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRecipeIndex not implemented")
}
func (UnimplementedBackendServer) Recipes(context.Context, *v1.RecipesRequest) (*v1.RecipesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recipes not implemented")
}
func (UnimplementedBackendServer) Recipe(context.Context, *v1.RecipeRequest) (*v1.RecipeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recipe not implemented")
}
func (UnimplementedBackendServer) SaveRecipe(context.Context, *v1.SaveRecipeRequest) (*v1.SaveRecipeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRecipe not implemented")
}
func (UnimplementedBackendServer) DeleteRecipe(context.Context, *v1.DeleteRecipeRequest) (*v1.DeleteRecipeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipe not implemented")
}
func (UnimplementedBackendServer) RecipeCategories(context.Context, *v1.CategoriesRequest) (*v1.CategoriesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecipeCategories not implemented")
}
func (UnimplementedBackendServer) RecipeCategory(context.Context, *v1.CategoryRequest) (*v1.CategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecipeCategory not implemented")
}
func (UnimplementedBackendServer) SaveRecipeCategory(context.Context, *v1.SaveCategoryRequest) (*v1.SaveCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRecipeCategory not implemented")
}
func (UnimplementedBackendServer) DeleteRecipeCategory(context.Context, *v1.DeleteCategoryRequest) (*v1.DeleteCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecipeCategory not implemented")
}
func (UnimplementedBackendServer) UploadUrl(context.Context, *UploadUrlRequest) (*UploadUrlReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadUrl not implemented")
}
func (UnimplementedBackendServer) mustEmbedUnimplementedBackendServer() {}

// UnsafeBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackendServer will
// result in compilation errors.
type UnsafeBackendServer interface {
	mustEmbedUnimplementedBackendServer()
}

func RegisterBackendServer(s grpc.ServiceRegistrar, srv BackendServer) {
	s.RegisterService(&Backend_ServiceDesc, srv)
}

func _Backend_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).User(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_RecipeIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).RecipeIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/RecipeIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).RecipeIndex(ctx, req.(*v1.IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_SaveRecipeIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SaveIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).SaveRecipeIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/SaveRecipeIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).SaveRecipeIndex(ctx, req.(*v1.SaveIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_Recipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.RecipesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Recipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/Recipes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Recipes(ctx, req.(*v1.RecipesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_Recipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.RecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Recipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/Recipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Recipe(ctx, req.(*v1.RecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_SaveRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SaveRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).SaveRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/SaveRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).SaveRecipe(ctx, req.(*v1.SaveRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_DeleteRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeleteRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeleteRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/DeleteRecipe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeleteRecipe(ctx, req.(*v1.DeleteRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_RecipeCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).RecipeCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/RecipeCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).RecipeCategories(ctx, req.(*v1.CategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_RecipeCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).RecipeCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/RecipeCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).RecipeCategory(ctx, req.(*v1.CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_SaveRecipeCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SaveCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).SaveRecipeCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/SaveRecipeCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).SaveRecipeCategory(ctx, req.(*v1.SaveCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_DeleteRecipeCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeleteRecipeCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/DeleteRecipeCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeleteRecipeCategory(ctx, req.(*v1.DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_UploadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).UploadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recipe.backend.v1.Backend/UploadUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).UploadUrl(ctx, req.(*UploadUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Backend_ServiceDesc is the grpc.ServiceDesc for Backend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recipe.backend.v1.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Backend_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Backend_Logout_Handler,
		},
		{
			MethodName: "User",
			Handler:    _Backend_User_Handler,
		},
		{
			MethodName: "RecipeIndex",
			Handler:    _Backend_RecipeIndex_Handler,
		},
		{
			MethodName: "SaveRecipeIndex",
			Handler:    _Backend_SaveRecipeIndex_Handler,
		},
		{
			MethodName: "Recipes",
			Handler:    _Backend_Recipes_Handler,
		},
		{
			MethodName: "Recipe",
			Handler:    _Backend_Recipe_Handler,
		},
		{
			MethodName: "SaveRecipe",
			Handler:    _Backend_SaveRecipe_Handler,
		},
		{
			MethodName: "DeleteRecipe",
			Handler:    _Backend_DeleteRecipe_Handler,
		},
		{
			MethodName: "RecipeCategories",
			Handler:    _Backend_RecipeCategories_Handler,
		},
		{
			MethodName: "RecipeCategory",
			Handler:    _Backend_RecipeCategory_Handler,
		},
		{
			MethodName: "SaveRecipeCategory",
			Handler:    _Backend_SaveRecipeCategory_Handler,
		},
		{
			MethodName: "DeleteRecipeCategory",
			Handler:    _Backend_DeleteRecipeCategory_Handler,
		},
		{
			MethodName: "UploadUrl",
			Handler:    _Backend_UploadUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/recipe_backend.proto",
}
