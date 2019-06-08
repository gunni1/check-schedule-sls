package main

import (
	"fmt"
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//TelegramNotificator is for bundleling the Telegram API with the ChatID of the notification target
type TelegramNotificator struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

//CreateTelegramNotificator - is a builder function to instantiate the Telegram API
func CreateTelegramNotificator(botAPIToken string, notificationTarget int64) TelegramNotificator {
	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		log.Panic(err)
	}
	return TelegramNotificator{Bot: bot, ChatID: notificationTarget}
}

//SendNotification - Sends a schedule change to the notifications target
func (notificator TelegramNotificator) SendNotification(change ScheduleChange) {
	log.Printf("Authorized on account %s", notificator.Bot.Self.UserName)
	text := fmt.Sprintf("Änderung im Stundenplan für kürzel %s am %s gefunden.\n", change.TeacherCode, change.DateConcerned)
	msg := tgbotapi.NewMessage(notificator.ChatID, text)
	notificator.Bot.Send(msg)
}
