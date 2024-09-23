package handlers

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type DoctorHandler struct {
	service interfaces.DoctorService
}

func (handler DoctorHandler) GetDoctors(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Fetch all doctor profiles
	doctors, err := handler.service.GetAllDoctors()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error fetching doctors profiles",
		})
		loggerZap.Error("Internal Server Error fetching doctors profiles")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}
	loggerZap.Info("Fetched all doctor profiles")

	var responseDoctors []models.APIResponseDoctor
	for _, doctor := range doctors {
		responseDoctors = append(responseDoctors, models.APIResponseDoctor{
			DoctorID:       doctor.UserID,
			Specialization: doctor.Specialization,
			Experience:     doctor.Experience,
			Rating:         doctor.Rating,
		})
	}

	// Prepare response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   responseDoctors,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

func (handler DoctorHandler) GetDoctor(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Fetch all doctor profiles
	doctor, err := handler.service.GetDoctorByID(mux.Vars(r)["doctor_id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusNotFound,
			Data:   http.StatusText(http.StatusNotFound),
		})
		loggerZap.Error("User not found")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	// Prepare response
	response := &models.APIResponseDoctor{
		DoctorID:       doctor.UserID,
		Specialization: doctor.Specialization,
		Experience:     doctor.Experience,
		Rating:         doctor.Rating,
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   response,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}

}

func NewDoctorHandler(service interfaces.DoctorService) *DoctorHandler {
	return &DoctorHandler{service}
}
