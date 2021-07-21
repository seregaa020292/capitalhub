package model

import "github.com/google/uuid"

// Session model
type Session struct {
	SessionID   string    `json:"-" redis:"session_id"`
	UserID      uuid.UUID `json:"-" redis:"user_id"`
	Fingerprint string    `json:"fingerprint" redis:"fingerprint"`
}
