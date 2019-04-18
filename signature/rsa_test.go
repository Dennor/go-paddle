package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type cryptoReader struct{}

func (c cryptoReader) Read(b []byte) (int, error) {
	return rand.Read(b)
}

func TestRSASHA1(t *testing.T) {
	reader := cryptoReader{}
	assert := assert.New(t)
	require := require.New(t)
	privateKey, err := rsa.GenerateKey(reader, 1024)
	require.NoError(err)
	require.NotNil(privateKey)
	message := []byte("test message")
	hashed := sha1.Sum(message)
	signature, err := rsa.SignPKCS1v15(reader, privateKey, crypto.SHA1, hashed[:])
	require.NoError(err)
	require.NotNil(signature)
	b64sig := make([]byte, base64.StdEncoding.EncodedLen(len(signature)))
	base64.StdEncoding.Encode(b64sig, signature)
	rsaDefault := RSA{
		PublicKey: &privateKey.PublicKey,
	}
	assert.NoError(rsaDefault.Verify(message, b64sig))
	rsasha1 := RSA{
		PublicKey: &privateKey.PublicKey,
		Hash:      crypto.SHA1,
	}
	assert.NoError(rsasha1.Verify(message, b64sig))
}
