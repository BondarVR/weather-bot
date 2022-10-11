package telegram

import (
	"fmt"
	"time"
)

func (b *Bot) handleParamForWeather(message Message) string {
	description := message.Weather[0].Description
	mainTemp := b.handleTemperature(message.Main.Temp)
	feelsLike := b.handleTemperature(message.Main.FeelsLike)
	pressure := b.handlePressure(message.Main.Pressure)
	humidity := message.Main.Humidity
	windSpeed := message.Wind.Speed
	clouds := message.Clouds.All
	sunrise := b.handleDate(message.Sys.Sunrise)
	sunset := b.handleDate(message.Sys.Sunset)
	nameLocation := message.Name

	text := fmt.Sprintf("*Локація:* %s\n"+
		"*Стан погоди:* _%s_\n"+
		"*Температура:* _%.1f градусів_\n"+
		"*Відчувається як:* _%.1f градусів_\n"+
		"*Атмосферний тиск:* _%.1f мм.рт.ст._\n"+
		"*Вологість повітря у відсотках:* _%d_\n"+
		"*Швидкість вітру :* _%.2f м/с_\n"+
		"*Хмарність у відсотках:* _%d_\n"+
		"*Світанок:* _%s_\n"+
		"*Закат:* _%s_\n",
		nameLocation,
		description,
		mainTemp,
		feelsLike,
		pressure,
		humidity,
		windSpeed,
		clouds,
		sunrise,
		sunset)

	return text
}

func (b *Bot) handleTemperature(temp float32) float32 {
	handleTemp := temp - 273.15
	return handleTemp
}

func (b *Bot) handleDate(date int) string {
	dateConvert := int64(date)
	timeT := time.Unix(dateConvert, 0)
	timeF := timeT.Format("15:04:05")
	return timeF
}

func (b *Bot) handlePressure(pressure int) float32 {
	formatPressure := float32(pressure) / 1.333
	return formatPressure
}
