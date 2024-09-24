package handlers

import (
	"bytes"
	"doctor-patient-cli/tests/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"doctor-patient-cli/handlers"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tokens"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReviewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockReviewService(ctrl)
	reviewHandler := handlers.NewReviewHandler(mockService)

	// Reset the claims extractor to the default before each test
	tokens.SetClaimsExtractor(tokens.ExtractClaims)

	t.Run("GetAllReview - Success", func(t *testing.T) {
		mockService.EXPECT().GetAllReviews().Return([]models.Review{
			{PatientID: "123", DoctorID: "1", Content: "Great doctor!", Rating: 5},
		}, nil)

		req := httptest.NewRequest("GET", "/reviews", nil)
		w := httptest.NewRecorder()

		reviewHandler.GetAllReview(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response models.APIResponse
		json.NewDecoder(res.Body).Decode(&response)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.NotEmpty(t, response.Data)
	})

	t.Run("GetAllReview - Internal Server Error", func(t *testing.T) {
		mockService.EXPECT().GetAllReviews().Return(nil, assert.AnError)

		req := httptest.NewRequest("GET", "/reviews", nil)
		w := httptest.NewRecorder()

		reviewHandler.GetAllReview(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

		var response models.APIResponse
		json.NewDecoder(res.Body).Decode(&response)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})

	t.Run("CreateReview - Success", func(t *testing.T) {
		review := models.Review{PatientID: "123", DoctorID: "1", Content: "Great doctor!", Rating: 5}
		body, _ := json.Marshal(review)
		req := httptest.NewRequest("POST", "/reviews", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer valid_token")
		w := httptest.NewRecorder()

		claims := map[string]interface{}{"id": "123", "role": "patient"}
		tokens.SetClaimsExtractor(func(token string) (map[string]interface{}, error) {
			return claims, nil
		})

		mockService.EXPECT().AddReview(&review).Return(nil)

		reviewHandler.CreateReview(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response models.APIResponse
		json.NewDecoder(res.Body).Decode(&response)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.Equal(t, "Successfully added review", response.Data)
	})

	t.Run("CreateReview - Invalid Token", func(t *testing.T) {
		review := models.Review{PatientID: "123", DoctorID: "1", Content: "Great doctor!", Rating: 5}
		body, _ := json.Marshal(review)
		req := httptest.NewRequest("POST", "/reviews", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer invalid_token")
		w := httptest.NewRecorder()

		// Set the claims extractor to return an error
		tokens.SetClaimsExtractor(func(token string) (map[string]interface{}, error) {
			return nil, assert.AnError // Simulate an error
		})

		// Call the handler
		reviewHandler.CreateReview(w, req)

		// Validate the response
		res := w.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

		// Ensure AddReview was not called
		mockService.EXPECT().AddReview(gomock.Any()).Times(0)
	})

	t.Run("CreateReview - Unauthorized User", func(t *testing.T) {
		review := models.Review{PatientID: "456", DoctorID: "1", Content: "Great doctor!", Rating: 5}
		body, _ := json.Marshal(review)
		req := httptest.NewRequest("POST", "/reviews", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer valid_token")
		w := httptest.NewRecorder()

		claims := map[string]interface{}{"id": "123", "role": "patient"}
		tokens.SetClaimsExtractor(func(token string) (map[string]interface{}, error) {
			return claims, nil
		})

		reviewHandler.CreateReview(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	})

	t.Run("GetDoctorSpecificReviews - Success", func(t *testing.T) {
		mockService.EXPECT().GetReviewsByDoctorID("").Return([]models.Review{
			{PatientID: "123", DoctorID: "1", Content: "Great doctor!", Rating: 5},
		}, nil)

		req := httptest.NewRequest("GET", "/reviews/1", nil)
		w := httptest.NewRecorder()

		reviewHandler.GetDoctorSpecificReviews(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response models.APIResponse
		json.NewDecoder(res.Body).Decode(&response)
		assert.Equal(t, http.StatusOK, response.Status)
		assert.NotEmpty(t, response.Data)
	})

	t.Run("GetDoctorSpecificReviews - Internal Server Error", func(t *testing.T) {
		mockService.EXPECT().GetReviewsByDoctorID("").Return(nil, assert.AnError)

		req := httptest.NewRequest("GET", "/reviews/1", nil)
		w := httptest.NewRecorder()

		reviewHandler.GetDoctorSpecificReviews(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

		var response models.APIResponse
		json.NewDecoder(res.Body).Decode(&response)
		assert.Equal(t, http.StatusInternalServerError, response.Status)
	})
}
