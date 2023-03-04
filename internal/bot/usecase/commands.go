package bot

import (
	"fmt"
	"strings"

	"github.com/fishkaoff/tg-monitor/internal/bot/consts"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// TODO move error handling to handlers
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

func (b *Bot) addSiteCommand(chatID int64, site string) string {

	// validate url
	if !b.mw.CheckUrl(site) {
		return consts.NOTURL
	}

	// check for similar websites in db
	webSites, err := b.db.Get(chatID)
	if err != nil {
		b.sugar.Error(err)
		return consts.CANNOTGETSITES
	}

	if b.mw.CheckMatches(webSites, site) {
		return consts.SITEALREADYADDED
	}

	// add site
	err = b.db.Save(chatID, strings.TrimSpace(site))
	if err != nil {
		b.sugar.Error(err)
		return consts.SITENOTADDED
	}
	return consts.SITEADDED
}

func (b *Bot) getSite(chatid int64, site string) string {
	webSite, err := b.db.GetSite(chatid, site)
	if err != nil {
		return consts.SITENOTDELETED
	}

	if len(webSite) == 0 {
		return consts.SITENOTFOUND
	}

	return webSite
}

func (b *Bot) deleteSiteCommand(chatID int64, site string) string {
	trimmedSite := strings.TrimSpace(site)


	if !b.mw.CheckUrl(trimmedSite) {
		return consts.SITENOTDELETED
	}

	if b.getSite(chatID, trimmedSite) == consts.SITENOTFOUND {
		return consts.SITENOTFOUND
	}

	err := b.db.Delete(chatID, trimmedSite)
	if err != nil {
		b.sugar.Error(err)
		return consts.SITENOTDELETED
	}

	return consts.SITEDELETED + fmt.Sprintf("( %s)", trimmedSite)
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
