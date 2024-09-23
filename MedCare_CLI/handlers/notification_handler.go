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
	tokenID := claims["id"].(string)
	role := claims["role"].(string)

	// Get user_id from path variable
	userID := mux.Vars(r)["user_id"]

	// If the role is 'user', ensure the token ID matches the user ID
	if role == "doctor" || role == "patient" {
		if tokenID != userID {
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
	}

	// Fetch the user profile
	notifications, err := handler.service.GetNotificationsByUserID(userID)
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
	loggerZap.Info("Notification fetched successfully")

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   notifications,
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

//func (handler *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
//	loggerZap, _ := zap.NewProduction()
//	defer loggerZap.Sync()
//
//	w.Header().Set("Content-Type", "application/json")
//
//	// Fetch all user profiles
//	users, err := handler.service.GetAllUsers()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		err = json.NewEncoder(w).Encode(models.APIResponse{
//			Status: http.StatusInternalServerError,
//			Data:   "Error fetching user profiles",
//		})
//		loggerZap.Error("Internal Server Error fetching user profiles")
//		if err != nil {
//			loggerZap.Error("Encoding response")
//		}
//		return
//	}
//
//	var responseUsers []models.APIResponseUser
//	for _, user := range users {
//		responseUsers = append(responseUsers, models.APIResponseUser{
//			UserID:      user.UserID,
//			Name:        user.Name,
//			Age:         user.Age,
//			Gender:      user.Gender,
//			Email:       user.Email,
//			PhoneNumber: user.PhoneNumber,
//		})
//	}
//
//	// Prepare response
//	w.WriteHeader(http.StatusOK)
//	err = json.NewEncoder(w).Encode(models.APIResponse{
//		Status: http.StatusOK,
//		Data:   responseUsers,
//	})
//	if err != nil {
//		loggerZap.Error("Encoding response")
//	}
//}
