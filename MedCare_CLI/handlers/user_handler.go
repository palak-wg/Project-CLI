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

type UserHandler struct {
	service interfaces.UserService
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{service}
}

//func (handler *UserHandler) ViewProfile(w http.ResponseWriter, r *http.Request) {
//	loggerZap, _ := zap.NewProduction()
//	defer loggerZap.Sync()
//
//	w.Header().Set("Content-Type", "application/json")
//	bearerToken := r.Header.Get("Authorization")
//	claims, err := tokens.ExtractClaims(bearerToken)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		err = json.NewEncoder(w).Encode(models.APIResponse{
//			Status: http.StatusInternalServerError,
//			Data:   "Error extracting claims",
//		})
//		loggerZap.Error("Internal Server Error")
//		if err != nil {
//			loggerZap.Error("Encoding response")
//		}
//		return
//	}
//
//	id := claims["id"].(string)
//	user, err := handler.service.GetUserByID(id)
//	if err != nil {
//		w.WriteHeader(http.StatusUnauthorized)
//		err = json.NewEncoder(w).Encode(models.APIResponse{
//			Status: http.StatusUnauthorized,
//			Data:   http.StatusText(http.StatusUnauthorized),
//		})
//		loggerZap.Error("Unauthorized user")
//		if err != nil {
//			loggerZap.Error("Encoding response")
//		}
//		return
//	}
//
//	// Check if the user is an admin
//	if claims["role"].(string) == "admin" { // Assuming user struct has IsAdmin field
//		allUsers, err := handler.service.GetAllUsers()
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			err = json.NewEncoder(w).Encode(models.APIResponse{
//				Status: http.StatusInternalServerError,
//				Data:   "Error fetching all users",
//			})
//			loggerZap.Error("Internal Server Error fetching all users")
//			if err != nil {
//				loggerZap.Error("Encoding response")
//			}
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//		err = json.NewEncoder(w).Encode(allUsers)
//		if err != nil {
//			loggerZap.Error("Encoding response")
//		}
//		return
//	}
//
//	// Regular user profile response
//	response := &models.User{
//		UserID:      user.UserID,
//		Name:        user.Name,
//		Age:         user.Age,
//		Gender:      user.Gender,
//		Email:       user.Email,
//		PhoneNumber: user.PhoneNumber,
//	}
//	w.WriteHeader(http.StatusOK)
//	err = json.NewEncoder(w).Encode(response)
//	if err != nil {
//		loggerZap.Error("Encoding response")
//	}
//}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
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
	user, err := handler.service.GetUserByID(userID)
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
	response := &models.APIResponseUser{
		UserID:      user.UserID,
		Name:        user.Name,
		Age:         user.Age,
		Gender:      user.Gender,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
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
