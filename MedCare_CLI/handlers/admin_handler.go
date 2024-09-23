package handlers

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type AdminHandler struct {
	service interfaces.AdminService
}

func NewAdminHandler(service interfaces.AdminService) *AdminHandler {
	return &AdminHandler{service}
}

func (handler *AdminHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Fetch all user profiles
	users, err := handler.service.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching user profiles",
		})
		loggerZap.Error("Internal Server Error fetching user profiles")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	var responseUsers []models.APIResponseUser
	for _, user := range users {
		responseUsers = append(responseUsers, models.APIResponseUser{
			UserID:      user.UserID,
			Name:        user.Name,
			Age:         user.Age,
			Gender:      user.Gender,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		})
	}

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   responseUsers,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

func (handler *AdminHandler) GetPendingDoctors(w http.ResponseWriter, request *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Fetch all pending profiles
	pendingDoctors, err := handler.service.GetPendingDoctorRequests()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching pending requests",
		})
		loggerZap.Error("Internal Server Error fetching signup requests")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	var responseUsers []models.APIResponsePendingSignup
	for _, pendingDoctor := range pendingDoctors {
		responseUsers = append(responseUsers, models.APIResponsePendingSignup{
			ID:   pendingDoctor.UserID,
			Name: pendingDoctor.Name,
		})
	}
	loggerZap.Info("Fetched Pending IDs")

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   responseUsers,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

func (handler *AdminHandler) ApprovePendingDoctors(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Get user_id from body
	var doctor models.APIResponsePendingSignup
	err := json.NewDecoder(r.Body).Decode(&doctor)
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

	// Approving doctor signup for doctorID
	err = handler.service.ApproveDoctorSignup(doctor.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error approving doctor profile for signup",
		})
		loggerZap.Error("Internal Server Error approving doctor profile")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	loggerZap.Info(fmt.Sprintf("Doctor %s signup has been approved", doctor.ID))

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   fmt.Sprintf("Approved doctor %s signup", doctor.ID),
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}
