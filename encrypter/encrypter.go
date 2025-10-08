package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("SALT")

	if key == "" {
		panic("Не передан параметр в переменные окружения")
	}

	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plain []byte) []byte {
	aesGCM := enc.createGCM()

	nonce := make([]byte, aesGCM.NonceSize())
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}

	return aesGCM.Seal(nonce, nonce, plain, nil)
}

func (enc *Encrypter) Decrypt(encrypted []byte) []byte {
	aesGCM := enc.createGCM()

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encrypted[:nonceSize], encrypted[nonceSize:]

	plain, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return plain
}

func (enc *Encrypter) createGCM() cipher.AEAD {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	return aesGCM
}
