package toffee

import (
	pb "github.com/toffeepay/sdk-go/internal/gen/pay/v1"
)

// Re-export proto types so consumers only import the toffee package.

type (
	Item             = pb.Item
	Tax              = pb.Tax
	SessionMetadata  = pb.SessionMetadata
	PaymentExtraData = pb.PaymentExtraData
	Error            = pb.Error

	Session = pb.Session
	Payment = pb.Payment
	Refund  = pb.Refund

	CreateSessionRequest  = pb.CreateSessionRequest
	ListSessionsRequest   = pb.ListSessionsRequest
	ListSessionsResponse  = pb.ListSessionsResponse
	GetSessionStatusResponse = pb.GetSessionStatusResponse

	ListPaymentsRequest  = pb.ListPaymentsRequest
	ListPaymentsResponse = pb.ListPaymentsResponse

	CreateRefundRequest  = pb.CreateRefundRequest
	ListRefundsRequest   = pb.ListRefundsRequest
	ListRefundsResponse  = pb.ListRefundsResponse
)
