package middleware

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"unsafe"

	"github.com/dennor/go-paddle/events"
	"github.com/dennor/go-paddle/events/alerts"
	"github.com/dennor/go-paddle/events/subscription"
	"github.com/dennor/go-paddle/httperrors"
	"github.com/dennor/go-paddle/mime"
)

const (
	DefaultContextErrKey = "paddle-event-middleware-error"
	DefaultContextKey    = "paddle-event-middleware"
)

type buffer struct {
	bytes.Buffer
}

func (b *buffer) Close() error { return nil }

type EventConfig struct {
	Verifier        events.Verifier
	ContextKey      interface{}
	ContextErrKey   interface{}
	CopyBody        bool
	ContinueOnError bool
	SkipContext     bool
}

type Event struct {
	EventConfig
}

type bodyBufferPool struct {
	sync.Pool
}

func newBodyBufferPool() bodyBufferPool {
	return bodyBufferPool{
		Pool: sync.Pool{
			New: func() interface{} {
				return &buffer{}
			},
		},
	}
}

func (p *bodyBufferPool) Get() *buffer {
	buf := p.Pool.Get().(*buffer)
	buf.Reset()
	return buf
}

func (p *bodyBufferPool) Put(buf *buffer) {
	p.Pool.Put(buf)
}

var (
	bodyPool = newBodyBufferPool()
)

var (
	alertKey = []byte("alert_name=")
)

func eventNameFromURLEncoded(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	var j, start, end int
	for i := 0; i < len(b); i++ {
		j = 0
		if i != 0 && b[i] != '&' {
			continue
		}
		if b[i] == '&' {
			i++
		}
		for _, ac := range alertKey {
			if b[i] != ac || i >= len(b) {
				break
			}
			j++
			i++
		}
		if j == len(alertKey) {
			start = i
			for i < len(b) && b[i] != '&' {
				i++
			}
			end = i
		}
	}
	if start != 0 && end > start {
		return (*(*string)(unsafe.Pointer(&b)))[start:end]
	}
	return ""
}

type unmarshalFunc func(io.Reader, interface{}) error

var (
	unmarshalForm = events.UnmarshalForm
	unmarshalJSON = events.UnmarshalJSON
)

func unmarshalEvent(ename string, r io.Reader, f unmarshalFunc) (events.Event, error) {
	var e events.Event
	switch ename {
	case subscription.CancelledAlertName:
		e = new(subscription.Cancelled)
	case subscription.CreatedAlertName:
		e = new(subscription.Created)
	case subscription.PaymentFailedAlertName:
		e = new(subscription.PaymentFailed)
	case subscription.PaymentRefundedAlertName:
		e = new(subscription.PaymentRefunded)
	case subscription.PaymentSucceededAlertName:
		e = new(subscription.PaymentSucceeded)
	case subscription.UpdatedAlertName:
		e = new(subscription.Updated)
	case alerts.HighRiskTransactionCreatedAlertName:
		e = new(alerts.HighRiskTransactionCreated)
	case alerts.HighRiskTransactionUpdatedAlertName:
		e = new(alerts.HighRiskTransactionUpdated)
	case alerts.LockerProcessedAlertName:
		e = new(alerts.LockerProcessed)
	case alerts.NewAudienceMemberAlertName:
		e = new(alerts.NewAudienceMember)
	case alerts.PaymentDisputeClosedAlertName:
		e = new(alerts.PaymentDisputeClosed)
	case alerts.PaymentDisputeCreatedAlertName:
		e = new(alerts.PaymentDisputeCreated)
	case alerts.PaymentRefundedAlertName:
		e = new(alerts.PaymentRefunded)
	case alerts.PaymentSucceededAlertName:
		e = new(alerts.PaymentSucceeded)
	case alerts.TransferCreatedAlertName:
		e = new(alerts.TransferCreated)
	case alerts.TransferPaidAlertName:
		e = new(alerts.TransferPaid)
	case alerts.UpdateAudienceMemberAlertName:
		e = new(alerts.UpdateAudienceMember)
	default:
		return nil, httperrors.NewBadRequestError(ename + " is not a supported event type")
	}
	return e, f(r, e)
}

func readEventFromRequest(req *http.Request, copyBody bool) (events.Event, error) {
	buf := bodyPool.Get()
	if _, err := io.Copy(buf, req.Body); err != nil {
		return nil, err
	}
	var ename string
	var f unmarshalFunc
	switch {
	case strings.HasPrefix(req.Header.Get(mime.ContentTypeHeader), mime.ApplicationForm):
		ename = eventNameFromURLEncoded(buf.Bytes())
		f = unmarshalForm
	default:
		return nil, httperrors.NewBadRequestError(req.Header.Get(mime.ContentTypeHeader) + " is not supported mime type")
	}
	var r io.Reader
	r = buf
	if copyBody {
		r = bytes.NewReader(buf.Bytes())
	}
	e, err := unmarshalEvent(ename, r, f)
	if copyBody {
		req.Body.Close()
		req.Body = buf
	} else {
		bodyPool.Put(buf)
	}
	return e, err
}

func (e *Event) onError(next http.Handler, rw http.ResponseWriter, req *http.Request, err error) {
	if !e.ContinueOnError {
		if httpError, ok := err.(httperrors.Error); ok {
			httpError.WriteTo(rw)
			return
		}
		http.Error(rw, err.Error(), 500)
		return
	}
	if !e.SkipContext {
		req = req.WithContext(context.WithValue(req.Context(), e.ContextErrKey, err))
	}
	next.ServeHTTP(rw, req)
}

func (e *Event) Handle(next http.Handler) http.Handler {
	if e.ContextErrKey == nil {
		e.ContextErrKey = DefaultContextErrKey
	}
	if e.ContextKey == nil {
		e.ContextKey = DefaultContextKey
	}
	getEvent := e.EventFromRequest()
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		paddleEvent, err := getEvent(req)
		if err != nil {
			e.onError(next, rw, req, err)
			return
		}
		if !e.SkipContext {
			req = req.WithContext(context.WithValue(req.Context(), e.ContextKey, paddleEvent))
		}
		next.ServeHTTP(rw, req)
	})
}

func (e *Event) EventFromRequest() func(req *http.Request) (events.Event, error) {
	verify := e.Verifier != nil
	return func(req *http.Request) (events.Event, error) {
		if !e.SkipContext {
			if ev, ok := req.Context().Value(e.ContextKey).(events.Event); ok && ev != nil {
				return ev, nil
			}
		}
		ev, err := readEventFromRequest(req, e.CopyBody)
		if err != nil {
			return nil, err
		}
		if verify {
			if err := e.Verifier.Verify(ev); err != nil {
				return nil, err
			}
		}
		return ev, nil
	}
}
