package utils

import (
	"fmt"
)

func SendEmail(to, subject, body string) {
	fmt.Printf("Sending email to: %s\nSubject: %s\nBody: %s\n", to, subject, body)
}
