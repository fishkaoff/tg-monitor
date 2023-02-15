package main


import (
	"fmt"
	"log"
	"github.com/fishkaoff/tg-mail/configs"
	Bot "github.com/fishkaoff/tg-mail/internal/bot/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func main() {
	// init config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to init config")
	}
	

	// init Google API
	
	// start bot
	bot, err := tgbotapi.NewBotAPI(config.TGTOKEN)
	if err != nil {
		log.Fatal("failed at bot`s auth")
	}

	tg := Bot.NewBot(bot)
	fmt.Printf("Authentificated with token %v", config.TGTOKEN)
	fmt.Println("starting bot")
	tg.Start()
}