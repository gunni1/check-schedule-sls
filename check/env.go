package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

)

const (
	userEnv    = "PAGE_USER"
	pwEnv      = "PAGE_PW"
	dateEnv    = "DATE"
	codeEnv    = "TEACHER_CODE"
	baseURLEnv = "BASE_URL"
)

var (
	dateFormat, _ = regexp.Compile("[0-9]{8}")
)

//CreateSchedulerConfigFromEnv - prepares all variables for the schedule request based on env variables
func CreateSchedulerConfigFromEnv() ScheduleClientConfig {
	username := parseEnvMandatory(userEnv)
	password := parseEnvMandatory(pwEnv)

	baseURL := os.Getenv(baseURLEnv)
	if baseURL == "" {
		baseURL = "https://www.stundenplan24.de/10124219/vplanle/vdaten/VplanLe"
	}
	date := os.Getenv(dateEnv)
	if date == "" {
		today := time.Now()
		date = fmt.Sprintf("%d%02d%02d", today.Year(), int(today.Month()), today.Day())
	}
	if !dateFormat.MatchString(date) {
		log.Fatalln("incorrect date format. Please format as YYYYMMDD")
	}
	return ScheduleClientConfig{User: username, Password: password, BaseURL: baseURL, Date: date}
}

//GetTeacherCode - receive the Code from the environment Variable
func GetTeacherCode() string {
	return parseEnvMandatory(codeEnv)
}

func parseEnvMandatory(variableKEy string) string {
	variableValue := os.Getenv(variableKEy)
	if variableValue == "" {
		log.Fatalln("Environment variable: " + variableKEy + " is empty")
	}
	return variableValue
}
