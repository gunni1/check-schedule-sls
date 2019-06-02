package main

import (
	"context"
	"encoding/xml"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) error {
	code := GetTeacherCode()
	config := CreateSchedulerConfigFromEnv()

	xmlResponse, _ := RequestSchedule(config)
	var document Schedule
	xml.Unmarshal(xmlResponse, &document)
	log.Println("Prüfe Plan für: " + document.Head.Titel)
	log.Println("Suche nach Kürzel: " + code)
	if strings.Contains(document.Head.Info.ChangesTeacher, code+";") {
		log.Println("Änderungen gefunden!")
	} else {
		log.Println("Keine relevanten Änderungen.")
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
