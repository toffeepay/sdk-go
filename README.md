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

// Checkout
session, err := client.Checkout(ctx, &toffee.CreateSessionRequest{
    Item:      &toffee.Item{Title: "Gem Pack 50", Price: 100, Currency: "USD"},
    GameId:    "your-game-id",
    UserId:    "user-id",
    ReturnUrl: "https://example.com/return",
})

fmt.Println(session.GetUrl()) // redirect the user here

// Refund
refund, err := client.Refund(ctx, &toffee.CreateRefundRequest{
    PaymentId: "pay_123",
})
```

### Idempotency

Pass an idempotency key to protect against duplicate operations:

```go
session, err := client.Checkout(ctx, &toffee.CreateSessionRequest{
    Item: &toffee.Item{Title: "Gem Pack 50", Price: 100, Currency: "USD"}, GameId: "game_1", UserId: "user_1", ReturnUrl: "https://example.com/return",
}, toffee.WithIdempotencyKey("unique-key-123"))

refund, err := client.Refund(ctx, &toffee.CreateRefundRequest{
    PaymentId: "pay_123",
}, toffee.WithIdempotencyKey("refund-key-456"))
```

### Sandbox

```go
client := toffee.New(toffee.Config{
    AccessToken: "your-sandbox-token",
    Environment: toffee.Sandbox,
})
```

## Resources

### Sessions

```go
session, err := client.Sessions.Create(ctx, &toffee.CreateSessionRequest{...})
session, err := client.Sessions.Get(ctx, "sess_123")
status, err  := client.Sessions.Status(ctx, "sess_123")
list, err    := client.Sessions.List(ctx, &toffee.ListSessionsRequest{GameId: "game_1"})
err          := client.Sessions.Cancel(ctx, "sess_123")
```

### Payments

```go
payment, err := client.Payments.Get(ctx, "pay_123")
list, err    := client.Payments.List(ctx, &toffee.ListPaymentsRequest{GameId: "game_1"})
err          := client.Payments.Complete(ctx, "pay_123")
err          := client.Payments.Cancel(ctx, "pay_123")
```

### Refunds

```go
refund, err := client.Refunds.Create(ctx, &toffee.CreateRefundRequest{PaymentId: "pay_123"})
refund, err := client.Refunds.Get(ctx, "ref_123")
list, err   := client.Refunds.List(ctx, &toffee.ListRefundsRequest{PaymentId: &paymentId})
```

## Code Generation

Proto stubs are vendored in `internal/gen/`. To regenerate:

```bash
buf generate
```

Requires [buf CLI](https://buf.build/docs/installation). Source protos: [buf.build/toffeepay/toffee](https://buf.build/toffeepay/toffee).

## License

[MIT](LICENSE)
