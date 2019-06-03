package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

// CheckSchedule is our lambda handler invoked by the `lambda.Start` function call
func CheckSchedule(ctx context.Context) error {
	code := GetTeacherCode()
	config := CreateSchedulerConfigFromEnv()

	//HTTP Call
	scheduleClient := ScheduleClient{Config: config, Client: http.Client{}}
	xmlResponse, httpError := scheduleClient.RequestSchedule(date)
	if httpError != nil {
		log.Println("could not get schedule: " + httpError.Error())
		return httpError
	}
	//Response XML verarbeiten
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

//GetFutureWeekdays creates the Date-Strings for the schedule requests.
//Format of the Strings is YYYYDDMM. Weekends are skipped
func GetFutureWeekdays(today time.Time, intoFuture int) []string {
	futureWeekDays := make([]string, 0)
	nextDate := today
	counter := intoFuture
	for {
		nextDate = nextDate.AddDate(0, 0, 1)
		if nextDate.Weekday() == time.Saturday || nextDate.Weekday() == time.Sunday {
			continue
		}
		dateStr := fmt.Sprintf("%d%02d%02d", nextDate.Year(), int(nextDate.Month()), nextDate.Day())
		futureWeekDays = append(futureWeekDays, dateStr)

		counter--
		if counter < 1 {
			break
		}
	}
	return futureWeekDays
}
