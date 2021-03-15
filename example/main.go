package main

import (
	"fmt"
	"github.com/gilang-as/endecrypt"
)

func main() {
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
