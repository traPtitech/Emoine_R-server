// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: emoine_r/v1/general_api.proto

package emoine_rv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// GeneralAPIServiceName is the fully-qualified name of the GeneralAPIService service.
	GeneralAPIServiceName = "emoine_r.v1.GeneralAPIService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GeneralAPIServiceGetEventsProcedure is the fully-qualified name of the GeneralAPIService's
	// GetEvents RPC.
	GeneralAPIServiceGetEventsProcedure = "/emoine_r.v1.GeneralAPIService/GetEvents"
	// GeneralAPIServiceGetEventProcedure is the fully-qualified name of the GeneralAPIService's
	// GetEvent RPC.
	GeneralAPIServiceGetEventProcedure = "/emoine_r.v1.GeneralAPIService/GetEvent"
	// GeneralAPIServiceGetEventCommentsProcedure is the fully-qualified name of the GeneralAPIService's
	// GetEventComments RPC.
	GeneralAPIServiceGetEventCommentsProcedure = "/emoine_r.v1.GeneralAPIService/GetEventComments"
	// GeneralAPIServiceGetEventReactionsProcedure is the fully-qualified name of the
	// GeneralAPIService's GetEventReactions RPC.
	GeneralAPIServiceGetEventReactionsProcedure = "/emoine_r.v1.GeneralAPIService/GetEventReactions"
	// GeneralAPIServiceConnectToEventStreamProcedure is the fully-qualified name of the
	// GeneralAPIService's ConnectToEventStream RPC.
	GeneralAPIServiceConnectToEventStreamProcedure = "/emoine_r.v1.GeneralAPIService/ConnectToEventStream"
	// GeneralAPIServiceSendCommentProcedure is the fully-qualified name of the GeneralAPIService's
	// SendComment RPC.
	GeneralAPIServiceSendCommentProcedure = "/emoine_r.v1.GeneralAPIService/SendComment"
	// GeneralAPIServiceSendReactionProcedure is the fully-qualified name of the GeneralAPIService's
	// SendReaction RPC.
	GeneralAPIServiceSendReactionProcedure = "/emoine_r.v1.GeneralAPIService/SendReaction"
)

// GeneralAPIServiceClient is a client for the emoine_r.v1.GeneralAPIService service.
type GeneralAPIServiceClient interface {
	// イベント一覧を取得します
	GetEvents(context.Context, *connect.Request[v1.GetEventsRequest]) (*connect.Response[v1.GetEventsResponse], error)
	// 該当するイベントを取得します
	GetEvent(context.Context, *connect.Request[v1.GetEventRequest]) (*connect.Response[v1.GetEventResponse], error)
	// 該当するイベントのコメント一覧を取得します
	GetEventComments(context.Context, *connect.Request[v1.GetEventCommentsRequest]) (*connect.Response[v1.GetEventCommentsResponse], error)
	// 該当するイベントのリアクション一覧を取得します
	GetEventReactions(context.Context, *connect.Request[v1.GetEventReactionsRequest]) (*connect.Response[v1.GetEventReactionsResponse], error)
	// イベントのストリームに接続します
	ConnectToEventStream(context.Context, *connect.Request[v1.ConnectToEventStreamRequest]) (*connect.ServerStreamForClient[v1.ConnectToEventStreamResponse], error)
	// (コメントはイベントのストリームに反映されます)
	SendComment(context.Context, *connect.Request[v1.SendCommentRequest]) (*connect.Response[v1.SendCommentResponse], error)
	// (リアクションはイベントのストリームに反映されます)
	SendReaction(context.Context, *connect.Request[v1.SendReactionRequest]) (*connect.Response[v1.SendReactionResponse], error)
}

// NewGeneralAPIServiceClient constructs a client for the emoine_r.v1.GeneralAPIService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGeneralAPIServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GeneralAPIServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &generalAPIServiceClient{
		getEvents: connect.NewClient[v1.GetEventsRequest, v1.GetEventsResponse](
			httpClient,
			baseURL+GeneralAPIServiceGetEventsProcedure,
			opts...,
		),
		getEvent: connect.NewClient[v1.GetEventRequest, v1.GetEventResponse](
			httpClient,
			baseURL+GeneralAPIServiceGetEventProcedure,
			opts...,
		),
		getEventComments: connect.NewClient[v1.GetEventCommentsRequest, v1.GetEventCommentsResponse](
			httpClient,
			baseURL+GeneralAPIServiceGetEventCommentsProcedure,
			opts...,
		),
		getEventReactions: connect.NewClient[v1.GetEventReactionsRequest, v1.GetEventReactionsResponse](
			httpClient,
			baseURL+GeneralAPIServiceGetEventReactionsProcedure,
			opts...,
		),
		connectToEventStream: connect.NewClient[v1.ConnectToEventStreamRequest, v1.ConnectToEventStreamResponse](
			httpClient,
			baseURL+GeneralAPIServiceConnectToEventStreamProcedure,
			opts...,
		),
		sendComment: connect.NewClient[v1.SendCommentRequest, v1.SendCommentResponse](
			httpClient,
			baseURL+GeneralAPIServiceSendCommentProcedure,
			opts...,
		),
		sendReaction: connect.NewClient[v1.SendReactionRequest, v1.SendReactionResponse](
			httpClient,
			baseURL+GeneralAPIServiceSendReactionProcedure,
			opts...,
		),
	}
}

