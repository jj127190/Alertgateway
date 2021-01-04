package main

import (
	"Alertgateway/common/share"
	"Alertgateway/dao"
	"Alertgateway/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

// alertmidtel.blingabc.com

func main() {
	fmt.Println(" ... Platef starting ...")
	defer share.Logger.FileClose()
	defer dao.GDB.Close()
	share.Logger.Warn("[Tel Alert system start....]")
	gin.SetMode(gin.ReleaseMode)
	Rcontext := gin.Default()
	Rcontext.LoadHTMLGlob("templates/*")
	Rcontext.Static("/assets", "./assets")
	router.Distribute(Rcontext)
	Rcontext.Run(share.Conf.Run.StartPort)

}
