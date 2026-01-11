package utils

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"encoding/base64"
// 	"errors"
// 	"io"
// 	"os"
// )

// func EncryptUUID(id string, key []byte) (string, error) {
// 	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
// 		return "", errors.New("invalid key size: must be 16, 24, or 32 bytes")
// 	}

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return "", err
// 	}

// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return "", err
// 	}

// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
// 		return "", err
// 	}

// 	// encrypt the 16-byte UUID
// 	ct := gcm.Seal(nil, nonce, id[:], nil)

// 	// prefix nonce so we can decrypt
// 	out := append(nonce, ct...)
// 	return base64.RawURLEncoding.EncodeToString(out), nil
// }

// func DecryptToUUID(encoded string, key []byte) (string, error) {
// 	var zero string
// 	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
// 		return zero, errors.New("invalid key size: must be 16, 24, or 32 bytes")
// 	}

// 	data, err := base64.RawURLEncoding.DecodeString(encoded)
// 	if err != nil {
// 		return zero, err
// 	}

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return zero, err
// 	}

// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return zero, err
// 	}

// 	if len(data) < gcm.NonceSize() {
// 		return zero, errors.New("ciphertext too short")
// 	}

// 	nonce := data[:gcm.NonceSize()]
// 	ct := data[gcm.NonceSize():]

// 	pt, err := gcm.Open(nil, nonce, ct, nil)
// 	if err != nil {
// 		return zero, err
// 	}

// 	if len(pt) != 16 {
// 		return zero, errors.New("invalid plaintext length for UUID")
// 	}

// 	var id string
// 	copy(id[:], pt[:16])
// 	return id, nil
// }

// // GetEncryptKey returns the encryption key from the ENCRYPT_KEY env var.
// // It accepts either raw bytes or base64 (URL/raw) encoded string.
// func GetEncryptKey() ([]byte, error) {
// 	s := os.Getenv("ENCRYPT_KEY")
// 	if s == "" {
// 		return nil, errors.New("ENCRYPT_KEY not set")
// 	}

// 	// try raw bytes
// 	if ks := []byte(s); len(ks) == 16 || len(ks) == 24 || len(ks) == 32 {
// 		return ks, nil
// 	}

// 	// try URL-safe base64
// 	if b, err := base64.RawURLEncoding.DecodeString(s); err == nil {
// 		if len(b) == 16 || len(b) == 24 || len(b) == 32 {
// 			return b, nil
// 		}
// 	}

// 	// try Std encoding
// 	if b, err := base64.StdEncoding.DecodeString(s); err == nil {
// 		if len(b) == 16 || len(b) == 24 || len(b) == 32 {
// 			return b, nil
// 		}
// 	}

// 	return nil, errors.New("ENCRYPT_KEY has invalid length or format")
// }
