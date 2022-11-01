package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	// saving what stdout is before running a test
	stdOut := os.Stdout

	// creating out own stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	go printSomething("epsilon", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	// cast to string
	output := string(result)

	// set things back as before test
	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("expected to find epsilon, but it is not there")
	}

}
