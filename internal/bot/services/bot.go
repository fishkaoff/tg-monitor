package bot

import (
	"log"
	"os"

	"github.com/fishkaoff/tg-monitor/internal/metric"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	metrik *metric.Metric
}

type Boter interface {
	Start() error
	Stop()
	HandleUpdates(updates tgbotapi.UpdatesChannel)
	HandleMessage(message tgbotapi.Update)
	HandleCommand(command tgbotapi.Update)
	SendMessage(text string, command tgbotapi.Update)
	GetMetricCommand() string
	RenderStats(stats map[string]int) string
}

func NewBot(bot *tgbotapi.BotAPI, metrik *metric.Metric) *Bot {
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
