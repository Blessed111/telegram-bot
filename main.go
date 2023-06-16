package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleIncomingMessage(update tgbotapi.Update) {
	bot, err := tgbotapi.NewBotAPI("5954727378:AAEDoSRsyai4-fRwZrNoiOvQFj5AcMBbq3Q")
	if err != nil {
		log.Fatal(err)
	}
	message := update.Message

	if message != nil && message.IsCommand() && message.Command() == "r" {
		// Generate a random number between 1 and 100
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(10) + 1
		username := update.Message.From.FirstName
		if username == "Ospanov" {
			replyText := fmt.Sprintf("%s rolled 10!", username)
			reply := tgbotapi.NewMessage(message.Chat.ID, replyText)
			_, err := bot.Send(reply)
			if err != nil {
				log.Println(err)
			}
		} else {
			// Construct the reply message with the user's name and the random number
			replyText := fmt.Sprintf("%s rolled %d!", username, num)
			reply := tgbotapi.NewMessage(message.Chat.ID, replyText)

			// Send the reply message
			_, err := bot.Send(reply)
			if err != nil {
				log.Println(err)
			}
		}

	}
}

func lunch(bot *tgbotapi.BotAPI) {
	chatID := int64(809582341)

	//message := tgbotapi.NewMessage(chatID, "LUNCH_TIME!!!")

	// Установка времени, когда нужно отправить сообщение
	//now := time.Now()
	//t := time.Date(now.Year(), now.Month(), now.Day(), 18, 8, 0, 0, time.Local)
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 18, 10, 0, 0, time.Now().Location())
	for {
		if time.Now().After(t) {
			message := tgbotapi.NewMessage(chatID, "Hello, World!")
			bot.Send(message)
			t = t.Add(24 * time.Hour)
		}
		time.Sleep(1 * time.Minute)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("5954727378:AAEDoSRsyai4-fRwZrNoiOvQFj5AcMBbq3Q")
	if err != nil {
		log.Fatal(err)
	}
	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 10

	// ID чата, в который будут отправляться сообщения

	// Начинаем слушать обновления

	// updates, err := bot.GetUpdatesChan(u)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for update := range updates {
	// 	// Проверяем, является ли сообщение командой "/roll"
	// 	if update.Message != nil {
	// 		// Обрабатываем команду "/roll"
	// 		handleIncomingMessage(update)

	// 	}
	// }
	lunch(bot)
}
