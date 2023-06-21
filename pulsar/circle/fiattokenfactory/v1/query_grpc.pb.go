// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: circle/fiattokenfactory/v1/query.proto

package fiattokenfactoryv1

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
	Query_Params_FullMethodName              = "/circle.fiattokenfactory.v1.Query/Params"
	Query_Blacklisted_FullMethodName         = "/circle.fiattokenfactory.v1.Query/Blacklisted"
	Query_BlacklistedAll_FullMethodName      = "/circle.fiattokenfactory.v1.Query/BlacklistedAll"
	Query_Paused_FullMethodName              = "/circle.fiattokenfactory.v1.Query/Paused"
	Query_MasterMinter_FullMethodName        = "/circle.fiattokenfactory.v1.Query/MasterMinter"
	Query_Minters_FullMethodName             = "/circle.fiattokenfactory.v1.Query/Minters"
	Query_MintersAll_FullMethodName          = "/circle.fiattokenfactory.v1.Query/MintersAll"
	Query_Pauser_FullMethodName              = "/circle.fiattokenfactory.v1.Query/Pauser"
	Query_Blacklister_FullMethodName         = "/circle.fiattokenfactory.v1.Query/Blacklister"
	Query_Owner_FullMethodName               = "/circle.fiattokenfactory.v1.Query/Owner"
	Query_MinterController_FullMethodName    = "/circle.fiattokenfactory.v1.Query/MinterController"
	Query_MinterControllerAll_FullMethodName = "/circle.fiattokenfactory.v1.Query/MinterControllerAll"
	Query_MintingDenom_FullMethodName        = "/circle.fiattokenfactory.v1.Query/MintingDenom"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a Blacklisted by index.
	Blacklisted(ctx context.Context, in *QueryGetBlacklistedRequest, opts ...grpc.CallOption) (*QueryGetBlacklistedResponse, error)
	// Queries a list of Blacklisted items.
	BlacklistedAll(ctx context.Context, in *QueryAllBlacklistedRequest, opts ...grpc.CallOption) (*QueryAllBlacklistedResponse, error)
	// Queries a Paused by index.
	Paused(ctx context.Context, in *QueryGetPausedRequest, opts ...grpc.CallOption) (*QueryGetPausedResponse, error)
	// Queries a MasterMinter by index.
	MasterMinter(ctx context.Context, in *QueryGetMasterMinterRequest, opts ...grpc.CallOption) (*QueryGetMasterMinterResponse, error)
	// Queries a Minters by index.
	Minters(ctx context.Context, in *QueryGetMintersRequest, opts ...grpc.CallOption) (*QueryGetMintersResponse, error)
	// Queries a list of Minters items.
	MintersAll(ctx context.Context, in *QueryAllMintersRequest, opts ...grpc.CallOption) (*QueryAllMintersResponse, error)
	// Queries a Pauser by index.
	Pauser(ctx context.Context, in *QueryGetPauserRequest, opts ...grpc.CallOption) (*QueryGetPauserResponse, error)
	// Queries a Blacklister by index.
	Blacklister(ctx context.Context, in *QueryGetBlacklisterRequest, opts ...grpc.CallOption) (*QueryGetBlacklisterResponse, error)
	// Queries a Owner by index.
	Owner(ctx context.Context, in *QueryGetOwnerRequest, opts ...grpc.CallOption) (*QueryGetOwnerResponse, error)
	// Queries a MinterController by index.
	MinterController(ctx context.Context, in *QueryGetMinterControllerRequest, opts ...grpc.CallOption) (*QueryGetMinterControllerResponse, error)
	// Queries a list of MinterController items.
	MinterControllerAll(ctx context.Context, in *QueryAllMinterControllerRequest, opts ...grpc.CallOption) (*QueryAllMinterControllerResponse, error)
	// Queries a MintingDenom by index.
	MintingDenom(ctx context.Context, in *QueryGetMintingDenomRequest, opts ...grpc.CallOption) (*QueryGetMintingDenomResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Blacklisted(ctx context.Context, in *QueryGetBlacklistedRequest, opts ...grpc.CallOption) (*QueryGetBlacklistedResponse, error) {
	out := new(QueryGetBlacklistedResponse)
	err := c.cc.Invoke(ctx, Query_Blacklisted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BlacklistedAll(ctx context.Context, in *QueryAllBlacklistedRequest, opts ...grpc.CallOption) (*QueryAllBlacklistedResponse, error) {
	out := new(QueryAllBlacklistedResponse)
	err := c.cc.Invoke(ctx, Query_BlacklistedAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Paused(ctx context.Context, in *QueryGetPausedRequest, opts ...grpc.CallOption) (*QueryGetPausedResponse, error) {
	out := new(QueryGetPausedResponse)
	err := c.cc.Invoke(ctx, Query_Paused_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MasterMinter(ctx context.Context, in *QueryGetMasterMinterRequest, opts ...grpc.CallOption) (*QueryGetMasterMinterResponse, error) {
	out := new(QueryGetMasterMinterResponse)
	err := c.cc.Invoke(ctx, Query_MasterMinter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Minters(ctx context.Context, in *QueryGetMintersRequest, opts ...grpc.CallOption) (*QueryGetMintersResponse, error) {
	out := new(QueryGetMintersResponse)
	err := c.cc.Invoke(ctx, Query_Minters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MintersAll(ctx context.Context, in *QueryAllMintersRequest, opts ...grpc.CallOption) (*QueryAllMintersResponse, error) {
	out := new(QueryAllMintersResponse)
	err := c.cc.Invoke(ctx, Query_MintersAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pauser(ctx context.Context, in *QueryGetPauserRequest, opts ...grpc.CallOption) (*QueryGetPauserResponse, error) {
	out := new(QueryGetPauserResponse)
	err := c.cc.Invoke(ctx, Query_Pauser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Blacklister(ctx context.Context, in *QueryGetBlacklisterRequest, opts ...grpc.CallOption) (*QueryGetBlacklisterResponse, error) {
	out := new(QueryGetBlacklisterResponse)
	err := c.cc.Invoke(ctx, Query_Blacklister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Owner(ctx context.Context, in *QueryGetOwnerRequest, opts ...grpc.CallOption) (*QueryGetOwnerResponse, error) {
	out := new(QueryGetOwnerResponse)
	err := c.cc.Invoke(ctx, Query_Owner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MinterController(ctx context.Context, in *QueryGetMinterControllerRequest, opts ...grpc.CallOption) (*QueryGetMinterControllerResponse, error) {
	out := new(QueryGetMinterControllerResponse)
	err := c.cc.Invoke(ctx, Query_MinterController_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MinterControllerAll(ctx context.Context, in *QueryAllMinterControllerRequest, opts ...grpc.CallOption) (*QueryAllMinterControllerResponse, error) {
	out := new(QueryAllMinterControllerResponse)
	err := c.cc.Invoke(ctx, Query_MinterControllerAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MintingDenom(ctx context.Context, in *QueryGetMintingDenomRequest, opts ...grpc.CallOption) (*QueryGetMintingDenomResponse, error) {
	out := new(QueryGetMintingDenomResponse)
	err := c.cc.Invoke(ctx, Query_MintingDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a Blacklisted by index.
	Blacklisted(context.Context, *QueryGetBlacklistedRequest) (*QueryGetBlacklistedResponse, error)
	// Queries a list of Blacklisted items.
	BlacklistedAll(context.Context, *QueryAllBlacklistedRequest) (*QueryAllBlacklistedResponse, error)
	// Queries a Paused by index.
	Paused(context.Context, *QueryGetPausedRequest) (*QueryGetPausedResponse, error)
	// Queries a MasterMinter by index.
	MasterMinter(context.Context, *QueryGetMasterMinterRequest) (*QueryGetMasterMinterResponse, error)
	// Queries a Minters by index.
	Minters(context.Context, *QueryGetMintersRequest) (*QueryGetMintersResponse, error)
	// Queries a list of Minters items.
	MintersAll(context.Context, *QueryAllMintersRequest) (*QueryAllMintersResponse, error)
	// Queries a Pauser by index.
	Pauser(context.Context, *QueryGetPauserRequest) (*QueryGetPauserResponse, error)
	// Queries a Blacklister by index.
	Blacklister(context.Context, *QueryGetBlacklisterRequest) (*QueryGetBlacklisterResponse, error)
	// Queries a Owner by index.
	Owner(context.Context, *QueryGetOwnerRequest) (*QueryGetOwnerResponse, error)
	// Queries a MinterController by index.
	MinterController(context.Context, *QueryGetMinterControllerRequest) (*QueryGetMinterControllerResponse, error)
	// Queries a list of MinterController items.
	MinterControllerAll(context.Context, *QueryAllMinterControllerRequest) (*QueryAllMinterControllerResponse, error)
	// Queries a MintingDenom by index.
	MintingDenom(context.Context, *QueryGetMintingDenomRequest) (*QueryGetMintingDenomResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Blacklisted(context.Context, *QueryGetBlacklistedRequest) (*QueryGetBlacklistedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Blacklisted not implemented")
}
func (UnimplementedQueryServer) BlacklistedAll(context.Context, *QueryAllBlacklistedRequest) (*QueryAllBlacklistedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlacklistedAll not implemented")
}
func (UnimplementedQueryServer) Paused(context.Context, *QueryGetPausedRequest) (*QueryGetPausedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Paused not implemented")
}
func (UnimplementedQueryServer) MasterMinter(context.Context, *QueryGetMasterMinterRequest) (*QueryGetMasterMinterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MasterMinter not implemented")
}
func (UnimplementedQueryServer) Minters(context.Context, *QueryGetMintersRequest) (*QueryGetMintersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Minters not implemented")
}
func (UnimplementedQueryServer) MintersAll(context.Context, *QueryAllMintersRequest) (*QueryAllMintersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintersAll not implemented")
}
func (UnimplementedQueryServer) Pauser(context.Context, *QueryGetPauserRequest) (*QueryGetPauserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pauser not implemented")
}
func (UnimplementedQueryServer) Blacklister(context.Context, *QueryGetBlacklisterRequest) (*QueryGetBlacklisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Blacklister not implemented")
}
func (UnimplementedQueryServer) Owner(context.Context, *QueryGetOwnerRequest) (*QueryGetOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Owner not implemented")
}
func (UnimplementedQueryServer) MinterController(context.Context, *QueryGetMinterControllerRequest) (*QueryGetMinterControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinterController not implemented")
}
func (UnimplementedQueryServer) MinterControllerAll(context.Context, *QueryAllMinterControllerRequest) (*QueryAllMinterControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinterControllerAll not implemented")
}
func (UnimplementedQueryServer) MintingDenom(context.Context, *QueryGetMintingDenomRequest) (*QueryGetMintingDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintingDenom not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Blacklisted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetBlacklistedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Blacklisted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Blacklisted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Blacklisted(ctx, req.(*QueryGetBlacklistedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BlacklistedAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBlacklistedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BlacklistedAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BlacklistedAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BlacklistedAll(ctx, req.(*QueryAllBlacklistedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Paused_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPausedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Paused(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Paused_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Paused(ctx, req.(*QueryGetPausedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MasterMinter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetMasterMinterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MasterMinter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MasterMinter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MasterMinter(ctx, req.(*QueryGetMasterMinterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Minters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetMintersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Minters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Minters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Minters(ctx, req.(*QueryGetMintersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MintersAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllMintersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MintersAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MintersAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MintersAll(ctx, req.(*QueryAllMintersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pauser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPauserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pauser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Pauser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pauser(ctx, req.(*QueryGetPauserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Blacklister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetBlacklisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Blacklister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Blacklister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Blacklister(ctx, req.(*QueryGetBlacklisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Owner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Owner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Owner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Owner(ctx, req.(*QueryGetOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MinterController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetMinterControllerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MinterController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MinterController_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MinterController(ctx, req.(*QueryGetMinterControllerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MinterControllerAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllMinterControllerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MinterControllerAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MinterControllerAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MinterControllerAll(ctx, req.(*QueryAllMinterControllerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MintingDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetMintingDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MintingDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_MintingDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MintingDenom(ctx, req.(*QueryGetMintingDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "circle.fiattokenfactory.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Blacklisted",
			Handler:    _Query_Blacklisted_Handler,
		},
		{
			MethodName: "BlacklistedAll",
			Handler:    _Query_BlacklistedAll_Handler,
		},
		{
			MethodName: "Paused",
			Handler:    _Query_Paused_Handler,
		},
		{
			MethodName: "MasterMinter",
			Handler:    _Query_MasterMinter_Handler,
		},
		{
			MethodName: "Minters",
			Handler:    _Query_Minters_Handler,
		},
		{
			MethodName: "MintersAll",
			Handler:    _Query_MintersAll_Handler,
		},
		{
			MethodName: "Pauser",
			Handler:    _Query_Pauser_Handler,
		},
		{
			MethodName: "Blacklister",
			Handler:    _Query_Blacklister_Handler,
		},
		{
			MethodName: "Owner",
			Handler:    _Query_Owner_Handler,
		},
		{
			MethodName: "MinterController",
			Handler:    _Query_MinterController_Handler,
		},
		{
			MethodName: "MinterControllerAll",
			Handler:    _Query_MinterControllerAll_Handler,
		},
		{
			MethodName: "MintingDenom",
			Handler:    _Query_MintingDenom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "circle/fiattokenfactory/v1/query.proto",
}
