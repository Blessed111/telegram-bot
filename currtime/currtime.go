package currtime

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMessageAtTime(token string, chatID int64, text string, t time.Time) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(chatID, text)

	// Установка времени, когда нужно отправить сообщение
	now := time.Now()
	d := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, time.Local)
	duration := d.Sub(now)

	// Ожидание до указанного времени
	<-time.After(duration)

	// Отправка сообщения
	_, err = bot.Send(message)
	if err != nil {
		return err
	}

	return nil
}
