package subscription

import (
	"encoding/json"
	"testing"

	"github.com/dennor/go-paddle/events"
	"github.com/dennor/go-paddle/events/test"
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/go-paddle/signature"
	"github.com/dennor/urldecode"
	"github.com/stretchr/testify/assert"
)

func subscriptionPaymentRefundedData() struct {
	d   test.Data
	spr PaymentRefunded
} {
	d := test.Sign(map[string]string{
		"alert_id":                  "1024",
		"alert_name":                "subscription_payment_succeeded",
		"amount":                    "1.23",
		"balance_currency":          "PLN",
		"balance_earnings_decrease": "4.56",
		"balance_fee_refund":        "7.89",
		"balance_gross_refund":      "10.1112",
		"balance_tax_refund":        "13.1415",
		"checkout_id":               "1-c8a82616c183ad6-377f00add1",
		"currency":                  "PLN",
		"earnings_decrease":         "16.1718",
		"email":                     "jan@kowalski.net",
		"event_time":                "2019-04-15 07:37:53",
		"fee_refund":                "19.2021",
		"gross_refund":              "22.2324",
		"initial_payment":           "1",
		"instalments":               "1",
		"marketing_consent":         "1",
		"order_id":                  "unique-id-of-order",
		"status":                    "status",
		"passthrough":               "Example String",
		"quantity":                  "1",
		"refund_reason":             "",
		"refund_type":               "full",
		"subscription_id":           "1",
		"subscription_payment_id":   "5",
		"subscription_plan_id":      "10",
		"tax_refund":                "25.2627",
		"unit_price":                "49.99",
		"user_id":                   "10",
	})
	var mc types.MarketingConsent
	mc.UnmarshalText([]byte(d.M["marketing_consent"]))
	spr := PaymentRefunded{
		AlertID:                 int(test.IntFromString(d.M["alert_id"])),
		AlertName:               d.M["alert_name"],
		Amount:                  test.CurrencyValueFromString(d.M["amount"]),
		BalanceCurrency:         d.M["balance_currency"],
		BalanceEarningsDecrease: test.CurrencyValueFromString(d.M["balance_earnings_decrease"]),
		BalanceFeeRefund:        test.CurrencyValueFromString(d.M["balance_fee_refund"]),
		BalanceGrossRefund:      test.CurrencyValueFromString(d.M["balance_gross_refund"]),
		BalanceTaxRefund:        test.CurrencyValueFromString(d.M["balance_tax_refund"]),
		CheckoutID:              d.M["checkout_id"],
		Currency:                d.M["currency"],
		EarningsDecrease:        test.CurrencyValueFromString(d.M["earnings_decrease"]),
		Email:                   d.M["email"],
		EventTime:               &types.Datetime{test.ParseTime(types.DatetimeFormat, d.M["event_time"])},
		FeeRefund:               test.CurrencyValueFromString(d.M["fee_refund"]),
		GrossRefund:             test.CurrencyValueFromString(d.M["gross_refund"]),
		InitialPayment:          int(test.IntFromString(d.M["initial_payment"])),
		Instalments:             int(test.IntFromString(d.M["instalments"])),
		MarketingConsent:        &mc,
		OrderID:                 d.M["order_id"],
		Status:                  d.M["status"],
		Passthrough:             d.M["passthrough"],
		Quantity:                int(test.IntFromString(d.M["quantity"])),
		RefundReason:            d.M["refund_reason"],
		RefundType:              d.M["refund_type"],
		SubscriptionID:          int(test.IntFromString(d.M["subscription_id"])),
		SubscriptionPaymentID:   int(test.IntFromString(d.M["subscription_payment_id"])),
		SubscriptionPlanID:      int(test.IntFromString(d.M["subscription_plan_id"])),
		TaxRefund:               test.CurrencyValueFromString(d.M["tax_refund"]),
		UnitPrice:               test.CurrencyValueFromString(d.M["unit_price"]),
		UserID:                  int(test.IntFromString(d.M["user_id"])),
		PSignature:              d.M["p_signature"],
	}
	return struct {
		d   test.Data
		spr PaymentRefunded
	}{d, spr}
}

func TestPaymentRefunded(t *testing.T) {
	data := subscriptionPaymentRefundedData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual PaymentRefunded
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.spr, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual PaymentRefunded
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.spr, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.spr.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.spr))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &PaymentRefunded{})
	})
}

func BenchmarkPaymentRefunded(b *testing.B) {
	data := subscriptionPaymentRefundedData()
	var spr PaymentRefunded
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &spr)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &spr)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.spr.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.spr)
		}
	})
}
