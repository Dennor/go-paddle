package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const PaymentFailedAlertName = "subscription_payment_failed"

// PaymentFailed refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_payment_failed
type PaymentFailed struct {
	AlertID               int                     `json:"alert_id,string"`
	AlertName             string                  `json:"alert_name"`
	Amount                *types.CurrencyValue    `json:"amount,string"`
	AttemptNumber         int                     `json:"attempt_number,string,omitempty"`
	CancelURL             string                  `json:"cancel_url"`
	CheckoutID            string                  `json:"checkout_id"`
	Currency              string                  `json:"currency"`
	Email                 string                  `json:"email"`
	EventTime             *types.Datetime         `json:"event_time,string"`
	HardFailure           *types.PhpBool          `json:"hard_failure,string,omitempty"`
	Instalments           int                     `json:"instalments,string,omitempty"`
	MarketingConsent      *types.MarketingConsent `json:"marketing_consent,string"`
	NextRetryDate         *types.Date             `json:"next_retry_date,string"`
	OrderID               string                  `json:"order_id,omitempty"`
	Passthrough           string                  `json:"passthrough"`
	Quantity              int                     `json:"quantity,string"`
	Status                string                  `json:"status"`
	SubscriptionID        int                     `json:"subscription_id,string"`
	SubscriptionPaymentID int                     `json:"subscription_payment_id,string,omitempty"`
	SubscriptionPlanID    int                     `json:"subscription_plan_id,string"`
	UnitPrice             *types.CurrencyValue    `json:"unit_price,string"`
	UpdateURL             string                  `json:"update_url"`
	UserID                int                     `json:"user_id,string"`
	PSignature            string                  `json:"p_signature" php:"-"`
}

func (s *PaymentFailed) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *PaymentFailed) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
