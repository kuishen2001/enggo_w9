package main

import (
	"testing"
)

func TestQuery(t *testing.T) {
	app := NewApp()

	queryResult := app.Query("default")

	if queryResult != "Your query for 'default' returned: 'used within a switch statement to specify a block of code that should be executed when none of the case conditions match the value being switched'\n" {
		t.Errorf("Query result is incorrect: got %q, want %q", queryResult, "Your query for 'default' returned: 'used within a switch statement to specify a block of code that should be executed when none of the case conditions match the value being switched'\n")
	}
}
