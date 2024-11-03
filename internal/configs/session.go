package configs

import "time"

const (
	MAX_AGE_SESSION time.Duration = 3 * 30 * 24 * time.Hour
	PATH            string        = "/"
	DOMAIN          string        = "localhost"
	SECURE          bool          = false
	HTTP_ONLY       bool          = true
	MAX_AGE_TOKEN   time.Duration = 2 * time.Minute
)
