package main

import (
	"fmt"
	"os"

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

func StartApp()error{
		//initializing authentication support
		log.Info().Msg("main started : initializing with the authentication support")

		//reading private key file
		privatePemFile, err := os.ReadFile(`C:\Users\somas\Desktop\job-portal-apis\private.pem`)
		if err!=nil{
			log.Info().Msg("Error in reading private Key file")
			return fmt.Errorf("error in reading private key file : %w",err)
		}

		privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePemFile)
		if err!=nil{
			log.Info().Msg("Error in reading private Key")
			return fmt.Errorf("error in parsing private key : %w",err)
		}

		publicPemFile, err := os.ReadFile(`C:\Users\somas\Desktop\job-portal-apis\pubkey.pem`)
		if err!=nil {
			log.Info().Msg("Error in reading public Key filer")
			return fmt.Errorf("error in reading public key file : %w",err)
		}

		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPemFile)
		if err!=nil{
			log.Info().Msg("error in parsing public key")
			return fmt.Errorf("error in paring public key")
		}

		a,err := 



}