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

func subscriptionCreatedData() (subscriptionCreatedData []struct {
	d  test.Data
	sc Created
}) {
	for _, data := range []map[string]string{
		map[string]string{
			"alert_id":             "1024",
			"alert_name":           "subscription_created",
			"cancel_url":           "https://checkout.paddle.com/subscription/cancel?user=5&subscription=4&hash=a4dca832089cc76fc05da732664d971c6a7c8840",
			"checkout_id":          "1-c8a82616c183ad6-377f00add1",
			"currency":             "GBP",
			"email":                "makenzie89@example.net",
			"event_time":           "2019-04-15 07:37:53",
			"linked_subscriptions": "",
			"marketing_consent":    "1",
			"next_bill_date":       "2019-05-14",
			"passthrough":          "Example String",
			"quantity":             "60",
			"status":               "active",
			"subscription_id":      "1",
			"subscription_plan_id": "5",
			"unit_price":           "49.99",
			"update_url":           "https://checkout.paddle.com/subscription/update?user=4&subscription=2&hash=a0aef1af98b11ef5d220751a77d0eda187f836d4",
		},
		map[string]string{
			"alert_id":             "1024",
			"alert_name":           "subscription_created",
			"cancel_url":           "https://checkout.paddle.com/subscription/cancel?user=5&subscription=4&hash=a4dca832089cc76fc05da732664d971c6a7c8840",
			"checkout_id":          "1-c8a82616c183ad6-377f00add1",
			"currency":             "GBP",
			"email":                "makenzie89@example.net",
			"event_time":           "2019-04-15 07:37:53",
			"linked_subscriptions": "",
			"marketing_consent":    "0",
			"next_bill_date":       "2019-05-14",
			"passthrough":          "Example String",
			"quantity":             "60",
			"status":               "active",
			"subscription_id":      "1",
			"subscription_plan_id": "5",
			"unit_price":           "49.99",
			"update_url":           "https://checkout.paddle.com/subscription/update?user=4&subscription=2&hash=a0aef1af98b11ef5d220751a77d0eda187f836d4",
		},
		map[string]string{
			"alert_id":             "1024",
			"alert_name":           "subscription_created",
			"cancel_url":           "https://checkout.paddle.com/subscription/cancel?user=5&subscription=4&hash=a4dca832089cc76fc05da732664d971c6a7c8840",
			"checkout_id":          "1-c8a82616c183ad6-377f00add1",
			"currency":             "GBP",
			"email":                "makenzie89@example.net",
			"event_time":           "2019-04-15 07:37:53",
			"linked_subscriptions": "",
			"marketing_consent":    "",
			"next_bill_date":       "2019-05-14",
			"passthrough":          "Example String",
			"quantity":             "60",
			"status":               "active",
			"subscription_id":      "1",
			"subscription_plan_id": "5",
			"unit_price":           "49.99",
			"update_url":           "https://checkout.paddle.com/subscription/update?user=4&subscription=2&hash=a0aef1af98b11ef5d220751a77d0eda187f836d4",
		},
	} {
		d := test.Sign(data)
		var mc types.MarketingConsent
		mc.UnmarshalText([]byte(d.M["marketing_consent"]))
		sc := Created{
			AlertID:             int(test.IntFromString(d.M["alert_id"])),
			AlertName:           d.M["alert_name"],
			CancelURL:           d.M["cancel_url"],
			CheckoutID:          d.M["checkout_id"],
			Currency:            d.M["currency"],
			Email:               d.M["email"],
			EventTime:           &types.Datetime{test.ParseTime(types.DatetimeFormat, d.M["event_time"])},
			LinkedSubscriptions: d.M["linked_subscriptions"],
			MarketingConsent:    &mc,
			NextBillDate:        &types.Date{test.ParseTime(types.DateFormat, d.M["next_bill_date"]), false},
			Passthrough:         d.M["passthrough"],
			Quantity:            int(test.IntFromString(d.M["quantity"])),
			Status:              d.M["status"],
			SubscriptionID:      int(test.IntFromString(d.M["subscription_id"])),
			SubscriptionPlanID:  int(test.IntFromString(d.M["subscription_plan_id"])),
			UnitPrice:           test.CurrencyValueFromString("49.99"),
			UpdateURL:           d.M["update_url"],
			PSignature:          d.M["p_signature"],
		}
		subscriptionCreatedData = append(subscriptionCreatedData, struct {
			d  test.Data
			sc Created
		}{d, sc})
	}
	return
}

func TestCreated(t *testing.T) {
	data := subscriptionCreatedData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		for _, d := range data {
			var actual Created
			assert.NoError(json.Unmarshal([]byte(d.d.JSON), &actual))
			assert.Equal(d.sc, actual)
		}
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		for _, d := range data {
			var actual Created
			assert.NoError(urldecode.Unmarshal([]byte(d.d.URL), &actual))
			assert.Equal(d.sc, actual)
		}
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		for _, d := range data {
			b, err := d.sc.Serialize()
			assert.NoError(err)
			assert.Equal(d.d.PHP, string(b))
		}
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		for _, d := range data {
			assert.NoError(events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&d.sc))
		}
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &Created{})
	})
}

func BenchmarkCreated(b *testing.B) {
	data := subscriptionCreatedData()
	var sc Created
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data[0].d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &sc)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data[0].d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &sc)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data[0].sc.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data[0].sc)
		}
	})
}
