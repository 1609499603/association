package system

import (
	"association/common/response"
	"association/global"
	models "association/modules"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"strconv"
	"time"
)

// AssociationList 查询redis中社团信息，如果redis中没有社团信息，则查询数据库后放入redis
func AssociationList(c *gin.Context) {
	pageNo, _ := strconv.ParseInt(c.Query("page_no"), 10, 32)
	name := c.Query("name")
	key := name
	if name == "" {
		key = "all"
	}
	bytes, _ := global.ASS_REDIS.Get(context.Background(), key+":association").Bytes()
	if len(bytes) != 0 {
		m := make(map[string]interface{})
		_ = json.Unmarshal(bytes, &m)

		response.OkWithDetailed(gin.H{
			"data":    m["association"+strconv.FormatInt(pageNo, 10)],
			"pageAll": cast.ToString(m["index"]),
			"number":  cast.ToString(m["number"]),
		}, "成功", c)
	} else {
		assPutRedis(name)
		bytes, _ := global.ASS_REDIS.Get(context.Background(), key+":association").Bytes()
		m := make(map[string]interface{})
		_ = json.Unmarshal(bytes, &m)
		response.OkWithDetailed(gin.H{
			"data":    m["association"+strconv.FormatInt(pageNo, 10)],
			"pageAll": cast.ToString(m["index"]),
			"number":  cast.ToString(m["number"]),
		}, "成功", c)
	}
}

// 若name为空则是全部数据存入缓存，若name不为空则是检索name后存入缓存
func assPutRedis(name string) {

	var key string
	var all []models.Association
	var number int64
	if name != "" {
		key = name + ":association"
		all, number = homePageService.AssociationName(name)
	} else {
		key = "all:association"
		all, number = homePageService.AssociationNumber()
	}
	pageSize := global.ASS_CONFIG.System.PageSize

	var i int64
	var j int64 = 1
	var k int64 = 0
	m := make(map[string]interface{})

	if number > pageSize {
		for i = pageSize; i <= number; i += pageSize {
			m["association"+strconv.FormatInt(j, 10)] = all[k:i]
			j++
			k += pageSize
		}
		if number%pageSize != 0 {
			m["association"+strconv.FormatInt(j, 10)] = all[i-pageSize:]
		}

	} else {
		m["association"+strconv.FormatInt(j, 10)] = all
	}

	m["index"] = j
	m["number"] = number
	marshal, _ := json.Marshal(m)

	global.ASS_REDIS.Set(context.Background(), key, marshal, time.Hour*24*365*100)
}

// MyAssociation 我的社团
func MyAssociation(c *gin.Context) {

}

func SendNotice(c *gin.Context) {
	
}
