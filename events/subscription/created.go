package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const CreatedAlertName = "subscription_created"

// Created refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_created
type Created struct {
	AlertID            int                     `json:"alert_id,string"`
	AlertName          string                  `json:"alert_name"`
	CancelURL          string                  `json:"cancel_url"`
	CheckoutID         string                  `json:"checkout_id"`
	Currency           string                  `json:"currency"`
	Email              string                  `json:"email"`
	EventTime          *types.Datetime         `json:"event_time,string"`
	MarketingConsent   *types.MarketingConsent `json:"marketing_consent,string"`
	NextBillDate       *types.Date             `json:"next_bill_date,string"`
	Passthrough        string                  `json:"passthrough"`
	Quantity           int                     `json:"quantity,string"`
	SourcePage         string                  `json:"source_page,omitempty"`
	Status             string                  `json:"status"`
	SubscriptionID     int                     `json:"subscription_id,string"`
	SubscriptionPlanID int                     `json:"subscription_plan_id,string"`
	UnitPrice          *types.CurrencyValue    `json:"unit_price,string"`
	UpdateURL          string                  `json:"update_url"`
	UserID             int                     `json:"user_id,string,omitempty"`
	PSignature         string                  `json:"p_signature" php:"-"`
}

func (s *Created) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *Created) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
