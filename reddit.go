package main

import (
	"fmt"
	"regexp"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vartanbeno/go-reddit/reddit"
)

// match is a function that checks whether any keyword exists in the string
func match(str string, strSlice []string) bool {
	for _, s := range strSlice {
		reg := regexp.MustCompile(fmt.Sprintf("(?i)%v", regexp.QuoteMeta(s)))
		if reg.MatchString(str) {
			return true
		}
	}

	return false
}

// filterAndPost filters post with the config filters and posts them on telegram channels
func filterAndPost(bot *tgbotapi.BotAPI, cfg config, post *reddit.Post) {
	if len(cfg.Filters.Keywords) > 0 && !match(post.Title, cfg.Filters.Keywords) {
		log.Debugf("Post %v filtered, no keywords match", post.ID)
		return
	}
	if len(cfg.Filters.MediaWhitelist) > 0 && post.URL != "" && !match(post.URL, cfg.Filters.MediaWhitelist) {
		log.Debugf("Post %v filtered, no media whitelist urls match", post.ID)
		return
	}

	for _, ch := range cfg.Telegram.ChannelIDs {
		if isChMember(bot, ch) {
			msg := tgbotapi.NewMessage(ch, post.Title+"\n\n"+post.URL)
			go bot.Send(msg)
		}
	}
}
