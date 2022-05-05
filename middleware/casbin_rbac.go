package middleware

import (
	"association/common/response"
	"association/global"
	"association/utils"
	"association/utils/casbinMatch"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func CasbinHander() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, err := utils.GetClaims(c)
		if err != nil {
			global.ASS_LOG.Error("CasbinHander JWT Parse failed,error:" + err.Error())
		}
		obj := c.Request.URL.Path
		act := c.Request.Method
		sub := strconv.Itoa(waitUse.BaseClaims.Authority)
		a, _ := gormadapter.NewAdapterByDB(global.ASS_DB)
		e, _ := casbin.NewSyncedEnforcer(global.ASS_CONFIG.Casbin.ModelPath, a)
		e.AddFunction("KeyMatch", casbinMatch.KeyMatchFunc)
		ok, err := e.Enforce(obj, act, sub)
		if err != nil {
			global.ASS_LOG.Error("Casbin enforce failed,error:" + err.Error())
		}
		if ok {
			c.Next()
		} else {
			global.ASS_LOG.Info("obj:" + obj + "\n act:" + act + "\n sub:" + act)
			response.FailWithMessage("权限不足", c)
			c.Abort()
			return
		}
	}
}
