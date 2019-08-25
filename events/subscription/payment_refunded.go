package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const PaymentRefundedAlertName = "subscription_payment_refunded"

// PaymentRefunded refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_payment_refunded
type PaymentRefunded struct {
	AlertID                 int                     `json:"alert_id,string"`
	AlertName               string                  `json:"alert_name"`
	Amount                  *types.CurrencyValue    `json:"amount,string"`
	BalanceCurrency         string                  `json:"balance_currency"`
	BalanceEarningsDecrease *types.CurrencyValue    `json:"balance_earnings_decrease,string"`
	BalanceFeeRefund        *types.CurrencyValue    `json:"balance_fee_refund,string"`
	BalanceGrossRefund      *types.CurrencyValue    `json:"balance_gross_refund,string"`
	BalanceTaxRefund        *types.CurrencyValue    `json:"balance_tax_refund,string"`
	CheckoutID              string                  `json:"checkout_id"`
	Currency                string                  `json:"currency"`
	EarningsDecrease        *types.CurrencyValue    `json:"earnings_decrease,string"`
	Email                   string                  `json:"email"`
	EventTime               *types.Datetime         `json:"event_time,string"`
	FeeRefund               *types.CurrencyValue    `json:"fee_refund,string"`
	GrossRefund             *types.CurrencyValue    `json:"gross_refund,string"`
	InitialPayment          int                     `json:"initial_payment"`
	Instalments             int                     `json:"instalments,string"`
	MarketingConsent        *types.MarketingConsent `json:"marketing_consent"`
	OrderID                 string                  `json:"order_id"`
	Passthrough             string                  `json:"passthrough"`
	Quantity                int                     `json:"quantity,string"`
	RefundType              string                  `json:"refund_type"`
	SubscriptionID          int                     `json:"subscription_id,string"`
	SubscriptionPaymentID   int                     `json:"subscription_payment_id,string"`
	TaxRefund               *types.CurrencyValue    `json:"tax_refund,string"`
	UnitPrice               *types.CurrencyValue    `json:"unit_price,string"`
	UserID                  int                     `json:"user_id,string"`
	PSignature              string                  `json:"p_signature" php:"-"`
}

func (s *PaymentRefunded) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *PaymentRefunded) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
