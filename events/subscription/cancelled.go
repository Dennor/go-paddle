package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const CancelledAlertName = "subscription_cancelled"

// Cancelled refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_cancelled
type Cancelled struct {
	AlertID                   int                     `json:"alert_id,string"`
	AlertName                 string                  `json:"alert_name"`
	CancellationEffectiveDate *types.Date             `json:"cancellation_effective_date,string"`
	CheckoutID                string                  `json:"checkout_id"`
	Currency                  string                  `json:"currency"`
	CustomData                string                  `json:"custom_data"`
	Email                     string                  `json:"email"`
	EventTime                 *types.Datetime         `json:"event_time,string"`
	LinkedSubscriptions       string                  `json:"linked_subscriptions"`
	MarketingConsent          *types.MarketingConsent `json:"marketing_consent,string"`
	Passthrough               string                  `json:"passthrough"`
	Quantity                  int                     `json:"quantity,string"`
	Status                    string                  `json:"status"`
	SubscriptionID            int                     `json:"subscription_id,string"`
	SubscriptionPlanID        int                     `json:"subscription_plan_id,string"`
	UnitPrice                 *types.CurrencyValue    `json:"unit_price,string"`
	UserID                    int                     `json:"user_id,string"`
	PSignature                string                  `json:"p_signature" php:"-"`
}

func (s *Cancelled) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *Cancelled) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
