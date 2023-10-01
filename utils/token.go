package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// SecretKey should be a long, random string used to sign the JWT token.
var SecretKey = []byte("your-secret-key")

// GetUserIDFromToken retrieves the user ID from a JWT token stored in the request context.
func GetUserIDFromToken(c echo.Context) (int, error) {
	// Get the JWT token string from the Authorization header
	tokenString := c.Request().Header.Get("Authorization")

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method and set the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		return SecretKey, nil
	})

	if err != nil {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	// Check if the token is valid and has valid claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(float64)
		if !ok {
			return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		return int(userID), nil
	}

	return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
}

// GenerateJWTToken generates a JWT token for a user with the given userID.
func GenerateJWTToken(userID int) (string, error) {
	// Define the claims for the JWT token
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (24 hours)
		"iat":    time.Now().Unix(),                     // Token issue time
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
