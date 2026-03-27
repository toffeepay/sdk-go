package toffee

import (
	"context"

	"connectrpc.com/connect"

	pb "github.com/toffeepay/sdk-go/internal/gen/pay/v1"
	"github.com/toffeepay/sdk-go/internal/gen/pay/v1/v1connect"
)

// RefundService provides operations on refunds.
type RefundService struct {
	client v1connect.RefundServiceClient
}

// Create initiates a new refund.
func (r *RefundService) Create(ctx context.Context, req *CreateRefundRequest, opts ...RequestOption) (*Refund, error) {
	o := applyOptions(opts)
	cr := connect.NewRequest(req)
	applyHeaders(cr, o)
	resp, err := r.client.CreateRefund(ctx, cr)
	if err != nil {
		return nil, err
	}
	return resp.Msg.GetRefund(), nil
}

// Get retrieves a refund by ID.
func (r *RefundService) Get(ctx context.Context, id string) (*Refund, error) {
	resp, err := r.client.GetRefund(ctx, connect.NewRequest(&pb.GetRefundRequest{Id: id}))
	if err != nil {
		return nil, err
	}
	return resp.Msg.GetRefund(), nil
}

// List returns a paginated list of refunds.
func (r *RefundService) List(ctx context.Context, req *ListRefundsRequest) (*ListRefundsResponse, error) {
	resp, err := r.client.ListRefunds(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}
