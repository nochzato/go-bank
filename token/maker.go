package token

import "time"

// Maker interface is used to manage tokens.
type Maker interface {
	// CreateToken creates a new token for a specific username and duration.
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken verifies the token and returns it's payload.
	VerifyToken(token string) (*Payload, error)
}
