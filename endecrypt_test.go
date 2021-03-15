package endecrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndecrypt(t *testing.T) {
	SecretKey := "ExampleSecretKey123!!!"
	SecretIV := "ExampleSecretIV000!!!!."
	a := EndecryptConfig(SecretKey, SecretIV)
	plainText := "Hello World"
	cipherText := a.Encrypt(plainText)
	decrypt := a.Decrypt(cipherText)
	assert.Equal(t, plainText, decrypt)
}
