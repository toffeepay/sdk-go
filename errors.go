package toffee

import (
	"errors"
	"fmt"

	"connectrpc.com/connect"
)

// ToffeeError is returned when the API responds with an application-level error.
type ToffeeError struct {
	Code       string
	Message    string
	StatusCode int
}

func (e *ToffeeError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// AsToffeeError unwraps a ConnectRPC error into a [ToffeeError], if possible.
// Returns nil if err is not a connect error.
func AsToffeeError(err error) *ToffeeError {
	var connectErr *connect.Error
	if !errors.As(err, &connectErr) {
		return nil
	}
	for _, detail := range connectErr.Details() {
		msg, err := detail.Value()
		if err != nil {
			continue
		}
		if e, ok := msg.(*Error); ok {
			return &ToffeeError{
				Code:       e.GetCode(),
				Message:    e.GetMessage(),
				StatusCode: int(connectErr.Code()),
			}
		}
	}
	return &ToffeeError{
		Code:       connectErr.Code().String(),
		Message:    connectErr.Message(),
		StatusCode: int(connectErr.Code()),
	}
}
