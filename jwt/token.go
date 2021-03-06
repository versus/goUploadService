package jwt

import (
	"errors"
	"time"

	"github.com/brianvoe/sjwt"
)

//GenerateToken is function for generate JWT token
func GenerateToken(secretKey string, validDays int) string {

	if validDays < 1 {
		validDays = 1
	}

	claims := sjwt.New() // Issuer of the token
	claims.SetIssuer("Ifconfigco")
	/*
		claims.SetTokenID()                                                            // UUID generated
		claims.SetSubject("Bearer Token")                                              // Subject of the token
		claims.SetAudience([]string{"Prometeus"})                                      // Audience the toke is for
		claims.SetIssuedAt(time.Now())                                                 // IssuedAt in time, value is set in unix
	*/
	claims.SetNotBeforeAt(time.Now())                                              // Token valid now
	claims.SetExpiresAt(time.Now().Add(time.Hour * 24 * time.Duration(validDays))) // Token expires in 24 hours
	jwt := claims.Generate([]byte(secretKey))
	return jwt
}

//ValidateToken is function for validate JWT token
func ValidateToken(secretKey string, token string) error {

	if !sjwt.Verify(token, []byte(secretKey)) {
		return errors.New("Token isn't valid")
	}

	claims, err := sjwt.Parse(token)
	if err != nil {
		return err
	}
	//TODO: add check claims fields if we need

	// Validate will check(if set) Expiration At and Not Before At dates
	err = claims.Validate()
	if err != nil {
		return err
	}

	return nil

}
