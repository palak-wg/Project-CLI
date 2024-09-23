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

type PatientHandler struct {
	service interfaces.PatientService
}

func NewPatientHandler(service interfaces.PatientService) *PatientHandler {
	return &PatientHandler{service}
}

func (handler *PatientHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	var user models.Patient
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

	bearerToken := r.Header.Get("Authorization")
	claims, err := tokens.ExtractClaims(bearerToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusUnauthorized,
			Data:   http.StatusText(http.StatusUnauthorized),
		})
		loggerZap.Error("Extracting Claims")
	}
	user.UserID = claims["id"].(string)

	err = handler.service.UpdatePatientDetails(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusBadRequest,
			Data:   http.StatusText(http.StatusBadRequest),
		})
		loggerZap.Error("Error Updating profile")
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   fmt.Sprintf("Updated %s profile", user.Name),
	})
}
