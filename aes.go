package endecrypt

import (
	AES "crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func AesConfig(Passphrase string) Aes {
	return &aes{
		Passphrase: Passphrase,
	}
}

type aes struct {
	Passphrase string
}

type Aes interface {
	Encrypt(plainText string) string
	Decrypt(cipherText string) string
}

func GenerateAesKey() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	return hex.EncodeToString(bytes)
}

func (a *aes) Encrypt(stringToEncrypt string) (encryptedString string) {
	key, _ := hex.DecodeString(a.Passphrase)
	plaintext := []byte(stringToEncrypt)
	block, err := AES.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func (a *aes) Decrypt(encryptedString string) (decryptedString string) {
	key, _ := hex.DecodeString(a.Passphrase)
	enc, _ := hex.DecodeString(encryptedString)
	block, err := AES.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%s", plaintext)
}