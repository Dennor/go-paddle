package types

import "time"

const (
	TimeFormatYYYYMMDD       = "2006-01-02"
	TimeFormatYYYYMMDDHHmmSS = "2006-01-02 15:04:05"
)

type TimeYYYYMMDD struct {
	time.Time
}

func (t *TimeYYYYMMDD) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.Parse(`"`+TimeFormatYYYYMMDD+`"`, string(data))
	return err
}

func (t *TimeYYYYMMDD) UnmarshalText(data []byte) error {
	var err error
	t.Time, err = time.Parse(TimeFormatYYYYMMDD, string(data))
	return err
}

func (t TimeYYYYMMDD) String() string {
	return t.Format(TimeFormatYYYYMMDD)
}

type TimeYYYYMMDDHHmmSS struct {
	time.Time
}

func (t *TimeYYYYMMDDHHmmSS) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.Parse(`"`+TimeFormatYYYYMMDDHHmmSS+`"`, string(data))
	return err
}

func (t *TimeYYYYMMDDHHmmSS) UnmarshalText(data []byte) error {
	var err error
	t.Time, err = time.Parse(TimeFormatYYYYMMDDHHmmSS, string(data))
	return err
}

func (t TimeYYYYMMDDHHmmSS) String() string {
	return t.Format(TimeFormatYYYYMMDDHHmmSS)
}

type MarketingConsent uint8

const (
	REFUSED MarketingConsent = iota
	GRANTED
)

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
