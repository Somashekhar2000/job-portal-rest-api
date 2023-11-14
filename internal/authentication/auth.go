package authentication

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	privateKety *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

type Authenticaton interface{
	GenerateToken(claims jwt.RegisteredClaims)(string error)
	ValidateToken(token string)(jwt.RegisteredClaims,error)
}

func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey)(Authenticaton,error){
	if privateKey==nil && publicKey==nil{
		return nil,errors.New("public and private are nil")
	}

	return &Auth{
		privateKety: privateKey,
		publicKey: publicKey,
	},nil
}