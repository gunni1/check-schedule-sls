package main

import (
	"log"
	"os"
)

const (
	userEnv    = "PAGE_USER"
	pwEnv      = "PAGE_PW"
	codeEnv    = "TEACHER_CODE"
	baseURLEnv = "BASE_URL"
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

func parseEnvMandatory(variableKEy string) string {
	variableValue := os.Getenv(variableKEy)
	if variableValue == "" {
		log.Fatalln("Environment variable: " + variableKEy + " is empty")
	}
	return variableValue
}
