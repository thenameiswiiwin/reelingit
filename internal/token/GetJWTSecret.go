package token

import (
	"os"

	"github.com/thenameiswiiwin/reelingit/internal/logger"
)

func GetJWTSecret(logger logger.Logger) string {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret-for-dev"
		logger.Info("JWT_SECRET environment variable is not set, using default secret for development purposes")
	} else {
		logger.Info("Using JWT_SECRET from environment variable")
	}
	return jwtSecret
}
