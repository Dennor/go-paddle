package router

import (
	"net/http"

	"github.com/dennor/go-paddle/events"
	"github.com/dennor/go-paddle/events/alerts"
	"github.com/dennor/go-paddle/events/subscription"
	"github.com/dennor/go-paddle/httperrors"
	"github.com/dennor/go-paddle/middleware"
	"github.com/dennor/go-paddle/signature"
)

type Config struct {
	Verifier                        events.Verifier
	CopyBody                        bool
	AlertHighRiskTransactionCreated AlertHighRiskTransactionCreated
	AlertHighRiskTransactionUpdated AlertHighRiskTransactionUpdated
	AlertLockerProcessed            AlertLockerProcessed
	AlertNewAudienceMember          AlertNewAudienceMember
	AlertPaymentDisputeClosed       AlertPaymentDisputeClosed
	AlertPaymentDisputeCreated      AlertPaymentDisputeCreated
	AlertPaymentRefunded            AlertPaymentRefunded
	AlertPaymentSucceeded           AlertPaymentSucceeded
	AlertTransferCreated            AlertTransferCreated
	AlertTransferPaid               AlertTransferPaid
	AlertUpdateAudienceMember       AlertUpdateAudienceMember
	SubscriptionCancelled           SubscriptionCancelled
	SubscriptionCreated             SubscriptionCreated
	SubscriptionPaymentFailed       SubscriptionPaymentFailed
	SubscriptionPaymentRefunded     SubscriptionPaymentRefunded
	SubscriptionPaymentSucceeded    SubscriptionPaymentSucceeded
	SubscriptionUpdated             SubscriptionUpdated
}

type Router struct {
	Config
	alertHighRiskTransactionCreated AlertHighRiskTransactionCreated
	alertHighRiskTransactionUpdated AlertHighRiskTransactionUpdated
	alertLockerProcessed            AlertLockerProcessed
	alertNewAudienceMember          AlertNewAudienceMember
	alertPaymentDisputeClosed       AlertPaymentDisputeClosed
	alertPaymentDisputeCreated      AlertPaymentDisputeCreated
	alertPaymentRefunded            AlertPaymentRefunded
	alertPaymentSucceeded           AlertPaymentSucceeded
	alertTransferCreated            AlertTransferCreated
	alertTransferPaid               AlertTransferPaid
	alertUpdateAudienceMember       AlertUpdateAudienceMember
	subscriptionCancelled           SubscriptionCancelled
	subscriptionCreated             SubscriptionCreated
	subscriptionPaymentFailed       SubscriptionPaymentFailed
	subscriptionPaymentRefunded     SubscriptionPaymentRefunded
	subscriptionPaymentSucceeded    SubscriptionPaymentSucceeded
	subscriptionUpdated             SubscriptionUpdated
	ev                              middleware.Event
}

