package toffee

import (
	"context"

	"connectrpc.com/connect"

	pb "github.com/toffeepay/sdk-go/internal/gen/pay/v1"
	"github.com/toffeepay/sdk-go/internal/gen/pay/v1/v1connect"
)

// SessionService provides operations on payment sessions.
type SessionService struct {
	client v1connect.PaymentServiceClient
}

// Create initiates a new payment session.
func (s *SessionService) Create(ctx context.Context, req *CreateSessionRequest, opts ...RequestOption) (*Session, error) {
	o := applyOptions(opts)
	r := connect.NewRequest(req)
	applyHeaders(r, o)
	resp, err := s.client.CreateSession(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.Msg.GetSession(), nil
}

// Get retrieves a session by ID.
func (s *SessionService) Get(ctx context.Context, id string) (*Session, error) {
	resp, err := s.client.GetSession(ctx, connect.NewRequest(&pb.GetSessionRequest{Id: id}))
	if err != nil {
		return nil, err
	}
	return resp.Msg.GetSession(), nil
}

// Status retrieves just the status of a session.
func (s *SessionService) Status(ctx context.Context, id string) (*GetSessionStatusResponse, error) {
	resp, err := s.client.GetSessionStatus(ctx, connect.NewRequest(&pb.GetSessionStatusRequest{Id: id}))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

// List returns a paginated list of sessions.
func (s *SessionService) List(ctx context.Context, req *ListSessionsRequest) (*ListSessionsResponse, error) {
	resp, err := s.client.ListSessions(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

// Cancel cancels a pending session.
func (s *SessionService) Cancel(ctx context.Context, id string) error {
	_, err := s.client.CancelSession(ctx, connect.NewRequest(&pb.CancelSessionRequest{Id: id}))
	return err
}
