// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: emoine_r/v1/admin_api.proto

package emoine_rv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/traPtitech/Emoine_R/pkg/bufgen/emoine_r/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AdminAPIServiceName is the fully-qualified name of the AdminAPIService service.
	AdminAPIServiceName = "emoine_r.v1.AdminAPIService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AdminAPIServiceCreateMeetingProcedure is the fully-qualified name of the AdminAPIService's
	// CreateMeeting RPC.
	AdminAPIServiceCreateMeetingProcedure = "/emoine_r.v1.AdminAPIService/CreateMeeting"
	// AdminAPIServiceUpdateMeetingProcedure is the fully-qualified name of the AdminAPIService's
	// UpdateMeeting RPC.
	AdminAPIServiceUpdateMeetingProcedure = "/emoine_r.v1.AdminAPIService/UpdateMeeting"
	// AdminAPIServiceDeleteMeetingProcedure is the fully-qualified name of the AdminAPIService's
	// DeleteMeeting RPC.
	AdminAPIServiceDeleteMeetingProcedure = "/emoine_r.v1.AdminAPIService/DeleteMeeting"
	// AdminAPIServiceGetTokensProcedure is the fully-qualified name of the AdminAPIService's GetTokens
	// RPC.
	AdminAPIServiceGetTokensProcedure = "/emoine_r.v1.AdminAPIService/GetTokens"
	// AdminAPIServiceCreateTokenProcedure is the fully-qualified name of the AdminAPIService's
	// CreateToken RPC.
	AdminAPIServiceCreateTokenProcedure = "/emoine_r.v1.AdminAPIService/CreateToken"
	// AdminAPIServiceUpdateTokenProcedure is the fully-qualified name of the AdminAPIService's
	// UpdateToken RPC.
	AdminAPIServiceUpdateTokenProcedure = "/emoine_r.v1.AdminAPIService/UpdateToken"
)

