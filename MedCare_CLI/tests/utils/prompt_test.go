package utils

import (
	"doctor-patient-cli/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptPassword(t *testing.T) {
	// Save original stdin and stdout
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	// Create pipes to simulate user input and capture output
	r, w, _ := os.Pipe()
	os.Stdin = r
	os.Stdout = w

	// Simulate user input
	go func() {
		w.Write([]byte("\n"))
		w.Close()
	}()

	// Call the function to test
	password := utils.PromptPassword("Enter password:")

	// Check the result
	assert.Equal(t, "", password)
}
