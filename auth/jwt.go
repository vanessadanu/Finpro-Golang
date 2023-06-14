package auth

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type BlacklistedToken struct {
	ID        uint `gorm:"primaryKey"`
	Token     string
	ExpiresAt time.Time
}

var jwtSecret = []byte("secret-key")

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err.Error())
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.NewValidationError("Invalid token claims", jwt.ValidationErrorClaimsInvalid)
	}

	return claims, nil
}

func ExtractToken(authorizationHeader string) string {
	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return ""
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	return token
}

// func GenerateExpiredToken(tokenString string) (string, error) {
// 	// Parse the existing token to extract the claims
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("failed to parse token: %v", err)
// 	}

// 	// Create a new token with a short expiration time (e.g., 1 second)
// 	expiredClaims := jwt.MapClaims{
// 		"exp": time.Now().Add(time.Second).Unix(),
// 	}

// 	// Sign the new token with the secret key
// 	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
// 	signedToken, err := newToken.SignedString(jwtSecret)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to generate expired token: %v", err)
// 	}

// 	return signedToken, nil
// }
