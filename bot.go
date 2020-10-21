package loggingBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
	"strings"
)

var bot *tgbotapi.BotAPI

func startBot() {
	token := os.Getenv("LOGGING_BOT_TOKEN")
	if token == "" {
		log.Fatal("token not found")
	}
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("bot init failed:", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func sendMsgToBot(msg string) {
	admins := strings.Split(os.Getenv("LOGGING_ADMIN_CHAT"), ",")
	for _, admin := range admins {
		adminId, err := strconv.ParseInt(admin, 10, 64)
		if err != nil {
			continue
		}
		msg := tgbotapi.NewMessage(adminId, msg)

		bot.Send(msg)
	}
}
