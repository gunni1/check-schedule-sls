package main

import (
	b64 "encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

//ScheduleClientConfig stores everything needed to execute the schedule request
type ScheduleClientConfig struct {
	User     string
	Password string
	BaseURL  string
}

// ScheduleClient holds the http client for the http request execution and its config.
type ScheduleClient struct {
	Client http.Client
	Config ScheduleClientConfig
}

//RequestSchedule executes a http request to the provider of schedule changes
func (client ScheduleClient) RequestSchedule(date string) ([]byte, error) {
	requestURL := client.Config.BaseURL + date + ".xml"
	request, _ := http.NewRequest("GET", requestURL, nil)
	request.Header.Add("authorization", encodeAsBasicAuth(client.Config.User, client.Config.Password))
	log.Println("executing request to " + requestURL)
	resp, respErr := client.Client.Do(request)
	if respErr != nil {
		log.Println(respErr.Error())
		return nil, errors.New("HTTP Error: " + respErr.Error())
	}
	data, parseBodyErr := ioutil.ReadAll(resp.Body)
	if parseBodyErr == nil {
		return data, nil
	} else {
		return nil, errors.New("Parse Response Body Error: " + parseBodyErr.Error())
	}
}

func encodeAsBasicAuth(user string, password string) string {
	return "Basic " + b64.URLEncoding.EncodeToString([]byte(user+":"+password))
}
