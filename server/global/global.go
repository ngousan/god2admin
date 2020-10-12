package global

import (
	"god2admin/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server
	DB     *gorm.DB
	LOGGER *logrus.Logger
	REDIS  *redis.Client
	VIPER  *viper.Viper
)
