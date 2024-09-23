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

	_ "github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReviewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockReviewService(ctrl)
	handler := handlers.NewReviewHandler(mockService)

	t.Run("GetAllReview - Success", func(t *testing.T) {
		reviews := []models.Review{{DoctorID: "1", PatientID: "123", Content: "Great doctor!"}}
		mockService.EXPECT().GetAllReviews().Return(reviews, nil)

		req := httptest.NewRequest("GET", "/reviews", nil)
		rr := httptest.NewRecorder()

		handler.GetAllReview(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, reviews, response.Data)
	})

	t.Run("GetAllReview - Internal Server Error", func(t *testing.T) {
		mockService.EXPECT().GetAllReviews().Return(nil, errors.New("service error"))

		req := httptest.NewRequest("GET", "/reviews", nil)
		rr := httptest.NewRecorder()

		handler.GetAllReview(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})

	t.Run("CreateReview - Success", func(t *testing.T) {
		review := models.Review{DoctorID: "1", PatientID: "123", Content: "Great doctor!"}
		mockService.EXPECT().AddReview(gomock.Any()).Return(nil)

		reqBody, _ := json.Marshal(review)
		req := httptest.NewRequest("POST", "/reviews", &ReadCloser{bytes.NewReader(reqBody)})
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient"))

		rr := httptest.NewRecorder()

		handler.CreateReview(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "Successfully added review", response.Data)
	})

	t.Run("CreateReview - Invalid Token", func(t *testing.T) {
		review := models.Review{DoctorID: "1", PatientID: "123", Content: "Great doctor!"}
		reqBody, _ := json.Marshal(review)
		req := httptest.NewRequest("POST", "/reviews", &ReadCloser{bytes.NewReader(reqBody)})
		req.Header.Set("Authorization", "Bearer invalid_token")

		rr := httptest.NewRecorder()

		handler.CreateReview(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})

	t.Run("CreateReview - Unauthorized User", func(t *testing.T) {
		review := models.Review{DoctorID: "1", PatientID: "456", Content: "Great doctor!"} // Different PatientID
		reqBody, _ := json.Marshal(review)
		req := httptest.NewRequest("POST", "/reviews", &ReadCloser{bytes.NewReader(reqBody)})
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient"))

		rr := httptest.NewRecorder()

		handler.CreateReview(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, response.Status)
	})

	t.Run("CreateReview - Service Error", func(t *testing.T) {
		review := models.Review{DoctorID: "1", PatientID: "123", Content: "Great doctor!"}
		mockService.EXPECT().AddReview(gomock.Any()).Return(errors.New("service error"))

		reqBody, _ := json.Marshal(review)
		req := httptest.NewRequest("POST", "/reviews", &ReadCloser{bytes.NewReader(reqBody)})
		req.Header.Set("Authorization", "Bearer "+generateToken("123", "patient"))

		rr := httptest.NewRecorder()

		handler.CreateReview(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})

	t.Run("GetDoctorSpecificReviews - Success", func(t *testing.T) {
		reviews := []models.Review{{DoctorID: "1", PatientID: "123", Content: "Great doctor!"}}
		mockService.EXPECT().GetReviewsByDoctorID("1").Return(reviews, nil)

		req := httptest.NewRequest("GET", "/reviews/1", nil)
		rr := httptest.NewRecorder()

		handler.GetDoctorSpecificReviews(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, reviews, response.Data)
	})

	t.Run("GetDoctorSpecificReviews - Internal Server Error", func(t *testing.T) {
		mockService.EXPECT().GetReviewsByDoctorID("1").Return(nil, errors.New("service error"))

		req := httptest.NewRequest("GET", "/reviews/1", nil)
		rr := httptest.NewRecorder()

		handler.GetDoctorSpecificReviews(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		var response models.APIResponse
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})
}
