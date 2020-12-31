package config

import "time"

type Mysql struct {
	LogMode       bool
	Config        string
	MysqlUser     string
	MysqlPassword string
	MysqlHost     string
	MysqlName     string
	MysqlPrefix   string
	MaxLifetime   time.Duration
}
