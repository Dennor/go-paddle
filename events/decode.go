package events

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/dennor/go-paddle/mime"
	"github.com/dennor/urldecode"
)

type schemaDecoder struct {
	urldecode.Decoder
	bytes.Buffer
}

func (s *schemaDecoder) Decode(v interface{}, r io.Reader) error {
	if _, err := io.Copy(&s.Buffer, r); err != nil {
		return err
	}
	return s.Decoder.Decode(v)
}

type schemaDecoderPool struct {
	sync.Pool
}

func newSchemaDecoderPool() schemaDecoderPool {
	return schemaDecoderPool{
		Pool: sync.Pool{
			New: func() interface{} {
				sdec := &schemaDecoder{}
				decoder := urldecode.NewDecoder(&sdec.Buffer)
				sdec.Decoder = *decoder
				return sdec
			},
		},
	}
}

func (s *schemaDecoderPool) Get() *schemaDecoder {
	dec := s.Pool.Get().(*schemaDecoder)
	dec.Buffer.Reset()
	return dec
}

func (s *schemaDecoderPool) Put(d *schemaDecoder) {
	s.Pool.Put(d)
}

var (
	sp = newSchemaDecoderPool()
)

func reqContentType(req *http.Request) string {
	ct := req.Header.Get(mime.ContentTypeHeader)
	if ct == "" {
		return ""
	}
	return strings.Split(ct, ",")[0]
}

func UnmarshalForm(r io.Reader, v interface{}) error {
	dec := sp.Get()
	err := dec.Decode(v, r)
	sp.Put(dec)
	return err
}

func UnmarshalJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
