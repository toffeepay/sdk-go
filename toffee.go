// Package toffee provides a Go SDK for the ToffeePay payment API.
//
// Usage:
//
//	client := toffee.New(toffee.Config{
//	    AccessToken: "your-access-token",
//	    Environment: toffee.Sandbox,
//	})
//
//	session, err := client.Checkout(ctx, &toffee.CreateSessionRequest{
//	    GameId: "game-1",
//	    UserId: "user-1",
//	    Item:   &toffee.Item{Title: "100 Coins", Price: 999, Currency: "USD"},
//	})
package toffee

import (
	"context"
	"net/http"

	"connectrpc.com/connect"

	payconnect "github.com/toffeepay/sdk-go/internal/gen/pay/v1/v1connect"
	walletconnect "github.com/toffeepay/sdk-go/internal/gen/wallet/v1/v1connect"
)

const (
	ProductionURL = "https://api.toffeepay.com"
	SandboxURL    = "https://api.sandbox.toffeepay.com"
)

// Environment selects between production and sandbox APIs.
type Environment string

const (
	Production Environment = "production"
	Sandbox    Environment = "sandbox"
)

// Config holds the configuration for creating a [Client].
type Config struct {
	// AccessToken is the bearer token for authenticating with the ToffeePay API.
	AccessToken string

	// Environment selects between production and sandbox. Defaults to Production.
	Environment Environment

	// HTTPClient overrides the default HTTP client.
	HTTPClient *http.Client

	// BaseURL overrides the environment-derived base URL.
	BaseURL string
}

// Client provides access to the ToffeePay API.
type Client struct {
	Sessions *SessionService
	Payments *PaymentService
	Refunds  *RefundService
	Accounts *AccountService
	Deposits *DepositService
}

// New creates a new ToffeePay [Client].
func New(cfg Config) *Client {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		switch cfg.Environment {
		case Sandbox:
			baseURL = SandboxURL
		default:
			baseURL = ProductionURL
		}
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	opts := []connect.ClientOption{
		connect.WithInterceptors(authInterceptor(cfg.AccessToken)),
	}

	paymentClient := payconnect.NewPaymentServiceClient(httpClient, baseURL, opts...)
	refundClient := payconnect.NewRefundServiceClient(httpClient, baseURL, opts...)
	accountClient := walletconnect.NewAccountServiceClient(httpClient, baseURL, opts...)

	return &Client{
		Sessions: &SessionService{client: paymentClient},
		Payments: &PaymentService{client: paymentClient},
		Refunds:  &RefundService{client: refundClient},
		Accounts: &AccountService{client: accountClient},
		Deposits: &DepositService{client: accountClient},
	}
}

// Checkout is a convenience method that creates a payment session and returns the [Session] directly.
func (c *Client) Checkout(ctx context.Context, req *CreateSessionRequest, opts ...RequestOption) (*Session, error) {
	resp, err := c.Sessions.Create(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return resp.GetSession(), nil
}

// Refund is a convenience method that creates a refund and returns the [Refund] directly.
func (c *Client) Refund(ctx context.Context, req *CreateRefundRequest, opts ...RequestOption) (*Refund, error) {
	resp, err := c.Refunds.Create(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return resp.GetRefund(), nil
}

func authInterceptor(token string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req.Header().Set("Authorization", "Bearer "+token)
			return next(ctx, req)
		}
	}
}
