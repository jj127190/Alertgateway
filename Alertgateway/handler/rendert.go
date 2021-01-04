package handler

import (
	"Alertgateway/common/share"
	"Alertgateway/dao"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

//index
func IndexHandler(c *gin.Context) {
	username := c.Query("username")
	c.HTML(http.StatusOK, "index.html", gin.H{"username": username})
}

//首页
func WelcomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{"index": "index"})
}

//登录
func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"loggin": "登录"})
}

//404
func F404Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", gin.H{"index": "index"})
}

///////////////////// Alertmaint 报警维护相关页面

func Alertmaint(c *gin.Context) {
	c.HTML(http.StatusOK, "Alertmaint.html", nil)

}

///////////////////// Alertmaint 报警通知相关页面

func Alertopeople(c *gin.Context) {
	c.HTML(http.StatusOK, "Alertopeople.html", nil)

}

func Altemplaterender(c *gin.Context) {
	datetimestr := make(map[string]string, 10)
	res := dao.GetVoicetem()
	for _, v := range *res {
		if v.IsCanUsed {
			datetimestr[v.LimitTimeobj] = v.LimitTime
		} else {
			datetimestr[v.LimitTimeobj] = ""
		}

	}
	c.HTML(http.StatusOK, "alertempoper.html", gin.H{"datetimestr": datetimestr})

}

//添加报警人
func AddAlertPeoRender(c *gin.Context) {
	c.HTML(http.StatusOK, "alertPeopleAdd.html", nil)
	// c.HTML(http.StatusOK, "vpn_user_add_ldap.html", gin.H{"vpnuseradd": "vpnuseradd"})

}

//添加报警模板
func AddAlertTempRender(c *gin.Context) {
	c.HTML(http.StatusOK, "alertTempAdd.html", nil)
	// c.HTML(http.StatusOK, "vpn_user_add_ldap.html", gin.H{"vpnuseradd": "vpnuseradd"})

}

//更改电话号码
func UpdatenumPertem(c *gin.Context) {

	cnumber := c.DefaultQuery("cnumber", "")
	belongname := c.DefaultQuery("belongname", "")

	if len(belongname) > 0 && len(cnumber) > 0 {

		c.HTML(http.StatusOK, "alerusertelchange.html", gin.H{"cnumber": cnumber, "belongname": belongname})
	}

}

//报警模板添加报警人的组
func Alertpeopaddrento(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	//把当前模板与报警人关联起来
	res := dao.GetCurrAlertPeople(id)
	jsonRes, err := json.Marshal(res)
	if err != nil {
		share.Logger.Warn("序列化失败!")

	}

	c.HTML(http.StatusOK, "addtemalertpeople.html", gin.H{"jsonRes": string(jsonRes), "tempid": id})
}

// 用户管理
func AdminListHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "admin_list.html", gin.H{"index": "index"})
}
