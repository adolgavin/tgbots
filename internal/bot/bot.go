package bot

import (
	"log"
	"tgbots/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotInterface interface {
	Init(apiKey string) error
	HandleUpdates(config *config.Config)
}

type bot struct {
	api *tgbotapi.BotAPI
}

func (b *bot) Init(apiKey string) error {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return err
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	b.api = bot

	return nil
}
