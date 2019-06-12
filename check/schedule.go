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
	Body    Body     `xml:"haupt"`
}

//Head is some meta information about schedule changes
type Head struct {
	Titel      string `xml:"titel"`
	UploadDate string `xml:"datum"`
	Info       Info   `xml:"kopfinfo"`
}

//Body - is the main part of the schedule change list
type Body struct {
	Description []Description `xml:"aktion"`
}

//Description - is a detail description of a schedule change
type Description struct {
	Stunde     string `xml:"stunde"`
	Fach       string `xml:"fach"`
	Lehrer     string `xml:"lehrer"`
	Klasse     string `xml:"klasse"`
	VertFach   string `xml:"vfach"`
	VertLehrer string `xml:"vlehrer"`
	VertRaum   string `xml:"vraum"`
	Info       string `xml:"info"`
}

//Info is called kopfinfo in xml. It holds overview information about current schedule changes
type Info struct {
	ChangesTeacher string `xml:"aenderungl"`
}

//ScheduleChange is a summary of data about a change in the schedule
type ScheduleChange struct {
	DateConcerned string
	IssueDate     string
	TeacherCode   string
	Descriptions  []Description
}

//FindChange checks the XML-Document for relevant schedule changes
func (schedule Schedule) FindChange(teacherCode string) (ScheduleChange, error) {
	if strings.Contains(schedule.Head.Info.ChangesTeacher, teacherCode+";") {
		return ScheduleChange{
			DateConcerned: schedule.Head.Titel,
			IssueDate:     schedule.Head.UploadDate,
			TeacherCode:   teacherCode,
			Descriptions:  schedule.findDescriptions(teacherCode),
		}, nil
	}
	return ScheduleChange{}, errors.New("no schedule change found for code: " + teacherCode)
}

//Look in all descriptions for a teacher code and return the results
//where the code is either in "lehrer" or in "vlehrer"
func (schedule Schedule) findDescriptions(teacherCode string) []Description {
	descriptions := make([]Description, 0)
	for _, descr := range schedule.Body.Description {
		if strings.Contains(descr.Lehrer, teacherCode) || strings.Contains(descr.VertLehrer, teacherCode) {
			descriptions = append(descriptions, descr)
		}
	}
	return descriptions
}

//Hash the change as simple fnv
func (change ScheduleChange) Hash() uint32 {
	stringRep := change.DateConcerned + change.IssueDate + change.TeacherCode
	h := fnv.New32a()
	h.Write([]byte(stringRep))
	return h.Sum32()
}
