package handlers

import (
	"doctor-patient-cli/interfaces"
	"doctor-patient-cli/middlewares"
	"doctor-patient-cli/models"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type MessageHandler struct {
	service interfaces.MessageService
}

func NewMessageHandler(service interfaces.MessageService) *MessageHandler {
	return &MessageHandler{service}
}

func (handler *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	userID, _ := r.Context().Value(middlewares.UserIdKey).(string)

	messages, err := handler.service.GetUnreadMessages(userID)
	fromId := r.URL.Query().Get("from_id")

	if fromId != "" {
		messages, err = handler.service.GetUnreadMessagesById(fromId, userID)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error getting messages",
		})
		loggerZap.Error("Error getting messages")
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   messages,
	})
	loggerZap.Info("Returning messages")
}

func (handler *MessageHandler) AddMessage(w http.ResponseWriter, r *http.Request) {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	w.Header().Set("Content-Type", "application/json")

	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
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

	userID, _ := r.Context().Value(middlewares.UserIdKey).(string)

	if userID != message.Sender {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusBadRequest,
			Data:   http.StatusText(http.StatusBadRequest),
		})
		loggerZap.Error("Cannot send message")
	}

	err = handler.service.SendMessage(message.Sender, message.Receiver, message.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(models.APIResponse{
			Status: http.StatusInternalServerError,
			Data:   "Error sending message",
		})
		loggerZap.Error("Error sending message")
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(models.APIResponse{
		Status: http.StatusOK,
		Data:   "Message sent successfully",
	})
}
