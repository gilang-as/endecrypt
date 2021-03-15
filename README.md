### Instalation
```bash
go get github.com/gilang-as/endecrypt
```
### Example
```go
package main

import (
	"fmt"
	"github.com/gilang-as/endecrypt"
)

func main() {
	endecryptExample()
	aesExample()
}

func endecryptExample()  {
	fmt.Println("==========[ Endecrypt Example ]==========")
	SecretKey := "ExampleSecretKey123!!!"
	SecretIV := "ExampleSecretIV000!!!!."
	a := endecrypt.EndecryptConfig(SecretKey, SecretIV)

	plainText := "Hello World"
	fmt.Println("Before Encrypt : ", plainText)

	cipherText := a.Encrypt(plainText)
	fmt.Println("After Encrypt : ", cipherText)

	decrypt := a.Decrypt(cipherText)
	fmt.Println("After Decrypt : ", decrypt)
}

func aesExample()  {
	fmt.Println("==========[ AES Example ]==========")
	passphrase := endecrypt.GenerateAesKey()
	fmt.Println("Passphrase : ", passphrase)
	c := endecrypt.AesConfig(passphrase)
	plainText := "Hello World"
	fmt.Println("Before Encrypt : ", plainText)
	cipherText := c.Encrypt(plainText)
	fmt.Println("After Encrypt : ", cipherText)
	decrypt := c.Decrypt(cipherText)
	fmt.Println("After Decrypt : ", decrypt)
}
```