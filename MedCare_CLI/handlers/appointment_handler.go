package handlers

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/middlewares"
	"doctor-patient-cli/models"
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

	role, _ := r.Context().Value(middlewares.RoleKey).(string)
	id, _ := r.Context().Value(middlewares.UserIdKey).(string)

	// If the role is 'patient', appointments shouldn't be shown to him
	if role == "patient" {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusForbidden,
			Data:   "Access denied",
		})
		loggerZap.Error("Access denied for user")
		return
	}

	if role == "admin" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(models.APIResponse{
				Status: http.StatusBadRequest,
				Data:   http.StatusText(http.StatusBadRequest),
			})
			return
		}
		id = user.UserID
	}

	// Fetch the user profile
	appointments, err := handler.service.GetAppointmentsByDoctorID(id)

	approve := r.URL.Query().Get("approve")

	if approve != "" {
		appointments, err = handler.service.GetPendingAppointmentsByDoctorID(id)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching appointments",
		})
		loggerZap.Error("Error fetching appointments")
		return
	}

	// Prepare response
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   appointments,
	})
}

func (handler *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	role, _ := r.Context().Value(middlewares.RoleKey).(string)
	id, _ := r.Context().Value(middlewares.UserIdKey).(string)

	var appointment models.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error decoding request body",
		})
		loggerZap.Error("Decoding request body")
		return
	}

	if role == "patient" || role == "doctor" {
		if appointment.PatientID != id {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(models.APIResponse{
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
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error sending appointment",
		})
		loggerZap.Error("Error sending appointment")
		return
	}

	// Add success response here
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   "Appointment created successfully",
	})
}

func (handler *AppointmentHandler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	role, _ := r.Context().Value(middlewares.RoleKey).(string)

	var appointment models.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusBadRequest,
			Data:   http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if role == "patient" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusBadRequest,
			Data:   http.StatusText(http.StatusBadRequest),
		})
		loggerZap.Error("Cannot approve appointment")
		return
	}

	err = handler.service.ApproveAppointment(appointment.AppointmentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(models.APIResponse{
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
