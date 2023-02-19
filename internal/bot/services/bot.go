package bot

import (
	"os"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/fishkaoff/tg-monitor/internal/metrik"
	"log"
)



type Bot struct {
	bot *tgbotapi.BotAPI
	metrik *metrik.Metrik
}

type Boter interface {
	Start() error 
	Stop() 
	HandleUpdates(updates tgbotapi.UpdatesChannel)
	HandleMessage(message tgbotapi.Update)
	HandleCommand(command tgbotapi.Update) 
	SendMessage(text string, command tgbotapi.Update) 
}



func NewBot(bot *tgbotapi.BotAPI, metrik *metrik.Metrik) *Bot {
	return &Bot{bot, metrik}
}

func (b *Bot) Start() error {
    updateConfig := tgbotapi.NewUpdate(0)

    
    updateConfig.Timeout = 30

    updates := b.bot.GetUpdatesChan(updateConfig)

    b.HandleUpdates(updates)

	return nil

}


func (b *Bot) SendMessage(text string, command tgbotapi.Update) {
	msg := tgbotapi.NewMessage(command.Message.Chat.ID, text)

	msg.ReplyToMessageID = command.Message.MessageID

	if _, err := b.bot.Send(msg); err != nil {
		log.Print(err)
	}
}

func (b *Bot) Stop() {
	os.Exit(1)
}


