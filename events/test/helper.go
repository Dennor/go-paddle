package test

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dennor/go-paddle/events/types"
	"github.com/shopspring/decimal"
)

var Key = func() *rsa.PrivateKey {
	k, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	return k
}()

type Data struct {
	JSON string
	URL  string
	PHP  string
	M    map[string]string
}

func Sign(m map[string]string, bools ...map[string]bool) Data {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	for _, mb := range bools {
		for k := range mb {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	d := Data{M: make(map[string]string, len(keys))}
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("a:%d:{", len(keys)))
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf(`s:%d:"%s";`, len(k), k))
		if v, ok := m[k]; ok {
			builder.WriteString(fmt.Sprintf(`s:%d:"%s";`, len(v), v))
			continue
		}
		for _, mb := range bools {
			if b, ok := mb[k]; ok {
				if b {
					builder.WriteString(`s:1:"1";`)
				} else {
					builder.WriteString(`s:1:"0";`)
				}
				break
			}
		}
	}
	builder.WriteString("}")
	d.PHP = builder.String()
	digest := sha1.Sum([]byte(d.PHP))
	signatureRaw, err := Key.Sign(rand.Reader, digest[:], crypto.SHA1)
	if err != nil {
		panic(err)
	}
	signature := base64.StdEncoding.EncodeToString(signatureRaw)
	builder.Reset()
	builder.WriteString("{")
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf(`"%s":`, k))
		if v, ok := m[k]; ok {
			builder.WriteString(fmt.Sprintf(`"%s",`, v))
			continue
		}
		for _, mb := range bools {
			if b, ok := mb[k]; ok {
				if b {
					builder.WriteString(`"true",`)
				} else {
					builder.WriteString(`"false",`)
				}
				break
			}
		}
	}
	builder.WriteString(`"p_signature":`)
	builder.WriteString(fmt.Sprintf(`"%s"`, signature))
	builder.WriteString("}")
	d.JSON = builder.String()
	builder.Reset()
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("%s=", url.QueryEscape(k)))
		if v, ok := m[k]; ok {
			builder.WriteString(fmt.Sprintf("%s&", url.QueryEscape(v)))
			continue
		}
		for _, mb := range bools {
			if b, ok := mb[k]; ok {
				if b {
					builder.WriteString("1&")
				} else {
					builder.WriteString("0&")
				}
				break
			}
		}
	}
	builder.WriteString("p_signature=")
	builder.WriteString(url.QueryEscape(signature))
	d.URL = builder.String()
	for _, k := range keys {
		if v, ok := m[k]; ok {
			d.M[k] = v
		}
		for _, mb := range bools {
			if b, ok := mb[k]; ok {
				d.M[k] = strconv.FormatBool(b)
				break
			}
		}
	}
	d.M["p_signature"] = signature
	return d
}

func ParseTime(layout, value string) time.Time {
	tt, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return tt
}

func DecimalFromString(s string) *decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	return &d
}

func CurrencyValueFromString(s string) *types.CurrencyValue {
	var cv types.CurrencyValue
	err := cv.UnmarshalText([]byte(s))
	if err != nil {
		panic(err)
	}
	return &cv
}

func IntFromString(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func BoolFromString(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return b
}
