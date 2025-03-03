package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// Encipher AES-128/GCM + BASE64算 加密
func Encipher(plaintext string, key string) (string, error) {

	decodeString, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err1 := aes.NewCipher(decodeString)
	if err1 != nil {
		return "", err1
	}

	gcm, err2 := cipher.NewGCM(block)
	if err2 != nil {
		return "", err2
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err3 := io.ReadFull(rand.Reader, nonce); err3 != nil {
		return "", err3
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	result := fmt.Sprintf("{\"data\": \"%s\"}", base64.StdEncoding.EncodeToString(ciphertext))

	return result, nil
}
