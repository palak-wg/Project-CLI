package utils

import (
	"doctor-patient-cli/utils"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		password string
	}{
		{"Password123!"},
		{"simplePass"},
		{"another$Tr0ngPass"},
		{"12345678"},
		{""},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			hash := utils.HashPassword(test.password)
			if hash == "" {
				t.Errorf("HashPassword(%s) returned an empty hash", test.password)
			}
			if !utils.CheckPasswordHash(test.password, hash) {
				t.Errorf("CheckPasswordHash(%s, %s) = false; expected true", test.password, hash)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	tests := []struct {
		password  string
		incorrect string
	}{
		{"Password123!", "WrongPassword"},
		{"simplePass", "SimplePass"},
		{"another$Tr0ngPass", "anotherStrongPass"},
		{"12345678", "87654321"},
		{"", "notEmpty"},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			hash := utils.HashPassword(test.password)
			if !utils.CheckPasswordHash(test.password, hash) {
				t.Errorf("CheckPasswordHash(%s, %s) = false; expected true", test.password, hash)
			}
			if utils.CheckPasswordHash(test.incorrect, hash) {
				t.Errorf("CheckPasswordHash(%s, %s) = true; expected false", test.incorrect, hash)
			}
		})
	}
}
