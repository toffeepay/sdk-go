package toffee

import (
	"context"

	"connectrpc.com/connect"

	"github.com/toffeepay/sdk-go/internal/gen/wallet/v1/v1connect"
)

// AccountService provides operations on wallet accounts.
type AccountService struct {
	client v1connect.AccountServiceClient
}

// Get retrieves an account by ID.
func (a *AccountService) Get(ctx context.Context, req *GetAccountRequest) (*GetAccountResponse, error) {
	resp, err := a.client.GetAccount(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

// List returns a paginated list of accounts.
func (a *AccountService) List(ctx context.Context, req *ListAccountsRequest) (*ListAccountsResponse, error) {
	resp, err := a.client.ListAccounts(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}