// generalAPIServiceClient implements GeneralAPIServiceClient.
type generalAPIServiceClient struct {
	getEvents            *connect.Client[v1.GetEventsRequest, v1.GetEventsResponse]
	getEvent             *connect.Client[v1.GetEventRequest, v1.GetEventResponse]
	getEventComments     *connect.Client[v1.GetEventCommentsRequest, v1.GetEventCommentsResponse]
	getEventReactions    *connect.Client[v1.GetEventReactionsRequest, v1.GetEventReactionsResponse]
	connectToEventStream *connect.Client[v1.ConnectToEventStreamRequest, v1.ConnectToEventStreamResponse]
	sendComment          *connect.Client[v1.SendCommentRequest, v1.SendCommentResponse]
	sendReaction         *connect.Client[v1.SendReactionRequest, v1.SendReactionResponse]
}

// GetEvents calls emoine_r.v1.GeneralAPIService.GetEvents.
func (c *generalAPIServiceClient) GetEvents(ctx context.Context, req *connect.Request[v1.GetEventsRequest]) (*connect.Response[v1.GetEventsResponse], error) {
	return c.getEvents.CallUnary(ctx, req)
}

// GetEvent calls emoine_r.v1.GeneralAPIService.GetEvent.
func (c *generalAPIServiceClient) GetEvent(ctx context.Context, req *connect.Request[v1.GetEventRequest]) (*connect.Response[v1.GetEventResponse], error) {
	return c.getEvent.CallUnary(ctx, req)
}

// GetEventComments calls emoine_r.v1.GeneralAPIService.GetEventComments.
func (c *generalAPIServiceClient) GetEventComments(ctx context.Context, req *connect.Request[v1.GetEventCommentsRequest]) (*connect.Response[v1.GetEventCommentsResponse], error) {
	return c.getEventComments.CallUnary(ctx, req)
}

// GetEventReactions calls emoine_r.v1.GeneralAPIService.GetEventReactions.
func (c *generalAPIServiceClient) GetEventReactions(ctx context.Context, req *connect.Request[v1.GetEventReactionsRequest]) (*connect.Response[v1.GetEventReactionsResponse], error) {
	return c.getEventReactions.CallUnary(ctx, req)
}

// ConnectToEventStream calls emoine_r.v1.GeneralAPIService.ConnectToEventStream.
func (c *generalAPIServiceClient) ConnectToEventStream(ctx context.Context, req *connect.Request[v1.ConnectToEventStreamRequest]) (*connect.ServerStreamForClient[v1.ConnectToEventStreamResponse], error) {
	return c.connectToEventStream.CallServerStream(ctx, req)
}

// SendComment calls emoine_r.v1.GeneralAPIService.SendComment.
func (c *generalAPIServiceClient) SendComment(ctx context.Context, req *connect.Request[v1.SendCommentRequest]) (*connect.Response[v1.SendCommentResponse], error) {
	return c.sendComment.CallUnary(ctx, req)
}

// SendReaction calls emoine_r.v1.GeneralAPIService.SendReaction.
func (c *generalAPIServiceClient) SendReaction(ctx context.Context, req *connect.Request[v1.SendReactionRequest]) (*connect.Response[v1.SendReactionResponse], error) {
	return c.sendReaction.CallUnary(ctx, req)
}

// GeneralAPIServiceHandler is an implementation of the emoine_r.v1.GeneralAPIService service.
type GeneralAPIServiceHandler interface {
	// イベント一覧を取得します
	GetEvents(context.Context, *connect.Request[v1.GetEventsRequest]) (*connect.Response[v1.GetEventsResponse], error)
	// 該当するイベントを取得します
	GetEvent(context.Context, *connect.Request[v1.GetEventRequest]) (*connect.Response[v1.GetEventResponse], error)
	// 該当するイベントのコメント一覧を取得します
	GetEventComments(context.Context, *connect.Request[v1.GetEventCommentsRequest]) (*connect.Response[v1.GetEventCommentsResponse], error)
	// 該当するイベントのリアクション一覧を取得します
	GetEventReactions(context.Context, *connect.Request[v1.GetEventReactionsRequest]) (*connect.Response[v1.GetEventReactionsResponse], error)
	// イベントのストリームに接続します
	ConnectToEventStream(context.Context, *connect.Request[v1.ConnectToEventStreamRequest], *connect.ServerStream[v1.ConnectToEventStreamResponse]) error
	// (コメントはイベントのストリームに反映されます)
	SendComment(context.Context, *connect.Request[v1.SendCommentRequest]) (*connect.Response[v1.SendCommentResponse], error)
	// (リアクションはイベントのストリームに反映されます)
	SendReaction(context.Context, *connect.Request[v1.SendReactionRequest]) (*connect.Response[v1.SendReactionResponse], error)
}

// NewGeneralAPIServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGeneralAPIServiceHandler(svc GeneralAPIServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	generalAPIServiceGetEventsHandler := connect.NewUnaryHandler(
		GeneralAPIServiceGetEventsProcedure,
		svc.GetEvents,
		opts...,
	)
	generalAPIServiceGetEventHandler := connect.NewUnaryHandler(
		GeneralAPIServiceGetEventProcedure,
		svc.GetEvent,
		opts...,
	)
	generalAPIServiceGetEventCommentsHandler := connect.NewUnaryHandler(
		GeneralAPIServiceGetEventCommentsProcedure,
		svc.GetEventComments,
		opts...,
	)
	generalAPIServiceGetEventReactionsHandler := connect.NewUnaryHandler(
		GeneralAPIServiceGetEventReactionsProcedure,
		svc.GetEventReactions,
		opts...,
	)
	generalAPIServiceConnectToEventStreamHandler := connect.NewServerStreamHandler(
		GeneralAPIServiceConnectToEventStreamProcedure,
		svc.ConnectToEventStream,
		opts...,
	)
	generalAPIServiceSendCommentHandler := connect.NewUnaryHandler(
		GeneralAPIServiceSendCommentProcedure,
		svc.SendComment,
		opts...,
	)
	generalAPIServiceSendReactionHandler := connect.NewUnaryHandler(
		GeneralAPIServiceSendReactionProcedure,
		svc.SendReaction,
		opts...,
	)
	return "/emoine_r.v1.GeneralAPIService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GeneralAPIServiceGetEventsProcedure:
			generalAPIServiceGetEventsHandler.ServeHTTP(w, r)
		case GeneralAPIServiceGetEventProcedure:
			generalAPIServiceGetEventHandler.ServeHTTP(w, r)
		case GeneralAPIServiceGetEventCommentsProcedure:
			generalAPIServiceGetEventCommentsHandler.ServeHTTP(w, r)
		case GeneralAPIServiceGetEventReactionsProcedure:
			generalAPIServiceGetEventReactionsHandler.ServeHTTP(w, r)
		case GeneralAPIServiceConnectToEventStreamProcedure:
			generalAPIServiceConnectToEventStreamHandler.ServeHTTP(w, r)
		case GeneralAPIServiceSendCommentProcedure:
			generalAPIServiceSendCommentHandler.ServeHTTP(w, r)
		case GeneralAPIServiceSendReactionProcedure:
			generalAPIServiceSendReactionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGeneralAPIServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGeneralAPIServiceHandler struct{}

func (UnimplementedGeneralAPIServiceHandler) GetEvents(context.Context, *connect.Request[v1.GetEventsRequest]) (*connect.Response[v1.GetEventsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("emoine_r.v1.GeneralAPIService.GetEvents is not implemented"))
}

func (UnimplementedGeneralAPIServiceHandler) GetEvent(context.Context, *connect.Request[v1.GetEventRequest]) (*connect.Response[v1.GetEventResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("emoine_r.v1.GeneralAPIService.GetEvent is not implemented"))
}

func (UnimplementedGeneralAPIServiceHandler) GetEventComments(context.Context, *connect.Request[v1.GetEventCommentsRequest]) (*connect.Response[v1.GetEventCommentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("emoine_r.v1.GeneralAPIService.GetEventComments is not implemented"))
}

func (UnimplementedGeneralAPIServiceHandler) GetEventReactions(context.Context, *connect.Request[v1.GetEventReactionsRequest]) (*connect.Response[v1.GetEventReactionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("emoine_r.v1.GeneralAPIService.GetEventReactions is not implemented"))
}

func (UnimplementedGeneralAPIServiceHandler) ConnectToEventStream(context.Context, *connect.Request[v1.ConnectToEventStreamRequest], *connect.ServerStream[v1.ConnectToEventStreamResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("emoine_r.v1.GeneralAPIService.ConnectToEventStream is not implemented"))
}

func (UnimplementedGeneralAPIServiceHandler) SendComment(context.Context, *connect.Request[v1.SendCommentRequest]) (*connect.Response[v1.SendCommentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("emoine_r.v1.GeneralAPIService.SendComment is not implemented"))
}

func (UnimplementedGeneralAPIServiceHandler) SendReaction(context.Context, *connect.Request[v1.SendReactionRequest]) (*connect.Response[v1.SendReactionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("emoine_r.v1.GeneralAPIService.SendReaction is not implemented"))
}
