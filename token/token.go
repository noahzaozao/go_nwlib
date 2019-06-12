package token

import (
	"github.com/dgrijalva/jwt-go"
	gerror "github.com/noahzaozao/go_nwlib/error"
	"time"
)

type User interface {
	GetIdStr() string
	GetUuid() string
}

func Encode(user User, issuer string, subject string, secretKey string) (string, error) {
	expireToken := time.Now().Add(time.Minute * 15).Unix()

	// Create the Claims
	standardClaims := jwt.StandardClaims{
		ExpiresAt: expireToken,
		Issuer:    issuer,
		Subject:   subject,
		Audience:  user.GetUuid(),
		Id:        user.GetIdStr(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardClaims)

	// Sign token and return
	return token.SignedString([]byte(secretKey))
}

func Decode(tokenString string, secretKey string) (*jwt.StandardClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (
			interface{}, error) {
			return []byte(secretKey), nil
		})

	if token != nil {
		// Validate the token and return the custom claims
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			return claims, nil
		} else {
			return nil, err
		}
	} else {
		return nil, gerror.GeneralRaiseError("token is nil")
	}
}
