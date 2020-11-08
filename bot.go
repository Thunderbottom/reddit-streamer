package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// isChWhitelisted is a function that checks whether the
// channel ID exists in the channel ID whitelist
func isChWhitelisted(ch int64, chIDs []int64) bool {
	for _, c := range chIDs {
		if ch == c {
			return true
		}
	}
	return false
}

// isChMember is a function that checks whether the bot
// is a current member of the specified chat
func isChMember(bot *tgbotapi.BotAPI, ch int64) bool {
	chatMember, err := bot.GetChatMember(tgbotapi.ChatConfigWithUser{
		ChatID: ch,
		UserID: bot.Self.ID,
	})
	if err != nil {
		log.Errorf("The check for chat membership failed: %v", err)
		return false
	}
	if chatMember.IsMember() {
		return true
	}
	return false
}

// handleBotUpdates is a function that handles updates from
// the telegram bot api
func handleBotUpdates(ud tgbotapi.Update, bot *tgbotapi.BotAPI, cfg config) {
	// completely ignore non-updates and private messages
	if ud.Message != nil && !ud.Message.Chat.IsPrivate() {
		// check if new members were added to the group
		if ud.Message.NewChatMembers != nil {
			for _, ncm := range *ud.Message.NewChatMembers {
				// check if the bot was added to a group
				if ncm.IsBot && ncm.ID == bot.Self.ID {
					// leave the chat if it is not in the channel whitelist
					if !isChWhitelisted(ud.Message.Chat.ID, cfg.Telegram.ChannelIDs) {
						log.Infof("Channel %v not found in the whitelist, exiting", ud.Message.Chat.ID)
						bot.LeaveChat(ud.Message.Chat.ChatConfig())
					}
					break
				}
			}
		}
	}

	// TODO: handle group messages for bot
}

// initBot is a function that generates and returns
// an instance of the telegram bot api
func initBot(cfg config) *tgbotapi.BotAPI {
	if cfg.Telegram.APIKey == "" {
		log.Fatal("Telegram API Token not found.")
	} else if len(cfg.Telegram.ChannelIDs) == 0 {
		log.Fatal("The bot requires Telegram Channel IDs to post content. Please configure channels before continuing.")
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.APIKey)
	if err != nil {
		log.Fatal("Failed to configure Telegram bot: ", err)
	}
	log.Infof("Authorised as @%v", bot.Self.UserName)

	for _, ch := range cfg.Telegram.ChannelIDs {
		if !isChMember(bot, ch) {
			log.Errorf("The bot is currently not a member of the chat: %v", ch)
		}
	}

	bot.Debug = cfg.General.Debug
	log.Debug("Bot is set to debug mode")

	return bot
}
