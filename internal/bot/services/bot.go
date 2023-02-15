package bot

import (
	"os"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Boter interface {
	Start() error 
	Stop() 
	HandleUpdates(updates tgbotapi.UpdatesChannel)
	ParseMessage(message *tgbotapi.Update)
}

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot}
}

func (b *Bot) Start() error {
    updateConfig := tgbotapi.NewUpdate(0)

    
    updateConfig.Timeout = 30

    updates := b.bot.GetUpdatesChan(updateConfig)

    b.HandleUpdates(updates)

	return nil

}

func (b *Bot) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
        
        if update.Message == nil {
            continue
        }

		b.HandleMessage(update)
    }
}


func (b *Bot) HandleMessage(message tgbotapi.Update) {
	msg := tgbotapi.NewMessage(message.Message.Chat.ID, message.Message.Text)
    

	msg.ReplyToMessageID = message.Message.MessageID

	
	if _, err := b.bot.Send(msg); err != nil {
		
		panic(err)
	}
}


func (b *Bot) Stop() {
	os.Exit(1)
}


