package alerts

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const PaymentDisputeClosedAlertName = "payment_dispute_closed"

// PaymentDisputeClosed refer to https://paddle.com/docs/reference-using-webhooks/#payment_dispute_closed
type PaymentDisputeClosed struct {
	AlertName        string                    `json:"alert_name"`
	Amount           *decimal.Decimal          `json:"amount,string"`
	CheckoutID       string                    `json:"checkout_id"`
	Currency         string                    `json:"currency"`
	Email            string                    `json:"email"`
	EventTime        *types.TimeYYYYMMDDHHmmSS `json:"event_time,string"`
	FeeUsd           *decimal.Decimal          `json:"fee_usd,string"`
	MarketingConsent types.MarketingConsent    `json:"marketing_consent,string"`
	OrderID          int                       `json:"order_id,string"`
	Passthrough      string                    `json:"passthrough"`
	Status           string                    `json:"status"`
	PSignature       string                    `json:"p_signature" php:"-"`
}

func (p *PaymentDisputeClosed) Serialize() ([]byte, error) {
	return phpserialize.Marshal(p)
}

func (p *PaymentDisputeClosed) Signature() ([]byte, error) {
	return []byte(p.PSignature), nil
}
