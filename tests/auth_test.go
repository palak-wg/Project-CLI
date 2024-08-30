package tests

import (
	"doctor-patient-cli/controllers"
	"doctor-patient-cli/models"
	"testing"
)

func TestSignup(t *testing.T) {
	// Example test case for signup
	controllers.Signup()
	user, _ := models.GetUserByID("testuser")
	if user.Username != "testusername" {
		t.Errorf("Expected testusername, got %s", user.Username)
	}
}

func TestLogin(t *testing.T) {
	// Example test case for login
	user := controllers.Login()
	if user.UserID != "testuser" {
		t.Errorf("Expected testuser, got %s", user.UserID)
	}
}
