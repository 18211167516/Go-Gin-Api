package global

import (
	"go-api/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	CF  config.Config
	VP  *viper.Viper
	DB  *gorm.DB
	LOG *logrus.Logger
	SER Server
)

type LOGF = logrus.Fields

type Server interface {
	ListenAndServe() error
	Restart() error
	Shutdown()
}
