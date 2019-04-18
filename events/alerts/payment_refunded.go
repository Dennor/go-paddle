package alerts

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const PaymentRefundedAlertName = "payment_refunded"

// PaymentRefunded refer to https://paddle.com/docs/reference-using-webhooks/#payment_refunded
type PaymentRefunded struct {
	AlertName               string                    `json:"alert_name"`
	Amount                  *decimal.Decimal          `json:"amount,string"`
	BalanceCurrency         string                    `json:"balance_currency"`
	BalanceEarningsDecrease *decimal.Decimal          `json:"balance_earnings_decrease,string"`
	BalanceFeeRefund        *decimal.Decimal          `json:"balance_fee_refund,string"`
	BalanceGrossRefund      *decimal.Decimal          `json:"balance_gross_refund,string"`
	BalanceTaxRefund        *decimal.Decimal          `json:"balance_tax_refund,string"`
	CheckoutID              string                    `json:"checkout_id"`
	Currency                string                    `json:"currency"`
	EarningsDecrease        *decimal.Decimal          `json:"earnings_decrease,string"`
	Email                   string                    `json:"email"`
	EventTime               *types.TimeYYYYMMDDHHmmSS `json:"event_time,string"`
	FeeRefund               *decimal.Decimal          `json:"fee_refund,string"`
	GrossRefund             *decimal.Decimal          `json:"gross_refund,string"`
	MarketingConsent        types.MarketingConsent    `json:"marketing_consent,string"`
	OrderID                 string                    `json:"order_id"`
	Passthrough             string                    `json:"passthrough"`
	Quantity                int                       `json:"quantity,string"`
	RefundType              string                    `json:"refund_type"`
	TaxRefund               *decimal.Decimal          `json:"tax_refund,string"`
	PSignature              string                    `json:"p_signature" php:"-"`
}

func (s *PaymentRefunded) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *PaymentRefunded) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
