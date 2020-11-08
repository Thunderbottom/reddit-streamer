package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vartanbeno/go-reddit/reddit"
)

var ctx = context.Background()

func main() {
	cfg := getConfig()
	log = getLogger(cfg.General.Debug)

	if cfg.Filters.Subreddit == "" {
		log.Fatal("No subreddit configured under filters.")
	}

	bot := initBot(cfg)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Errorf("Failed to get bot updates channel: %v", err)
	}

	sig := make(chan os.Signal, 1)
	defer close(sig)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	if cfg.General.Interval <= 0 {
		cfg.General.Interval = 60
		log.Debugf("Stream interval not set or is negative, defaulting to 60 seconds")
	}

	log.Debugf("Starting Reddit Stream for subreddit: %v", cfg.Filters.Subreddit)
	if len(cfg.Filters.Keywords) == 0 {
		log.Warn("No keywords configured! This may lead to a channel spam.")
	}

	if len(cfg.Filters.MediaWhitelist) == 0 {
		log.Warn("No media whitelist configured! All links will be posted.")
	}

	log.Infof("Setting stream interval to %v seconds", cfg.General.Interval)
	posts, errs, stop := reddit.DefaultClient().Stream.Posts(cfg.Filters.Subreddit, reddit.StreamInterval(time.Second*time.Duration(cfg.General.Interval)), reddit.StreamDiscardInitial)
	defer stop()

	for {
		select {
		case post, ok := <-posts:
			if !ok {
				return
			}
			log.Debugf("Received post: %v", post.Title)
			go filterAndPost(bot, cfg, post)
		case err, ok := <-errs:
			if !ok {
				return
			}
			log.Infof("An Error occurred while streaming posts: %v", err)
		case rcvSig, ok := <-sig:
			if !ok {
				return
			}
			log.Infof("Received %v signal, stopping bot.", rcvSig)
			return
		case ud, ok := <-updates:
			if !ok {
				return
			}
			log.Infof("Received update from chat: %v", ud.Message.Chat.ID)
			handleBotUpdates(ud, bot, cfg)
		}
	}
}
