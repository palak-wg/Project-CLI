package handlers

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/middlewares"
	"doctor-patient-cli/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type NotificationHandler struct {
	service interfaces.NotificationService
}

func NewNotificationHandler(service interfaces.NotificationService) *NotificationHandler {
	return &NotificationHandler{service}
}

func (handler *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")
	role, _ := r.Context().Value(middlewares.RoleKey).(string)
	tokenID, _ := r.Context().Value(middlewares.UserIdKey).(string)

	// Get user_id from path variable
	userID := mux.Vars(r)["user_id"]

	// If the role is 'user', ensure the token ID matches the user ID
	if role == "doctor" || role == "patient" {
		if tokenID != userID {
			w.WriteHeader(http.StatusForbidden)
			_ = json.NewEncoder(w).Encode(models.APIResponse{
				Status: http.StatusForbidden,
				Data:   "Access denied",
			})
			loggerZap.Error("Access denied for user")
			return
		}
	}

	// Fetch notifications
	notifications, err := handler.service.GetNotificationsByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusNotFound,
			Data:   http.StatusText(http.StatusNotFound),
		})
		loggerZap.Error("User not found")
		return
	}
	loggerZap.Info("Notification fetched successfully")

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   notifications,
	})
}

func (handler *NotificationHandler) AllNotifications(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	// Fetch notifications
	notifications, err := handler.service.GetAllNotifications()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusNotFound,
			Data:   http.StatusText(http.StatusNotFound),
		})
		loggerZap.Error("User not found")
		return
	}
	loggerZap.Info("Notification fetched successfully")

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   notifications,
	})
}
