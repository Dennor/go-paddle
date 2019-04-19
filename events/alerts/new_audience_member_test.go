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

func subscriptionNewAudienceMemberData() struct {
	d   test.Data
	nam NewAudienceMember
} {
	d := test.Sign(map[string]string{
		"alert_name":        "new_audience_member",
		"created_at":        "2019-04-15 07:37:53",
		"email":             "jan@kowalski.net",
		"event_time":        "2019-04-15 07:37:53",
		"marketing_consent": "1",
		"products":          "12345,678910,1112131415",
		"source":            "Checkout",
		"subscribed":        "1",
		"user_id":           "1234",
	})
	products := AudienceMemberProducts{}
	products.UnmarshalText([]byte(d.M["products"]))
	var mc types.MarketingConsent
	mc.UnmarshalText([]byte(d.M["marketing_consent"]))
	nam := NewAudienceMember{
		AlertName:        d.M["alert_name"],
		CreatedAt:        &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["created_at"])},
		Email:            d.M["email"],
		EventTime:        &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		MarketingConsent: &mc,
		Products:         &products,
		Source:           d.M["source"],
		Subscribed:       int(test.IntFromString(d.M["subscribed"])),
		UserID:           int(test.IntFromString(d.M["user_id"])),
		PSignature:       d.M["p_signature"],
	}
	return struct {
		d   test.Data
		nam NewAudienceMember
	}{d, nam}
}

func TestNewAudienceMember(t *testing.T) {
	data := subscriptionNewAudienceMemberData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual NewAudienceMember
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.nam, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual NewAudienceMember
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.nam, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.nam.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.nam))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &NewAudienceMember{})
	})
}

func BenchmarkNewAudienceMember(b *testing.B) {
	data := subscriptionNewAudienceMemberData()
	var nam NewAudienceMember
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &nam)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &nam)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.nam.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.nam)
		}
	})
}
