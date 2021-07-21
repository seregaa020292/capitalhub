package csrf

import (
	"crypto/sha256"
	"encoding/base64"
	"io"

	"github.com/seregaa020292/capitalhub/pkg/logger"
)

const (
	CSRFHeader = "X-CSRF-Token"
)

// Create CSRF token
func MakeToken(sid string, csrfSalt string, logger logger.Logger) string {
	hash := sha256.New()
	if _, err := io.WriteString(hash, csrfSalt+sid); err != nil {
		logger.Errorf("Make CSRF Token", err)
	}
	return base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
}

// Validate CSRF token
func ValidateToken(token string, sid string, csrfSalt string, logger logger.Logger) bool {
	trueToken := MakeToken(sid, csrfSalt, logger)
	return token == trueToken
}
