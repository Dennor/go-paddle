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

func subscriptionTransferCreatedData() struct {
	d  test.Data
	tc TransferCreated
} {
	d := test.Sign(map[string]string{
		"alert_name": "transfer_paid",
		"amount":     "1.23",
		"currency":   "PLN",
		"event_time": "2019-04-15 07:37:53",
		"payout_id":  "2",
		"status":     "closed",
	})
	tc := TransferCreated{
		AlertName:  d.M["alert_name"],
		Amount:     test.DecimalFromString(d.M["amount"]),
		Currency:   d.M["currency"],
		EventTime:  &types.Datetime{test.ParseTime(types.DatetimeFormat, d.M["event_time"])},
		PayoutID:   int(test.IntFromString(d.M["payout_id"])),
		Status:     d.M["status"],
		PSignature: d.M["p_signature"],
	}
	return struct {
		d  test.Data
		tc TransferCreated
	}{d, tc}
}

func TestTransferCreated(t *testing.T) {
	data := subscriptionTransferCreatedData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual TransferCreated
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.tc, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual TransferCreated
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.tc, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.tc.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.tc))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &TransferCreated{})
	})
}

func BenchmarkTransferCreated(b *testing.B) {
	data := subscriptionTransferCreatedData()
	var tc TransferCreated
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &tc)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &tc)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.tc.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.tc)
		}
	})
}
