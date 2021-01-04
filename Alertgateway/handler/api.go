package handler

import (
	"Alertgateway/common/share"
	"Alertgateway/common/utils"
	"Alertgateway/dao"

	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	_timestamp_format = "2006-01-02 15:04:05"
)

func TelAlertmsg(c *gin.Context) {

	var tp dao.TtsParam
	tp = dao.TtsParam{}
	if err := c.BindJSON(&tp); err != nil {
		fmt.Println(err)
		utils.ReturnErr(c, "failed to call alert telnums!")

	} else {
		//判断是个模板,需要查看维护时间

		if bres := utils.TlsLimitTime(&tp); bres {
			am, cs, cn := dao.NewAlertObj(tp.AlertPf)
			if cn != nil {
				err = utils.AliyunCall(am.AccessKeyId, am.AccessSecret, am.TtsCode, cs.Csnumber, cn, tp)
				if err != nil {
					utils.ReturnErr(c, "failed to call alert telnums!")
				}
				utils.ReturnOK(c, "success!")
			} else {
				utils.ReturnErr(c, "voice temp is null!")
			}
		} else {
			utils.ReturnErr(c, "voice temp  name is err or limit time is curr!")
		}

	}
}

//添加报警人信息
func Addalertpeople(c *gin.Context) {
	cnumber := c.DefaultPostForm("cnumber", "")
	belongname := c.DefaultPostForm("belongname", "")
	belongdep := c.DefaultPostForm("belongdep", "")
	if len(cnumber) != 0 && len(belongname) != 0 && len(belongdep) != 0 {
		ctn := dao.CalledTelNum{
			Cnumber:    cnumber,
			BelongName: belongname,
			BelongDep:  belongdep,
			Oper:       "",
			IsCall:     true,
		}
		dao.AddalertpeopleDB(ctn)
		utils.ReturnOK(c, map[string]interface{}{
			"code": 200,
			"msg":  "success insert",
		})
	} else {
		utils.ReturnErr(c, "params is err")
	}

}

//更改是否维护的状态
func Changecnstats(c *gin.Context) {
	iscall := c.DefaultQuery("iscall", "")
	id := c.DefaultQuery("id", "")
	if len(iscall) != 0 && len(id) != 0 {
		dao.ChangeCnstat(iscall, id)
		utils.ReturnOK(c, "success!")
	} else {
		utils.ReturnErr(c, "params iscall is err")
	}
}

//删除
func Delctn(c *gin.Context) {
	id := c.DefaultPostForm("ID", "")
	if len(id) != 0 {

		err := dao.Delctnwithid(id)
		if err != nil {
			utils.ReturnErr(c, "del is failed!")
		} else {
			utils.ReturnOK(c, "success!")
		}

	} else {
		utils.ReturnErr(c, "params iscall is err")
	}

}

//更新电话号码

func Updatetel(c *gin.Context) {
	user := c.DefaultPostForm("L_user", "")
	tel := c.DefaultPostForm("L_tel", "")

	if len(user) > 0 && len(tel) > 0 {
		err := dao.Updatetel(user, tel)
		if err != nil {

			share.Logger.Warn("更新电话号码错误: ", err)
			utils.ReturnErr(c, "更新电话号码错误!")
		} else {
			utils.ReturnOK(c, "success!")
		}
	} else {
		utils.ReturnErr(c, "参数错误")
	}

}

//模板添加报警人
func Temaddalpeo(c *gin.Context) {
	res := c.DefaultPostForm("data", "")
	tempid := c.DefaultPostForm("tempid", "")
	if len(res) != 0 && len(tempid) != 0 {

		err := dao.SaveTemaddAlerPeo(res, tempid)
		if err != nil {
			utils.ReturnErr(c, "err,params is err!")
		}
		utils.ReturnOK(c, "success")

	} else {
		utils.ReturnErr(c, "err,params is err!")
	}
}

func Addalerttemp(c *gin.Context) {

	t_name := c.DefaultPostForm("T_name", "")
	t_content := c.DefaultPostForm("T_content", "")
	if len(t_name) != 0 {
		dao.TempAddCon(t_name, t_content)
		utils.ReturnOK(c, "success")
	} else {
		utils.ReturnErr(c, "temp name is null ,params is err!")
	}
}

//删除模板
func Deltempmsg(c *gin.Context) {
	// v, ok := c.Get("caccount")
	// if ok {
	// 	fmt.Println("当前的账号", v)
	// } else {
	// 	fmt.Println("找不到 caccount...")
	// }
	id := c.DefaultPostForm("ID", "")
	dao.DelTempMsg(id)

	share.Logger.Warn("删除模板!id:")
	utils.ReturnOK(c, "success")

}
