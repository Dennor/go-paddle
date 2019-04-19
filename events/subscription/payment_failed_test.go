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

func subscriptionPaymentFailedData() struct {
	d   test.Data
	spf PaymentFailed
} {
	d := test.Sign(map[string]string{
		"alert_name":           "subscription_payment_failed",
		"amount":               "1.23",
		"cancel_url":           "https://checkout.paddle.com/subscription/cancel?user=5&subscription=4&hash=a4dca832089cc76fc05da732664d971c6a7c8840",
		"checkout_id":          "1-c8a82616c183ad6-377f00add1",
		"currency":             "GBP",
		"email":                "makenzie89@example.net",
		"event_time":           "2019-04-15 07:37:53",
		"marketing_consent":    "1",
		"next_retry_date":      "2019-05-14",
		"passthrough":          "Example String",
		"quantity":             "60",
		"status":               "active",
		"subscription_id":      "1",
		"subscription_plan_id": "5",
		"unit_price":           "49.99",
		"update_url":           "https://checkout.paddle.com/subscription/update?user=4&subscription=2&hash=a0aef1af98b11ef5d220751a77d0eda187f836d4",
	}, map[string]bool{"hard_failure": true})
	pbool := test.BoolFromString(d.M["hard_failure"])
	var mc types.MarketingConsent
	mc.UnmarshalText([]byte(d.M["marketing_consent"]))
	spf := PaymentFailed{
		AlertName:          d.M["alert_name"],
		Amount:             test.DecimalFromString(d.M["amount"]),
		CancelURL:          d.M["cancel_url"],
		CheckoutID:         d.M["checkout_id"],
		Currency:           d.M["currency"],
		Email:              d.M["email"],
		EventTime:          &types.Datetime{test.ParseTime(types.DatetimeFormat, d.M["event_time"])},
		HardFailure:        (*types.PhpBool)(&pbool),
		MarketingConsent:   &mc,
		NextRetryDate:      &types.Date{test.ParseTime(types.DateFormat, d.M["next_retry_date"])},
		Passthrough:        d.M["passthrough"],
		Quantity:           int(test.IntFromString(d.M["quantity"])),
		Status:             d.M["status"],
		SubscriptionID:     int(test.IntFromString(d.M["subscription_id"])),
		SubscriptionPlanID: int(test.IntFromString(d.M["subscription_plan_id"])),
		UnitPrice:          test.DecimalFromString("49.99"),
		UpdateURL:          d.M["update_url"],
		PSignature:         d.M["p_signature"],
	}
	return struct {
		d   test.Data
		spf PaymentFailed
	}{d, spf}
}

func TestPaymentFailed(t *testing.T) {
	data := subscriptionPaymentFailedData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual PaymentFailed
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.spf, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual PaymentFailed
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.spf, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.spf.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.spf))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &PaymentFailed{})
	})
}

func BenchmarkPaymentFailed(b *testing.B) {
	data := subscriptionPaymentFailedData()
	var spf PaymentFailed
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &spf)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &spf)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.spf.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.spf)
		}
	})
}
