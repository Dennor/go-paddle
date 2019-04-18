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

func subscriptionCancelledData() struct {
	d  test.Data
	sc Cancelled
} {
	d := test.Sign(map[string]string{
		"alert_name":                  "subscription_cancelled",
		"cancellation_effective_date": "2019-05-14",
		"checkout_id":                 "1-c8a82616c183ad6-377f00add1",
		"currency":                    "GBP",
		"email":                       "makenzie89@example.net",
		"event_time":                  "2019-04-15 07:37:53",
		"marketing_consent":           "1",
		"passthrough":                 "Example String",
		"quantity":                    "1",
		"status":                      "active",
		"subscription_id":             "1",
		"subscription_plan_id":        "5",
		"unit_price":                  "49.99",
		"user_id":                     "10",
	})
	sc := Cancelled{
		AlertName:                 d.M["alert_name"],
		CancellationEffectiveDate: &types.TimeYYYYMMDD{test.ParseTime(types.TimeFormatYYYYMMDD, d.M["cancellation_effective_date"])},
		CheckoutID:                d.M["checkout_id"],
		Currency:                  d.M["currency"],
		Email:                     d.M["email"],
		EventTime:                 &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		MarketingConsent:          types.MarketingConsent(test.IntFromString(d.M["marketing_consent"])),
		Passthrough:               d.M["passthrough"],
		Quantity:                  int(test.IntFromString(d.M["quantity"])),
		Status:                    d.M["status"],
		SubscriptionID:            int(test.IntFromString(d.M["subscription_id"])),
		SubscriptionPlanID:        int(test.IntFromString(d.M["subscription_plan_id"])),
		UnitPrice:                 test.DecimalFromString(d.M["unit_price"]),
		UserID:                    int(test.IntFromString(d.M["user_id"])),
		PSignature:                d.M["p_signature"],
	}
	return struct {
		d  test.Data
		sc Cancelled
	}{d, sc}
}

func TestCancelled(t *testing.T) {
	data := subscriptionCancelledData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual Cancelled
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.sc, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual Cancelled
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.sc, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.sc.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.sc))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &Cancelled{})
	})
}

func BenchmarkCancelled(b *testing.B) {
	data := subscriptionCancelledData()
	var sc Cancelled
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &sc)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &sc)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.sc.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.sc)
		}
	})
}
