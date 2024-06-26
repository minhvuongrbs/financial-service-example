// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: promotion/service.proto

package promotion

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
	Promotion_JoinCampaign_FullMethodName = "/promotion.service.Promotion/JoinCampaign"
)

// PromotionClient is the client API for Promotion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromotionClient interface {
	JoinCampaign(ctx context.Context, in *JoinCampaignRequest, opts ...grpc.CallOption) (*JoinCampaignReply, error)
}

type promotionClient struct {
	cc grpc.ClientConnInterface
}

func NewPromotionClient(cc grpc.ClientConnInterface) PromotionClient {
	return &promotionClient{cc}
}

func (c *promotionClient) JoinCampaign(ctx context.Context, in *JoinCampaignRequest, opts ...grpc.CallOption) (*JoinCampaignReply, error) {
	out := new(JoinCampaignReply)
	err := c.cc.Invoke(ctx, Promotion_JoinCampaign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionServer is the server API for Promotion service.
// All implementations must embed UnimplementedPromotionServer
// for forward compatibility
type PromotionServer interface {
	JoinCampaign(context.Context, *JoinCampaignRequest) (*JoinCampaignReply, error)
	mustEmbedUnimplementedPromotionServer()
}

// UnimplementedPromotionServer must be embedded to have forward compatible implementations.
type UnimplementedPromotionServer struct {
}

func (UnimplementedPromotionServer) JoinCampaign(context.Context, *JoinCampaignRequest) (*JoinCampaignReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCampaign not implemented")
}
func (UnimplementedPromotionServer) mustEmbedUnimplementedPromotionServer() {}

// UnsafePromotionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromotionServer will
// result in compilation errors.
type UnsafePromotionServer interface {
	mustEmbedUnimplementedPromotionServer()
}

func RegisterPromotionServer(s grpc.ServiceRegistrar, srv PromotionServer) {
	s.RegisterService(&Promotion_ServiceDesc, srv)
}

func _Promotion_JoinCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServer).JoinCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promotion_JoinCampaign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServer).JoinCampaign(ctx, req.(*JoinCampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Promotion_ServiceDesc is the grpc.ServiceDesc for Promotion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Promotion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "promotion.service.Promotion",
	HandlerType: (*PromotionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinCampaign",
			Handler:    _Promotion_JoinCampaign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promotion/service.proto",
}
