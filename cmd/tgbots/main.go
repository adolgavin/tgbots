package main

import (
	"tgbots/internal/app"
	"tgbots/internal/bot"
	"tgbots/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	app := app.New(cfg, &bot.ChatGPTBot{})
	app.Run()
}
