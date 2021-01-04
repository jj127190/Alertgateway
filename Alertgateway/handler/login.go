package handler

import (
	//"Alertgateway/common/share"
	// "Alertgateway/dao"
	"Alertgateway/common/ldap"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 0表示立即销毁cookie

// -1表示关闭浏览器销毁cookie

// 大于1是设置的销毁时间，单位是秒，例如1小时后销毁，就可以设置60*60
func LogOut(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/welcome")
	c.SetCookie("VASessionId", "logout", 1, "/", "localhost", false, true)
	// c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Welcome(c *gin.Context) {
	cookie, err := c.Cookie("VASessionId")
	if err != nil {
		fmt.Println("cookie get err:", err)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
	c.SetCookie("VASessionId", "logout", 3600, "/", "localhost", false, true)

	c.HTML(http.StatusOK, "login.html", gin.H{"loggin": "登录"})

}

func LoginDu(c *gin.Context) {

	username := c.DefaultPostForm("username", "username_nil")
	password := c.DefaultPostForm("password", "password_nil")

	if username != "username_nil" && password != "password_nil" {

		err := ldap.LdapVerifica(username, password)

		if err != nil {

			c.HTML(http.StatusOK, "error_return_index.html", gin.H{"msg": "账号或密码输入错误！"})
			return
		} else {

			passwd, err := ldap.GetLdaPass(username)

			if err != nil {
				fmt.Println("获取密码出错!ERR:", err)
				return
			}
			if passwd == password {

				c.SetCookie("VASessionId", "logging", 3600, "/", "", false, true)
				// c.SetCookie("VASessionId", "logging", 3600, "/", "localhost", false, true)
				// fmt.Println("此处设置账号到环境变量中...", username)
				c.Set("caccount", username)
				//c.HTML(http.StatusOK, "index.html", gin.H{"index": "index"})
				c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/rendto/index?username=%s", username))
			} else {

				c.HTML(http.StatusOK, "error_return_index.html", gin.H{"msg": "账号或密码输入错误！"})
				return
			}
		}

	}

}
