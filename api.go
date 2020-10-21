package loggingBot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ptime "github.com/yaa110/go-persian-calendar"
	"os"
	"time"
)

func startWebServer() {
	r := gin.Default()
	r.POST("/error", getError)
	apiPort := os.Getenv("LOGGING_API_PORT")
	if apiPort == "" {
		apiPort = ":8080"
	} else {
		apiPort = ":" + apiPort
	}
	go r.Run(apiPort)
	fmt.Println("run gin")
}

func getError(c *gin.Context) {
	errText := c.PostForm("text")
	appId := c.PostForm("app")
	if appId == "" {
		fmt.Println("empty app id")
		return
	}

	pt := ptime.New(time.Now())
	msg := fmt.Sprintf("New error from %s:\n\n%s\n\n date: %s", appId, errText, pt.Format("yyyy/MM/dd HH:mm:ss Z z"))

	sendMsgToBot(msg)
}
