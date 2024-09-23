package middlewares

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/tokens"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func AdminMiddleware(next http.Handler) http.Handler {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

		// Get role from claims
		role := claims["role"].(string)

		// Check if the user is an admin
		if role != "admin" {
			w.WriteHeader(http.StatusForbidden)
			err = json.NewEncoder(w).Encode(models.APIResponse{
				Status: http.StatusForbidden,
				Data:   "Access denied",
			})
			loggerZap.Error("Access denied for non-admin user")
			if err != nil {
				loggerZap.Error("Encoding response")
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}
