package main


import (
	"fmt"
	"log"
	"github.com/fishkaoff/tg-monitor/configs"
	"github.com/fishkaoff/tg-monitor/internal/metrik"
	Bot "github.com/fishkaoff/tg-monitor/internal/bot/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)



func main() {
	WEBSITES := []string{"https://yandex.ru", "https://google.com",  "https://asdasdsaddsa.sv"}


	// init config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to init config")
	}
	
	
	// start bot
	bot, err := tgbotapi.NewBotAPI(config.TGTOKEN)
	if err != nil {
		log.Fatal("failed at bot`s auth")
	}

	// IoC
	mtr := metrik.NewMetrik(WEBSITES)
	tg := Bot.NewBot(bot, mtr)
	
	fmt.Printf("Authentificated with token %v", config.TGTOKEN)
	fmt.Println("starting bot")
	tg.Start()
}