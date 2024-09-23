package handlers

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tokens"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type ReviewHandler struct {
	service interfaces.ReviewService
}

func NewReviewHandler(service interfaces.ReviewService) *ReviewHandler {
	return &ReviewHandler{service}
}

func (handler *ReviewHandler) GetAllReview(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Fetch all reviews
	reviews, err := handler.service.GetAllReviews()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching reviews",
		})
		loggerZap.Error("Internal Server Error fetching reviews")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}
	loggerZap.Info("Successfully fetched reviews")

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   reviews,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

func (handler *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Get review from body
	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error decoding body",
		})
		loggerZap.Error("Internal Server Error decoding body")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	// Checking the claims to the info provided
	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.GetClaims(bearerToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error extracting claims",
		})
		loggerZap.Error("Internal Server Error")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	// Get user ID and role from claims
	tokenID := claims["id"].(string)
	role := claims["role"].(string)

	if (tokenID != review.PatientID) || (role != "patient") {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusUnauthorized,
			Data:   "Can't create review",
		})
		loggerZap.Error("Can't create review by being someone else")
		return
	}

	// Adding review to doctorID
	err = handler.service.AddReview(&review)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error adding review",
		})
		loggerZap.Error("Internal Server Error adding review")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}
	loggerZap.Info("Successfully added review")

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   "Successfully added review",
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

func (handler *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {}

func (handler *ReviewHandler) GetDoctorSpecificReviews(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Fetch all reviews
	vars := mux.Vars(r)
	doctorID := vars["doctor_id"]
	reviews, err := handler.service.GetReviewsByDoctorID(doctorID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching reviews",
		})
		loggerZap.Error("Internal Server Error fetching reviews")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}
	loggerZap.Info("Successfully fetched reviews")

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   reviews,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}
