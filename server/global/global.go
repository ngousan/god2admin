package global

import (
	"god2admin/config"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
	LOGGER *logrus.Logger
	VIPER  *viper.Viper
	// GVA_REDIS  *redis.Client
	// GVA_LOG    *oplogging.Logger
)
