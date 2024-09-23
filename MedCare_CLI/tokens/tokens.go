package tokens

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

func GenerateToken(role string, username string, isApprove bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      username,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
		"role":    role,
		"approve": isApprove,
	})
	signedToken, err := token.SignedString([]byte(os.Getenv("Secret")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(bearerToken string) bool {
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("Secret")), nil
	})
	if err != nil {
		return false
	}
	return parsedToken.Valid
}

//func ExtractClaims(bearerToken string) (jwt.MapClaims, error) {
//	token := strings.TrimPrefix(bearerToken, "Bearer ")
//	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(os.Getenv("Secret")), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
//		return claims, nil
//	}
//	return nil, errors.New("invalid token")
//}

// ClaimsExtractor defines a type for a function that extracts claims from a token.
type ClaimsExtractor func(token string) (map[string]interface{}, error)

// The default claims extractor implementation.
var extractClaims ClaimsExtractor = ExtractClaims

// SetClaimsExtractor allows you to override the default claims extractor (useful for testing).
func SetClaimsExtractor(extractor ClaimsExtractor) {
	extractClaims = extractor
}

// ExtractClaims extracts claims from a given bearer token.
func ExtractClaims(bearerToken string) (map[string]interface{}, error) {
	// Trim the "Bearer " prefix from the token
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token uses HMAC signing
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("Secret")), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract the claims if the token is valid
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GetClaims is a wrapper around the current extractClaims function (either default or overridden).
func GetClaims(token string) (map[string]interface{}, error) {
	return extractClaims(token)
}
