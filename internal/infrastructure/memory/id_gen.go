package memory

import (
	"crypto/rand"
	"encoding/base64"
)

type RandIDGen struct{}

// Generate returns 8-chars URL-safe code.
func (RandIDGen) Generate() (string, error) {
	var b [6]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b[:]), nil
}
