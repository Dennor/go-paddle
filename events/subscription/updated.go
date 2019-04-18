package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const UpdatedAlertName = "subscription_updated"

// Updated refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_updated
type Updated struct {
	AlertName          string                    `json:"alert_name"`
	CancelUrl          string                    `json:"cancel_url"`
	CheckoutID         string                    `json:"checkout_id"`
	Email              string                    `json:"email"`
	EventTime          *types.TimeYYYYMMDDHHmmSS `json:"event_time,string"`
	MarketingConsent   types.MarketingConsent    `json:"marketing_consent,string"`
	NewPrice           *decimal.Decimal          `json:"new_price,string"`
	NewQuantity        int                       `json:"new_quantity,string"`
	NewUnitPrice       *decimal.Decimal          `json:"new_unit_price,string"`
	NextBillDate       *types.TimeYYYYMMDD       `json:"next_bill_date,string"`
	OldPrice           *decimal.Decimal          `json:"old_price,string"`
	OldQuantity        int                       `json:"old_quantity,string"`
	OldUnitPrice       *decimal.Decimal          `json:"old_unit_price,string"`
	Passthrough        string                    `json:"passthrough"`
	Status             string                    `json:"status"`
	SubscriptionID     int                       `json:"subscription_id,string"`
	SubscriptionPlanID int                       `json:"subscription_plan_id,string"`
	UpdateURL          string                    `json:"update_url"`
	PSignature         string                    `json:"p_signature" php:"-"`
}

func (s *Updated) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *Updated) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
