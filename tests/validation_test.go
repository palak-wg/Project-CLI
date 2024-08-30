package tests

import (
	"doctor-patient-cli/utils"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"invalid-email", false},
		{"user@domain.co", true},
		{"user@domain.c", false},
	}

	for _, test := range tests {
		t.Run(test.email, func(t *testing.T) {
			result := utils.ValidateEmail(test.email)
			if result != test.expected {
				t.Errorf("ValidateEmail(%s) = %v; expected %v", test.email, result, test.expected)
			}
		})
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	tests := []struct {
		phone    string
		expected bool
	}{
		{"+1234567890", true},
		{"1234567890", true},
		{"+123456789012345", true},
		{"12345", false},
		{"abc1234567", false},
	}

	for _, test := range tests {
		t.Run(test.phone, func(t *testing.T) {
			result := utils.ValidatePhoneNumber(test.phone)
			if result != test.expected {
				t.Errorf("ValidatePhoneNumber(%s) = %v; expected %v", test.phone, result, test.expected)
			}
		})
	}
}

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		username string
		expected bool
	}{
		{"JohnDoe", true},
		{"jd", false},
		{"John Doe", true},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.username, func(t *testing.T) {
			result := utils.ValidateUsername(test.username)
			if result != test.expected {
				t.Errorf("ValidateUsername(%s) = %v; expected %v", test.username, result, test.expected)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"Password1!", true},
		{"password", false},
		{"PASSWORD1", false},
		{"Pass1", false},
		{"Password!", false},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			result := utils.ValidatePassword(test.password)
			if result != test.expected {
				t.Errorf("ValidatePassword(%s) = %v; expected %v", test.password, result, test.expected)
			}
		})
	}
}

func TestValidateUserID(t *testing.T) {
	tests := []struct {
		userID   string
		expected bool
	}{
		{"User123", true},
		{"abc", true},
		{"ab", false},
		{"12345678901234567", false},
		{"user_123", false},
	}

	for _, test := range tests {
		t.Run(test.userID, func(t *testing.T) {
			result := utils.ValidateUserID(test.userID)
			if result != test.expected {
				t.Errorf("ValidateUserID(%s) = %v; expected %v", test.userID, result, test.expected)
			}
		})
	}
}

func TestValidateAge(t *testing.T) {
	tests := []struct {
		age      int
		expected bool
	}{
		{25, true},
		{-1, false},
		{0, true},
		{150, true},
	}

	for _, test := range tests {
		t.Run(string(rune(test.age)), func(t *testing.T) {
			result := utils.ValidateAge(test.age)
			if result != test.expected {
				t.Errorf("ValidateAge(%d) = %v; expected %v", test.age, result, test.expected)
			}
		})
	}
}

func TestValidateGender(t *testing.T) {
	tests := []struct {
		gender   string
		expected bool
	}{
		{"male", true},
		{"female", true},
		{"other", true},
		{"unknown", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.gender, func(t *testing.T) {
			result := utils.ValidateGender(test.gender)
			if result != test.expected {
				t.Errorf("ValidateGender(%s) = %v; expected %v", test.gender, result, test.expected)
			}
		})
	}
}

func TestValidateRole(t *testing.T) {
	tests := []struct {
		role     string
		expected bool
	}{
		{"patient", true},
		{"doctor", true},
		{"admin", false},
		{"guest", false},
	}

	for _, test := range tests {
		t.Run(test.role, func(t *testing.T) {
			result := utils.ValidateRole(test.role)
			if result != test.expected {
				t.Errorf("ValidateRole(%s) = %v; expected %v", test.role, result, test.expected)
			}
		})
	}
}
