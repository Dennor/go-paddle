package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const PaymentFailedAlertName = "subscription_payment_failed"

// PaymentFailed refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_payment_failed
type PaymentFailed struct {
	AlertName          string                  `json:"alert_name"`
	Amount             *decimal.Decimal        `json:"amount,string"`
	CancelUrl          string                  `json:"cancel_url"`
	CheckoutID         string                  `json:"checkout_id"`
	Currency           string                  `json:"currency"`
	Email              string                  `json:"email"`
	EventTime          *types.Datetime         `json:"event_time,string"`
	HardFailure        *types.PhpBool          `json:"hard_failure,string"`
	MarketingConsent   *types.MarketingConsent `json:"marketing_consent,string"`
	NextRetryDate      *types.Date             `json:"next_retry_date,string"`
	Passthrough        string                  `json:"passthrough"`
	Quantity           int                     `json:"quantity,string"`
	Status             string                  `json:"status"`
	SubscriptionID     int                     `json:"subscription_id,string"`
	SubscriptionPlanID int                     `json:"subscription_plan_id,string"`
	UnitPrice          *decimal.Decimal        `json:"unit_price,string"`
	UpdateURL          string                  `json:"update_url"`
	PSignature         string                  `json:"p_signature" php:"-"`
}

func (s *PaymentFailed) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *PaymentFailed) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
