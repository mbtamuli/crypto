package telegram

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mbtamuli/telegram-bot/sqlite"
)

func Handler(store *sqlite.Store) {
	botAPIToken, err := getAPIToken("TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	username := "mbtamuli"

	user, err := store.UserByUsername(username)
	if err != nil {
		log.Fatal(err)
	}

	chatIDString := user.ChatID
	chatID, _ := strconv.ParseInt(chatIDString.String, 10, 64)

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Hello, %s!", username))
	bot.Send(msg)
}

func getAPIToken(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("environment variable `%s` not set", key)
	}
	return val, nil
}
