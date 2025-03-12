package middlewares

import (
	"context"
	"doctor-patient-cli/models"
	"doctor-patient-cli/tokens"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type contextKey string

var UserIdKey = contextKey("userId")
var RoleKey = contextKey("role")

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

		bearerToken := r.Header.Get("Authorization")
		claims, err := tokens.GetClaims(bearerToken)
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
		id := claims["id"].(string)
		role := claims["role"].(string)

		ctx := context.WithValue(r.Context(), UserIdKey, id)
		ctx = context.WithValue(ctx, RoleKey, role)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
