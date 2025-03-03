package util

import (
	"crypto/sha256"
	"fmt"
)

// Sign SHA256 签名
func Sign(plaintext string) (string, error) {
	h := sha256.New()
	h.Write([]byte(plaintext))
	sum := h.Sum(nil)

	return fmt.Sprintf("%x", sum), nil
}
