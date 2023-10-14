package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

const (
	secret = "sdf$WSDFO04gSDes"
)

func Encrypt(plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	cipherText := aesgcm.Seal(nonce, nonce, plainText, nil)
	return cipherText, nil
}

func Decrypt(cipherText []byte) ([]byte, error) {
	aes, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
