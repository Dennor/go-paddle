package alerts

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const TransferCreatedAlertName = "transfer_created"

// TransferCreated refer to https://paddle.com/docs/reference-using-webhooks/#transfer_created
type TransferCreated struct {
	AlertName  string                    `json:"alert_name"`
	Amount     *decimal.Decimal          `json:"amount,string"`
	Currency   string                    `json:"currency"`
	EventTime  *types.TimeYYYYMMDDHHmmSS `json:"event_time,string"`
	PayoutID   int                       `json:"payout_id,string"`
	Status     string                    `json:"status"`
	PSignature string                    `json:"p_signature" php:"-"`
}

func (t *TransferCreated) Serialize() ([]byte, error) {
	return phpserialize.Marshal(t)
}

func (t *TransferCreated) Signature() ([]byte, error) {
	return []byte(t.PSignature), nil
}
