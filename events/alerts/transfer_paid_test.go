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

func subscriptionTransferPaidData() struct {
	d  test.Data
	tp TransferPaid
} {
	d := test.Sign(map[string]string{
		"alert_name": "transfer_paid",
		"amount":     "1.23",
		"currency":   "PLN",
		"event_time": "2019-04-15 07:37:53",
		"payout_id":  "2",
		"status":     "closed",
	})
	tp := TransferPaid{
		AlertName:  d.M["alert_name"],
		Amount:     test.DecimalFromString(d.M["amount"]),
		Currency:   d.M["currency"],
		EventTime:  &types.TimeYYYYMMDDHHmmSS{test.ParseTime(types.TimeFormatYYYYMMDDHHmmSS, d.M["event_time"])},
		PayoutID:   int(test.IntFromString(d.M["payout_id"])),
		Status:     d.M["status"],
		PSignature: d.M["p_signature"],
	}
	return struct {
		d  test.Data
		tp TransferPaid
	}{d, tp}
}

func TestTransferPaid(t *testing.T) {
	data := subscriptionTransferPaidData()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		assert := assert.New(t)
		var actual TransferPaid
		assert.NoError(json.Unmarshal([]byte(data.d.JSON), &actual))
		assert.Equal(data.tp, actual)
	})

	t.Run("UnmarshalURL", func(t *testing.T) {
		assert := assert.New(t)
		var actual TransferPaid
		assert.NoError(urldecode.Unmarshal([]byte(data.d.URL), &actual))
		assert.Equal(data.tp, actual)
	})

	t.Run("Serialize", func(t *testing.T) {
		assert := assert.New(t)
		b, err := data.tp.Serialize()
		assert.NoError(err)
		assert.Equal(data.d.PHP, string(b))
	})

	t.Run("Verify", func(t *testing.T) {
		assert := assert.New(t)
		assert.NoError(events.RSAVerifier(signature.RSA{
			PublicKey: &test.Key.PublicKey,
		}).Verify(&data.tp))
	})

	t.Run("ImplementsEvent", func(t *testing.T) {
		assert := assert.New(t)
		assert.Implements((*events.Event)(nil), &TransferPaid{})
	})
}

func BenchmarkTransferPaid(b *testing.B) {
	data := subscriptionTransferPaidData()
	var tp TransferPaid
	b.Run("UnmarshalJSON", func(b *testing.B) {
		payload := []byte(data.d.JSON)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			json.Unmarshal(payload, &tp)
		}
	})
	b.Run("UnmarshalURL", func(b *testing.B) {
		payload := []byte(data.d.URL)
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			urldecode.Unmarshal(payload, &tp)
		}
	})

	b.Run("Serialize", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			data.tp.Serialize()
		}
	})

	b.Run("Verify", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			events.RSAVerifier(signature.RSA{
				PublicKey: &test.Key.PublicKey,
			}).Verify(&data.tp)
		}
	})
}
