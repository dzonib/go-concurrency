package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	wg.Add(1)

	go updateMessage("hello thor", &wg)

	wg.Wait()

	if msg != "hello thor" {
		t.Errorf("Expected hello thor, but got %s", msg)
	}
}

func Test_printMessage(t *testing.T) {
	// saving what stdout is before running a test
	stdOut := os.Stdout

	// creating out own stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	msg = "king kong"

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	// cast to string
	output := string(result)

	// set things back as before test
	os.Stdout = stdOut

	if !strings.Contains(output, "king kong") {
		t.Errorf("expected to find %s, but got %s", msg, output)
	}
}
