package handlers

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

func (handler *AppointmentHandler) GetAppointments(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")
	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
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
	id := claims["id"].(string)
	role := claims["role"].(string)

	// If the role is 'patient', appointments shouldn't be shown to him
	if role == "patient" {
		w.WriteHeader(http.StatusForbidden)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusForbidden,
			Data:   "Access denied",
		})
		loggerZap.Error("Access denied for user")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	if role == "admin" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(models.APIResponse{
				Status: http.StatusBadRequest,
				Data:   http.StatusText(http.StatusBadRequest),
			})
			if err != nil {
				loggerZap.Error("Encoding response")
			}

			return
		}
		id = user.UserID
	}

	// Fetch the user profile
	appointments, err := handler.service.GetAppointmentsByDoctorID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching appointments",
		})
		loggerZap.Error("Error fetching appointments")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   appointments,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

func (handler *AppointmentHandler) GetPendingAppointments(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")
	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
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
	ID := claims["id"].(string)
	role := claims["role"].(string)

	// If the role is 'patient', appointments shouldn't be shown to him
	if role == "patient" {
		w.WriteHeader(http.StatusForbidden)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusForbidden,
			Data:   "Access denied",
		})
		loggerZap.Error("Access denied for user")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	if role == "admin" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(models.APIResponse{
				Status: http.StatusBadRequest,
				Data:   http.StatusText(http.StatusBadRequest),
			})
			if err != nil {
				loggerZap.Error("Encoding response")
			}

			return
		}
		ID = user.UserID
	}

	// Fetch the user profile
	appointments, err := handler.service.GetPendingAppointmentsByDoctorID(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching appointments",
		})
		loggerZap.Error("Error fetching appointments")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   appointments,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

func (handler *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
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
	id := claims["id"].(string)
	role := claims["role"].(string)

	var appointment models.Appointment
	err = json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error decoding request body",
		})
		loggerZap.Error("Decoding request body")
		if err != nil {
			loggerZap.Error("Encoding response")
		}

		return
	}

	if role == "patient" || role == "doctor" {
		if appointment.PatientID != id {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(models.APIResponse{
				Status: http.StatusBadRequest,
				Data:   http.StatusText(http.StatusBadRequest),
			})
			loggerZap.Error("Cannot create appointment")
			return
		}
	}

	err = handler.service.SendAppointmentRequest(appointment.PatientID, appointment.DoctorID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error sending appointment",
		})
		loggerZap.Error("Error sending appointment")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
	}
}

func (handler *AppointmentHandler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
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

	var appointment models.Appointment
	err = json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusBadRequest,
			Data:   http.StatusText(http.StatusBadRequest),
		})
		if err != nil {
			loggerZap.Error("Encoding response")
		}

		return
	}

	// Get user ID and role from claims
	role := claims["role"].(string)

	if role == "patient" {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusBadRequest,
			Data:   http.StatusText(http.StatusBadRequest),
		})
		loggerZap.Error("Cannot approve appointment")
		return
	}

	err = handler.service.ApproveAppointment(appointment.AppointmentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error approving appointment",
		})
		loggerZap.Error("Error approving appointment")
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   fmt.Sprintf("Appointment %v approved successfully", appointment.AppointmentID),
	})
}
