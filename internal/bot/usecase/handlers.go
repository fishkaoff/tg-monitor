package bot

import (
	"fmt"

	"github.com/fishkaoff/tg-monitor/internal/bot/consts"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update)
			continue
		}
		b.handleMessage(update)
	}
}

func (b *Bot) handleMessage(message tgbotapi.Update) {
	// user waiting for add site
	if b.usersStatus[message.Message.Chat.ID] == 1 {
		result := b.addSiteCommand(message.Message.Chat.ID, message.Message.Text)
		b.sugar.Info(fmt.Sprintf("request: Handled command /addsite from user: %v; respponse: %s", message.Message.Chat.ID, result))
		b.sendMessage(result, message)
		return
	}

	// user waiting for delete site
	if b.usersStatus[message.Message.Chat.ID] == 2 {
		b.sendMessage(b.deleteSiteCommand(message.Message.Chat.ID, message.Message.Text), message)
		b.sugar.Info(fmt.Sprintf("request: Handled command /addsite from user: %v; sending response", message.Message.Chat.ID))
		return
	}

	b.sendMessage(consts.UNKMOWNCOMMAND, message)
}

func (b *Bot) handleCommand(command tgbotapi.Update) {

	// commands in commands.go
	switch command.Message.Command() {
	case consts.GETMETRICCOMMAND:
		sites := b.getMetricCommand(command)
		b.sugar.Info(fmt.Sprintf("request: Handled command /status from user: %v; sending response", command.Message.Chat.ID))
		b.sendMessage(sites, command)
		break
	case consts.ADDSITECOMMAND:
		b.sendMessage(consts.SENDDATA, command)
		// 1 means wait for site, which should add, 2 means for site which should delete, 0 not waiting for data
		b.usersStatus[command.Message.Chat.ID] = 1
	case consts.DELETESITECOMMAND:
		b.sendMessage(consts.SENDDATA, command)
		// 1 means wait for site, which should add, 2 means for site which should delete, 0 not waiting for data
		b.usersStatus[command.Message.Chat.ID] = 2
		break
	default:
		b.sendMessage(consts.UNKMOWNCOMMAND, command)
	}
}
