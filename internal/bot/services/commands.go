package bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/fishkaoff/tg-monitor/internal/bot/consts"
)



func (b *Bot) getMetricCommand(command tgbotapi.Update) string {
	webSites, err := b.db.Get(command.Message.Chat.ID)
	if err != nil {
		b.sugar.Error(err)
		b.sendMessage(consts.CANNOTGETSITES, command)
		return ""
	}

	if len(webSites) == 0 {
		b.sendMessage(consts.SITESNOTFOUND, command)
	}
	stats := b.metric.CheckSites(webSites)
	return b.renderStats(stats)
}

func (b *Bot) addSite(chatID int64, site string) error {
	err := b.db.Save(chatID, strings.TrimSpace(site))
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) deleteSite(chatID int64, site string) error {
	err := b.db.Delete(chatID, strings.TrimSpace(site))
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) renderStats(stats map[string]int) string {
	var result string
	var status string

	for key, value := range stats {
		if value != 200 {
			status = consts.SITEUNAWAILABLE
		} else {
			status = consts.SITEAWAILABLE
		}
		result += fmt.Sprintf("%s : \n ➖code: %v \n ➖status: %s \n", key, value, status)
	}

	return result
}
