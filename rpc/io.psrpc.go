// Code generated by protoc-gen-psrpc v0.5.1, DO NOT EDIT.
// source: rpc/io.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/rand"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import google_protobuf "google.golang.org/protobuf/types/known/emptypb"
import livekit2 "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_5

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	// egress
	CreateEgress(ctx context.Context, req *livekit2.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	UpdateEgress(ctx context.Context, req *livekit2.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)

	GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error)

	ListEgress(ctx context.Context, req *livekit2.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit2.ListEgressResponse, error)

	// ingress
	GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error)

	UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error)
}

// ===========================
// IOInfo ServerImpl Interface
// ===========================

type IOInfoServerImpl interface {
	// egress
	CreateEgress(context.Context, *livekit2.EgressInfo) (*google_protobuf.Empty, error)

	UpdateEgress(context.Context, *livekit2.EgressInfo) (*google_protobuf.Empty, error)

	GetEgress(context.Context, *GetEgressRequest) (*livekit2.EgressInfo, error)

	ListEgress(context.Context, *livekit2.ListEgressRequest) (*livekit2.ListEgressResponse, error)

	// ingress
	GetIngressInfo(context.Context, *GetIngressInfoRequest) (*GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *UpdateIngressStateRequest) (*google_protobuf.Empty, error)
}

// =======================
// IOInfo Server Interface
// =======================

type IOInfoServer interface {

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// IOInfo Client
// =============

type iOInfoClient struct {
	client *client.RPCClient
}

// NewIOInfoClient creates a psrpc client that implements the IOInfoClient interface.
func NewIOInfoClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IOInfoClient, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	sd.RegisterMethod("GetEgress", false, false, true, true)
	sd.RegisterMethod("ListEgress", false, false, true, true)
	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	sd.RegisterMethod("UpdateIngressState", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &iOInfoClient{
		client: rpcClient,
	}, nil
}

func (c *iOInfoClient) CreateEgress(ctx context.Context, req *livekit2.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "CreateEgress", nil, req, opts...)
}

func (c *iOInfoClient) UpdateEgress(ctx context.Context, req *livekit2.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetEgress(ctx context.Context, req *GetEgressRequest, opts ...psrpc.RequestOption) (*livekit2.EgressInfo, error) {
	return client.RequestSingle[*livekit2.EgressInfo](ctx, c.client, "GetEgress", nil, req, opts...)
}

func (c *iOInfoClient) ListEgress(ctx context.Context, req *livekit2.ListEgressRequest, opts ...psrpc.RequestOption) (*livekit2.ListEgressResponse, error) {
	return client.RequestSingle[*livekit2.ListEgressResponse](ctx, c.client, "ListEgress", nil, req, opts...)
}

func (c *iOInfoClient) GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error) {
	return client.RequestSingle[*GetIngressInfoResponse](ctx, c.client, "GetIngressInfo", nil, req, opts...)
}

func (c *iOInfoClient) UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf.Empty, error) {
	return client.RequestSingle[*google_protobuf.Empty](ctx, c.client, "UpdateIngressState", nil, req, opts...)
}

// =============
// IOInfo Server
// =============

type iOInfoServer struct {
	svc IOInfoServerImpl
	rpc *server.RPCServer
}

// NewIOInfoServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIOInfoServer(svc IOInfoServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IOInfoServer, error) {
	sd := &info.ServiceDefinition{
		Name: "IOInfo",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CreateEgress", false, false, true, true)
	var err error
	err = server.RegisterHandler(s, "CreateEgress", nil, svc.CreateEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateEgress", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateEgress", nil, svc.UpdateEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetEgress", false, false, true, true)
	err = server.RegisterHandler(s, "GetEgress", nil, svc.GetEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("ListEgress", false, false, true, true)
	err = server.RegisterHandler(s, "ListEgress", nil, svc.ListEgress, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("GetIngressInfo", false, false, true, true)
	err = server.RegisterHandler(s, "GetIngressInfo", nil, svc.GetIngressInfo, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("UpdateIngressState", false, false, true, true)
	err = server.RegisterHandler(s, "UpdateIngressState", nil, svc.UpdateIngressState, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &iOInfoServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *iOInfoServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *iOInfoServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor3 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x67, 0xad, 0x5b, 0xcc, 0xdb, 0x45, 0x64, 0x6c, 0x96, 0x9a, 0xa0, 0x48, 0x4e, 0x8b, 0xca,
	0x04, 0xea, 0xc1, 0x83, 0x37, 0x25, 0x2c, 0xc1, 0x05, 0x61, 0xa5, 0x17, 0x2f, 0x21, 0x4d, 0x5e,
	0xe3, 0x90, 0x34, 0x33, 0xce, 0x4c, 0x5a, 0xfa, 0x59, 0xfc, 0xb2, 0x92, 0xcc, 0x24, 0xad, 0xfd,
	0x03, 0xb2, 0xa7, 0x30, 0xbf, 0x7f, 0xf3, 0xde, 0xfc, 0x02, 0xd7, 0x52, 0x64, 0x21, 0xe3, 0x54,
	0x48, 0xae, 0x39, 0x19, 0x49, 0x91, 0x79, 0x93, 0x8a, 0xad, 0xb1, 0x64, 0x3a, 0xc1, 0x42, 0xa2,
	0x52, 0x86, 0xf2, 0xdc, 0x1e, 0x65, 0xf5, 0x3e, 0xec, 0x17, 0x9c, 0x17, 0x15, 0x86, 0xdd, 0x69,
	0xd1, 0x2c, 0x43, 0x5c, 0x09, 0xbd, 0x35, 0x64, 0x10, 0xc2, 0x8b, 0x3b, 0xd4, 0x51, 0xa7, 0x7f,
	0xc0, 0xdf, 0x0d, 0x2a, 0x4d, 0x7c, 0x70, 0x4c, 0x6e, 0xc2, 0xf2, 0xe9, 0xc5, 0xdb, 0x8b, 0x5b,
	0xe7, 0xe1, 0x99, 0x01, 0xe2, 0x3c, 0x98, 0x83, 0x7b, 0x87, 0x3a, 0x36, 0x37, 0xc4, 0xf5, 0x92,
	0xf7, 0xae, 0xd7, 0x00, 0xf6, 0xde, 0x9d, 0xcd, 0xb1, 0x48, 0x9c, 0xb7, 0xb4, 0xd2, 0x12, 0xd3,
	0x55, 0x52, 0xe2, 0x76, 0xfa, 0xc4, 0xd0, 0x06, 0xf9, 0x86, 0xdb, 0x80, 0xc3, 0xcd, 0x61, 0xac,
	0x12, 0xbc, 0x56, 0x48, 0x6e, 0xe1, 0x29, 0xab, 0x97, 0xbc, 0x4b, 0xbc, 0x9a, 0x4d, 0xa8, 0x5d,
	0x92, 0xee, 0x6b, 0x3b, 0x05, 0x99, 0xc0, 0xa5, 0xe6, 0x25, 0xd6, 0x36, 0xdd, 0x1c, 0x88, 0x0b,
	0xe3, 0x8d, 0x4a, 0x1a, 0x59, 0x4d, 0x47, 0x06, 0xde, 0xa8, 0xb9, 0xac, 0x82, 0x02, 0x5e, 0xcd,
	0x45, 0x9e, 0x6a, 0xb4, 0x39, 0x3f, 0x74, 0xaa, 0xf1, 0x3f, 0x77, 0x79, 0x0f, 0x97, 0xaa, 0x95,
	0x77, 0x17, 0x5d, 0xcd, 0xdc, 0xc3, 0x99, 0x4c, 0x96, 0xd1, 0xcc, 0xfe, 0x8c, 0x60, 0x1c, 0x7f,
	0x6f, 0xc7, 0x24, 0x9f, 0xe1, 0xfa, 0xab, 0xc4, 0x54, 0xa3, 0x79, 0x6f, 0xf2, 0x72, 0x30, 0x46,
	0xc3, 0x2e, 0xde, 0x0d, 0x35, 0x7d, 0xd1, 0xbe, 0x2f, 0x1a, 0xb5, 0x7d, 0xb5, 0x66, 0x33, 0xf0,
	0x63, 0xcc, 0x9f, 0xc0, 0x19, 0x6a, 0x26, 0x2e, 0x95, 0x22, 0xa3, 0x87, 0xb5, 0x7b, 0xa7, 0x02,
	0x49, 0x04, 0x70, 0xcf, 0x54, 0xef, 0xf4, 0x06, 0xc9, 0x0e, 0xec, 0xed, 0xfe, 0x49, 0xce, 0x96,
	0x18, 0xc3, 0xf3, 0x7f, 0xeb, 0x25, 0x5e, 0x3f, 0xc4, 0xf1, 0xaf, 0xe4, 0xf9, 0x27, 0x39, 0x1b,
	0x75, 0x0f, 0xe4, 0xb8, 0x38, 0xf2, 0xa6, 0xb3, 0x9c, 0x6d, 0xf4, 0xdc, 0xc3, 0x7c, 0xf9, 0xf0,
	0xf3, 0x5d, 0xc1, 0xf4, 0xaf, 0x66, 0x41, 0x33, 0xbe, 0x0a, 0xed, 0x06, 0xc3, 0x57, 0x94, 0x45,
	0xa8, 0x50, 0xae, 0x59, 0x86, 0xa1, 0x14, 0xd9, 0x62, 0xdc, 0xb9, 0x3f, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0x9b, 0xd6, 0xcd, 0x93, 0x03, 0x00, 0x00,
}
