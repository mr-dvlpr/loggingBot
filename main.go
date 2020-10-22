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
	BotToken  string
	TgUserIds string
	ApiHost   string
	AppName   string
	ApiPort   int
}

var mandatory *LoggingInfo
var initiate bool = false

func Start(mi *LoggingInfo) {
	mandatory = mi
	if mandatory == nil {
		log.Panic("Mandatory Info properties not set. please initialize loggingBot first")
	}

	if mandatory.ApiHost == "" {
		log.Fatal("API host not set. please initialize it first")
	}

	if mandatory.BotToken == "" {
		log.Panic("Telegram bot token not set. please initialize it first")
	}

	initiate = true

	mi.startBot()
	mi.startWebServer()
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
