package middleware

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/dennor/go-paddle/events"
	"github.com/dennor/go-paddle/events/alerts"
	"github.com/dennor/go-paddle/events/subscription"
	"github.com/dennor/go-paddle/mime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEventNameFromURLEncoded(t *testing.T) {
	assert := assert.New(t)
	data := []struct {
		query        []byte
		expectedName string
	}{
		{
			query:        []byte("alert_name=alertName"),
			expectedName: "alertName",
		},
		{
			query:        []byte("alert_name=alertName&other_parameter=otherParameter"),
			expectedName: "alertName",
		},
		{
			query:        []byte("beforeParameter=beforeParameter&alert_name=alertName&other_parameter=otherParameter"),
			expectedName: "alertName",
		},
	}
	for _, tt := range data {
		assert.Equal(tt.expectedName, eventNameFromURLEncoded(tt.query), "query was %s", string(tt.query))
	}
}

var benchmarkEventNameDoNotOptimize string

func BenchmarkEventNameFromURLEncoded(b *testing.B) {
	data := []struct {
		query        []byte
		expectedName string
		name         string
	}{
		{
			query:        []byte("alert_name=alertName"),
			expectedName: "alertName",
			name:         "OnlyAlert",
		},
		{
			query:        []byte("alert_name=alertName&other_parameter=otherParameter"),
			expectedName: "alertName",
			name:         "AlertWithFollowing",
		},
		{
			query:        []byte("beforeParameter=beforeParameter&alert_name=alertName&other_parameter=otherParameter"),
			expectedName: "alertName",
			name:         "AlertWithBefore",
		},
	}
	for _, bb := range data {
		b.Run(bb.name, func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				benchmarkEventNameDoNotOptimize = eventNameFromURLEncoded(bb.query)
			}
		})
	}
}

type mockHandler struct {
	mock.Mock
}

func (m *mockHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	m.Called(rw, req)
}

type mockResponseWriter struct {
	mock.Mock
}

func (m *mockResponseWriter) Header() http.Header {
	h := m.Called().Get(0)
	if h == nil {
		return nil
	}
	return h.(http.Header)
}

func (m *mockResponseWriter) Write(b []byte) (int, error) {
	called := m.Called(b)
	return called.Int(0), called.Error(1)
}

func (m *mockResponseWriter) WriteHeader(i int) {
	m.Called(i)
}

func TestEventMiddleware(t *testing.T) {
	t.Run("CallsNextWithEvent", func(t *testing.T) {
		data := []struct {
			query []byte
			event events.Event
		}{
			{
				query: []byte("alert_name=subscription_cancelled"),
				event: &subscription.Cancelled{
					AlertName: "subscription_cancelled",
				},
			},
			{
				query: []byte("alert_name=subscription_created"),
				event: &subscription.Created{
					AlertName: "subscription_created",
				},
			},
			{
				query: []byte("alert_name=subscription_payment_failed"),
				event: &subscription.PaymentFailed{
					AlertName: "subscription_payment_failed",
				},
			},
			{
				query: []byte("alert_name=subscription_payment_refunded"),
				event: &subscription.PaymentRefunded{
					AlertName: "subscription_payment_refunded",
				},
			},
			{
				query: []byte("alert_name=subscription_payment_succeeded"),
				event: &subscription.PaymentSucceeded{
					AlertName: "subscription_payment_succeeded",
				},
			},
			{
				query: []byte("alert_name=subscription_updated"),
				event: &subscription.Updated{
					AlertName: "subscription_updated",
				},
			},
			{
				query: []byte("alert_name=high_risk_transaction_created"),
				event: &alerts.HighRiskTransactionCreated{
					AlertName: "high_risk_transaction_created",
				},
			},
			{
				query: []byte("alert_name=high_risk_transaction_updated"),
				event: &alerts.HighRiskTransactionUpdated{
					AlertName: "high_risk_transaction_updated",
				},
			},
			{
				query: []byte("alert_name=locker_processed"),
				event: &alerts.LockerProcessed{
					AlertName: "locker_processed",
				},
			},
			{
				query: []byte("alert_name=new_audience_member"),
				event: &alerts.NewAudienceMember{
					AlertName: "new_audience_member",
				},
			},
			{
				query: []byte("alert_name=payment_dispute_closed"),
				event: &alerts.PaymentDisputeClosed{
					AlertName: "payment_dispute_closed",
				},
			},
			{
				query: []byte("alert_name=payment_dispute_created"),
				event: &alerts.PaymentDisputeCreated{
					AlertName: "payment_dispute_created",
				},
			},
			{
				query: []byte("alert_name=payment_refunded"),
				event: &alerts.PaymentRefunded{
					AlertName: "payment_refunded",
				},
			},
			{
				query: []byte("alert_name=payment_succeeded"),
				event: &alerts.PaymentSucceeded{
					AlertName: "payment_succeeded",
				},
			},
			{
				query: []byte("alert_name=transfer_created"),
				event: &alerts.TransferCreated{
					AlertName: "transfer_created",
				},
			},
			{
				query: []byte("alert_name=transfer_paid"),
				event: &alerts.TransferPaid{
					AlertName: "transfer_paid",
				},
			},
			{
				query: []byte("alert_name=update_audience_member"),
				event: &alerts.UpdateAudienceMember{
					AlertName: "update_audience_member",
				},
			},
		}
		for _, tt := range data {
			req := &http.Request{
				Header: make(http.Header),
			}
			req.Header.Set(mime.ContentTypeHeader, mime.ApplicationForm)
			req.Body = ioutil.NopCloser(bytes.NewReader(tt.query))
			expectedReq := new(http.Request)
			*expectedReq = *req
			expectedReq = expectedReq.WithContext(context.WithValue(expectedReq.Context(), DefaultContextKey, tt.event))
			mockWriter := new(mockResponseWriter)
			mockHandler := new(mockHandler)
			mockHandler.On("ServeHTTP", mockWriter, expectedReq).Once()
			(&Event{}).Handle(mockHandler).ServeHTTP(mockWriter, req)
		}
	})
}

