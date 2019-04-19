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

func subscriptionHighRiskTransactionCreatedData() struct {
	d    test.Data
	hrtc HighRiskTransactionCreated
} {
	d := test.Sign(map[string]string{
		"alert_name":             "high_risk_transaction_created",
		"case_id":                "1",
		"checkout_id":            "1-c8a82616c183ad6-377f00add1",
		"created_at":             "2019-04-15 07:37:53",
		"customer_email_address": "jan@kowalski.net",
		"customer_user_id":       "1",
		"event_time":             "2019-04-15 07:37:53",
		"marketing_consent":      "1",
		"passthrough":            "Example String",
		"product_id":             "2",
		"risk_score":             "99.99",
		"status":                 "pending",
	})
	var mc types.MarketingConsent
	mc.UnmarshalText([]byte(d.M["marketing_consent"]))
	hrtc := HighRiskTransactionCreated{
		AlertName:            d.M["alert_name"],
		CaseID:               int(test.IntFromString(d.M["case_id"])),
		CheckoutID:           d.M["checkout_id"],
		CreatedAt:            &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["created_at"])},
		CustomerEmailAddress: d.M["customer_email_address"],
		CustomerUserID:       int(test.IntFromString(d.M["customer_user_id"])),
		EventTime:            &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		MarketingConsent:     &mc,
		Passthrough:          d.M["passthrough"],
		ProductID:            int(test.IntFromString(d.M["product_id"])),
		RiskScore:            test.DecimalFromString(d.M["risk_score"]),
		Status:               d.M["status"],
		PSignature:           d.M["p_signature"],
	}
	return struct {
		d    test.Data
		hrtc HighRiskTransactionCreated
	}{d, hrtc}
}

func TestHighRiskTransactionCreated(t *testing.T) {
	data := subscriptionHighRiskTransactionCreatedData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual HighRiskTransactionCreated
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.hrtc, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual HighRiskTransactionCreated
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.hrtc, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.hrtc.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.hrtc))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &HighRiskTransactionCreated{})
	})
}

func BenchmarkHighRiskTransactionCreated(b *testing.B) {
	data := subscriptionHighRiskTransactionCreatedData()
	var hrtc HighRiskTransactionCreated
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &hrtc)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &hrtc)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.hrtc.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.hrtc)
		}
	})
}
