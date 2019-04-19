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

func subscriptionPaymentSucceededData() struct {
	d   test.Data
	sps PaymentSucceeded
} {
	d := test.Sign(map[string]string{
		"alert_name":           "subscription_payment_succeeded",
		"balance_currency":     "PLN",
		"balance_earnings":     "1.23",
		"balance_fee":          "4.56",
		"balance_gross":        "7.89",
		"balance_tax":          "10.1112",
		"checkout_id":          "1-c8a82616c183ad6-377f00add1",
		"country":              "PL",
		"coupon":               "secret-coupon-code",
		"currency":             "PLN",
		"customer_name":        "Jan Kowalski",
		"earnings":             "13.1415",
		"email":                "jan@kowalski.net",
		"event_time":           "2019-04-15 07:37:53",
		"fee":                  "16.1718",
		"initial_payment":      "1",
		"instalments":          "1",
		"marketing_consent":    "1",
		"next_bill_date":       "2019-04-15",
		"order_id":             "unique-id-of-order",
		"passthrough":          "Example String",
		"payment_method":       "credit card",
		"payment_tax":          "19.2021",
		"plan_name":            "premium",
		"quantity":             "1",
		"receipt_url":          "https://example.org/a=b&c=d",
		"sale_gross":           "22.2324",
		"status":               "active",
		"subscription_id":      "1",
		"subscription_plan_id": "5",
		"unit_price":           "49.99",
		"user_id":              "10",
	})
	var mc types.MarketingConsent
	mc.UnmarshalText([]byte(d.M["marketing_consent"]))
	sps := PaymentSucceeded{
		AlertName:          d.M["alert_name"],
		BalanceCurrency:    d.M["balance_currency"],
		BalanceEarnings:    test.DecimalFromString(d.M["balance_earnings"]),
		BalanceFee:         test.DecimalFromString(d.M["balance_fee"]),
		BalanceGross:       test.DecimalFromString(d.M["balance_gross"]),
		BalanceTax:         test.DecimalFromString(d.M["balance_tax"]),
		CheckoutID:         d.M["checkout_id"],
		Country:            d.M["country"],
		Coupon:             d.M["coupon"],
		Currency:           d.M["currency"],
		CustomerName:       d.M["customer_name"],
		Earnings:           test.DecimalFromString(d.M["earnings"]),
		Email:              d.M["email"],
		EventTime:          &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		Fee:                test.DecimalFromString(d.M["fee"]),
		InitialPayment:     int(test.IntFromString(d.M["initial_payment"])),
		Instalments:        int(test.IntFromString(d.M["instalments"])),
		MarketingConsent:   &mc,
		NextBillDate:       &types.TimeYYYYMMDD{test.ParseTime(types.TimeFormatYYYYMMDD, d.M["next_bill_date"])},
		OrderID:            d.M["order_id"],
		Passthrough:        d.M["passthrough"],
		PaymentMethod:      d.M["payment_method"],
		PaymentTax:         test.DecimalFromString(d.M["payment_tax"]),
		PlanName:           d.M["plan_name"],
		Quantity:           int(test.IntFromString(d.M["quantity"])),
		ReceiptUrl:         d.M["receipt_url"],
		SaleGross:          test.DecimalFromString(d.M["sale_gross"]),
		Status:             d.M["status"],
		SubscriptionID:     int(test.IntFromString(d.M["subscription_id"])),
		SubscriptionPlanID: int(test.IntFromString(d.M["subscription_plan_id"])),
		UnitPrice:          test.DecimalFromString(d.M["unit_price"]),
		UserID:             int(test.IntFromString(d.M["user_id"])),
		PSignature:         d.M["p_signature"],
	}
	return struct {
		d   test.Data
		sps PaymentSucceeded
	}{d, sps}
}

func TestPaymentSucceeded(t *testing.T) {
	data := subscriptionPaymentSucceededData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual PaymentSucceeded
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.sps, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual PaymentSucceeded
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.sps, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.sps.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.sps))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &PaymentSucceeded{})
	})
}

func BenchmarkPaymentSucceeded(b *testing.B) {
	data := subscriptionPaymentSucceededData()
	var sps PaymentSucceeded
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &sps)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &sps)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.sps.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.sps)
		}
	})
}
