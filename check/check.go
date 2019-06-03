package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

// CheckSchedule is our lambda handler invoked by the `lambda.Start` function call
func CheckSchedule(ctx context.Context) error {
	code := GetTeacherCode()
	config := CreateSchedulerConfigFromEnv()

	today := time.Now()
	date := fmt.Sprintf("%d%02d%02d", today.Year(), int(today.Month()), today.Day())

	scheduleClient := ScheduleClient{Config: config, Client: http.Client{}}
	xmlResponse, _ := scheduleClient.RequestSchedule(date)
	var schedule Schedule
	xml.Unmarshal(xmlResponse, &schedule)
	scheduleChange, err := schedule.FindChange(code)

	if err == nil {
		fmt.Println("No schedule changes found for teacher: " + code)
		return nil
	}

	hash := scheduleChange.Hash()
	fmt.Println("Found change, hashed: " + string(hash))
	//Prüfen, ob Events bereits bekannt sind (db?)
	//Events publizieren

	return nil
}
