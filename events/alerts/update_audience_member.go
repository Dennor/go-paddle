package alerts

import (
	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const UpdateAudienceMemberAlertName = "update_audience_member"

// UpdateAudienceMember refer to https://paddle.com/docs/reference-using-webhooks/#update_audience_member
type UpdateAudienceMember struct {
	AlertName           string                  `json:"alert_name"`
	EventTime           *types.Datetime         `json:"event_time,string"`
	NewCustomerEmail    string                  `json:"new_customer_email"`
	NewMarketingConsent *types.MarketingConsent `json:"new_marketing_consent,string"`
	OldCustomerEmail    string                  `json:"old_customer_email"`
	OldMarketingConsent *types.MarketingConsent `json:"old_marketing_consent,string"`
	Products            *AudienceMemberProducts `json:"products,string"`
	Source              string                  `json:"source"`
	UpdatedAt           *types.Datetime         `json:"updated_at,string"`
	UserID              int                     `json:"user_id,string"`
	PSignature          string                  `json:"p_signature" php:"-"`
}

func (u *UpdateAudienceMember) Serialize() ([]byte, error) {
	return phpserialize.Marshal(u)
}

func (u *UpdateAudienceMember) Signature() ([]byte, error) {
	return []byte(u.PSignature), nil
}
