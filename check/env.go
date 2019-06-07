package main

import (
	"log"
	"os"
	"strconv"
)

const (
	userEnv     = "PAGE_USER"
	pwEnv       = "PAGE_PW"
	codeEnv     = "TEACHER_CODE"
	baseURLEnv  = "BASE_URL"
	sqsQueueEnv = "SQS_QUEUE"
	daysEnv     = "DAYS_COUNT"
)

//CreateSchedulerConfigFromEnv - prepares all variables for the schedule request based on env variables
func CreateSchedulerConfigFromEnv() ScheduleClientConfig {
	username := parseEnvMandatory(userEnv)
	password := parseEnvMandatory(pwEnv)

	baseURL := os.Getenv(baseURLEnv)
	if baseURL == "" {
		baseURL = "https://www.stundenplan24.de/10124219/vplanle/vdaten/VplanLe"
	}
	return ScheduleClientConfig{User: username, Password: password, BaseURL: baseURL}
}

//GetTeacherCode - receive the Code from the environment Variable
func GetTeacherCode() string {
	return parseEnvMandatory(codeEnv)
}

//GetSQSQueueURL - receive a SQS Queue URL, where the notification should be published to from env var
func GetSQSQueueURL() string {
	return parseEnvMandatory(sqsQueueEnv)
}

//GetDaysCount - receive the configured count of days in the future to check. Default is 2.
func GetDaysCount() int {
	varValue := os.Getenv(daysEnv)
	if days, err := strconv.Atoi(varValue); err == nil {
		return days
	}
	return 2
}

func parseEnvMandatory(variableKEy string) string {
	variableValue := os.Getenv(variableKEy)
	if variableValue == "" {
		log.Fatalln("Environment variable: " + variableKEy + " is empty")
	}
	return variableValue
}
