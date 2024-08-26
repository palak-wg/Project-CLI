package tests

import (
	"bytes"
	"doctor-patient-cli/cmd"
	"doctor-patient-cli/controllers"
	"doctor-patient-cli/models"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartApp(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput []string
		mockSetup      func()
	}{
		{
			name:  "Exit Option",
			input: "3\n",
			expectedOutput: []string{
				"Welcome To The MedCare",
				"Please choose an option",
				"Enter your choice:",
			},
			mockSetup: func() {},
		},
		{
			name:  "Invalid Option",
			input: "invalid\n3\n",
			expectedOutput: []string{
				"Invalid choice. Please try again.",
				"Welcome To The MedCare",
				"Please choose an option",
			},
			mockSetup: func() {},
		},
		{
			name:  "Signup Option",
			input: "2\n3\n",
			expectedOutput: []string{
				"Signup function called",
			},
			mockSetup: func() {
				// Mock the Signup function
				oldSignup := controllers.Signup
				controllers.Signup = func() {
					fmt.Println("Signup function called")
				}
				t.Cleanup(func() { controllers.Signup = oldSignup }) // Reset after the test
			},
		},
		{
			name:  "Login as Admin",
			input: "1\n3\n",
			expectedOutput: []string{
				"Admin Menu function called",
			},
			mockSetup: func() {
				// Mock the Login and AdminMenu functions
				oldLogin := controllers.Login
				oldAdminMenu := controllers.AdminMenu

				controllers.Login = func() models.User {
					return models.User{UserType: "admin"}
				}
				controllers.AdminMenu = func(user models.User) {
					fmt.Println("Admin Menu function called")
				}

				t.Cleanup(func() {
					controllers.Login = oldLogin
					controllers.AdminMenu = oldAdminMenu
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock user input and capture output
			input := bytes.NewBufferString(tt.input)
			output := &bytes.Buffer{}
			oldStdin := os.Stdin
			oldStdout := os.Stdout
			os.Stdin = input
			os.Stdout = output
			defer func() {
				os.Stdin = oldStdin
				os.Stdout = oldStdout
			}()

			// Setup mocks
			tt.mockSetup()

			// Run the app
			main.StartApp()

			// Check output
			for _, expected := range tt.expectedOutput {
				assert.Contains(t, output.String(), expected)
			}
		})
	}
}