func (r Router) Handler() http.Handler {
	r.ev.Verifier = r.Config.Verifier
	r.ev.SkipContext = true
	r.ev.CopyBody = r.CopyBody
	r.alertHighRiskTransactionCreated = r.AlertHighRiskTransactionCreated
	if r.alertHighRiskTransactionCreated == nil {
		r.alertHighRiskTransactionCreated = alertHighRiskTransactionCreatedNotFound
	}
	r.alertHighRiskTransactionUpdated = r.AlertHighRiskTransactionUpdated
	if r.alertHighRiskTransactionUpdated == nil {
		r.alertHighRiskTransactionUpdated = alertHighRiskTransactionUpdatedNotFound
	}
	r.alertLockerProcessed = r.AlertLockerProcessed
	if r.alertLockerProcessed == nil {
		r.alertLockerProcessed = alertLockerProcessedNotFound
	}
	r.alertNewAudienceMember = r.AlertNewAudienceMember
	if r.alertNewAudienceMember == nil {
		r.alertNewAudienceMember = alertNewAudienceMemberNotFound
	}
	r.alertPaymentDisputeClosed = r.AlertPaymentDisputeClosed
	if r.alertPaymentDisputeClosed == nil {
		r.alertPaymentDisputeClosed = alertPaymentDisputeClosedNotFound
	}
	r.alertPaymentDisputeCreated = r.AlertPaymentDisputeCreated
	if r.alertPaymentDisputeCreated == nil {
		r.alertPaymentDisputeCreated = alertPaymentDisputeCreatedNotFound
	}
	r.alertPaymentRefunded = r.AlertPaymentRefunded
	if r.alertPaymentRefunded == nil {
		r.alertPaymentRefunded = alertPaymentRefundedNotFound
	}
	r.alertPaymentSucceeded = r.AlertPaymentSucceeded
	if r.alertPaymentSucceeded == nil {
		r.alertPaymentSucceeded = alertPaymentSucceededNotFound
	}
	r.alertTransferCreated = r.AlertTransferCreated
	if r.alertTransferCreated == nil {
		r.alertTransferCreated = alertTransferCreatedNotFound
	}
	r.alertTransferPaid = r.AlertTransferPaid
	if r.alertTransferPaid == nil {
		r.alertTransferPaid = alertTransferPaidNotFound
	}
	r.alertUpdateAudienceMember = r.AlertUpdateAudienceMember
	if r.alertUpdateAudienceMember == nil {
		r.alertUpdateAudienceMember = alertUpdateAudienceMemberNotFound
	}
	r.subscriptionCancelled = r.SubscriptionCancelled
	if r.subscriptionCancelled == nil {
		r.subscriptionCancelled = susbcriptionCancelledNotFound
	}
	r.subscriptionCreated = r.SubscriptionCreated
	if r.subscriptionCreated == nil {
		r.subscriptionCreated = susbcriptionCreatedNotFound
	}
	r.subscriptionPaymentFailed = r.SubscriptionPaymentFailed
	if r.subscriptionPaymentFailed == nil {
		r.subscriptionPaymentFailed = susbcriptionPaymentFailedNotFound
	}
	r.subscriptionPaymentRefunded = r.SubscriptionPaymentRefunded
	if r.subscriptionPaymentRefunded == nil {
		r.subscriptionPaymentRefunded = susbcriptionPaymentRefundedNotFound
	}
	r.subscriptionPaymentSucceeded = r.SubscriptionPaymentSucceeded
	if r.subscriptionPaymentSucceeded == nil {
		r.subscriptionPaymentSucceeded = susbcriptionPaymentSucceededNotFound
	}
	r.subscriptionUpdated = r.SubscriptionUpdated
	if r.subscriptionUpdated == nil {
		r.subscriptionUpdated = susbcriptionUpdatedNotFound
	}
	getEvent := r.ev.EventFromRequest()
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ev, err := getEvent(req)
		if err != nil {
			var httpError httperrors.Error
			switch terr := err.(type) {
			case httperrors.Error:
				httpError = terr
			case signature.VerificationError:
				httpError = httperrors.NewBadRequestError(err.Error())
			default:
				httpError = httperrors.NewBadRequestError(err.Error())
			}
			httpError.WriteTo(rw)
			return
		}
		switch tev := ev.(type) {
		case *alerts.HighRiskTransactionCreated:
			r.alertHighRiskTransactionCreated.ServeHTTP(tev, rw, req)
		case *alerts.HighRiskTransactionUpdated:
			r.alertHighRiskTransactionUpdated.ServeHTTP(tev, rw, req)
		case *alerts.LockerProcessed:
			r.alertLockerProcessed.ServeHTTP(tev, rw, req)
		case *alerts.NewAudienceMember:
			r.alertNewAudienceMember.ServeHTTP(tev, rw, req)
		case *alerts.PaymentDisputeClosed:
			r.alertPaymentDisputeClosed.ServeHTTP(tev, rw, req)
		case *alerts.PaymentDisputeCreated:
			r.alertPaymentDisputeCreated.ServeHTTP(tev, rw, req)
		case *alerts.PaymentRefunded:
			r.alertPaymentRefunded.ServeHTTP(tev, rw, req)
		case *alerts.PaymentSucceeded:
			r.alertPaymentSucceeded.ServeHTTP(tev, rw, req)
		case *alerts.TransferCreated:
			r.alertTransferCreated.ServeHTTP(tev, rw, req)
		case *alerts.TransferPaid:
			r.alertTransferPaid.ServeHTTP(tev, rw, req)
		case *alerts.UpdateAudienceMember:
			r.alertUpdateAudienceMember.ServeHTTP(tev, rw, req)
		case *subscription.Cancelled:
			r.subscriptionCancelled.ServeHTTP(tev, rw, req)
		case *subscription.Created:
			r.subscriptionCreated.ServeHTTP(tev, rw, req)
		case *subscription.PaymentFailed:
			r.subscriptionPaymentFailed.ServeHTTP(tev, rw, req)
		case *subscription.PaymentRefunded:
			r.subscriptionPaymentRefunded.ServeHTTP(tev, rw, req)
		case *subscription.PaymentSucceeded:
			r.subscriptionPaymentSucceeded.ServeHTTP(tev, rw, req)
		case *subscription.Updated:
			r.subscriptionUpdated.ServeHTTP(tev, rw, req)
		}
	})
}

func NewRouter(c Config) Router {
	return Router{Config: c}
}
