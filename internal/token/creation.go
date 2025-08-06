package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thenameiswiiwin/reelingit/internal/logger"
	"github.com/thenameiswiiwin/reelingit/internal/models"
)

func CreateJWT(user models.User, logger logger.Logger) string {
	jwtSecret := GetJWTSecret(logger)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.Error("Failed to sign JWT token", err)
		return ""
	}

	return tokenString
}
