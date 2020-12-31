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
)

type LOGF = logrus.Fields
