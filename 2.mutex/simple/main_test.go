package main

import (
	"testing"
)

// go test -race .
// this fails
// you can check if race conditions are good like this
func Test_updateMessage(t *testing.T) {

	msg = "Hello tea"

	wg.Add(2)

	go updateMessage("Bad girl 2!", &mutex)
	go updateMessage("Bad girl!", &mutex)

	wg.Wait()

	if msg != "Bad girl 2!" {
		t.Errorf("Expected Bad girl!, but got %s", msg)
	}
}
