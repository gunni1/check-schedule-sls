package main

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var scheduleXML = `<vp>
<kopf>
	<titel>Donnerstag, 13. Juni 2019 (A-Woche)</titel>
	<datum>12.06.2019, 09:23</datum>
	<kopfinfo>
	<aenderungl>AA; bd; BS; DD; KF; LO; PI; SC; SR; UH</aenderungl>
	</kopfinfo>
 </kopf>
<haupt>
	<aktion>
		<stunde>3-4</stunde>
		<fach>EN</fach>
		<lehrer>BB</lehrer>
		<klasse>.7-1</klasse>
		<vfach fageaendert="ae">CH</vfach>
		<vlehrer legeaendert="ae">FF</vlehrer>
		<vraum rageaendert="ae">215</vraum>
		<info>geändert</info>
	</aktion>
	<aktion>
		<stunde>5</stunde>
		<fach>MA</fach>
		<lehrer>AA</lehrer>
		<klasse>.5-3</klasse>
		<vfach fageaendert="ae">MU</vfach>
		<vlehrer legeaendert="ae">DD</vlehrer>
		<vraum rageaendert="ae">011</vraum>
		<info>geändert</info>
	</aktion>
 </haupt>
 </vp>
 `

func TestFindScheduleChangeDetails(t *testing.T) {
	var schedule Schedule
	xml.Unmarshal([]byte(scheduleXML), &schedule)
	fmt.Printf("parsed: %v", schedule)

	change, err := schedule.FindChange("AA")
	if err != nil {
		t.Errorf("XML unmarshaling error: %v", err.Error())
	}
	if len(change.Descriptions) != 1 {
		t.Errorf("Expected exact one description: %v", change.Descriptions)
	}
	descr := change.Descriptions[0]
	if descr.Stunde != "5" {
		t.Errorf("Stunde 5 erwartet, erhalten: %v", descr.Stunde)
	}
	if descr.Fach != "MA" {
		t.Errorf("MA erwartet: %v", descr.Fach)
	}
	if descr.Klasse != ".5-3" {
		t.Errorf("Klasse .5-3 erwartet: %v", descr.Klasse)
	}
	if descr.VertFach != "MU" {
		t.Errorf("Vert Fach MU erwartet: %v", descr.VertFach)
	}
	if descr.VertLehrer != "DD" {
		t.Errorf("Vert Lehrer DD erwartet: %v", descr.VertLehrer)
	}
	if descr.VertRaum != "011" {
		t.Errorf("Vert Raum 011 erwartet: %v", descr.VertRaum)
	}
}

func TestHashUsesDateConcern(t *testing.T) {
	change1 := ScheduleChange{
		DateConcerned: "aaa",
		IssueDate:     "bbb",
		TeacherCode:   "ccc",
	}
	change2 := ScheduleChange{
		DateConcerned: "aaaa",
		IssueDate:     "bbb",
		TeacherCode:   "ccc",
	}
	if change1.Hash() == change2.Hash() {
		t.Error("hash calculation not uses dateConcern")
	}
}

func TestHashUsesIssueDate(t *testing.T) {
	change1 := ScheduleChange{
		DateConcerned: "aaa",
		IssueDate:     "bbb",
		TeacherCode:   "ccc",
	}
	change2 := ScheduleChange{
		DateConcerned: "aaa",
		IssueDate:     "bbbb",
		TeacherCode:   "ccc",
	}
	if change1.Hash() == change2.Hash() {
		t.Error("hash calculation not uses IssueDate")
	}
}

func TestHashUsesTeacherCode(t *testing.T) {
	change1 := ScheduleChange{
		DateConcerned: "aaa",
		IssueDate:     "bbb",
		TeacherCode:   "ccc",
	}
	change2 := ScheduleChange{
		DateConcerned: "aaa",
		IssueDate:     "bbb",
		TeacherCode:   "cccc",
	}
	if change1.Hash() == change2.Hash() {
		t.Error("hash calculation not uses TeacherCode")
	}
}
