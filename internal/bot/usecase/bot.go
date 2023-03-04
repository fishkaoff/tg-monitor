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
	sugar       *zap.SugaredLogger
	metric      metric.Metricer
	db          sites.SiteRepository
	mw          Middlware
	usersStatus map[int64]int
}

type Middlware interface {
	CheckUrl(URL string) bool
	CheckMatches(webSites []string, site string) bool
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
	deleteSiteCommand(chatID int64, site string) error
	addSiteCommand(chatID int64, site string) error
	getSite(chatid int64, site string) string 
}

func NewBot(bot *tgbotapi.BotAPI, metric metric.Metricer, db sites.SiteRepository, sugar *zap.SugaredLogger, mw Middlware) *Bot {
	return &Bot{
		bot,
		sugar,
		metric,
		db,
		mw,
		make(map[int64]int),
	}
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
