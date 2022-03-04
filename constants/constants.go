package constants

import (
	"time"
)

const (
	AuthenticationTimeout = time.Hour * 24 * 2
	IsAuthenticatedKey    = "is_authenticated"
	UserIDKey             = "user_id"
)
