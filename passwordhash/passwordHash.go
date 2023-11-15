package passwordhash

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("error in hashing password")
		return "", fmt.Errorf("error in hashing password : %w", err)
	}
	return string(hashedPassword), nil
}
