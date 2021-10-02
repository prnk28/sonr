// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package node

import (
	context "context"
	common "github.com/sonr-io/core/internal/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	// Node Methods
	// Signing Method Request for Data
	Supply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error)
	// Verification Method Request for Signed Data
	Edit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
	// Fetch method finds data from Key/Value store
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	// Respond Method to an Invite with Decision
	Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error)
	// Respond Method to an Invite with Decision
	Respond(ctx context.Context, in *RespondRequest, opts ...grpc.CallOption) (*RespondResponse, error)
	// Search Method to find a Peer by SName or PeerID
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Stat Method returns the Node Stats
	Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	// Events Streams
	// Returns a stream of StatusEvents
	OnNodeStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnNodeStatusClient, error)
	// Returns a stream of Lobby Refresh Events
	OnLobbyRefresh(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnLobbyRefreshClient, error)
	// Returns a stream of DecisionEvent's for Accepted Invites
	OnTransferAccepted(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferAcceptedClient, error)
	// Returns a stream of DecisionEvent's for Rejected Invites
	OnTransferDeclined(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferDeclinedClient, error)
	// Returns a stream of DecisionEvent's for Invites
	OnTransferInvite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferInviteClient, error)
	// Returns a stream of ProgressEvent's for Sessions
	OnTransferProgress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferProgressClient, error)
	// Returns a stream of Completed Transfers
	OnTransferComplete(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferCompleteClient, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Supply(ctx context.Context, in *SupplyRequest, opts ...grpc.CallOption) (*SupplyResponse, error) {
	out := new(SupplyResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.NodeService/Supply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Edit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	out := new(EditResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.NodeService/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.NodeService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Share(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.NodeService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Respond(ctx context.Context, in *RespondRequest, opts ...grpc.CallOption) (*RespondResponse, error) {
	out := new(RespondResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.NodeService/Respond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.NodeService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Stat(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.NodeService/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) OnNodeStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnNodeStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[0], "/sonr.node.NodeService/OnNodeStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnNodeStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnNodeStatusClient interface {
	Recv() (*common.StatusEvent, error)
	grpc.ClientStream
}

type nodeServiceOnNodeStatusClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnNodeStatusClient) Recv() (*common.StatusEvent, error) {
	m := new(common.StatusEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnLobbyRefresh(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnLobbyRefreshClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[1], "/sonr.node.NodeService/OnLobbyRefresh", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnLobbyRefreshClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnLobbyRefreshClient interface {
	Recv() (*common.RefreshEvent, error)
	grpc.ClientStream
}

type nodeServiceOnLobbyRefreshClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnLobbyRefreshClient) Recv() (*common.RefreshEvent, error) {
	m := new(common.RefreshEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnTransferAccepted(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferAcceptedClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[2], "/sonr.node.NodeService/OnTransferAccepted", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnTransferAcceptedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnTransferAcceptedClient interface {
	Recv() (*common.DecisionEvent, error)
	grpc.ClientStream
}

type nodeServiceOnTransferAcceptedClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnTransferAcceptedClient) Recv() (*common.DecisionEvent, error) {
	m := new(common.DecisionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnTransferDeclined(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferDeclinedClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[3], "/sonr.node.NodeService/OnTransferDeclined", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnTransferDeclinedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnTransferDeclinedClient interface {
	Recv() (*common.DecisionEvent, error)
	grpc.ClientStream
}

type nodeServiceOnTransferDeclinedClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnTransferDeclinedClient) Recv() (*common.DecisionEvent, error) {
	m := new(common.DecisionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnTransferInvite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferInviteClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[4], "/sonr.node.NodeService/OnTransferInvite", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnTransferInviteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnTransferInviteClient interface {
	Recv() (*common.InviteEvent, error)
	grpc.ClientStream
}

type nodeServiceOnTransferInviteClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnTransferInviteClient) Recv() (*common.InviteEvent, error) {
	m := new(common.InviteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnTransferProgress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[5], "/sonr.node.NodeService/OnTransferProgress", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnTransferProgressClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnTransferProgressClient interface {
	Recv() (*common.ProgressEvent, error)
	grpc.ClientStream
}

type nodeServiceOnTransferProgressClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnTransferProgressClient) Recv() (*common.ProgressEvent, error) {
	m := new(common.ProgressEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) OnTransferComplete(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeService_OnTransferCompleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[6], "/sonr.node.NodeService/OnTransferComplete", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceOnTransferCompleteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_OnTransferCompleteClient interface {
	Recv() (*common.CompleteEvent, error)
	grpc.ClientStream
}

type nodeServiceOnTransferCompleteClient struct {
	grpc.ClientStream
}

func (x *nodeServiceOnTransferCompleteClient) Recv() (*common.CompleteEvent, error) {
	m := new(common.CompleteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	// Node Methods
	// Signing Method Request for Data
	Supply(context.Context, *SupplyRequest) (*SupplyResponse, error)
	// Verification Method Request for Signed Data
	Edit(context.Context, *EditRequest) (*EditResponse, error)
	// Fetch method finds data from Key/Value store
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	// Respond Method to an Invite with Decision
	Share(context.Context, *ShareRequest) (*ShareResponse, error)
	// Respond Method to an Invite with Decision
	Respond(context.Context, *RespondRequest) (*RespondResponse, error)
	// Search Method to find a Peer by SName or PeerID
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// Stat Method returns the Node Stats
	Stat(context.Context, *StatRequest) (*StatResponse, error)
	// Events Streams
	// Returns a stream of StatusEvents
	OnNodeStatus(*Empty, NodeService_OnNodeStatusServer) error
	// Returns a stream of Lobby Refresh Events
	OnLobbyRefresh(*Empty, NodeService_OnLobbyRefreshServer) error
	// Returns a stream of DecisionEvent's for Accepted Invites
	OnTransferAccepted(*Empty, NodeService_OnTransferAcceptedServer) error
	// Returns a stream of DecisionEvent's for Rejected Invites
	OnTransferDeclined(*Empty, NodeService_OnTransferDeclinedServer) error
	// Returns a stream of DecisionEvent's for Invites
	OnTransferInvite(*Empty, NodeService_OnTransferInviteServer) error
	// Returns a stream of ProgressEvent's for Sessions
	OnTransferProgress(*Empty, NodeService_OnTransferProgressServer) error
	// Returns a stream of Completed Transfers
	OnTransferComplete(*Empty, NodeService_OnTransferCompleteServer) error
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) Supply(context.Context, *SupplyRequest) (*SupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
}
func (UnimplementedNodeServiceServer) Edit(context.Context, *EditRequest) (*EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedNodeServiceServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedNodeServiceServer) Share(context.Context, *ShareRequest) (*ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Share not implemented")
}
func (UnimplementedNodeServiceServer) Respond(context.Context, *RespondRequest) (*RespondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Respond not implemented")
}
func (UnimplementedNodeServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedNodeServiceServer) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedNodeServiceServer) OnNodeStatus(*Empty, NodeService_OnNodeStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method OnNodeStatus not implemented")
}
func (UnimplementedNodeServiceServer) OnLobbyRefresh(*Empty, NodeService_OnLobbyRefreshServer) error {
	return status.Errorf(codes.Unimplemented, "method OnLobbyRefresh not implemented")
}
func (UnimplementedNodeServiceServer) OnTransferAccepted(*Empty, NodeService_OnTransferAcceptedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferAccepted not implemented")
}
func (UnimplementedNodeServiceServer) OnTransferDeclined(*Empty, NodeService_OnTransferDeclinedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferDeclined not implemented")
}
func (UnimplementedNodeServiceServer) OnTransferInvite(*Empty, NodeService_OnTransferInviteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferInvite not implemented")
}
func (UnimplementedNodeServiceServer) OnTransferProgress(*Empty, NodeService_OnTransferProgressServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferProgress not implemented")
}
func (UnimplementedNodeServiceServer) OnTransferComplete(*Empty, NodeService_OnTransferCompleteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransferComplete not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Supply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.NodeService/Supply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Supply(ctx, req.(*SupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.NodeService/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Edit(ctx, req.(*EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.NodeService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.NodeService/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Share(ctx, req.(*ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Respond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Respond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.NodeService/Respond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Respond(ctx, req.(*RespondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.NodeService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.NodeService/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Stat(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_OnNodeStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnNodeStatus(m, &nodeServiceOnNodeStatusServer{stream})
}

type NodeService_OnNodeStatusServer interface {
	Send(*common.StatusEvent) error
	grpc.ServerStream
}

type nodeServiceOnNodeStatusServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnNodeStatusServer) Send(m *common.StatusEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnLobbyRefresh_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnLobbyRefresh(m, &nodeServiceOnLobbyRefreshServer{stream})
}

type NodeService_OnLobbyRefreshServer interface {
	Send(*common.RefreshEvent) error
	grpc.ServerStream
}

type nodeServiceOnLobbyRefreshServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnLobbyRefreshServer) Send(m *common.RefreshEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnTransferAccepted_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnTransferAccepted(m, &nodeServiceOnTransferAcceptedServer{stream})
}

type NodeService_OnTransferAcceptedServer interface {
	Send(*common.DecisionEvent) error
	grpc.ServerStream
}

type nodeServiceOnTransferAcceptedServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnTransferAcceptedServer) Send(m *common.DecisionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnTransferDeclined_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnTransferDeclined(m, &nodeServiceOnTransferDeclinedServer{stream})
}

type NodeService_OnTransferDeclinedServer interface {
	Send(*common.DecisionEvent) error
	grpc.ServerStream
}

type nodeServiceOnTransferDeclinedServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnTransferDeclinedServer) Send(m *common.DecisionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnTransferInvite_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnTransferInvite(m, &nodeServiceOnTransferInviteServer{stream})
}

type NodeService_OnTransferInviteServer interface {
	Send(*common.InviteEvent) error
	grpc.ServerStream
}

type nodeServiceOnTransferInviteServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnTransferInviteServer) Send(m *common.InviteEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnTransferProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnTransferProgress(m, &nodeServiceOnTransferProgressServer{stream})
}

type NodeService_OnTransferProgressServer interface {
	Send(*common.ProgressEvent) error
	grpc.ServerStream
}

type nodeServiceOnTransferProgressServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnTransferProgressServer) Send(m *common.ProgressEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_OnTransferComplete_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).OnTransferComplete(m, &nodeServiceOnTransferCompleteServer{stream})
}

type NodeService_OnTransferCompleteServer interface {
	Send(*common.CompleteEvent) error
	grpc.ServerStream
}

type nodeServiceOnTransferCompleteServer struct {
	grpc.ServerStream
}

func (x *nodeServiceOnTransferCompleteServer) Send(m *common.CompleteEvent) error {
	return x.ServerStream.SendMsg(m)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sonr.node.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Supply",
			Handler:    _NodeService_Supply_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _NodeService_Edit_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _NodeService_Fetch_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _NodeService_Share_Handler,
		},
		{
			MethodName: "Respond",
			Handler:    _NodeService_Respond_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _NodeService_Search_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _NodeService_Stat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnNodeStatus",
			Handler:       _NodeService_OnNodeStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnLobbyRefresh",
			Handler:       _NodeService_OnLobbyRefresh_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferAccepted",
			Handler:       _NodeService_OnTransferAccepted_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferDeclined",
			Handler:       _NodeService_OnTransferDeclined_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferInvite",
			Handler:       _NodeService_OnTransferInvite_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferProgress",
			Handler:       _NodeService_OnTransferProgress_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransferComplete",
			Handler:       _NodeService_OnTransferComplete_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/client/rpc.proto",
}