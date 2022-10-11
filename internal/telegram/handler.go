package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

const (
	commandStart = "start"
	text         = "Цей бот вміє віддавати інформацію про погоду за вашою геопозицією :)\n\n" +
		"Коротка інструкція як користуватися: \n" +
		"1. Натисніть на скріпку праворуч внизу екрану\n" +
		"2. Оберіть у нижньому меню ʼГеопозицияʼ\n" +
		"3. Оберіть перший пункт ʼОтправить свою геопозициюʼ\n" +
		"4. Хай щастить!"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	default:
		return errors.New("Invalid type of command! ")
	}
	return nil
}

func (b *Bot) handleLocation(message *tgbotapi.Message) error {
	lon := message.Location.Longitude
	lat := message.Location.Latitude
	if err := b.handleSendMessage(lon, lat, message); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleSendMessage(lon, lat float64, message *tgbotapi.Message) error {
	msg, err := b.getWeatherInfo(lon, lat, message)
	if err != nil {
		return err
	}
	msg.ParseMode = "markdown"
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
