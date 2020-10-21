package loggingBot

import (
	"github.com/joho/godotenv"
	"log"
)

//
//func main() {
//	Start()
//}

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	startWebServer()
	startBot()
}
