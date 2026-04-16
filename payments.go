package toffee

import (
	"context"

	"connectrpc.com/connect"

	"github.com/toffeepay/sdk-go/internal/gen/pay/v1/v1connect"
)

// PaymentService provides operations on payments.
type PaymentService struct {
	client v1connect.PaymentServiceClient
}

// Get retrieves a payment by ID.
func (p *PaymentService) Get(ctx context.Context, req *GetPaymentRequest) (*GetPaymentResponse, error) {
	resp, err := p.client.GetPayment(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
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
func (p *PaymentService) Complete(ctx context.Context, req *CompletePaymentRequest) (*CompletePaymentResponse, error) {
	resp, err := p.client.CompletePayment(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

// Cancel cancels an authorized payment.
func (p *PaymentService) Cancel(ctx context.Context, req *CancelPaymentRequest) (*CancelPaymentResponse, error) {
	resp, err := p.client.CancelPayment(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}
