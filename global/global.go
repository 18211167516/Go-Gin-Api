package global

import (
	"embed"
	"go-api/config"

	cron "github.com/18211167516/robfig-cron/v3"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CF      config.Config
	VP      *viper.Viper
	DB      *gorm.DB
	LOG     *zap.Logger
	CRONLOG *zap.SugaredLogger
	CRON    *cron.Cron
	SER     Server
	Verify  *validator.Validate
	FS      embed.FS
)

type Server interface {
	ListenAndServe() error
	Restart() error
	Shutdown()
}
