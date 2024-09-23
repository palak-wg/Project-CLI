package handlers_test

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tokens"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type AppointmentHandler struct {
	service interfaces.AppointmentService
}

func NewAppointmentHandler(service interfaces.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service}
}

// encodeResponse centralizes response encoding and error handling
func encodeResponse(w http.ResponseWriter, status int, data interface{}, logger *zap.Logger) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(models.APIResponse{
		Status: status,
		Data:   data,
	})
	if err != nil {
		logger.Error("Error encoding response", zap.Error(err))
	}
}

// GetAppointments handles retrieving appointments for the doctor or admin
func (handler *AppointmentHandler) GetAppointments(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	w.Header().Set("Content-Type", "application/json")
	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
	if err != nil {
		logger.Error("Error extracting claims", zap.Error(err))
		encodeResponse(w, http.StatusInternalServerError, "Error extracting claims", logger)
		return
	}

	// Get user ID and role from claims
	id := claims["id"].(string)
	role := claims["role"].(string)

	if role == "patient" {
		logger.Error("Access denied for patient")
		encodeResponse(w, http.StatusForbidden, "Access denied", logger)
		return
	}

	if role == "admin" {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			logger.Error("Bad request", zap.Error(err))
			encodeResponse(w, http.StatusBadRequest, "Invalid request", logger)
			return
		}
		id = user.UserID
	}

	appointments, err := handler.service.GetAppointmentsByDoctorID(id)
	if err != nil {
		logger.Error("Error fetching appointments", zap.Error(err))
		encodeResponse(w, http.StatusInternalServerError, "Error fetching appointments", logger)
		return
	}

	encodeResponse(w, http.StatusOK, appointments, logger)
}

// GetPendingAppointments handles fetching pending appointments for the doctor or admin
func (handler *AppointmentHandler) GetPendingAppointments(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	w.Header().Set("Content-Type", "application/json")
	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
	if err != nil {
		logger.Error("Error extracting claims", zap.Error(err))
		encodeResponse(w, http.StatusInternalServerError, "Error extracting claims", logger)
		return
	}

	// Get user ID and role from claims
	id := claims["id"].(string)
	role := claims["role"].(string)

	if role == "patient" {
		logger.Error("Access denied for patient")
		encodeResponse(w, http.StatusForbidden, "Access denied", logger)
		return
	}

	if role == "admin" {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			logger.Error("Bad request", zap.Error(err))
			encodeResponse(w, http.StatusBadRequest, "Invalid request", logger)
			return
		}
		id = user.UserID
	}

	appointments, err := handler.service.GetPendingAppointmentsByDoctorID(id)
	if err != nil {
		logger.Error("Error fetching appointments", zap.Error(err))
		encodeResponse(w, http.StatusInternalServerError, "Error fetching appointments", logger)
		return
	}

	encodeResponse(w, http.StatusOK, appointments, logger)
}

// CreateAppointment handles creating a new appointment
func (handler *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	w.Header().Set("Content-Type", "application/json")

	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
	if err != nil {
		logger.Error("Error extracting claims", zap.Error(err))
		encodeResponse(w, http.StatusInternalServerError, "Error extracting claims", logger)
		return
	}

	// Get user ID and role from claims
	id := claims["id"].(string)
	role := claims["role"].(string)

	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		logger.Error("Bad request", zap.Error(err))
		encodeResponse(w, http.StatusBadRequest, "Invalid request", logger)
		return
	}

	if role == "patient" && appointment.PatientID != id {
		encodeResponse(w, http.StatusBadRequest, "Cannot create appointment for another patient", logger)
		return
	}

	err = handler.service.SendAppointmentRequest(appointment.PatientID, appointment.DoctorID)
	if err != nil {
		logger.Error("Error sending appointment", zap.Error(err))
		encodeResponse(w, http.StatusInternalServerError, "Error sending appointment", logger)
		return
	}

	encodeResponse(w, http.StatusOK, "Appointment created successfully", logger)
}

// UpdateAppointment handles approving an appointment
func (handler *AppointmentHandler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	w.Header().Set("Content-Type", "application/json")

	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
	if err != nil {
		logger.Error("Error extracting claims", zap.Error(err))
		encodeResponse(w, http.StatusInternalServerError, "Error extracting claims", logger)
		return
	}

	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		logger.Error("Bad request", zap.Error(err))
		encodeResponse(w, http.StatusBadRequest, "Invalid request", logger)
		return
	}

	// Get role from claims
	role := claims["role"].(string)

	if role == "patient" {
		encodeResponse(w, http.StatusForbidden, "Access denied for patient", logger)
		return
	}

	err = handler.service.ApproveAppointment(appointment.AppointmentID)
	if err != nil {
		logger.Error("Error approving appointment", zap.Error(err))
		encodeResponse(w, http.StatusInternalServerError, "Error approving appointment", logger)
		return
	}

	encodeResponse(w, http.StatusOK, fmt.Sprintf("Appointment %v approved successfully", appointment.AppointmentID), logger)
}
