package auth

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vanessadanu/Finpro-Golang.git/database"
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

func AddToBlacklist(tokenString string) error {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return err
	}

	expirationTime := time.Unix(int64(token.Claims.(jwt.MapClaims)["exp"].(float64)), 0)

	blacklistedToken := BlacklistedToken{
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}

	err = database.DB.Create(&blacklistedToken).Error
	if err != nil {
		return err
	}

	return nil
}

func IsTokenBlacklisted(tokenString string) bool {
	var count int64
	database.DB.Model(&BlacklistedToken{}).Where("token = ?", tokenString).Count(&count)

	return count > 0
}
