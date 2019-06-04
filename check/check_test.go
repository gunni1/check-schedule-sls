package main

import (
	"testing"
	"time"
)

var (
	friday, _   = time.Parse(time.RFC3339, "2019-06-07T22:08:41+00:00")
	monday, _   = time.Parse(time.RFC3339, "2019-06-03T22:08:41+00:00")
	thursday, _ = time.Parse(time.RFC3339, "2019-06-06T22:08:41+00:00")
)

func TestGetFutureWeekdays(t *testing.T) {
	tables := []struct {
		ExecutionDate  time.Time
		DaysIntoFuture int
		ExpectedResult []string
	}{
		{friday, 1, []string{"20190610"}},
		{monday, 3, []string{"20190604", "20190605", "20190606"}},
		{thursday, 2, []string{"20190607", "20190610"}},
		{friday, 0, []string{"20190610"}},
		{friday, -1, []string{"20190610"}},
	}
	for _, table := range tables {
		result := GetFutureWeekdays(table.ExecutionDate, table.DaysIntoFuture)
		if !testEq(result, table.ExpectedResult) {
			t.Errorf("\nInput: %v, steps: %v\nExpected: \n%v\nGot:\n%v",
				table.ExecutionDate, table.DaysIntoFuture, table.ExpectedResult, result)
		}
	}
}

func testEq(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
