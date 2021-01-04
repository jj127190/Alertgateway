package router

import (
	//"fmt"

	"Alertgateway/common/middleware"
	"Alertgateway/common/utils"
	"Alertgateway/handler"

	"github.com/gin-gonic/gin"
)

func Distribute(r *gin.Engine) {

	r.NoRoute(utils.NoResponse)

	// 登录相关
	r.POST("/loggin", handler.LoginDu)
	r.POST("/", handler.LoginDu)
	// r.GET("/logout", handler.LogOut)
	r.GET("/logout", handler.LogOut)
	r.GET("/welcome", handler.Welcome)

	//页面渲染相关
	renderTem := r.Group("/rendto")
	renderTem.Use(middleware.LoginMiddler())
	{

		// index
		renderTem.GET("/index", handler.IndexHandler)
		// 首页
		renderTem.GET("/welcome", handler.WelcomeHandler)
		//登录
		renderTem.GET("/login", handler.LoginHandler)
		// 错误404
		renderTem.GET("/errorf", handler.F404Handler)

		//报警维护
		renderTem.GET("/alertmaint", handler.Alertmaint)
		//报警给某人
		renderTem.GET("/alertopeople", handler.Alertopeople)

		//报警模板操作
		renderTem.GET("/altemplaterender", handler.Altemplaterender)

		///添加报警人电话

		renderTem.GET("/addAlertPeoRender", handler.AddAlertPeoRender)

		//添加报警模板

		renderTem.GET("/addAlertTempRender", handler.AddAlertTempRender)
		//修改电话
		renderTem.GET("/updatenumrender", handler.UpdatenumPertem)

		//添加报警人组

		renderTem.GET("/alertpeopaddrento", handler.Alertpeopaddrento)

	}

	// api接口
	apiv := r.Group("/api/v1/")

	{

		apiv.POST("alert", handler.TelAlertmsg) //电话报警统一接口

	}

	//api table 接口

	apit := r.Group("/api/table/")

	{

		// apit.GET("templatealert", handler.GetAlertTemp) //获取报警模板
		apit.GET("tempalertinfo", handler.GetAlertTemp) //获取报警模板
		apit.GET("alertpeople", handler.GetAlertPeople) //获取报警人
		// apit.GET("tempalertinfo", handler.Tempalertinfo) //获取报警模板信息

	}

	apiajax := r.Group("/api/ajax/")
	{
		apiajax.POST("addalertpeople", handler.Addalertpeople) //添加报警人电话
		apiajax.GET("changecnstats", handler.Changecnstats)    //更改状态
		apiajax.POST("delctn", handler.Delctn)                 //删除
		apiajax.POST("updatetel", handler.Updatetel)           //更新电话号码

		//语音模板编辑参数内容
		apiajax.POST("editparms", handler.Editparms)
		//更改报警模板维护状态

		apiajax.GET("changevttempstats", handler.Changevttempstats)

		//报警模板维护时间内容更新入库

		apiajax.POST("alertemlaydatedone", handler.Alertemlaydatedone)

		//模板添加报警人
		apiajax.POST("temaddalpeo", handler.Temaddalpeo)

		//添加报警模板
		apiajax.POST("addalerttemp", handler.Addalerttemp)

		//删除报警模板
		apiajax.POST("deltempmsg", handler.Deltempmsg)

	}

}
