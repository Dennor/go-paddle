package alerts

import (
	"errors"
	"strconv"
	"strings"

	"github.com/dennor/go-paddle/events/types"
	"github.com/dennor/phpserialize"
)

const NewAudienceMemberAlertName = "new_audience_member"

type AudienceMemberProducts []int64

func (n *AudienceMemberProducts) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*n = nil
		return nil
	}
	l := 1
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			l++
		}
	}
	*n = make([]int64, l)
	i := 0
	for len(data) > 0 {
		if data[0] != ',' {
			if data[0] < '0' || data[0] > '9' {
				return errors.New("not a number")
			}
			(*n)[i] = (*n)[i] * 10
			(*n)[i] = (*n)[i] + int64(data[0]) - '0'
		} else {
			i++
		}
		data = data[1:]
	}
	return nil
}

func (n *AudienceMemberProducts) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return nil
	}
	return n.UnmarshalText(data[1 : len(data)-1])
}

func (n AudienceMemberProducts) String() string {
	sarr := make([]string, len(n))
	for i, pid := range n {
		sarr[i] = strconv.FormatInt(pid, 10)
	}
	return strings.Join(sarr, ",")
}

// NewAudienceMember refer to https://paddle.com/docs/reference-using-webhooks/#new_audience_member
type NewAudienceMember struct {
	AlertName        string                  `json:"alert_name"`
	CreatedAt        *types.Datetime         `json:"created_at,string"`
	Email            string                  `json:"email"`
	EventTime        *types.Datetime         `json:"event_time,string"`
	MarketingConsent *types.MarketingConsent `json:"marketing_consent,string"`
	Products         *AudienceMemberProducts `json:"products,string"`
	Source           string                  `json:"source"`
	Subscribed       int                     `json:"subscribed,string"`
	UserID           int                     `json:"user_id,string"`
	PSignature       string                  `json:"p_signature" php:"-"`
}

func (m *NewAudienceMember) Serialize() ([]byte, error) {
	return phpserialize.Marshal(m)
}

func (m *NewAudienceMember) Signature() ([]byte, error) {
	return []byte(m.PSignature), nil
}
