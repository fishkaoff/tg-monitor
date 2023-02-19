package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


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
	b.SendMessage("Use /status to check web sites", message)

}

func (b *Bot) HandleCommand(command tgbotapi.Update) {

	// commands in commands.go
	switch command.Message.Command() {
	case GETMETRICCOMMAND:
		b.SendMessage(b.GetMetricCommand(), command)
		break
	default:
		b.SendMessage("idk this command((", command)
	}
}

func (b *Bot) GetMetricCommand() string {
	stats := b.metrik.CheckSites()

	return b.RenderStats(stats)
}

func (b *Bot) RenderStats(stats map[string]int) string {
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
