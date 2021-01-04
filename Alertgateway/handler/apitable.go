package handler

import (
	"Alertgateway/common/share"
	"Alertgateway/common/utils"
	"Alertgateway/dao"

	"github.com/gin-gonic/gin"
)

func GetAlertPeople(c *gin.Context) {
	res := dao.GetAlertPeopleDB()
	utils.ReturnOK(c, map[string]interface{}{
		"code":  0,
		"msg":   "",
		"count": len(res),
		"data":  res,
	})
}
func GetAlertTemp(c *gin.Context) {
	res := dao.GetAlertTempDB()

	utils.ReturnOK(c, map[string]interface{}{
		"code":  0,
		"msg":   "",
		"count": len(res),
		"data":  res,
	})

}

//参数编辑
func Editparms(c *gin.Context) {
	editdata := c.DefaultPostForm("editdata", "")
	currid := c.DefaultPostForm("id", "")
	if len(editdata) != 0 && len(currid) != 0 {

		dao.SaveVtParms(editdata, currid)
		utils.ReturnOK(c, editdata)
	}
}

func Changevttempstats(c *gin.Context) {
	iscanused := c.DefaultQuery("iscanused", "")
	id := c.DefaultQuery("id", "")
	if len(iscanused) != 0 && len(id) != 0 {
		dao.ChangeTemTimebutton(id, iscanused)

		utils.ReturnOK(c, "success")
	} else {
		share.Logger.Warn("更改模板状态失败!")
		utils.ReturnErr(c, "failed params")
	}

}

func Alertemlaydatedone(c *gin.Context) {

	share.Logger.Warn("报警模板时间更新!")
	laydatev := c.DefaultPostForm("value", "")
	vtdatime := c.DefaultPostForm("vtdatime", "")
	if len(laydatev) != 0 && len(vtdatime) != 0 {
		dao.AlertTemplaydatechange(laydatev, vtdatime)
	}
	utils.ReturnOK(c, "success")
}
