package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"weather-bot/internal/config"
	"weather-bot/internal/logger"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	lgr *logger.LogrusLogger
	cfg *config.Config
}

func NewBot(bot *tgbotapi.BotAPI, lgr *logger.LogrusLogger, cfg *config.Config) *Bot {
	return &Bot{bot: bot,
		lgr: lgr,
		cfg: cfg}
}

// Start (starts the bot)
func (b *Bot) Start() error {
	b.lgr.Infof("Authorized on account %s", b.bot.Self.UserName)
	updates := b.initUpdatesChannel()
	if err := b.handleUpdates(updates); err != nil {
		return err
	}
	return nil
}

// handleUpdates processing updates
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.lgr.Fatal(err)
			}
			continue
		}
		if update.Message.Location != nil {
			if err := b.handleLocation(update.Message); err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

// initUpdatesChannel getting updates
func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
