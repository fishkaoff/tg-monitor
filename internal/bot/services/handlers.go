package bot


import (
	"log"
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const GETMETRIKCOMMAND = "status"


func (b *Bot) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
        
        if update.Message == nil {
            continue
        }

		if update.Message.IsCommand() {
			b.HandleCommand(update)
			continue
		} 
		b.HandleMessage(update)
    }
}


func (b *Bot) HandleMessage(message tgbotapi.Update) {
	msg := tgbotapi.NewMessage(message.Message.Chat.ID, "Use /status to check web sites")
    

	msg.ReplyToMessageID = message.Message.MessageID

	
	if _, err := b.bot.Send(msg); err != nil {
		
		log.Print(err)
	}

}


func (b *Bot) HandleCommand(command tgbotapi.Update) {
	switch command.Message.Command() {
	case GETMETRIKCOMMAND: 
		b.SendMessage(b.GetMetrikCommand(), command)
		break
	default: 
		b.SendMessage("idk this command((", command)
	}
}

func (b *Bot) GetMetrikCommand() string {
	stats := b.metrik.CheckSites()

	var result string
	var status string

	for key, value := range stats {
		if value != 200 {
			status = "Unavailable❌"
		} else {
			status = "Available✅"
		}
			result += fmt.Sprintf("%s: \n ➖code: %v \n ➖status: %s \n", key, value, status)
	}

	return result
}