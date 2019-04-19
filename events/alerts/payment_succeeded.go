package alerts

import (
	"net"

	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const PaymentSucceededAlertName = "payment_succeeded"

// PaymentSucceeded refer to https://paddle.com/docs/reference-using-webhooks/#payment_succeeded
type PaymentSucceeded struct {
	AlertName         string                    `json:"alert_name"`
	BalanceCurrency   string                    `json:"balance_currency"`
	BalanceEarnings   *decimal.Decimal          `json:"balance_earnings,string"`
	BalanceFee        *decimal.Decimal          `json:"balance_fee,string"`
	BalanceGross      *decimal.Decimal          `json:"balance_gross,string"`
	BalanceTax        *decimal.Decimal          `json:"balance_tax,string"`
	CheckoutID        string                    `json:"checkout_id"`
	Country           string                    `json:"country"`
	Coupon            string                    `json:"coupon"`
	Currency          string                    `json:"currency"`
	CustomerName      string                    `json:"customer_name"`
	Earnings          *decimal.Decimal          `json:"earnings,string"`
	Email             string                    `json:"email"`
	EventTime         *types.TimeYYYYMMDDHHmmSS `json:"event_time,string"`
	Fee               *decimal.Decimal          `json:"fee,string"`
	IP                *net.IP                   `json:"ip,string"`
	MarketingConsent  types.MarketingConsent    `json:"marketing_consent,string"`
	OrderID           string                    `json:"order_id"`
	Passthrough       string                    `json:"passthrough"`
	PaymentMethod     string                    `json:"payment_method"`
	PaymentTax        *decimal.Decimal          `json:"payment_tax,string"`
	ProductID         int                       `json:"product_id,string"`
	ProductName       string                    `json:"product_name"`
	Quantity          int                       `json:"quantity,string"`
	ReceiptUrl        string                    `json:"receipt_url"`
	SaleGross         *decimal.Decimal          `json:"sale_gross,string"`
	UsedPriceOverride bool                      `json:"used_price_override,string"`
	PSignature        string                    `json:"p_signature" php:"-"`
}

func (s *PaymentSucceeded) Serialize() ([]byte, error) {
	return phpserialize.Marshal(s)
}

func (s *PaymentSucceeded) Signature() ([]byte, error) {
	return []byte(s.PSignature), nil
}
