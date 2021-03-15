package endecrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAES(t *testing.T) {
	passphrase := GenerateAesKey()
	c := AesConfig(passphrase)
	plainText := "Hello World"
	cipherText := c.Encrypt(plainText)
	decrypt := c.Decrypt(cipherText)
	assert.Equal(t, plainText, decrypt)
}
