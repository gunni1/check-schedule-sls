package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	//TODO: Events auf Basis der Analyseergebnisse erstellen
	var document Schedule
	xml.Unmarshal(xmlResponse, &document)
	log.Println("Prüfe Plan für: " + document.Head.Titel)
	log.Println("Suche nach Kürzel: " + code)
	if strings.Contains(document.Head.Info.ChangesTeacher, code+";") {
		log.Println("Änderungen gefunden!")
	} else {
		log.Println("Keine relevanten Änderungen.")
	}
	//Prüfen, ob Events bereits bekannt sind (db?)
	//Events publizieren

	return nil
}
