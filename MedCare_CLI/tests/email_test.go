package tests

import (
	"bytes"
	"doctor-patient-cli/utils"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestSendEmail(t *testing.T) {
	tests := []struct {
		to      string
		subject string
		body    string
	}{
		{"test@example.com", "Test Subject", "This is a test body."},
		{"user@domain.com", "Welcome", "Welcome to our service!"},
		{"", "No Recipient", "This email has no recipient."},
		{"recipient@domain.com", "", "Empty subject line."},
		{"recipient@domain.com", "Empty Body", ""},
	}

	for _, test := range tests {
		t.Run(test.to, func(t *testing.T) {
			// Create a pipe to capture the output
			r, w, _ := os.Pipe()
			stdout := os.Stdout
			os.Stdout = w

			// Run the SendEmail function
			utils.SendEmail(test.to, test.subject, test.body)

			// Close the writer and restore stdout
			w.Close()
			os.Stdout = stdout

			// Read the captured output
			var output bytes.Buffer
			io.Copy(&output, r)

			// Expected output
			expectedOutput := fmt.Sprintf("Sending email to: %s\nSubject: %s\nBody: %s\n", test.to, test.subject, test.body)

			// Compare the output
			if output.String() != expectedOutput {
				t.Errorf("SendEmail(%s, %s, %s) produced output %q; expected %q",
					test.to, test.subject, test.body, output.String(), expectedOutput)
			}
		})
	}
}
