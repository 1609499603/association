package internal

import (
	"association/global"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBBASE interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
func (g *_gorm) Config() *gorm.Config {
	configGorm := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	var logMode DBBASE
	logMode = &global.ASS_CONFIG.Mysql

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		configGorm.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		configGorm.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		configGorm.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		configGorm.Logger = _default.LogMode(logger.Info)
	default:
		configGorm.Logger = _default.LogMode(logger.Info)
	}
	return configGorm
}
