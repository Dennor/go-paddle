package signature

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"hash"
)

// RSA verifies RSA signatures with user defined hashing
// If hashing func is not provided, sha1 is used by default.
type RSA struct {
	PublicKey *rsa.PublicKey
	Encoding  *base64.Encoding
	Hash      crypto.Hash
}

func (r RSA) hash() crypto.Hash {
	if r.Hash == 0 {
		return crypto.SHA1
	}
	return r.Hash
}

func (r RSA) hashFunc() hash.Hash {
	switch r.hash() {
	case crypto.SHA1:
		fallthrough
	default:
		return sha1.New()
	}
}

func (r RSA) encoding() *base64.Encoding {
	if r.Encoding == nil {
		return base64.StdEncoding
	}
	return r.Encoding
}

func (r RSA) getSignature(b64sig []byte) ([]byte, error) {
	if len(b64sig) == 0 {
		return nil, errors.New("empty signature")
	}
	signature := make([]byte, r.encoding().DecodedLen(len(b64sig)))
	n, err := r.encoding().Decode(signature, b64sig)
	return signature[:n], err
}

func (r RSA) hashData(data []byte) ([]byte, error) {
	h := r.hashFunc()
	if h == nil {
		return nil, errors.New("unimplemented hash")
	}
	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func (r RSA) Verify(data, b64signature []byte) error {
	sig, err := r.getSignature(b64signature)
	if err != nil {
		return err
	}
	hashed, err := r.hashData(data)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(r.PublicKey, r.hash(), hashed, sig)
}
