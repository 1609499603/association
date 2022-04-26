package system

import (
	"association/common/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AssociationList(c *gin.Context) {

	pageNo, _ := strconv.ParseInt(c.Query("page_no"), 10, 32)
	pageSize, _ := strconv.ParseInt(c.Query("page_size"), 10, 32)

	//pageNo - 1     *pagesize =pageno
	associationPage, err := homePageService.AssociationPage(int((pageNo-1)*pageSize), int(pageSize))
	if err != nil {
		response.FailWithMessage("主页社团查询错误", c)
		return
	}

	response.OkWithDetailed(associationPage, "成功", c)

}
