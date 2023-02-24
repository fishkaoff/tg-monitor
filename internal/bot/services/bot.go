package bot

import (
	"os"

	"github.com/fishkaoff/tg-monitor/internal/metric"
	sites "github.com/fishkaoff/tg-monitor/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	metric      metric.Metricer
	db          sites.SiteRepository
	sugar       *zap.SugaredLogger
	usersStatus map[int64]int
}

type Boter interface {
	Start() error
	Stop()
	sendMessage(text string, command tgbotapi.Update)
	handleUpdates(updates tgbotapi.UpdatesChannel)
	handleMessage(message tgbotapi.Update)
	handleCommand(command tgbotapi.Update)
	renderStats(stats map[string]int) string
	getMetricCommand(command tgbotapi.Update) string
	deleteSite(chatID int64, site string) error
	addSite(chatID int64, site string) error
}

func NewBot(bot *tgbotapi.BotAPI, metric metric.Metricer, db sites.SiteRepository, sugar *zap.SugaredLogger) *Bot {
	return &Bot{bot, metric, db, sugar, make(map[int64]int)}
}

func (b *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := b.bot.GetUpdatesChan(updateConfig)

	b.sugar.Info("Bot started")
	b.handleUpdates(updates)

}

func (b *Bot) sendMessage(text string, command tgbotapi.Update) {
	msg := tgbotapi.NewMessage(command.Message.Chat.ID, text)

	msg.ReplyToMessageID = command.Message.MessageID

	if _, err := b.bot.Send(msg); err != nil {
		b.sugar.Error(err)
	}
}

func (b *Bot) Stop() {
	os.Exit(1)
}
