package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ptime "github.com/yaa110/go-persian-calendar"
	"time"
)

func init() {
	r := gin.Default()
	r.POST("/error", getError)
	go r.Run()
	fmt.Println("run gin")
}

func getError(c *gin.Context){
	errText:=c.PostForm("text")
	appId:=c.PostForm("app")
	if appId == "" {
		return
	}

	pt:=ptime.New(time.Now())
	msg:=fmt.Sprintf("New error from %s:\n\n%s\n\n date: %s", appId, errText, pt.Format("yyyy/MM/dd HH:mm:ss Z z"))

	sendMsgToBot(msg)
}
