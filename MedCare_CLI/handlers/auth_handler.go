package handlers

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/tokens"
	"doctor-patient-cli/utils"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func (handler *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {

	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

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

	if id, err := handler.service.GetUserByID(user.UserID); id != nil && err == nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusBadRequest,
			Data:   "This username already exists.",
		})
		loggerZap.Info("Username already exists")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	user.Password = utils.HashPassword(user.Password)

	err = handler.service.CreateUser(&user)
	if err != nil {
		fmt.Println("create user error")
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   http.StatusText(http.StatusInternalServerError),
		})
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusCreated,
		Data:   http.StatusText(http.StatusCreated),
	})
	if err != nil {
		loggerZap.Error("Encoding response")
	}
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")
	var client models.User
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusBadRequest,
			Data:   http.StatusText(http.StatusBadRequest),
		})

		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}

	user, err := handler.service.GetUserByID(client.UserID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusUnauthorized,
			Data:   http.StatusText(http.StatusUnauthorized),
		})
		loggerZap.Error("Unauthorized user")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}
	if !utils.CheckPasswordHash(client.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusUnauthorized,
			Data:   http.StatusText(http.StatusUnauthorized),
		})
		loggerZap.Error("Unauthorized user")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}
	generatedToken, err := tokens.GenerateToken(user.UserType, user.UserID, user.IsApproved)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   http.StatusText(http.StatusInternalServerError),
		})
		loggerZap.Error("Internal Server Error")
		if err != nil {
			loggerZap.Error("Encoding response")
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	response := struct {
		Code  int    `json:"status_code"`
		Token string `json:"token"`
	}{
		Code:  http.StatusOK,
		Token: generatedToken,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		loggerZap.Error("Encoding response")
	} else {
		loggerZap.Info("Successfully logged in & token sent")
	}
}
