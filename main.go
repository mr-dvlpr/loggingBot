package loggingBot

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type LoggingInfo struct {
	ApiHost string
	AppName string
	ApiPort int
}

type serviceInfo struct {
	RunningPort int
	BotToken    string
	TgUserIds   string
}

var mandatory *LoggingInfo
var initiate bool = false

func StartService(mi *serviceInfo) {
	if mi == nil {
		log.Panic("Mandatory Info properties not set. please initialize loggingBot first")
	}

	if mi.RunningPort < 0 {
		log.Fatal("API host not set. please initialize it first")
	}

	if mi.BotToken == "" {
		log.Panic("Telegram bot token not set. please initialize it first")
	}

	mi.startBot()
	mi.startWebServer()
}

func Startlogger(mi *LoggingInfo) {
	mandatory = mi
	if mi == nil {
		log.Panic("Mandatory Info properties not set. please initialize loggingBot first")
	}

	if mi.ApiPort < 0 {
		log.Fatal("API host not set. please initialize it first")
	}

	if mi.ApiHost == "" {
		log.Panic("API host not set. please initialize it first")
	}

	initiate = true
}

func SendError(msg string) {
	if !initiate {
		log.Panic("LoggingBot not Initiated")
	}
	apiHost := mandatory.ApiHost
	if apiHost == "" {
		return
	}
	apiPort := strconv.Itoa(mandatory.ApiPort)
	if apiPort == "" {
		apiPort = "8080"
	}
	apiUrl := fmt.Sprintf("%s:%s", apiHost, apiPort)
	data := url.Values{}
	data.Set("app", mandatory.AppName)
	data.Set("text", msg)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = "/error/"
	urlStr := u.String() // "https://api.com/error/"

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
}
