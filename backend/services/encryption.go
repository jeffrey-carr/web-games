package services

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// Encryption is the service that encrypts/decrypts data
type Encryption interface {
	Encrypt(data any) (string, error)
	Decrypt(ciphertext string, obj any) error
}

type encryption struct {
	gcm cipher.AEAD
}

// NewEncryption creates a new Encryption service
func NewEncryption(key cipher.AEAD) Encryption {
	return &encryption{
		gcm: key,
	}
}

// Encrypt encrypts the plaintext to a hex string
func (e *encryption) Encrypt(data any) (string, error) {
	plaintext, err := json.Marshal(data)
	nonce, err := e.generateNonce()
	if err != nil {
		return "", err
	}

	ciphertext := e.gcm.Seal(nonce, nonce, plaintext, nil)
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the ciphertext to its string state
func (e *encryption) Decrypt(ciphertext string, obj any) error {
	bytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return errors.Wrap(err, "error decoding cipher hex")
	}

	nonce := bytes[:e.gcm.NonceSize()]
	encryptedData := bytes[e.gcm.NonceSize():]

	data, err := e.gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}

	return nil
}

func (e *encryption) generateNonce() ([]byte, error) {
	nonce := make([]byte, e.gcm.NonceSize())
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return nonce, nil
}
