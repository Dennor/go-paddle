package subscription

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const PaymentSucceededAlertName = "subscription_payment_succeeded"

// PaymentSucceeded refer to https://paddle.com/docs/subscriptions-event-reference/#subscription_payment_succeeded
type PaymentSucceeded struct {
	AlertName          string                  `json:"alert_name"`
	BalanceCurrency    string                  `json:"balance_currency"`
	BalanceEarnings    *decimal.Decimal        `json:"balance_earnings,string"`
	BalanceFee         *decimal.Decimal        `json:"balance_fee,string"`
	BalanceGross       *decimal.Decimal        `json:"balance_gross,string"`
	BalanceTax         *decimal.Decimal        `json:"balance_tax,string"`
	CheckoutID         string                  `json:"checkout_id"`
	Country            string                  `json:"country"`
	Coupon             string                  `json:"coupon"`
	Currency           string                  `json:"currency"`
	CustomerName       string                  `json:"customer_name"`
	Earnings           *decimal.Decimal        `json:"earnings,string"`
	Email              string                  `json:"email"`
	EventTime          *types.Datetime         `json:"event_time,string"`
	Fee                *decimal.Decimal        `json:"fee,string"`
	InitialPayment     int                     `json:"initial_payment,string"`
	Instalments        int                     `json:"instalments,string"`
	MarketingConsent   *types.MarketingConsent `json:"marketing_consent,string"`
	NextBillDate       *types.Date             `json:"next_bill_date,string"`
	OrderID            string                  `json:"order_id"`
	Passthrough        string                  `json:"passthrough"`
	PaymentMethod      string                  `json:"payment_method"`
	PaymentTax         *decimal.Decimal        `json:"payment_tax,string"`
	PlanName           string                  `json:"plan_name"`
	Quantity           int                     `json:"quantity,string"`
	ReceiptUrl         string                  `json:"receipt_url"`
	SaleGross          *decimal.Decimal        `json:"sale_gross,string"`
	Status             string                  `json:"status"`
	SubscriptionID     int                     `json:"subscription_id,string"`
	SubscriptionPlanID int                     `json:"subscription_plan_id,string"`
	UnitPrice          *decimal.Decimal        `json:"unit_price,string"`
	UserID             int                     `json:"user_id,string"`
	PSignature         string                  `json:"p_signature" php:"-"`
}

func (s *PaymentSucceeded) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *PaymentSucceeded) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
