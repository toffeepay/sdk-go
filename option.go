package toffee

import "connectrpc.com/connect"

// RequestOption configures individual RPC calls.
type RequestOption func(*requestOptions)

type requestOptions struct {
	idempotencyKey string
}

// WithIdempotencyKey sets the Idempotency-Key header on the request.
// Supported on session creation and refund creation.
func WithIdempotencyKey(key string) RequestOption {
	return func(o *requestOptions) {
		o.idempotencyKey = key
	}
}

func applyOptions(opts []RequestOption) requestOptions {
	var o requestOptions
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func applyHeaders[T any](req *connect.Request[T], o requestOptions) {
	if o.idempotencyKey != "" {
		req.Header().Set("Idempotency-Key", o.idempotencyKey)
	}
}
