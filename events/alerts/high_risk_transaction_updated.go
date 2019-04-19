package alerts

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
	"github.com/shopspring/decimal"
)

const HighRiskTransactionUpdatedAlertName = "high_risk_transaction_updated"

// HighRiskTranasctionUpdated refer to https://paddle.com/docs/reference-using-webhooks/#high_risk_transaction_updated
type HighRiskTransactionUpdated struct {
	AlertName            string                  `json:"alert_name"`
	CaseID               int                     `json:"case_id,string"`
	CheckoutID           string                  `json:"checkout_id"`
	CreatedAt            *types.Datetime         `json:"created_at,string"`
	CustomerEmailAddress string                  `json:"customer_email_address"`
	CustomerUserID       int                     `json:"customer_user_id,string"`
	EventTime            *types.Datetime         `json:"event_time,string"`
	MarketingConsent     *types.MarketingConsent `json:"marketing_consent,string"`
	Passthrough          string                  `json:"passthrough"`
	ProductID            int                     `json:"product_id,string"`
	RiskScore            *decimal.Decimal        `json:"risk_score,string"`
	Status               string                  `json:"status"`
	PSignature           string                  `json:"p_signature" php:"-"`
}

func (h *HighRiskTransactionUpdated) Serialize() ([]byte, error) {
	return phpserialize.Marshal(h)
}

func (h *HighRiskTransactionUpdated) Signature() ([]byte, error) {
	return []byte(h.PSignature), nil
}
