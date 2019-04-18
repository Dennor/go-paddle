package router

import (
	"net/http"

	"github.com/dennor/go-paddle/events/alerts"
	"github.com/dennor/go-paddle/events/subscription"
)

type AlertHighRiskTransactionCreated interface {
	ServeHTTP(*alerts.HighRiskTransactionCreated, http.ResponseWriter, *http.Request)
}

type AlertHighRiskTransactionCreatedFunc func(*alerts.HighRiskTransactionCreated, http.ResponseWriter, *http.Request)

func (f AlertHighRiskTransactionCreatedFunc) ServeHTTP(e *alerts.HighRiskTransactionCreated, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertHighRiskTransactionUpdated interface {
	ServeHTTP(*alerts.HighRiskTransactionUpdated, http.ResponseWriter, *http.Request)
}

type AlertHighRiskTransactionUpdatedFunc func(*alerts.HighRiskTransactionUpdated, http.ResponseWriter, *http.Request)

func (f AlertHighRiskTransactionUpdatedFunc) ServeHTTP(e *alerts.HighRiskTransactionUpdated, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertLockerProcessed interface {
	ServeHTTP(*alerts.LockerProcessed, http.ResponseWriter, *http.Request)
}

type AlertLockerProcessedFunc func(*alerts.LockerProcessed, http.ResponseWriter, *http.Request)

func (f AlertLockerProcessedFunc) ServeHTTP(e *alerts.LockerProcessed, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertNewAudienceMember interface {
	ServeHTTP(*alerts.NewAudienceMember, http.ResponseWriter, *http.Request)
}

type AlertNewAudienceMemberFunc func(*alerts.NewAudienceMember, http.ResponseWriter, *http.Request)

func (f AlertNewAudienceMemberFunc) ServeHTTP(e *alerts.NewAudienceMember, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertPaymentDisputeClosed interface {
	ServeHTTP(*alerts.PaymentDisputeClosed, http.ResponseWriter, *http.Request)
}

type AlertPaymentDisputeClosedFunc func(*alerts.PaymentDisputeClosed, http.ResponseWriter, *http.Request)

func (f AlertPaymentDisputeClosedFunc) ServeHTTP(e *alerts.PaymentDisputeClosed, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertPaymentDisputeCreated interface {
	ServeHTTP(*alerts.PaymentDisputeCreated, http.ResponseWriter, *http.Request)
}

type AlertPaymentDisputeCreatedFunc func(*alerts.PaymentDisputeCreated, http.ResponseWriter, *http.Request)

func (f AlertPaymentDisputeCreatedFunc) ServeHTTP(e *alerts.PaymentDisputeCreated, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertPaymentRefunded interface {
	ServeHTTP(*alerts.PaymentRefunded, http.ResponseWriter, *http.Request)
}

type AlertPaymentRefundedFunc func(*alerts.PaymentRefunded, http.ResponseWriter, *http.Request)

func (f AlertPaymentRefundedFunc) ServeHTTP(e *alerts.PaymentRefunded, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertPaymentSucceeded interface {
	ServeHTTP(*alerts.PaymentSucceeded, http.ResponseWriter, *http.Request)
}

type AlertPaymentSucceededFunc func(*alerts.PaymentSucceeded, http.ResponseWriter, *http.Request)

func (f AlertPaymentSucceededFunc) ServeHTTP(e *alerts.PaymentSucceeded, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertTransferCreated interface {
	ServeHTTP(*alerts.TransferCreated, http.ResponseWriter, *http.Request)
}

type AlertTransferCreatedFunc func(*alerts.TransferCreated, http.ResponseWriter, *http.Request)

func (f AlertTransferCreatedFunc) ServeHTTP(e *alerts.TransferCreated, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertTransferPaid interface {
	ServeHTTP(*alerts.TransferPaid, http.ResponseWriter, *http.Request)
}

type AlertTransferPaidFunc func(*alerts.TransferPaid, http.ResponseWriter, *http.Request)

func (f AlertTransferPaidFunc) ServeHTTP(e *alerts.TransferPaid, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type AlertUpdateAudienceMember interface {
	ServeHTTP(*alerts.UpdateAudienceMember, http.ResponseWriter, *http.Request)
}

type AlertUpdateAudienceMemberFunc func(*alerts.UpdateAudienceMember, http.ResponseWriter, *http.Request)

func (f AlertUpdateAudienceMemberFunc) ServeHTTP(e *alerts.UpdateAudienceMember, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type SubscriptionCancelled interface {
	ServeHTTP(*subscription.Cancelled, http.ResponseWriter, *http.Request)
}

type SubscriptionCancelledFunc func(*subscription.Cancelled, http.ResponseWriter, *http.Request)

func (f SubscriptionCancelledFunc) ServeHTTP(e *subscription.Cancelled, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type SubscriptionCreated interface {
	ServeHTTP(*subscription.Created, http.ResponseWriter, *http.Request)
}

type SubscriptionCreatedFunc func(*subscription.Created, http.ResponseWriter, *http.Request)

func (f SubscriptionCreatedFunc) ServeHTTP(e *subscription.Created, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type SubscriptionPaymentFailed interface {
	ServeHTTP(*subscription.PaymentFailed, http.ResponseWriter, *http.Request)
}

type SubscriptionPaymentFailedFunc func(*subscription.PaymentFailed, http.ResponseWriter, *http.Request)

func (f SubscriptionPaymentFailedFunc) ServeHTTP(e *subscription.PaymentFailed, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type SubscriptionPaymentRefunded interface {
	ServeHTTP(*subscription.PaymentRefunded, http.ResponseWriter, *http.Request)
}

type SubscriptionPaymentRefundedFunc func(*subscription.PaymentRefunded, http.ResponseWriter, *http.Request)

func (f SubscriptionPaymentRefundedFunc) ServeHTTP(e *subscription.PaymentRefunded, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type SubscriptionPaymentSucceeded interface {
	ServeHTTP(*subscription.PaymentSucceeded, http.ResponseWriter, *http.Request)
}

type SubscriptionPaymentSucceededFunc func(*subscription.PaymentSucceeded, http.ResponseWriter, *http.Request)

func (f SubscriptionPaymentSucceededFunc) ServeHTTP(e *subscription.PaymentSucceeded, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

type SubscriptionUpdated interface {
	ServeHTTP(*subscription.Updated, http.ResponseWriter, *http.Request)
}

type SubscriptionUpdatedFunc func(*subscription.Updated, http.ResponseWriter, *http.Request)

func (f SubscriptionUpdatedFunc) ServeHTTP(e *subscription.Updated, rw http.ResponseWriter, req *http.Request) {
	f(e, rw, req)
}

func handlerNotFound(rw http.ResponseWriter, ename string) {
	http.Error(rw, "missing handler for event "+ename, http.StatusNotFound)
}

var (
	alertHighRiskTransactionCreatedNotFound = AlertHighRiskTransactionCreatedFunc(func(e *alerts.HighRiskTransactionCreated, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.HighRiskTransactionCreatedAlertName)
	})

	alertHighRiskTransactionUpdatedNotFound = AlertHighRiskTransactionUpdatedFunc(func(e *alerts.HighRiskTransactionUpdated, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.HighRiskTransactionUpdatedAlertName)
	})

	alertLockerProcessedNotFound = AlertLockerProcessedFunc(func(e *alerts.LockerProcessed, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.LockerProcessedAlertName)
	})

	alertNewAudienceMemberNotFound = AlertNewAudienceMemberFunc(func(e *alerts.NewAudienceMember, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.NewAudienceMemberAlertName)
	})

	alertPaymentDisputeClosedNotFound = AlertPaymentDisputeClosedFunc(func(e *alerts.PaymentDisputeClosed, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.PaymentDisputeClosedAlertName)
	})

	alertPaymentDisputeCreatedNotFound = AlertPaymentDisputeCreatedFunc(func(e *alerts.PaymentDisputeCreated, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.PaymentDisputeCreatedAlertName)
	})

	alertPaymentRefundedNotFound = AlertPaymentRefundedFunc(func(e *alerts.PaymentRefunded, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.PaymentRefundedAlertName)
	})

	alertPaymentSucceededNotFound = AlertPaymentSucceededFunc(func(e *alerts.PaymentSucceeded, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.PaymentSucceededAlertName)
	})

	alertTransferCreatedNotFound = AlertTransferCreatedFunc(func(e *alerts.TransferCreated, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.TransferCreatedAlertName)
	})

	alertTransferPaidNotFound = AlertTransferPaidFunc(func(e *alerts.TransferPaid, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.TransferPaidAlertName)
	})

	alertUpdateAudienceMemberNotFound = AlertUpdateAudienceMemberFunc(func(e *alerts.UpdateAudienceMember, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, alerts.UpdateAudienceMemberAlertName)
	})

	susbcriptionCancelledNotFound = SubscriptionCancelledFunc(func(e *subscription.Cancelled, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, subscription.CancelledAlertName)
	})

	susbcriptionCreatedNotFound = SubscriptionCreatedFunc(func(e *subscription.Created, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, subscription.CreatedAlertName)
	})

	susbcriptionPaymentFailedNotFound = SubscriptionPaymentFailedFunc(func(e *subscription.PaymentFailed, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, subscription.PaymentFailedAlertName)
	})

	susbcriptionPaymentRefundedNotFound = SubscriptionPaymentRefundedFunc(func(e *subscription.PaymentRefunded, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, subscription.PaymentRefundedAlertName)
	})

	susbcriptionPaymentSucceededNotFound = SubscriptionPaymentSucceededFunc(func(e *subscription.PaymentSucceeded, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, subscription.PaymentSucceededAlertName)
	})

	susbcriptionUpdatedNotFound = SubscriptionUpdatedFunc(func(e *subscription.Updated, rw http.ResponseWriter, req *http.Request) {
		handlerNotFound(rw, subscription.UpdatedAlertName)
	})
)
