package bot

import (
	"log"
	"tgbots/internal/chatgptapi"
	"tgbots/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ChatGPTBot struct {
	bot
}

func (b *ChatGPTBot) HandleUpdates(config *config.Config) {
	gptAPI := chatgptapi.New(config.ChatGPTAPIKey)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		log.Println("some errors: %v", err)
		return
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		prompt := update.Message.Text
		response := gptAPI.GetResponse(prompt)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		b.api.Send(msg)
	}
}
