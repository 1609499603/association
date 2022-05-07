package system

import (
	"association/common/response"
	"association/global"
	models "association/modules"
	"association/modules/dto"
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

// MyAssociation 我的社团
func MyAssociation(c *gin.Context) {
	getId, _ := c.Get("id")
	id := cast.ToUint(getId)
	ass, err := homePageService.SelectAssociationById(id)
	if err != nil {
		response.FailWithMessage("暂无法访问", c)
		return
	}
	if ass.ID == 0 {
		response.FailWithMessage("请先加入社团", c)
		return
	}
	response.OkWithData(gin.H{
		"data": ass,
	}, c)
}

// MyAssociationUser 我的社团用户
func MyAssociationUser(c *gin.Context) {
	pageNo, _ := strconv.ParseInt(c.Query("page_no"), 10, 32)
	associationId := c.Query("association_id")
	key := associationId
	bytes, _ := global.ASS_REDIS.Get(context.Background(), key).Bytes()

	if len(bytes) != 0 {
		m := make(map[string]interface{})
		_ = json.Unmarshal(bytes, &m)
		response.OkWithData(gin.H{
			"data":    m["user"+strconv.FormatInt(pageNo, 10)],
			"pageAll": cast.ToString(m["index"]),
			"number":  cast.ToString(m["number"]),
		}, c)
	} else {
		assUserPutRedis(associationId)
		bytes, _ := global.ASS_REDIS.Get(context.Background(), key).Bytes()
		m := make(map[string]interface{})
		_ = json.Unmarshal(bytes, &m)
		response.OkWithData(gin.H{
			"data":    m["user"+strconv.FormatInt(pageNo, 10)],
			"pageAll": cast.ToString(m["index"]),
			"number":  cast.ToString(m["number"]),
		}, c)
	}

}

// Exit 退出社团
func Exit(c *gin.Context) {
	getId, _ := c.Get("id")
	id := cast.ToUint(getId)
	user, err := homePageService.SelectUserById(id)
	if err != nil {
		response.FailWithMessage("获取社团数据失败，请重试", c)
		return
	}
	if user.AssociationId == 0 {
		response.FailWithMessage("未加入社团", c)
		return
	}
	err = homePageService.UpdateAssociationIdById(id, 0)
	if err != nil {
		response.FailWithMessage("退出社团失败，请重试", c)
		return
	}
	global.ASS_REDIS.Del(context.Background(), strconv.Itoa(int(user.AssociationId)))
	response.OkWithMessage("退出社团成功", c)
}

// Join 申请加入社团
func Join(c *gin.Context) {
	homeDto := new(dto.AssociationId)
	err := c.ShouldBindJSON(homeDto)
	if err != nil {
		response.FailWithMessage("json 格式错误", c)
		return
	}
	getAssId := homeDto.Id
	key := "examine" + getAssId
	getId, _ := c.Get("id")
	id := cast.ToString(getId)
	parseUint, _ := strconv.ParseUint(id, 10, 32)
	user, _ := homePageService.SelectUserById(uint(parseUint))
	if user.AssociationId != 0 {
		response.FailWithMessage("您已加入社团,请先退出", c)
		return
	}
	result, _ := global.ASS_REDIS.HGet(context.Background(), key, id).Result()
	if result != "" {
		response.OkWithMessage("请不要重复申请", c)
		return
	}
	global.ASS_REDIS.HSet(context.Background(), key, id, "申请加入")
	response.OkWithMessage("提交申请成功，等待审批", c)

}

// Examine 社团内审批页面显示
func Examine(c *gin.Context) {
	homeDto := new(dto.AssociationId)
	err := c.ShouldBindJSON(homeDto)
	if err != nil {
		response.FailWithMessage("json 格式错误", c)
		return
	}
	getAssId := homeDto.Id
	key := "examine" + getAssId
	i, err := global.ASS_REDIS.HGetAll(context.Background(), key).Result()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(i, c)
}

// AgreeAssociation 同意加入社团
func AgreeAssociation(c *gin.Context) {
	homeDto := new(dto.UserAssociation)
	err := c.ShouldBindJSON(homeDto)
	if err != nil {
		response.FailWithMessage("json 格式错误", c)
		return
	}
	getAssId := homeDto.AssociationId
	getUserId := homeDto.Id
	key := "examine" + getAssId
	userId, _ := strconv.ParseUint(getUserId, 10, 32)
	assId, _ := strconv.ParseUint(getAssId, 10, 32)
	user, _ := homePageService.SelectUserById(uint(userId))
	if user.AssociationId != 0 {
		response.FailWithMessage("该用户已加入社团", c)
		return
	}

	updateErr := homePageService.UpdateAssociationIdById(uint(userId), uint(assId))
	if updateErr != nil {
		response.FailWithMessage("加入社团失败，系统错误", c)
		return
	}
	redisErr := global.ASS_REDIS.HDel(context.Background(), key, getUserId).Err()
	if redisErr != nil {
		response.FailWithMessage("系统错误", c)
		return
	}
	//删除原有社团成员缓存
	global.ASS_REDIS.Del(context.Background(), strconv.Itoa(int(user.AssociationId)))

	response.OkWithMessage("加入社团成功", c)
}

// Notice 我的通知列表(社团）
func Notice(c *gin.Context) {
	pageNo, _ := strconv.ParseInt(c.Query("page_no"), 10, 32)
	getId, _ := c.Get("id")
	id := cast.ToUint(getId)
	user, _ := homePageService.SelectUserById(id)
	AssociationId := strconv.FormatUint(uint64(user.AssociationId), 10)
	key := "notice" + AssociationId
	bytes, _ := global.ASS_REDIS.Get(context.Background(), key).Bytes()
	if len(bytes) != 0 {
		m := make(map[string]interface{})
		_ = json.Unmarshal(bytes, &m)
		response.OkWithData(gin.H{
			"data":    m["notice"+strconv.FormatInt(pageNo, 10)],
			"pageAll": cast.ToString(m["index"]),
			"number":  cast.ToString(m["number"]),
		}, c)
	} else {
		noticeRedis(AssociationId)
		bytes, _ := global.ASS_REDIS.Get(context.Background(), key).Bytes()
		m := make(map[string]interface{})
		_ = json.Unmarshal(bytes, &m)
		response.OkWithData(gin.H{
			"data":    m["notice"+strconv.FormatInt(pageNo, 10)],
			"pageAll": cast.ToString(m["index"]),
			"number":  cast.ToString(m["number"]),
		}, c)
	}
}

// SendNotice 发送通知
func SendNotice(c *gin.Context) {
	homeDto := new(dto.SendNoticeDto)
	err := c.ShouldBindJSON(homeDto)
	if err != nil {
		response.FailWithMessage("json 格式错误", c)
		return
	}
	id, _ := strconv.ParseInt(homeDto.AssociationId, 10, 64)
	notice := models.Notice{
		Title:    homeDto.Title,
		Content:  homeDto.Content,
		IsSystem: int(id),
	}
	sendErr := homePageService.CreateNotice(notice)
	if sendErr != nil {
		response.FailWithMessage("发送失败", c)
		return
	}
	key := "notice" + homeDto.AssociationId
	global.ASS_REDIS.Del(context.Background(), key)
	response.OkWithMessage("发送成功", c)

}

//通知列表存入redis
func noticeRedis(associationId string) {
	key := "notice" + associationId
	var all []models.Notice
	var number int64

	all, number = homePageService.SelectNoticeByAssociationId(associationId)

	pageSize := global.ASS_CONFIG.System.PageSize
	var i int64
	var j int64 = 1
	var k int64 = 0
	m := make(map[string]interface{})

	if number > pageSize {
		for i = pageSize; i <= number; i += pageSize {
			m["notice"+strconv.FormatInt(j, 10)] = all[k:i]
			j++
			k += pageSize
		}
		if number%pageSize != 0 {
			m["notice"+strconv.FormatInt(j, 10)] = all[i-pageSize:]
		}

	} else {
		m["notice"+strconv.FormatInt(j, 10)] = all
	}

	m["index"] = j
	m["number"] = number
	marshal, _ := json.Marshal(m)

	global.ASS_REDIS.Set(context.Background(), key, marshal, time.Hour*24*365*100)

}

// AssUserPutRedis 社团所有用户存入redis
func assUserPutRedis(associationId string) {

	key := associationId
	var all []dto.AssociationContent
	var number int64

	all, number = homePageService.SelectUserByAssociationId(associationId)

	pageSize := global.ASS_CONFIG.System.PageSize

	var i int64
	var j int64 = 1
	var k int64 = 0
	m := make(map[string]interface{})

	if number > pageSize {
		for i = pageSize; i <= number; i += pageSize {
			m["user"+strconv.FormatInt(j, 10)] = all[k:i]
			j++
			k += pageSize
		}
		if number%pageSize != 0 {
			m["user"+strconv.FormatInt(j, 10)] = all[i-pageSize:]
		}

	} else {
		m["user"+strconv.FormatInt(j, 10)] = all
	}

	m["index"] = j
	m["number"] = number
	marshal, _ := json.Marshal(m)

	global.ASS_REDIS.Set(context.Background(), key, marshal, time.Hour*24*365*100)
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
