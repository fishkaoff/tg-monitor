package bot

import (
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
	if b.usersStatus[message.Message.Chat.ID] == 1 {
		err := b.addSite(message.Message.Chat.ID, message.Message.Text)
		if err != nil {
			b.sugar.Error(err)
			b.sendMessage(consts.SITENOTADDED, message)
			return
		}
		b.sendMessage(consts.SITEADDED, message)
		delete(b.usersStatus, message.Message.Chat.ID)

	} else if b.usersStatus[message.Message.Chat.ID] == 2 {
		err := b.deleteSite(message.Message.Chat.ID, message.Message.Text)
		if err != nil {
			b.sugar.Error(err)
			b.sendMessage(consts.SITENOTDELETED, message)
			return
		}
		b.sendMessage(consts.SITEDELETED, message)
		delete(b.usersStatus, message.Message.Chat.ID)
	} else {
		b.sendMessage(consts.UNKMOWNCOMMAND, message)
	}
}

func (b *Bot) handleCommand(command tgbotapi.Update) {

	// commands in commands.go
	switch command.Message.Command() {
	case consts.GETMETRICCOMMAND:
		sites := b.getMetricCommand(command)
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
