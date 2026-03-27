package toffee

import (
	"context"

	"connectrpc.com/connect"

	pb "github.com/toffeepay/sdk-go/internal/gen/pay/v1"
	"github.com/toffeepay/sdk-go/internal/gen/pay/v1/v1connect"
)

// PaymentService provides operations on payments.
type PaymentService struct {
	client v1connect.PaymentServiceClient
}

// Get retrieves a payment by ID.
func (p *PaymentService) Get(ctx context.Context, id string, withExtraData ...bool) (*Payment, error) {
	req := &pb.GetPaymentRequest{Id: id}
	if len(withExtraData) > 0 && withExtraData[0] {
		v := true
		req.WithExtraData = &v
	}
	resp, err := p.client.GetPayment(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg.GetPayment(), nil
}

// List returns a paginated list of payments.
func (p *PaymentService) List(ctx context.Context, req *ListPaymentsRequest) (*ListPaymentsResponse, error) {
	resp, err := p.client.ListPayments(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

// Complete captures an authorized payment.
func (p *PaymentService) Complete(ctx context.Context, id string) error {
	_, err := p.client.CompletePayment(ctx, connect.NewRequest(&pb.CompletePaymentRequest{Id: id}))
	return err
}

// Cancel cancels an authorized payment.
func (p *PaymentService) Cancel(ctx context.Context, id string) error {
	_, err := p.client.CancelPayment(ctx, connect.NewRequest(&pb.CancelPaymentRequest{Id: id}))
	return err
}