type benchHandler struct{}

func (b *benchHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {}

type benchResponseWriter struct{}

func (b *benchResponseWriter) Header() http.Header {
	return nil
}

func (b *benchResponseWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func (b *benchResponseWriter) WriteHeader(i int) {}

type reader struct {
	bytes.Reader
}

func (r *reader) Close() error { return nil }
func (r *reader) Reset()       { r.Seek(0, io.SeekStart) }

func BenchmarkEventMidleware(b *testing.B) {
	data := []struct {
		r             reader
		benchmarkName string
	}{
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=subscription_cancelled")),
			},
			benchmarkName: "SubscriptionCancelled",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=subscription_created")),
			},
			benchmarkName: "SubscriptionCreated",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=subscription_payment_failed")),
			},
			benchmarkName: "SubscriptionPaymentFailed",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=subscription_payment_refunded")),
			},
			benchmarkName: "SubscriptionPaymentRefunded",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=subscription_payment_succeeded")),
			},
			benchmarkName: "SubscriptionPaymentSucceeded",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=subscription_updated")),
			},
			benchmarkName: "SubscriptionUpdated",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=high_risk_transaction_created")),
			},
			benchmarkName: "HighRiskTransactionCreated",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=high_risk_transaction_updated")),
			},
			benchmarkName: "HighRiskTransactionUpdated",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=locker_processed")),
			},
			benchmarkName: "LockerProcessed",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=new_audience_member")),
			},
			benchmarkName: "NewAudienceMember",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=payment_dispute_closed")),
			},
			benchmarkName: "PaymentDisputeClosed",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=payment_dispute_created")),
			},
			benchmarkName: "PaymentDisputeCreated",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=payment_refunded")),
			},
			benchmarkName: "PaymentRefunded",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=payment_succeeded")),
			},
			benchmarkName: "PaymentSucceeded",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=transfer_created")),
			},
			benchmarkName: "TransferCreated",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=transfer_paid")),
			},
			benchmarkName: "TransferPaid",
		},
		{
			r: reader{
				Reader: *bytes.NewReader([]byte("alert_name=update_audience_member")),
			},
			benchmarkName: "UpdateAudienceMember",
		},
	}
	event := Event{}
	for _, tt := range data {
		req := &http.Request{
			Header: make(http.Header),
		}
		req.Header.Set(mime.ContentTypeHeader, mime.ApplicationForm)
		req.Body = ioutil.NopCloser(&tt.r)
		handler := event.Handle(&benchHandler{})
		rw := benchResponseWriter{}
		b.Run(tt.benchmarkName, func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				tt.r.Reset()
				handler.ServeHTTP(&rw, req)
			}
		})
	}
}
