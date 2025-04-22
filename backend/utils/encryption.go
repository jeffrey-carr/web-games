package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// GenerateGCMKey generates a new GCM key
func GenerateGCMKey() (cipher.AEAD, error) {
	key := make([]byte, 32)

	_, err := rand.Reader.Read(key)
	if err != nil {
		return nil, err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(blockCipher)
}
