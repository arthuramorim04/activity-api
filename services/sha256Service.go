package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Encode(s string) string {
	passSHA256 := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", passSHA256)
}
