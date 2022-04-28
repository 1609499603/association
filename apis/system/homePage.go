package system

import (
	"association/common/response"
	"association/global"
	models "association/modules"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func AssociationFirst(c *gin.Context) {
	pageSize := global.ASS_CONFIG.System.PageSize
	all, number := homePageService.AssociationNumber()
	var i int64
	var j int64 = 1
	var k int64 = 0
	m := make(map[string][]models.Association)
	key := "all:association"

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

	marshal, _ := json.Marshal(m)

	global.ASS_REDIS.Set(context.Background(), key, marshal, time.Hour*24*365*100)
	response.OkWithDetailed(gin.H{
		"pageAll": j,
		"number":  number,
	}, "ok", c)
}

func AssociationList(c *gin.Context) {
	pageNo, _ := strconv.ParseInt(c.Query("page_no"), 10, 32)

	bytes, err := global.ASS_REDIS.Get(context.Background(), "all:association").Bytes()
	if err != nil {
		response.Fail(c)
	}
	m := make(map[string][]models.Association)
	_ = json.Unmarshal(bytes, &m)

	response.OkWithDetailed(m["association"+strconv.FormatInt(pageNo, 10)], "成功", c)
}

// AssociationNameFirst 根据社团名称检索后加入缓存
func AssociationNameFirst(c *gin.Context) {
	name := c.Query("name")
	pageSize := global.ASS_CONFIG.System.PageSize
	key := name + ":association"
	m := make(map[string][]models.Association)
	all, number := homePageService.AssociationName(name)
	var i int64
	var j int64 = 1
	var k int64 = 0

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

	marshal, _ := json.Marshal(m)
	global.ASS_REDIS.Set(context.Background(), key, marshal, time.Hour*24*365*100)

	response.OkWithDetailed(gin.H{
		"pageAll": j,
		"number":  number,
	}, "ok", c)

}

func AssociationName(c *gin.Context) {
	pageNo, _ := strconv.ParseInt(c.Query("page_no"), 10, 32)
	name := c.Query("name")

	bytes, err := global.ASS_REDIS.Get(context.Background(), name+":association").Bytes()
	if err != nil {
		response.Fail(c)
	}
	m := make(map[string][]models.Association)
	_ = json.Unmarshal(bytes, &m)

	response.OkWithDetailed(m["association"+strconv.FormatInt(pageNo, 10)], "成功", c)
}
