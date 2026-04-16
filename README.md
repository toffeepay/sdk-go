# sdk-go

ToffeePay Go SDK.

## Install

```bash
go get github.com/toffeepay/sdk-go
```

## Usage

```go
import toffee "github.com/toffeepay/sdk-go"

client := toffee.New(toffee.Config{
    AccessToken: "your-access-token",
})

// Checkout — returns Session directly
session, err := client.Checkout(ctx, &toffee.CreateSessionRequest{
    GameId:    "your-game-id",
    UserId:    "user-id",
    Item:      &toffee.Item{Title: "Gem Pack 50", Price: 100, Currency: "USD"},
    ReturnUrl: "https://example.com/return",
})
fmt.Println(session.GetUrl()) // redirect the user here

// Refund — returns Refund directly
refund, err := client.Refund(ctx, &toffee.CreateRefundRequest{PaymentId: "pay_123"})
```

### Idempotency

```go
session, err := client.Checkout(ctx, &toffee.CreateSessionRequest{...},
    toffee.WithIdempotencyKey("unique-key-123"),
)

refund, err := client.Refund(ctx, &toffee.CreateRefundRequest{PaymentId: "pay_123"},
    toffee.WithIdempotencyKey("refund-key-456"),
)
```

### Sandbox

```go
client := toffee.New(toffee.Config{
    AccessToken: "your-sandbox-token",
    Environment: toffee.Sandbox,
})
```

## Resources

All resource methods accept request objects and return response objects.

### Sessions

```go
resp, err := client.Sessions.Create(ctx, &toffee.CreateSessionRequest{...})  // CreateSessionResponse
resp, err := client.Sessions.Get(ctx, &toffee.GetSessionRequest{Id: "sess_123"})          // GetSessionResponse
resp, err := client.Sessions.Status(ctx, &toffee.GetSessionStatusRequest{Id: "sess_123"}) // GetSessionStatusResponse
resp, err := client.Sessions.List(ctx, &toffee.ListSessionsRequest{GameId: "game_1"})     // ListSessionsResponse
resp, err := client.Sessions.Cancel(ctx, &toffee.CancelSessionRequest{Id: "sess_123"})   // CancelSessionResponse
```

### Payments

```go
resp, err := client.Payments.Get(ctx, &toffee.GetPaymentRequest{Id: "pay_123"})              // GetPaymentResponse
resp, err := client.Payments.List(ctx, &toffee.ListPaymentsRequest{GameId: "game_1"})         // ListPaymentsResponse
resp, err := client.Payments.Complete(ctx, &toffee.CompletePaymentRequest{Id: "pay_123"})     // CompletePaymentResponse
resp, err := client.Payments.Cancel(ctx, &toffee.CancelPaymentRequest{Id: "pay_123"})         // CancelPaymentResponse
```

### Refunds

```go
resp, err := client.Refunds.Create(ctx, &toffee.CreateRefundRequest{PaymentId: "pay_123"}) // CreateRefundResponse
resp, err := client.Refunds.Get(ctx, &toffee.GetRefundRequest{Id: "ref_123"})               // GetRefundResponse
resp, err := client.Refunds.List(ctx, &toffee.ListRefundsRequest{})                         // ListRefundsResponse
```

### Accounts

```go
resp, err := client.Accounts.Get(ctx, &toffee.GetAccountRequest{Id: "acc_123"})            // GetAccountResponse
resp, err := client.Accounts.List(ctx, &toffee.ListAccountsRequest{GameId: "game_1"})       // ListAccountsResponse
```

### Deposits

```go
resp, err := client.Deposits.Get(ctx, &toffee.GetDepositRequest{Id: "dep_123"})            // GetDepositResponse
resp, err := client.Deposits.List(ctx, &toffee.ListDepositsRequest{GameId: "game_1"})       // ListDepositsResponse
```

## Code Generation

Proto stubs are vendored in `internal/gen/`. To regenerate:

```bash
buf generate
```

Requires [buf CLI](https://buf.build/docs/installation). Source protos: [buf.build/toffeepay/toffee](https://buf.build/toffeepay/toffee).

## License

[MIT](LICENSE)
