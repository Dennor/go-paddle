package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const UpdatedAlertName = "subscription_updated"

// Updated refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_updated
type Updated struct {
	AlertID               int                     `json:"alert_id,string"`
	AlertName             string                  `json:"alert_name"`
	CancelURL             string                  `json:"cancel_url"`
	CheckoutID            string                  `json:"checkout_id"`
	Currency              string                  `json:"currency,omitempty"`
	CustomData            string                  `json:"custom_data"`
	Email                 string                  `json:"email"`
	EventTime             *types.Datetime         `json:"event_time,string"`
	LinkedSubscriptions   string                  `json:"linked_subscriptions"`
	MarketingConsent      *types.MarketingConsent `json:"marketing_consent,string"`
	NewPrice              *types.CurrencyValue    `json:"new_price,string"`
	NewQuantity           int                     `json:"new_quantity,string"`
	NewUnitPrice          *types.CurrencyValue    `json:"new_unit_price,string"`
	NextBillDate          *types.Date             `json:"next_bill_date,string"`
	OldNextBillDate       *types.Date             `json:"old_next_bill_date,string"`
	OldPrice              *types.CurrencyValue    `json:"old_price,string"`
	OldQuantity           int                     `json:"old_quantity,string"`
	OldStatus             string                  `json:"old_status"`
	OldSubscriptionPlanID int                     `json:"old_subscription_plan_id,string"`
	OldUnitPrice          *types.CurrencyValue    `json:"old_unit_price,string"`
	Passthrough           string                  `json:"passthrough"`
	Status                string                  `json:"status"`
	SubscriptionID        int                     `json:"subscription_id,string"`
	SubscriptionPlanID    int                     `json:"subscription_plan_id,string"`
	UpdateURL             string                  `json:"update_url"`
	UserID                int                     `json:"user_id,string,omitempty"`
	PSignature            string                  `json:"p_signature" php:"-"`
}

func (s *Updated) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *Updated) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
