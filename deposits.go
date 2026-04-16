package toffee

import (
	"context"

	"connectrpc.com/connect"

	"github.com/toffeepay/sdk-go/internal/gen/wallet/v1/v1connect"
)

// DepositService provides operations on wallet deposits.
type DepositService struct {
	client v1connect.AccountServiceClient
}

// Get retrieves a deposit by ID.
func (d *DepositService) Get(ctx context.Context, req *GetDepositRequest) (*GetDepositResponse, error) {
	resp, err := d.client.GetDeposit(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

// List returns a paginated list of deposits.
func (d *DepositService) List(ctx context.Context, req *ListDepositsRequest) (*ListDepositsResponse, error) {
	resp, err := d.client.ListDeposits(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}
