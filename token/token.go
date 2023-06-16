package token

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot, _ = tgbotapi.NewBotAPI("5954727378:AAEDoSRsyai4-fRwZrNoiOvQFj5AcMBbq3Q")


// package main

// import (
// 	"strings"
// 	"time"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// func main() {

// 	bot, err := tgbotapi.NewBotAPI("5954727378:AAEDoSRsyai4-fRwZrNoiOvQFj5AcMBbq3Q")
// 	if err != nil {
// 		panic(err)
// 	}

// 	//Устанавливаем время обновления
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 10
// 	updates, err := bot.GetUpdatesChan(u)

// 	for update := range updates {
// 		//now := time.Now()

// 		// if now.Hour() < 12 {
// 		// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Good morning!")
// 		// 	_, err := bot.Send(msg)
// 		// 	if err != nil {
// 		// 		log.Panic(err)
// 		// 	}
// 		// }

// 		if update.Message == nil {
// 			continue
// 		}

// 		messageText := update.Message.Text
// 		if strings.Contains(update.Message.Text, "hi") {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello!")
// 			bot.Send(msg)
// 		} else if strings.Contains(messageText, "hi") || strings.Contains(messageText, "hello") {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello!")
// 			bot.Send(msg)
// 		} else if strings.Contains(messageText, "time") {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, time.Now().String())
// 			bot.Send(msg)
// 		} else if strings.Contains(messageText, "bye") {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Goodbye!")
// 			bot.Send(msg)
// 		} else if strings.Contains(messageText, "Dias") {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Kotakbas")
// 			bot.Send(msg)
// 		} else if strings.Contains(messageText, "dias") {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Kotakbas")
// 			bot.Send(msg)
// 		} else if strings.Contains(messageText, "Диас") {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Kotakbas")
// 			bot.Send(msg)
// 		} else if strings.Contains(messageText, "диас") {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Kotakbas")
// 			bot.Send(msg)
// 		} else {
// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't understand what you're saying.")
// 			bot.Send(msg)
// 		}
// 	}

// }
