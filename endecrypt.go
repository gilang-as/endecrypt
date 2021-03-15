package endecrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func EndecryptConfig(SecretKey string, SecretIV string) EncryptDecrypt {
	return &endecrypt{
		SecretKey: SecretKey,
		SecretIV:  SecretIV,
	}
}

type endecrypt struct {
	SecretKey string
	SecretIV  string
}

type EncryptDecrypt interface {
	Encrypt(plainText string) string
	Decrypt(cipherText string) string
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func (a *endecrypt) Encrypt(plainText string) string {
	key := fmt.Sprintf("%x", sha256.Sum256([]byte(a.SecretKey)))
	key = key[:32]

	iv := fmt.Sprintf("%x", sha256.Sum256([]byte(a.SecretIV)))
	iv = iv[0:16]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	content := []byte(plainText)
	content = PKCS5Padding(content, 16)
	if len(content)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	encryptText := base64.StdEncoding.EncodeToString(crypted)
	return base64.StdEncoding.EncodeToString([]byte(encryptText))

}

func (a *endecrypt) Decrypt(cipherText string) string {
	key := fmt.Sprintf("%x", sha256.Sum256([]byte(a.SecretKey)))
	key = key[:32]

	iv := fmt.Sprintf("%x", sha256.Sum256([]byte(a.SecretIV)))
	iv = iv[0:16]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	data, err = base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		panic(err)
	}

	ecb := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(data))
	ecb.CryptBlocks(origData, data)
	origData = PKCS5UnPadding(origData)
	return string(origData)
}
