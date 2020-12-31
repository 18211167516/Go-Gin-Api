package config

import "time"

type App struct {
	Template      string
	PageSize      int
	JwtSecret     string
	JwtExpiresAt  time.Duration
	SigningMethod string
}
