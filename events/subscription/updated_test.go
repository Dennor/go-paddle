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

func subscriptionUpdatedData() struct {
	d  test.Data
	su Updated
} {
	d := test.Sign(map[string]string{
		"alert_name":           "subscription_updated",
		"cancel_url":           "https://checkout.paddle.com/subscription/cancel?user=5&subscription=4&hash=a4dca832089cc76fc05da732664d971c6a7c8840",
		"checkout_id":          "1-c8a82616c183ad6-377f00add1",
		"email":                "makenzie89@example.net",
		"event_time":           "2019-04-15 07:37:53",
		"marketing_consent":    "1",
		"new_price":            "79.98",
		"new_quantity":         "2",
		"new_unit_price":       "39.99",
		"next_bill_date":       "2019-05-14",
		"old_price":            "99.98",
		"old_quantity":         "1",
		"old_unit_price":       "49.99",
		"passthrough":          "Example String",
		"status":               "active",
		"subscription_id":      "1",
		"subscription_plan_id": "5",
		"update_url":           "https://checkout.paddle.com/subscription/update?user=4&subscription=2&hash=a0aef1af98b11ef5d220751a77d0eda187f836d4",
	})
	su := Updated{
		AlertName:          d.M["alert_name"],
		CancelUrl:          d.M["cancel_url"],
		CheckoutID:         d.M["checkout_id"],
		Email:              d.M["email"],
		EventTime:          &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		MarketingConsent:   types.MarketingConsent(test.IntFromString(d.M["marketing_consent"])),
		NewPrice:           test.DecimalFromString(d.M["new_price"]),
		NewQuantity:        int(test.IntFromString(d.M["new_quantity"])),
		NewUnitPrice:       test.DecimalFromString(d.M["new_unit_price"]),
		NextBillDate:       &types.TimeYYYYMMDD{test.ParseTime(types.TimeFormatYYYYMMDD, d.M["next_bill_date"])},
		OldPrice:           test.DecimalFromString(d.M["old_price"]),
		OldQuantity:        int(test.IntFromString(d.M["old_quantity"])),
		OldUnitPrice:       test.DecimalFromString(d.M["old_unit_price"]),
		Passthrough:        d.M["passthrough"],
		Status:             d.M["status"],
		SubscriptionID:     int(test.IntFromString(d.M["subscription_id"])),
		SubscriptionPlanID: int(test.IntFromString(d.M["subscription_plan_id"])),
		UpdateURL:          d.M["update_url"],
		PSignature:         d.M["p_signature"],
	}
	return struct {
		d  test.Data
		su Updated
	}{d, su}
}

func TestUpdated(t *testing.T) {
	data := subscriptionUpdatedData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual Updated
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.su, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual Updated
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.su, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.su.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.su))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &Updated{})
	})
}

func BenchmarkUpdated(b *testing.B) {
	data := subscriptionUpdatedData()
	var su Updated
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &su)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &su)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.su.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.su)
		}
	})
}
