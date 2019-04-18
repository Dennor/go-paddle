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

func subscriptionPaymentDisputeClosedData() struct {
	d   test.Data
	pdc PaymentDisputeClosed
} {
	d := test.Sign(map[string]string{
		"alert_name":        "payment_dispute_closed",
		"amount":            "1.23",
		"checkout_id":       "1-c8a82616c183ad6-377f00add1",
		"currency":          "PLN",
		"email":             "jan@kowalski.net",
		"event_time":        "2019-04-15 07:37:53",
		"fee_usd":           "4.56",
		"marketing_consent": "1",
		"order_id":          "2",
		"status":            "closed",
	})
	pdc := PaymentDisputeClosed{
		AlertName:        d.M["alert_name"],
		Amount:           test.DecimalFromString(d.M["amount"]),
		CheckoutID:       d.M["checkout_id"],
		Currency:         d.M["currency"],
		Email:            d.M["email"],
		EventTime:        &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		FeeUsd:           test.DecimalFromString(d.M["fee_usd"]),
		MarketingConsent: types.MarketingConsent(test.IntFromString(d.M["marketing_consent"])),
		OrderID:          int(test.IntFromString(d.M["order_id"])),
		Status:           d.M["status"],
		PSignature:       d.M["p_signature"],
	}
	return struct {
		d   test.Data
		pdc PaymentDisputeClosed
	}{d, pdc}
}

func TestPaymentDisputeClosed(t *testing.T) {
	data := subscriptionPaymentDisputeClosedData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual PaymentDisputeClosed
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.pdc, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual PaymentDisputeClosed
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.pdc, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.pdc.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.pdc))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &PaymentDisputeClosed{})
	})
}

func BenchmarkPaymentDisputeClosed(b *testing.B) {
	data := subscriptionPaymentDisputeClosedData()
	var pdc PaymentDisputeClosed
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &pdc)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &pdc)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.pdc.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.pdc)
		}
	})
}
