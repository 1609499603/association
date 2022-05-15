package models

import (
	"association/global"
	"go.uber.org/zap"
)

func CreateTable() error {
	AllTable := []interface{}{
		Teacher{},
		Student{},
		College{},
		User{},
		Notice{},
		Status{},
	}

	// 需要硬删除
	for _, table := range AllTable {

		if !global.ASS_DB.Migrator().HasTable(table) {
			err := global.ASS_DB.Debug().AutoMigrate(&table)
			if err != nil {
				global.ASS_LOG.Error("AutoMigrate table failed", zap.Error(err))
				return err
			}
		} else {
			global.ASS_LOG.Error("AutoMigrate table existed")
			global.ASS_LOG.Debug("AutoMigrate table existed")
			continue
		}
	}
	zap.L().Warn("All table create success")
	return nil
}
