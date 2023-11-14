package main

import (
	"fmt"
	"job-portal-api/database"
	"job-portal-api/handler"
	"job-portal-api/internal/authentication"
	"job-portal-api/repository"
	"job-portal-api/service"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

func main() {
	err := StartApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("Heloo this is our job-portal-app")
}

func StartApp() error {
	//initializing authentication support
	log.Info().Msg("main started : initializing with the authentication support")

	//reading private key file
	privatePemFile, err := os.ReadFile(`C:\Users\somas\Desktop\job-portal-apis\private.pem`)
	if err != nil {
		log.Info().Msg("Error in reading private Key file")
		return fmt.Errorf("error in reading private key file : %w", err)
	}

	//parsing private pem filr content to rsa private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePemFile)
	if err != nil {
		log.Info().Msg("Error in reading private Key")
		return fmt.Errorf("error in parsing private key : %w", err)
	}

	//reading public key file
	publicPemFile, err := os.ReadFile(`C:\Users\somas\Desktop\job-portal-apis\pubkey.pem`)
	if err != nil {
		log.Info().Msg("Error in reading public Key filer")
		return fmt.Errorf("error in reading public key file : %w", err)
	}

	//parsing public pem file content to rsa public key
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPemFile)
	if err != nil {
		log.Info().Msg("error in parsing public key")
		return fmt.Errorf("error in paring public key")
	}

	//callig factory function to get authentication
	a, err := authentication.NewAuth(privateKey, publicKey)
	if err != nil {
		log.Info().Msg("error in auth function ")
		return fmt.Errorf("error in auth function : %w", err)
	}

	//starting with dtatbase connection
	log.Info().Msg("main started : initializing the database")

	db, err := database.DatabaseConnection()
	if err != nil {
		log.Info().Msg("error while opening data base connection")
		return fmt.Errorf("error while opening data base connection : %w", err)
	}

	//initializing the repo layer
	repo, err := repository.NewRepository(db)
	if err != nil {
		log.Info().Msg("error while initializing the repository")
		return err
	}

	service, err := service.NewService(repo, a)
	if err != nil {
		log.Info().Msg("error while initializing service")
		return fmt.Errorf("error while initializing service : %w", err)
	}

	//initilazing http server
	api := http.Server{
		Addr:         "8087",
		ReadTimeout:  8000 * time.Second,
		WriteTimeout: 800 * time.Second,
		IdleTimeout:  800 * time.Second,
		Handler:      handler.SetupApi(a, service),
	}

}
