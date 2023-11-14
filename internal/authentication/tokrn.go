package authentication

import "github.com/golang-jwt/jwt/v5"

func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string error) {

}

func (a *Auth) ValidateToken(token string)(jwt.RegisteredClaims,error){
	
}