package main

import (
	"encoding/xml"
	"errors"
	"hash/fnv"
	"strings"
)

//Schedule is the top level of the response xml
type Schedule struct {
	XMLName xml.Name `xml:"vp"`
	Head    Head     `xml:"kopf"`
}

//Head is some meta information about schedule changes
type Head struct {
	Titel      string `xml:"titel"`
	UploadDate string `xml:"datum"`
	Info       Info   `xml:"kopfinfo"`
}

//Info is called kopfinfo in xml. It holds overview information about current schedule changes
type Info struct {
	ChangesTeacher string `xml:"aenderungl"`
}

//ScheduleChange is a summary of data about a change in the schedule
type ScheduleChange struct {
	DateConcerned string `json:"dateConcerned"`
	IssueDate     string `json:"issueDate"`
	TeacherCode   string `json:"teacherCode"`
}

//FindChange checks the XML-Document for relevant schedule changes
func (schedule Schedule) FindChange(teacherCode string) (ScheduleChange, error) {
	if strings.Contains(schedule.Head.Info.ChangesTeacher, teacherCode+";") {
		return ScheduleChange{
			DateConcerned: schedule.Head.Titel,
			IssueDate:     schedule.Head.UploadDate,
			TeacherCode:   teacherCode,
		}, nil
	}
	return ScheduleChange{}, errors.New("no schedule change found for code: " + teacherCode)
}

//Hash the change as simple fnv
func (change ScheduleChange) Hash() uint32 {
	stringRep := change.DateConcerned + change.IssueDate + change.TeacherCode
	h := fnv.New32a()
	h.Write([]byte(stringRep))
	return h.Sum32()
}
