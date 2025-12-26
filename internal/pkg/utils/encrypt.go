package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/google/uuid"
)

func EncryptUUID(id uuid.UUID, key []byte) (string, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("invalid key size: must be 16, 24, or 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// encrypt the 16-byte UUID
	ct := gcm.Seal(nil, nonce, id[:], nil)

	// prefix nonce so we can decrypt
	out := append(nonce, ct...)
	return base64.RawURLEncoding.EncodeToString(out), nil
}

func DecryptToUUID(encoded string, key []byte) (uuid.UUID, error) {
	var zero uuid.UUID
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return zero, errors.New("invalid key size: must be 16, 24, or 32 bytes")
	}

	data, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return zero, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return zero, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return zero, err
	}

	if len(data) < gcm.NonceSize() {
		return zero, errors.New("ciphertext too short")
	}

	nonce := data[:gcm.NonceSize()]
	ct := data[gcm.NonceSize():]

	pt, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return zero, err
	}

	if len(pt) != 16 {
		return zero, errors.New("invalid plaintext length for UUID")
	}

	var id uuid.UUID
	copy(id[:], pt[:16])
	return id, nil
}
