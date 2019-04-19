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

func subscriptionUpdateAudienceMemberData() struct {
	d   test.Data
	uam UpdateAudienceMember
} {
	d := test.Sign(map[string]string{
		"alert_name":            "new_audience_member",
		"event_time":            "2019-04-15 07:37:53",
		"new_customer_email":    "jan@kowalski.net",
		"new_marketing_consent": "1",
		"old_customer_email":    "jan@kowalski.net",
		"old_marketing_consent": "1",
		"products":              "12345,678910,1112131415",
		"source":                "Checkout",
		"updated_at":            "2019-04-15 07:37:53",
		"user_id":               "1234",
	})
	products := AudienceMemberProducts{}
	products.UnmarshalText([]byte(d.M["products"]))
	var nmc types.MarketingConsent
	nmc.UnmarshalText([]byte(d.M["new_marketing_consent"]))
	var omc types.MarketingConsent
	omc.UnmarshalText([]byte(d.M["old_marketing_consent"]))
	uam := UpdateAudienceMember{
		AlertName:           d.M["alert_name"],
		EventTime:           &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		NewCustomerEmail:    d.M["new_customer_email"],
		NewMarketingConsent: &nmc,
		OldCustomerEmail:    d.M["old_customer_email"],
		OldMarketingConsent: &omc,
		Products:            &products,
		Source:              d.M["source"],
		UpdatedAt:           &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["updated_at"])},
		UserID:              int(test.IntFromString(d.M["user_id"])),
		PSignature:          d.M["p_signature"],
	}
	return struct {
		d   test.Data
		uam UpdateAudienceMember
	}{d, uam}
}

func TestUpdateAudienceMember(t *testing.T) {
	data := subscriptionUpdateAudienceMemberData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual UpdateAudienceMember
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.uam, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual UpdateAudienceMember
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.uam, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.uam.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.uam))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &UpdateAudienceMember{})
	})
}

func BenchmarkUpdateAudienceMember(b *testing.B) {
	data := subscriptionUpdateAudienceMemberData()
	var uam UpdateAudienceMember
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &uam)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &uam)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.uam.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.uam)
		}
	})
}
