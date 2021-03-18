package config

type App struct {
	Template      string
	PageSize      int
	JwtSecret     string
	JwtExpiresAt  string
	SigningMethod string
}
