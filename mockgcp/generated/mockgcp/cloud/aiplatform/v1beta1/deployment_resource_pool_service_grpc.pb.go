// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/aiplatform/v1beta1/deployment_resource_pool_service.proto

package aiplatformpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DeploymentResourcePoolServiceClient is the client API for DeploymentResourcePoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeploymentResourcePoolServiceClient interface {
	// Create a DeploymentResourcePool.
	CreateDeploymentResourcePool(ctx context.Context, in *CreateDeploymentResourcePoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Get a DeploymentResourcePool.
	GetDeploymentResourcePool(ctx context.Context, in *GetDeploymentResourcePoolRequest, opts ...grpc.CallOption) (*DeploymentResourcePool, error)
	// List DeploymentResourcePools in a location.
	ListDeploymentResourcePools(ctx context.Context, in *ListDeploymentResourcePoolsRequest, opts ...grpc.CallOption) (*ListDeploymentResourcePoolsResponse, error)
	// Delete a DeploymentResourcePool.
	DeleteDeploymentResourcePool(ctx context.Context, in *DeleteDeploymentResourcePoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// List DeployedModels that have been deployed on this DeploymentResourcePool.
	QueryDeployedModels(ctx context.Context, in *QueryDeployedModelsRequest, opts ...grpc.CallOption) (*QueryDeployedModelsResponse, error)
}

type deploymentResourcePoolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentResourcePoolServiceClient(cc grpc.ClientConnInterface) DeploymentResourcePoolServiceClient {
	return &deploymentResourcePoolServiceClient{cc}
}

func (c *deploymentResourcePoolServiceClient) CreateDeploymentResourcePool(ctx context.Context, in *CreateDeploymentResourcePoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/CreateDeploymentResourcePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentResourcePoolServiceClient) GetDeploymentResourcePool(ctx context.Context, in *GetDeploymentResourcePoolRequest, opts ...grpc.CallOption) (*DeploymentResourcePool, error) {
	out := new(DeploymentResourcePool)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/GetDeploymentResourcePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentResourcePoolServiceClient) ListDeploymentResourcePools(ctx context.Context, in *ListDeploymentResourcePoolsRequest, opts ...grpc.CallOption) (*ListDeploymentResourcePoolsResponse, error) {
	out := new(ListDeploymentResourcePoolsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/ListDeploymentResourcePools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentResourcePoolServiceClient) DeleteDeploymentResourcePool(ctx context.Context, in *DeleteDeploymentResourcePoolRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/DeleteDeploymentResourcePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentResourcePoolServiceClient) QueryDeployedModels(ctx context.Context, in *QueryDeployedModelsRequest, opts ...grpc.CallOption) (*QueryDeployedModelsResponse, error) {
	out := new(QueryDeployedModelsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/QueryDeployedModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentResourcePoolServiceServer is the server API for DeploymentResourcePoolService service.
// All implementations must embed UnimplementedDeploymentResourcePoolServiceServer
// for forward compatibility
type DeploymentResourcePoolServiceServer interface {
	// Create a DeploymentResourcePool.
	CreateDeploymentResourcePool(context.Context, *CreateDeploymentResourcePoolRequest) (*longrunningpb.Operation, error)
	// Get a DeploymentResourcePool.
	GetDeploymentResourcePool(context.Context, *GetDeploymentResourcePoolRequest) (*DeploymentResourcePool, error)
	// List DeploymentResourcePools in a location.
	ListDeploymentResourcePools(context.Context, *ListDeploymentResourcePoolsRequest) (*ListDeploymentResourcePoolsResponse, error)
	// Delete a DeploymentResourcePool.
	DeleteDeploymentResourcePool(context.Context, *DeleteDeploymentResourcePoolRequest) (*longrunningpb.Operation, error)
	// List DeployedModels that have been deployed on this DeploymentResourcePool.
	QueryDeployedModels(context.Context, *QueryDeployedModelsRequest) (*QueryDeployedModelsResponse, error)
	mustEmbedUnimplementedDeploymentResourcePoolServiceServer()
}

// UnimplementedDeploymentResourcePoolServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeploymentResourcePoolServiceServer struct {
}

func (UnimplementedDeploymentResourcePoolServiceServer) CreateDeploymentResourcePool(context.Context, *CreateDeploymentResourcePoolRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeploymentResourcePool not implemented")
}
func (UnimplementedDeploymentResourcePoolServiceServer) GetDeploymentResourcePool(context.Context, *GetDeploymentResourcePoolRequest) (*DeploymentResourcePool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeploymentResourcePool not implemented")
}
func (UnimplementedDeploymentResourcePoolServiceServer) ListDeploymentResourcePools(context.Context, *ListDeploymentResourcePoolsRequest) (*ListDeploymentResourcePoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeploymentResourcePools not implemented")
}
func (UnimplementedDeploymentResourcePoolServiceServer) DeleteDeploymentResourcePool(context.Context, *DeleteDeploymentResourcePoolRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeploymentResourcePool not implemented")
}
func (UnimplementedDeploymentResourcePoolServiceServer) QueryDeployedModels(context.Context, *QueryDeployedModelsRequest) (*QueryDeployedModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeployedModels not implemented")
}
func (UnimplementedDeploymentResourcePoolServiceServer) mustEmbedUnimplementedDeploymentResourcePoolServiceServer() {
}

// UnsafeDeploymentResourcePoolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeploymentResourcePoolServiceServer will
// result in compilation errors.
type UnsafeDeploymentResourcePoolServiceServer interface {
	mustEmbedUnimplementedDeploymentResourcePoolServiceServer()
}

func RegisterDeploymentResourcePoolServiceServer(s grpc.ServiceRegistrar, srv DeploymentResourcePoolServiceServer) {
	s.RegisterService(&DeploymentResourcePoolService_ServiceDesc, srv)
}

func _DeploymentResourcePoolService_CreateDeploymentResourcePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeploymentResourcePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentResourcePoolServiceServer).CreateDeploymentResourcePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/CreateDeploymentResourcePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentResourcePoolServiceServer).CreateDeploymentResourcePool(ctx, req.(*CreateDeploymentResourcePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentResourcePoolService_GetDeploymentResourcePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentResourcePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentResourcePoolServiceServer).GetDeploymentResourcePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/GetDeploymentResourcePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentResourcePoolServiceServer).GetDeploymentResourcePool(ctx, req.(*GetDeploymentResourcePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentResourcePoolService_ListDeploymentResourcePools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeploymentResourcePoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentResourcePoolServiceServer).ListDeploymentResourcePools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/ListDeploymentResourcePools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentResourcePoolServiceServer).ListDeploymentResourcePools(ctx, req.(*ListDeploymentResourcePoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentResourcePoolService_DeleteDeploymentResourcePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeploymentResourcePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentResourcePoolServiceServer).DeleteDeploymentResourcePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/DeleteDeploymentResourcePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentResourcePoolServiceServer).DeleteDeploymentResourcePool(ctx, req.(*DeleteDeploymentResourcePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeploymentResourcePoolService_QueryDeployedModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeployedModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentResourcePoolServiceServer).QueryDeployedModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService/QueryDeployedModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentResourcePoolServiceServer).QueryDeployedModels(ctx, req.(*QueryDeployedModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeploymentResourcePoolService_ServiceDesc is the grpc.ServiceDesc for DeploymentResourcePoolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeploymentResourcePoolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.aiplatform.v1beta1.DeploymentResourcePoolService",
	HandlerType: (*DeploymentResourcePoolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeploymentResourcePool",
			Handler:    _DeploymentResourcePoolService_CreateDeploymentResourcePool_Handler,
		},
		{
			MethodName: "GetDeploymentResourcePool",
			Handler:    _DeploymentResourcePoolService_GetDeploymentResourcePool_Handler,
		},
		{
			MethodName: "ListDeploymentResourcePools",
			Handler:    _DeploymentResourcePoolService_ListDeploymentResourcePools_Handler,
		},
		{
			MethodName: "DeleteDeploymentResourcePool",
			Handler:    _DeploymentResourcePoolService_DeleteDeploymentResourcePool_Handler,
		},
		{
			MethodName: "QueryDeployedModels",
			Handler:    _DeploymentResourcePoolService_QueryDeployedModels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/aiplatform/v1beta1/deployment_resource_pool_service.proto",
}