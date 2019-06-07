package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	botAPITokenEnv        = "BOT_TOKEN"
	notificationTargetEnv = "NOTIFICATION_TARGET"
)

func main() {
	lambda.Start(handler)
}

type ScheduleChange struct {
	DateConcerned string `json:"dateConcerned"`
	IssueDate     string `json:"issueDate"`
	TeacherCode   string `json:"teacherCode"`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	botAPIToken := parseEnvMandatory(botAPITokenEnv)
	notificationTarget := parseEnvInt64Mandatory(notificationTargetEnv)
	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		log.Panic(err)
	}

	if len(sqsEvent.Records) == 0 {
		return errors.New("No SQS message passed to function")
	}

	for _, msg := range sqsEvent.Records {
		var change ScheduleChange
		unmarshalErr := json.Unmarshal([]byte(msg.Body), &change)
		if unmarshalErr != nil {
			log.Println("Error: " + unmarshalErr.Error())
			return unmarshalErr
		}
		text := fmt.Sprintf("Änderung im Stundenplan für kürzel %s am %s gefunden.\n", change.TeacherCode, change.DateConcerned)
		SendNotification(text, notificationTarget, bot)
	}

	return nil
}

func SendNotification(text string, chatId int64, bot *tgbotapi.BotAPI) {
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(chatId, text)

	bot.Send(msg)
}

func parseEnvInt64Mandatory(variableKey string) int64 {
	varValue := parseEnvMandatory(variableKey)
	intValue, err := strconv.ParseInt(varValue, 10, 64)
	if err != nil {
		log.Fatalln("Error: " + err.Error())
	}
	return intValue
}

func parseEnvMandatory(variableKEy string) string {
	variableValue := os.Getenv(variableKEy)
	if variableValue == "" {
		log.Fatalln("Environment variable: " + variableKEy + " is empty")
	}
	return variableValue
}
