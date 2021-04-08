package global

import (
	"embed"
	"go-api/config"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	CF     config.Config
	VP     *viper.Viper
	DB     *gorm.DB
	LOG    *logrus.Logger
	SER    Server
	Verify *validator.Validate
	FS    embed.FS
)

type LOGF = logrus.Fields

type Server interface {
	ListenAndServe() error
	Restart() error
	Shutdown()
}
