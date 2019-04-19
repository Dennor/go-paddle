package alerts

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const TransferPaidAlertName = "transfer_paid"

// TransferPaid refer to https://paddle.com/docs/reference-using-webhooks/#transfer_paid
type TransferPaid struct {
	AlertName  string           `json:"alert_name"`
	Amount     *decimal.Decimal `json:"amount,string"`
	Currency   string           `json:"currency"`
	EventTime  *types.Datetime  `json:"event_time,string"`
	PayoutID   int              `json:"payout_id,string"`
	Status     string           `json:"status"`
	PSignature string           `json:"p_signature" php:"-"`
}

func (t *TransferPaid) Serialize() ([]byte, error) {
	return phpserialize.Marshal(t)
}

func (t *TransferPaid) Signature() ([]byte, error) {
	return []byte(t.PSignature), nil
}
