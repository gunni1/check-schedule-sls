package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

//CheckScheduleResult is used to have a multiple return values in a channel
type CheckScheduleResult struct {
	ScheduleChange ScheduleChange
	Error          error
}

// CheckSchedule is our lambda handler invoked by the `lambda.Start` function call
func CheckSchedule(ctx context.Context) error {
	code := GetTeacherCode()
	config := CreateSchedulerConfigFromEnv()
	daysCount := GetDaysCount()

	notificator := CreateTelegramNotificator(GetBotToken(), GetNotificationTarget())

	//TODO: Anzahl der Tage wird ebenfalls ein Env-Parameter. Für Tests erstmal fix 2
	daysToCheck := GetFutureWeekdays(time.Now(), daysCount)

	checkResults := make(chan CheckScheduleResult, len(daysToCheck))
	//TODO: Pipelines einsetzen? Methode Liefert channel mit Ergebnis zum weiterarbeiten.
	// -> Es muss nicht auf die Synchronisation aller gewartet werden
	for _, day := range daysToCheck {
		go CheckScheduleChange(config, code, day, checkResults)
	}
	changes := make([]ScheduleChange, 0)
	for i := 0; i < len(daysToCheck); i++ {
		checkResult := <-checkResults
		if checkResult.Error == nil {
			changes = append(changes, checkResult.ScheduleChange)
		}
	}
	close(checkResults)

	fmt.Printf("changes: %v\n", changes)
	//TODO: Prüfen, ob Events bereits bekannt sind (db?)
	//hash := scheduleChange.Hash()
	//fmt.Println("Found change, hashed: " + string(hash))

	for _, change := range changes {
		go notificator.SendNotification(change)
	}

	return nil
}

//CheckScheduleChange is made for running in a goroutine.
//It Made a HTTP-Request to the provider of schedule changes and search the result for a specific teacher code
func CheckScheduleChange(config ScheduleClientConfig, teacherCode string, date string, changes chan CheckScheduleResult) {
	scheduleClient := ScheduleClient{Config: config, Client: http.Client{}}
	xmlResponse, httpError := scheduleClient.RequestSchedule(date)
	if httpError != nil {
		log.Println("could not get schedule: " + httpError.Error())
		changes <- CheckScheduleResult{Error: httpError}
		return
	}
	var schedule Schedule
	xml.Unmarshal(xmlResponse, &schedule)
	maybeChange, err := schedule.FindChange(teacherCode)
	changes <- CheckScheduleResult{ScheduleChange: maybeChange, Error: err}
}

//GetFutureWeekdays creates the Date-Strings for the schedule requests.
//Format of the Strings is YYYYDDMM. Weekends are skipped. ALways has at least one Result (next weekday).
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
