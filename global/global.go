package global

import (
	"embed"
	"go-api/config"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CF     config.Config
	VP     *viper.Viper
	DB     *gorm.DB
	LOG    *zap.Logger
	SER    Server
	Verify *validator.Validate
	FS     embed.FS
)

type Server interface {
	ListenAndServe() error
	Restart() error
	Shutdown()
}
