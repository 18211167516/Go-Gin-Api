package config

import "time"

type Server struct {
	HttpAddress  string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
