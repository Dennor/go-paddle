package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

const (
	DateFormat     = "2006-01-02"
	DatetimeFormat = "2006-01-02 15:04:05"
)

type Date struct {
	time.Time
	Empty bool
}

func (t *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.Parse(`"`+DateFormat+`"`, string(data))
	return err
}

func (t *Date) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		t.Empty = true
		return nil
	}
	var err error
	t.Time, err = time.Parse(DateFormat, string(data))
	return err
}

func (t Date) String() string {
	if t.Empty {
		return ""
	}
	return t.Format(DateFormat)
}

func (t Date) GoString() string {
	return t.String()
}

type Datetime struct {
	time.Time
}

func (t *Datetime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.Parse(`"`+DatetimeFormat+`"`, string(data))
	return err
}

func (t *Datetime) UnmarshalText(data []byte) error {
	var err error
	t.Time, err = time.Parse(DatetimeFormat, string(data))
	return err
}

func (t Datetime) String() string {
	return t.Format(DatetimeFormat)
}

func (t Datetime) GoString() string {
	return t.String()
}

type MarketingConsent int8

const abc = 1
const (
	UNSET MarketingConsent = iota
	REFUSED
	GRANTED
)

func (m *MarketingConsent) UnmarshalText(data []byte) error {
	switch len(data) {
	case 3:
		if data[0] == '"' && data[2] == '"' {
			data = data[1:2]
		}
		fallthrough
	case 1:
		switch data[0] {
		case '0':
			*m = REFUSED
		case '1':
			*m = GRANTED
		default:
			return errors.New("invalid marketing consent")
		}
	case 2:
		if data[0] != '"' || data[1] != '"' {
			return errors.New("invalid marketing consent")
		}
		fallthrough
	case 0:
		*m = UNSET
	default:
		return errors.New("invalid marketing consent")
	}

	return nil
}

func (m *MarketingConsent) UnmarshalJSON(data []byte) error {
	return m.UnmarshalText(data)
}

func (m MarketingConsent) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m MarketingConsent) String() string {
	switch m {
	case REFUSED:
		return "0"
	case GRANTED:
		return "1"
	}
	return ""
}

type PhpBool bool

var (
	phpTrue  = "1"
	phpFalse = "0"
)

func (b *PhpBool) String() string {
	if *b {
		return phpTrue
	}
	return phpFalse
}

type CurrencyValue struct {
	decimal.Decimal
	fixed int32
}

func (c *CurrencyValue) String() string {
	return c.StringFixed(c.fixed)
}

func (c *CurrencyValue) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *CurrencyValue) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return c.UnmarshalText([]byte(s))
}

func (c *CurrencyValue) UnmarshalText(b []byte) error {
	fixed := bytes.LastIndexByte(b, byte('.'))
	if fixed > 0 {
		c.fixed = int32(len(b[fixed+1:]))
	}
	return c.Decimal.UnmarshalText(b)
}
