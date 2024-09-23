package middlewares

import (
	"doctor-patient-cli/models"
	"doctor-patient-cli/tokens"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	loggerZap, _ := zap.NewProduction()
	defer loggerZap.Sync()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			err := json.NewEncoder(w).Encode(models.APIResponse{
				Status: http.StatusUnauthorized,
				Data:   http.StatusText(http.StatusUnauthorized),
			})
			loggerZap.Error("Authentication error", zap.Error(err))
			if err != nil {
				loggerZap.Error("Encoding response")
			}
			return
		}
		if !tokens.ValidateToken(authHeader) {
			w.WriteHeader(http.StatusUnauthorized)
			err := json.NewEncoder(w).Encode(models.APIResponse{
				Status: http.StatusUnauthorized,
				Data:   http.StatusText(http.StatusUnauthorized),
			})
			loggerZap.Error("Authentication error", zap.Error(err))
			if err != nil {
				loggerZap.Error("Encoding response")
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}
