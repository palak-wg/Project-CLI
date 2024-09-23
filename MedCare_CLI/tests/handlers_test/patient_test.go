package handlers_test

import (
	"bytes"
	"doctor-patient-cli/handlers"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tests/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type ReadCloser struct {
	*bytes.Reader
}

func (rc *ReadCloser) Close() error {
	return nil
}

func generateToken(userID string, role string) string {
	claims := jwt.MapClaims{
		"id":   userID,
		"role": role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secret")) // Use your actual secret here
	return signedToken
}

func TestPatientHandler_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPatientService(ctrl)
	handler := handlers.NewPatientHandler(mockService)

	t.Run("Success - Profile updated", func(t *testing.T) {
		mockService.EXPECT().UpdatePatientDetails(gomock.Any()).Return(nil)

		user := models.User{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest("PUT", "/patients/profile", &ReadCloser{bytes.NewReader(reqBody)})
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient"))

		rr := httptest.NewRecorder()

		handler.UpdateProfile(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "Updated John Doe profile", response.Data)
	})

	t.Run("Bad Request - Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest("PUT", "/patients/profile", &ReadCloser{bytes.NewReader([]byte("invalid_json"))})
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient"))

		rr := httptest.NewRecorder()

		handler.UpdateProfile(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
	})

	t.Run("Unauthorized - Invalid Token", func(t *testing.T) {
		user := models.User{
			Name:  "Jane Doe",
			Email: "jane@example.com",
		}

		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest("PUT", "/patients/profile", &ReadCloser{bytes.NewReader(reqBody)})
		req.Header.Set("Authorization", "Bearer invalid_token")

		rr := httptest.NewRecorder()

		handler.UpdateProfile(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, response.Status)
	})

	t.Run("Service Error - Unable to Update Profile", func(t *testing.T) {
		mockService.EXPECT().UpdatePatientDetails(gomock.Any()).Return(errors.New("service error"))

		user := models.User{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		reqBody, _ := json.Marshal(user)
		req := httptest.NewRequest("PUT", "/patients/profile", &ReadCloser{bytes.NewReader(reqBody)})
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient"))

		rr := httptest.NewRecorder()

		handler.UpdateProfile(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
	})

}
