package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const PaymentSucceededAlertName = "subscription_payment_succeeded"

// PaymentSucceeded refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_payment_succeeded
type PaymentSucceeded struct {
	AlertID            int                     `json:"alert_id,string"`
	AlertName          string                  `json:"alert_name"`
	BalanceCurrency    string                  `json:"balance_currency"`
	BalanceEarnings    *types.CurrencyValue    `json:"balance_earnings,string"`
	BalanceFee         *types.CurrencyValue    `json:"balance_fee,string"`
	BalanceGross       *types.CurrencyValue    `json:"balance_gross,string"`
	BalanceTax         *types.CurrencyValue    `json:"balance_tax,string"`
	CheckoutID         string                  `json:"checkout_id"`
	Country            string                  `json:"country"`
	Coupon             string                  `json:"coupon"`
	Currency           string                  `json:"currency"`
	CustomerName       string                  `json:"customer_name"`
	Earnings           *types.CurrencyValue    `json:"earnings,string"`
	Email              string                  `json:"email"`
	EventTime          *types.Datetime         `json:"event_time,string"`
	Fee                *types.CurrencyValue    `json:"fee,string"`
	InitialPayment     int                     `json:"initial_payment,string"`
	Instalments        int                     `json:"instalments,string"`
	MarketingConsent   *types.MarketingConsent `json:"marketing_consent,string"`
	NextBillDate       *types.Date             `json:"next_bill_date,string"`
	NextPaymentAmount  *types.CurrencyValue    `json:"next_payment_amount,string,omitempty"`
	OrderID            string                  `json:"order_id"`
	Passthrough        string                  `json:"passthrough"`
	PaymentMethod      string                  `json:"payment_method"`
	PaymentTax         *types.CurrencyValue    `json:"payment_tax,string"`
	PlanName           string                  `json:"plan_name"`
	Quantity           int                     `json:"quantity,string"`
	ReceiptURL         string                  `json:"receipt_url"`
	SaleGross          *types.CurrencyValue    `json:"sale_gross,string"`
	Status             string                  `json:"status"`
	SubscriptionID     int                     `json:"subscription_id,string"`
	SubscriptionPlanID int                     `json:"subscription_plan_id,string"`
	UnitPrice          *types.CurrencyValue    `json:"unit_price,string"`
	UserID             int                     `json:"user_id,string"`
	PSignature         string                  `json:"p_signature" php:"-"`
}

func (s *PaymentSucceeded) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *PaymentSucceeded) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
