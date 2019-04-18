package router

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/dennor/go-paddle/events/alerts"
	"github.com/dennor/go-paddle/events/subscription"
	"github.com/dennor/go-paddle/mime"
	"github.com/stretchr/testify/mock"
)

type mockAlertHighRiskTransactionCreated struct {
	mock.Mock
}

func (m *mockAlertHighRiskTransactionCreated) ServeHTTP(e *alerts.HighRiskTransactionCreated, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertHighRiskTransactionUpdated struct {
	mock.Mock
}

func (m *mockAlertHighRiskTransactionUpdated) ServeHTTP(e *alerts.HighRiskTransactionUpdated, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertLockerProcessed struct {
	mock.Mock
}

func (m *mockAlertLockerProcessed) ServeHTTP(e *alerts.LockerProcessed, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertNewAudienceMember struct {
	mock.Mock
}

func (m *mockAlertNewAudienceMember) ServeHTTP(e *alerts.NewAudienceMember, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertPaymentDisputeClosed struct {
	mock.Mock
}

func (m *mockAlertPaymentDisputeClosed) ServeHTTP(e *alerts.PaymentDisputeClosed, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertPaymentDisputeCreated struct {
	mock.Mock
}

func (m *mockAlertPaymentDisputeCreated) ServeHTTP(e *alerts.PaymentDisputeCreated, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertPaymentRefunded struct {
	mock.Mock
}

func (m *mockAlertPaymentRefunded) ServeHTTP(e *alerts.PaymentRefunded, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertPaymentSucceeded struct {
	mock.Mock
}

func (m *mockAlertPaymentSucceeded) ServeHTTP(e *alerts.PaymentSucceeded, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertTransferCreated struct {
	mock.Mock
}

func (m *mockAlertTransferCreated) ServeHTTP(e *alerts.TransferCreated, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertTransferPaid struct {
	mock.Mock
}

func (m *mockAlertTransferPaid) ServeHTTP(e *alerts.TransferPaid, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockAlertUpdateAudienceMember struct {
	mock.Mock
}

func (m *mockAlertUpdateAudienceMember) ServeHTTP(e *alerts.UpdateAudienceMember, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockSubscriptionCancelled struct {
	mock.Mock
}

func (m *mockSubscriptionCancelled) ServeHTTP(e *subscription.Cancelled, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockSubscriptionCreated struct {
	mock.Mock
}

func (m *mockSubscriptionCreated) ServeHTTP(e *subscription.Created, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockSubscriptionPaymentFailed struct {
	mock.Mock
}

func (m *mockSubscriptionPaymentFailed) ServeHTTP(e *subscription.PaymentFailed, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockSubscriptionPaymentRefunded struct {
	mock.Mock
}

func (m *mockSubscriptionPaymentRefunded) ServeHTTP(e *subscription.PaymentRefunded, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockSubscriptionPaymentSucceeded struct {
	mock.Mock
}

func (m *mockSubscriptionPaymentSucceeded) ServeHTTP(e *subscription.PaymentSucceeded, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
}

type mockSubscriptionUpdated struct {
	mock.Mock
}

func (m *mockSubscriptionUpdated) ServeHTTP(e *subscription.Updated, rw http.ResponseWriter, req *http.Request) {
	m.Called(e, rw, req)
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

func TestRouter(t *testing.T) {
	t.Run("CallsServeHTTPWithEvent", func(t *testing.T) {
		data := []struct {
			router Router
			query  []byte
		}{
			{
				router: Router{
					Config: Config{
						AlertHighRiskTransactionCreated: func() AlertHighRiskTransactionCreated {
							m := &mockAlertHighRiskTransactionCreated{}
							m.On("ServeHTTP", &alerts.HighRiskTransactionCreated{
								AlertName: "high_risk_transaction_created",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=high_risk_transaction_created"),
			},
			{
				router: Router{
					Config: Config{
						AlertHighRiskTransactionUpdated: func() AlertHighRiskTransactionUpdated {
							m := &mockAlertHighRiskTransactionUpdated{}
							m.On("ServeHTTP", &alerts.HighRiskTransactionUpdated{
								AlertName: "high_risk_transaction_updated",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=high_risk_transaction_updated"),
			},
			{
				router: Router{
					Config: Config{
						AlertLockerProcessed: func() AlertLockerProcessed {
							m := &mockAlertLockerProcessed{}
							m.On("ServeHTTP", &alerts.LockerProcessed{
								AlertName: "locker_processed",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=locker_processed"),
			},
			{
				router: Router{
					Config: Config{
						AlertNewAudienceMember: func() AlertNewAudienceMember {
							m := &mockAlertNewAudienceMember{}
							m.On("ServeHTTP", &alerts.NewAudienceMember{
								AlertName: "new_audience_member",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=new_audience_member"),
			},
			{
				router: Router{
					Config: Config{
						AlertPaymentDisputeClosed: func() AlertPaymentDisputeClosed {
							m := &mockAlertPaymentDisputeClosed{}
							m.On("ServeHTTP", &alerts.PaymentDisputeClosed{
								AlertName: "payment_dispute_closed",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=payment_dispute_closed"),
			},
			{
				router: Router{
					Config: Config{
						AlertPaymentDisputeCreated: func() AlertPaymentDisputeCreated {
							m := &mockAlertPaymentDisputeCreated{}
							m.On("ServeHTTP", &alerts.PaymentDisputeCreated{
								AlertName: "payment_dispute_created",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=payment_dispute_created"),
			},
			{
				router: Router{
					Config: Config{
						AlertPaymentRefunded: func() AlertPaymentRefunded {
							m := &mockAlertPaymentRefunded{}
							m.On("ServeHTTP", &alerts.PaymentRefunded{
								AlertName: "payment_refunded",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=payment_refunded"),
			},
			{
				router: Router{
					Config: Config{
						AlertPaymentSucceeded: func() AlertPaymentSucceeded {
							m := &mockAlertPaymentSucceeded{}
							m.On("ServeHTTP", &alerts.PaymentSucceeded{
								AlertName: "payment_succeeded",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=payment_succeeded"),
			},
			{
				router: Router{
					Config: Config{
						AlertTransferCreated: func() AlertTransferCreated {
							m := &mockAlertTransferCreated{}
							m.On("ServeHTTP", &alerts.TransferCreated{
								AlertName: "transfer_created",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=transfer_created"),
			},
			{
				router: Router{
					Config: Config{
						AlertTransferPaid: func() AlertTransferPaid {
							m := &mockAlertTransferPaid{}
							m.On("ServeHTTP", &alerts.TransferPaid{
								AlertName: "transfer_paid",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=transfer_paid"),
			},
			{
				router: Router{
					Config: Config{
						AlertUpdateAudienceMember: func() AlertUpdateAudienceMember {
							m := &mockAlertUpdateAudienceMember{}
							m.On("ServeHTTP", &alerts.UpdateAudienceMember{
								AlertName: "update_audience_member",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=update_audience_member"),
			},
			{
				router: Router{
					Config: Config{
						SubscriptionCancelled: func() SubscriptionCancelled {
							m := &mockSubscriptionCancelled{}
							m.On("ServeHTTP", &subscription.Cancelled{
								AlertName: "subscription_cancelled",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=subscription_cancelled"),
			},
			{
				router: Router{
					Config: Config{
						SubscriptionCreated: func() SubscriptionCreated {
							m := &mockSubscriptionCreated{}
							m.On("ServeHTTP", &subscription.Created{
								AlertName: "subscription_created",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=subscription_created"),
			},
			{
				router: Router{
					Config: Config{
						SubscriptionPaymentFailed: func() SubscriptionPaymentFailed {
							m := &mockSubscriptionPaymentFailed{}
							m.On("ServeHTTP", &subscription.PaymentFailed{
								AlertName: "subscription_payment_failed",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=subscription_payment_failed"),
			},
			{
				router: Router{
					Config: Config{
						SubscriptionPaymentRefunded: func() SubscriptionPaymentRefunded {
							m := &mockSubscriptionPaymentRefunded{}
							m.On("ServeHTTP", &subscription.PaymentRefunded{
								AlertName: "subscription_payment_refunded",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=subscription_payment_refunded"),
			},
			{
				router: Router{
					Config: Config{
						SubscriptionPaymentSucceeded: func() SubscriptionPaymentSucceeded {
							m := &mockSubscriptionPaymentSucceeded{}
							m.On("ServeHTTP", &subscription.PaymentSucceeded{
								AlertName: "subscription_payment_succeeded",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=subscription_payment_succeeded"),
			},
			{
				router: Router{
					Config: Config{
						SubscriptionUpdated: func() SubscriptionUpdated {
							m := &mockSubscriptionUpdated{}
							m.On("ServeHTTP", &subscription.Updated{
								AlertName: "subscription_updated",
							}, mock.AnythingOfType("*router.mockResponseWriter"), mock.AnythingOfType("*http.Request"))
							return m
						}(),
					},
				},
				query: []byte("alert_name=subscription_updated"),
			},
		}
		for _, tt := range data {
			req := &http.Request{
				Header: make(http.Header),
			}
			req.Header.Set(mime.ContentTypeHeader, mime.ApplicationForm)
			req.Body = ioutil.NopCloser(bytes.NewReader(tt.query))
			mockWriter := new(mockResponseWriter)
			tt.router.Handler().ServeHTTP(mockWriter, req)
		}
	})
}

type benchAlertHighRiskTransactionCreated struct{}

func (*benchAlertHighRiskTransactionCreated) ServeHTTP(e *alerts.HighRiskTransactionCreated, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertHighRiskTransactionUpdated struct {
}

func (*benchAlertHighRiskTransactionUpdated) ServeHTTP(e *alerts.HighRiskTransactionUpdated, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertLockerProcessed struct {
}

func (*benchAlertLockerProcessed) ServeHTTP(e *alerts.LockerProcessed, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertNewAudienceMember struct {
}

func (*benchAlertNewAudienceMember) ServeHTTP(e *alerts.NewAudienceMember, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertPaymentDisputeClosed struct {
}

func (*benchAlertPaymentDisputeClosed) ServeHTTP(e *alerts.PaymentDisputeClosed, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertPaymentDisputeCreated struct {
}

func (*benchAlertPaymentDisputeCreated) ServeHTTP(e *alerts.PaymentDisputeCreated, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertPaymentRefunded struct {
}

func (*benchAlertPaymentRefunded) ServeHTTP(e *alerts.PaymentRefunded, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertPaymentSucceeded struct {
}

func (*benchAlertPaymentSucceeded) ServeHTTP(e *alerts.PaymentSucceeded, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertTransferCreated struct {
}

func (*benchAlertTransferCreated) ServeHTTP(e *alerts.TransferCreated, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertTransferPaid struct {
}

func (*benchAlertTransferPaid) ServeHTTP(e *alerts.TransferPaid, rw http.ResponseWriter, req *http.Request) {

}

type benchAlertUpdateAudienceMember struct {
}

func (*benchAlertUpdateAudienceMember) ServeHTTP(e *alerts.UpdateAudienceMember, rw http.ResponseWriter, req *http.Request) {

}

type benchSubscriptionCancelled struct {
}

func (*benchSubscriptionCancelled) ServeHTTP(e *subscription.Cancelled, rw http.ResponseWriter, req *http.Request) {

}

type benchSubscriptionCreated struct {
}

func (*benchSubscriptionCreated) ServeHTTP(e *subscription.Created, rw http.ResponseWriter, req *http.Request) {

}

type benchSubscriptionPaymentFailed struct {
}

func (*benchSubscriptionPaymentFailed) ServeHTTP(e *subscription.PaymentFailed, rw http.ResponseWriter, req *http.Request) {

}

type benchSubscriptionPaymentRefunded struct {
}

func (*benchSubscriptionPaymentRefunded) ServeHTTP(e *subscription.PaymentRefunded, rw http.ResponseWriter, req *http.Request) {

}

type benchSubscriptionPaymentSucceeded struct {
}

func (*benchSubscriptionPaymentSucceeded) ServeHTTP(e *subscription.PaymentSucceeded, rw http.ResponseWriter, req *http.Request) {

}

type benchSubscriptionUpdated struct {
}

func (*benchSubscriptionUpdated) ServeHTTP(e *subscription.Updated, rw http.ResponseWriter, req *http.Request) {

}

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

func BenchmarkRouter(b *testing.B) {
	router := Router{
		Config: Config{
			AlertHighRiskTransactionCreated: &benchAlertHighRiskTransactionCreated{},
			AlertHighRiskTransactionUpdated: &benchAlertHighRiskTransactionUpdated{},
			AlertLockerProcessed:            &benchAlertLockerProcessed{},
			AlertNewAudienceMember:          &benchAlertNewAudienceMember{},
			AlertPaymentDisputeClosed:       &benchAlertPaymentDisputeClosed{},
			AlertPaymentDisputeCreated:      &benchAlertPaymentDisputeCreated{},
			AlertPaymentRefunded:            &benchAlertPaymentRefunded{},
			AlertPaymentSucceeded:           &benchAlertPaymentSucceeded{},
			AlertTransferCreated:            &benchAlertTransferCreated{},
			AlertTransferPaid:               &benchAlertTransferPaid{},
			AlertUpdateAudienceMember:       &benchAlertUpdateAudienceMember{},
			SubscriptionCancelled:           &benchSubscriptionCancelled{},
			SubscriptionCreated:             &benchSubscriptionCreated{},
			SubscriptionPaymentFailed:       &benchSubscriptionPaymentFailed{},
			SubscriptionPaymentRefunded:     &benchSubscriptionPaymentRefunded{},
			SubscriptionPaymentSucceeded:    &benchSubscriptionPaymentSucceeded{},
			SubscriptionUpdated:             &benchSubscriptionUpdated{},
		},
	}
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
	for _, tt := range data {
		req := &http.Request{
			Header: make(http.Header),
		}
		req.Header.Set(mime.ContentTypeHeader, mime.ApplicationForm)
		req.Body = ioutil.NopCloser(&tt.r)
		rw := benchResponseWriter{}
		handler := router.Handler()
		b.Run(tt.benchmarkName, func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				tt.r.Reset()
				handler.ServeHTTP(&rw, req)
			}
		})
	}
}
