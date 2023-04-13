package app

import (
	"log"
	"sync"
	"tgbots/internal/bot"
	"tgbots/internal/config"
)

type Application struct {
	config *config.Config
	bots   []bot.BotInterface
}

func New(cfg *config.Config, bots ...bot.BotInterface) *Application {
	return &Application{
		config: cfg,
		bots:   bots,
	}
}

func (a *Application) Run() {
	var wg sync.WaitGroup
	wg.Add(len(a.bots))

	for _, b := range a.bots {
		go func(b bot.BotInterface) {
			defer wg.Done()

			err := b.Init(a.config.TelegramAPIKey)
			if err != nil {
				log.Panic(err)
			}

			b.HandleUpdates(a.config)
		}(b)
	}

	wg.Wait()
}
