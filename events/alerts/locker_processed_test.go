package alerts

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

func lockerProcessedData() struct {
	d  test.Data
	lp LockerProcessed
} {
	d := test.Sign(map[string]string{
		"alert_name":        "locker_processed",
		"checkout_id":       "1-c8a82616c183ad6-377f00add1",
		"checkout_recovery": "1",
		"coupon":            "some coupon",
		"download":          "http://example.com/get/me?from=this&download=url",
		"email":             "makenzie89@example.net",
		"event_time":        "2019-04-15 07:37:53",
		"instructions":      "download and run!!",
		"license":           "MIT",
		"marketing_consent": "1",
		"order_id":          "12",
		"product_id":        "1",
		"quantity":          "1",
	})
	lp := LockerProcessed{
		AlertName:        d.M["alert_name"],
		CheckoutID:       d.M["checkout_id"],
		CheckoutRecovery: int(test.IntFromString(d.M["checkout_recovery"])),
		Coupon:           d.M["coupon"],
		Download:         d.M["download"],
		Email:            d.M["email"],
		EventTime:        &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		Instructions:     d.M["instructions"],
		License:          d.M["license"],
		MarketingConsent: types.MarketingConsent(test.IntFromString(d.M["marketing_consent"])),
		OrderID:          int(test.IntFromString(d.M["order_id"])),
		ProductID:        int(test.IntFromString(d.M["product_id"])),
		Quantity:         int(test.IntFromString(d.M["quantity"])),
		PSignature:       d.M["p_signature"],
	}
	return struct {
		d  test.Data
		lp LockerProcessed
	}{d, lp}
}

func TestLockerProcessed(t *testing.T) {
	data := lockerProcessedData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual LockerProcessed
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.lp, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual LockerProcessed
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.lp, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.lp.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.lp))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &LockerProcessed{})
	})
}

func BenchmarkLockerProcessed(b *testing.B) {
	data := lockerProcessedData()
	var lp LockerProcessed
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &lp)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &lp)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.lp.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.lp)
		}
	})
}
