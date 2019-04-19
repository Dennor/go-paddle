package alerts

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const LockerProcessedAlertName = "locker_processed"

// LockerProcessed refer to https://paddle.com/docs/reference-using-webhooks/#locker_processed
type LockerProcessed struct {
	AlertName        string                    `json:"alert_name"`
	CheckoutID       string                    `json:"checkout_id"`
	CheckoutRecovery int                       `json:"checkout_recovery,string"`
	Coupon           string                    `json:"coupon"`
	Download         string                    `json:"download"`
	Email            string                    `json:"email"`
	EventTime        *types.TimeYYYYMMDDHHmmSS `json:"event_time,string"`
	Instructions     string                    `json:"instructions"`
	License          string                    `json:"license"`
	MarketingConsent *types.MarketingConsent   `json:"marketing_consent,string"`
	OrderID          int                       `json:"order_id,string"`
	ProductID        int                       `json:"product_id,string"`
	Quantity         int                       `json:"quantity,string"`
	PSignature       string                    `json:"p_signature" php:"-"`
}

func (l *LockerProcessed) Serialize() ([]byte, error) {
	return phpserialize.Marshal(l)
}

func (l *LockerProcessed) Signature() ([]byte, error) {
	return []byte(l.PSignature), nil
}
