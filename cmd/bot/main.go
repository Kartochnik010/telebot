package main

import (
	"log"

	"telebot/config"
	"telebot/pkg/telegram"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
	bot, err := telegram.NewBot(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
