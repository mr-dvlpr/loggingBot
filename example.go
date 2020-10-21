package loggingBot

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//For sending error from your app to telegram bot using bellow code
//add API_URL to .env file
//github.com/joho/godotenv library is useful for loading local .env file
func sendError(msg string) {
	apiHost := os.Getenv("LOGGING_API_HOST")
	if apiHost == "" {
		return
	}
	apiPort := os.Getenv("LOGGING_API_PORT")
	if apiPort == "" {
		apiPort = "8080"
	}
	apiUrl := fmt.Sprintf("%s:%s", apiHost, apiPort)
	data := url.Values{}
	data.Set("app", "appName")
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
