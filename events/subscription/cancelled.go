package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const CancelledAlertName = "subscription_cancelled"

// Cancelled refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_cancelled
type Cancelled struct {
	AlertName                 string                  `json:"alert_name"`
	CancellationEffectiveDate *types.Date             `json:"cancellation_effective_date,string"`
	CheckoutID                string                  `json:"checkout_id"`
	Currency                  string                  `json:"currency"`
	Email                     string                  `json:"email"`
	EventTime                 *types.Datetime         `json:"event_time,string"`
	MarketingConsent          *types.MarketingConsent `json:"marketing_consent,string"`
	Passthrough               string                  `json:"passthrough"`
	Quantity                  int                     `json:"quantity,string"`
	Status                    string                  `json:"status"`
	SubscriptionID            int                     `json:"subscription_id,string"`
	SubscriptionPlanID        int                     `json:"subscription_plan_id,string"`
	UnitPrice                 *decimal.Decimal        `json:"unit_price,string"`
	UserID                    int                     `json:"user_id,string"`
	PSignature                string                  `json:"p_signature" php:"-"`
}

func (s *Cancelled) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *Cancelled) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
