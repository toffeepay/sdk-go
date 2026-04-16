package toffee

// Re-export proto types so consumers only import the toffee package.

import (
	wpb "github.com/toffeepay/sdk-go/internal/gen/wallet/v1"
	pb "github.com/toffeepay/sdk-go/internal/gen/pay/v1"
)

type (
	// pay/v1 common
	Error = pb.Error

	// pay/v1 shared types
	Item            = pb.Item
	Tax             = pb.Tax
	SessionMetadata = pb.SessionMetadata
	PaymentExtraData = pb.PaymentExtraData

	// pay/v1 domain objects
	Session = pb.Session
	Payment = pb.Payment
	Refund  = pb.Refund

	// pay/v1 session requests / responses
	CreateSessionRequest     = pb.CreateSessionRequest
	CreateSessionResponse    = pb.CreateSessionResponse
	GetSessionRequest        = pb.GetSessionRequest
	GetSessionResponse       = pb.GetSessionResponse
	GetSessionStatusRequest  = pb.GetSessionStatusRequest
	GetSessionStatusResponse = pb.GetSessionStatusResponse
	CancelSessionRequest     = pb.CancelSessionRequest
	CancelSessionResponse    = pb.CancelSessionResponse
	ListSessionsRequest      = pb.ListSessionsRequest
	ListSessionsResponse     = pb.ListSessionsResponse

	// pay/v1 payment requests / responses
	GetPaymentRequest        = pb.GetPaymentRequest
	GetPaymentResponse       = pb.GetPaymentResponse
	CompletePaymentRequest   = pb.CompletePaymentRequest
	CompletePaymentResponse  = pb.CompletePaymentResponse
	CancelPaymentRequest     = pb.CancelPaymentRequest
	CancelPaymentResponse    = pb.CancelPaymentResponse
	ListPaymentsRequest      = pb.ListPaymentsRequest
	ListPaymentsResponse     = pb.ListPaymentsResponse

	// pay/v1 refund requests / responses
	CreateRefundRequest  = pb.CreateRefundRequest
	CreateRefundResponse = pb.CreateRefundResponse
	GetRefundRequest     = pb.GetRefundRequest
	GetRefundResponse    = pb.GetRefundResponse
	ListRefundsRequest   = pb.ListRefundsRequest
	ListRefundsResponse  = pb.ListRefundsResponse

	// wallet/v1 domain objects
	Account  = wpb.Account
	Deposit  = wpb.Deposit
	Balance  = wpb.Balance
	GameUser = wpb.GameUser

	// wallet/v1 account requests / responses
	GetAccountRequest     = wpb.GetAccountRequest
	GetAccountResponse    = wpb.GetAccountResponse
	ListAccountsRequest   = wpb.ListAccountsRequest
	ListAccountsResponse  = wpb.ListAccountsResponse

	// wallet/v1 deposit requests / responses
	GetDepositRequest    = wpb.GetDepositRequest
	GetDepositResponse   = wpb.GetDepositResponse
	ListDepositsRequest  = wpb.ListDepositsRequest
	ListDepositsResponse = wpb.ListDepositsResponse
)
