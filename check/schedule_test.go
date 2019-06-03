package main

import "testing"

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

