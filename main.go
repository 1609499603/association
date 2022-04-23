package main

import (
	"association/global"
	"association/initialize"
	"association/initialize/mysql"
	"association/router"
	"association/utils"
	"association/utils/logger"
	"association/utils/snowflake"
)

func main() {

	//初始化viper
	global.ASS_VIPER = utils.Viper()
	//初始化日志
	global.ASS_LOG = logger.Zap()

	//初始化mysql数据库 程序结束前关闭sql连接
	global.ASS_DB = mysql.GormMysql()
	db, _ := global.ASS_DB.DB()
	defer db.Close()

	//初始化id生成时间
	SnowFlakeErr := snowflake.Init("2022-01-02 15:03:04", 1)
	if SnowFlakeErr != nil {
		return
	}
	//初始化redis
	initialize.Redis()

	//初始化路由
	router.Run()

}
