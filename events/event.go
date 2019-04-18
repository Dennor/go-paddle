package events

import (
	"github.com/dennor/go-paddle/signature"
)

type RSAVerifier signature.RSA

func (r RSAVerifier) Verify(e Event) error {
	data, err := e.Serialize()
	if err != nil {
		return signature.NewVerificationError(err)
	}
	sig, err := e.Signature()
	if err != nil {
		return signature.NewVerificationError(err)
	}
	if err := (signature.RSA)(r).Verify(data, sig); err != nil {
		return signature.NewVerificationError(err)
	}
	return nil
}

type Event interface {
	Serialize() ([]byte, error)
	Signature() ([]byte, error)
}

type Verifier interface {
	Verify(Event) error
}
