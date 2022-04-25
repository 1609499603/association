package global

import (
	"association/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ASS_DB     *gorm.DB
	ASS_LOG    *zap.Logger
	ASS_CONFIG config.Server
	ASS_VIPER  *viper.Viper
	ASS_REDIS  *redis.Client
)