// AdminAPIServiceClient is a client for the emoine_r.v1.AdminAPIService service.
type AdminAPIServiceClient interface {
	// 集会を作成します
	CreateMeeting(context.Context, *connect_go.Request[v1.CreateMeetingRequest]) (*connect_go.Response[v1.CreateMeetingResponse], error)
	// 集会情報を更新します
	UpdateMeeting(context.Context, *connect_go.Request[v1.UpdateMeetingRequest]) (*connect_go.Response[emptypb.Empty], error)
	// 集会を削除します
	DeleteMeeting(context.Context, *connect_go.Request[v1.DeleteMeetingRequest]) (*connect_go.Response[emptypb.Empty], error)
	// 該当する集会のトークン一覧を取得します
	GetTokens(context.Context, *connect_go.Request[v1.GetTokensRequest]) (*connect_go.Response[v1.GetTokensResponse], error)
	// 集会のトークンを作成します
	CreateToken(context.Context, *connect_go.Request[v1.CreateTokenRequest]) (*connect_go.Response[v1.CreateTokenResponse], error)
	// 集会のトークン情報を更新します
	UpdateToken(context.Context, *connect_go.Request[v1.UpdateTokenRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewAdminAPIServiceClient constructs a client for the emoine_r.v1.AdminAPIService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAdminAPIServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AdminAPIServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &adminAPIServiceClient{
		createMeeting: connect_go.NewClient[v1.CreateMeetingRequest, v1.CreateMeetingResponse](
			httpClient,
			baseURL+AdminAPIServiceCreateMeetingProcedure,
			opts...,
		),
		updateMeeting: connect_go.NewClient[v1.UpdateMeetingRequest, emptypb.Empty](
			httpClient,
			baseURL+AdminAPIServiceUpdateMeetingProcedure,
			opts...,
		),
		deleteMeeting: connect_go.NewClient[v1.DeleteMeetingRequest, emptypb.Empty](
			httpClient,
			baseURL+AdminAPIServiceDeleteMeetingProcedure,
			opts...,
		),
		getTokens: connect_go.NewClient[v1.GetTokensRequest, v1.GetTokensResponse](
			httpClient,
			baseURL+AdminAPIServiceGetTokensProcedure,
			opts...,
		),
		createToken: connect_go.NewClient[v1.CreateTokenRequest, v1.CreateTokenResponse](
			httpClient,
			baseURL+AdminAPIServiceCreateTokenProcedure,
			opts...,
		),
		updateToken: connect_go.NewClient[v1.UpdateTokenRequest, emptypb.Empty](
			httpClient,
			baseURL+AdminAPIServiceUpdateTokenProcedure,
			opts...,
		),
	}
}

// adminAPIServiceClient implements AdminAPIServiceClient.
type adminAPIServiceClient struct {
	createMeeting *connect_go.Client[v1.CreateMeetingRequest, v1.CreateMeetingResponse]
	updateMeeting *connect_go.Client[v1.UpdateMeetingRequest, emptypb.Empty]
	deleteMeeting *connect_go.Client[v1.DeleteMeetingRequest, emptypb.Empty]
	getTokens     *connect_go.Client[v1.GetTokensRequest, v1.GetTokensResponse]
	createToken   *connect_go.Client[v1.CreateTokenRequest, v1.CreateTokenResponse]
	updateToken   *connect_go.Client[v1.UpdateTokenRequest, emptypb.Empty]
}

// CreateMeeting calls emoine_r.v1.AdminAPIService.CreateMeeting.
func (c *adminAPIServiceClient) CreateMeeting(ctx context.Context, req *connect_go.Request[v1.CreateMeetingRequest]) (*connect_go.Response[v1.CreateMeetingResponse], error) {
	return c.createMeeting.CallUnary(ctx, req)
}

// UpdateMeeting calls emoine_r.v1.AdminAPIService.UpdateMeeting.
func (c *adminAPIServiceClient) UpdateMeeting(ctx context.Context, req *connect_go.Request[v1.UpdateMeetingRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updateMeeting.CallUnary(ctx, req)
}

// DeleteMeeting calls emoine_r.v1.AdminAPIService.DeleteMeeting.
func (c *adminAPIServiceClient) DeleteMeeting(ctx context.Context, req *connect_go.Request[v1.DeleteMeetingRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteMeeting.CallUnary(ctx, req)
}

// GetTokens calls emoine_r.v1.AdminAPIService.GetTokens.
func (c *adminAPIServiceClient) GetTokens(ctx context.Context, req *connect_go.Request[v1.GetTokensRequest]) (*connect_go.Response[v1.GetTokensResponse], error) {
	return c.getTokens.CallUnary(ctx, req)
}

// CreateToken calls emoine_r.v1.AdminAPIService.CreateToken.
func (c *adminAPIServiceClient) CreateToken(ctx context.Context, req *connect_go.Request[v1.CreateTokenRequest]) (*connect_go.Response[v1.CreateTokenResponse], error) {
	return c.createToken.CallUnary(ctx, req)
}

// UpdateToken calls emoine_r.v1.AdminAPIService.UpdateToken.
func (c *adminAPIServiceClient) UpdateToken(ctx context.Context, req *connect_go.Request[v1.UpdateTokenRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updateToken.CallUnary(ctx, req)
}

// AdminAPIServiceHandler is an implementation of the emoine_r.v1.AdminAPIService service.
type AdminAPIServiceHandler interface {
	// 集会を作成します
	CreateMeeting(context.Context, *connect_go.Request[v1.CreateMeetingRequest]) (*connect_go.Response[v1.CreateMeetingResponse], error)
	// 集会情報を更新します
	UpdateMeeting(context.Context, *connect_go.Request[v1.UpdateMeetingRequest]) (*connect_go.Response[emptypb.Empty], error)
	// 集会を削除します
	DeleteMeeting(context.Context, *connect_go.Request[v1.DeleteMeetingRequest]) (*connect_go.Response[emptypb.Empty], error)
	// 該当する集会のトークン一覧を取得します
	GetTokens(context.Context, *connect_go.Request[v1.GetTokensRequest]) (*connect_go.Response[v1.GetTokensResponse], error)
	// 集会のトークンを作成します
	CreateToken(context.Context, *connect_go.Request[v1.CreateTokenRequest]) (*connect_go.Response[v1.CreateTokenResponse], error)
	// 集会のトークン情報を更新します
	UpdateToken(context.Context, *connect_go.Request[v1.UpdateTokenRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewAdminAPIServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAdminAPIServiceHandler(svc AdminAPIServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	adminAPIServiceCreateMeetingHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceCreateMeetingProcedure,
		svc.CreateMeeting,
		opts...,
	)
	adminAPIServiceUpdateMeetingHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceUpdateMeetingProcedure,
		svc.UpdateMeeting,
		opts...,
	)
	adminAPIServiceDeleteMeetingHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceDeleteMeetingProcedure,
		svc.DeleteMeeting,
		opts...,
	)
	adminAPIServiceGetTokensHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceGetTokensProcedure,
		svc.GetTokens,
		opts...,
	)
	adminAPIServiceCreateTokenHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceCreateTokenProcedure,
		svc.CreateToken,
		opts...,
	)
	adminAPIServiceUpdateTokenHandler := connect_go.NewUnaryHandler(
		AdminAPIServiceUpdateTokenProcedure,
		svc.UpdateToken,
		opts...,
	)
	return "/emoine_r.v1.AdminAPIService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AdminAPIServiceCreateMeetingProcedure:
			adminAPIServiceCreateMeetingHandler.ServeHTTP(w, r)
		case AdminAPIServiceUpdateMeetingProcedure:
			adminAPIServiceUpdateMeetingHandler.ServeHTTP(w, r)
		case AdminAPIServiceDeleteMeetingProcedure:
			adminAPIServiceDeleteMeetingHandler.ServeHTTP(w, r)
		case AdminAPIServiceGetTokensProcedure:
			adminAPIServiceGetTokensHandler.ServeHTTP(w, r)
		case AdminAPIServiceCreateTokenProcedure:
			adminAPIServiceCreateTokenHandler.ServeHTTP(w, r)
		case AdminAPIServiceUpdateTokenProcedure:
			adminAPIServiceUpdateTokenHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAdminAPIServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAdminAPIServiceHandler struct{}

func (UnimplementedAdminAPIServiceHandler) CreateMeeting(context.Context, *connect_go.Request[v1.CreateMeetingRequest]) (*connect_go.Response[v1.CreateMeetingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.CreateMeeting is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) UpdateMeeting(context.Context, *connect_go.Request[v1.UpdateMeetingRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.UpdateMeeting is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) DeleteMeeting(context.Context, *connect_go.Request[v1.DeleteMeetingRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.DeleteMeeting is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) GetTokens(context.Context, *connect_go.Request[v1.GetTokensRequest]) (*connect_go.Response[v1.GetTokensResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.GetTokens is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) CreateToken(context.Context, *connect_go.Request[v1.CreateTokenRequest]) (*connect_go.Response[v1.CreateTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.CreateToken is not implemented"))
}

func (UnimplementedAdminAPIServiceHandler) UpdateToken(context.Context, *connect_go.Request[v1.UpdateTokenRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("emoine_r.v1.AdminAPIService.UpdateToken is not implemented"))
}
