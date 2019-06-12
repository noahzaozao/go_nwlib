package token

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	gerror "github.com/noahzaozao/go_nwlib/error"
	"strings"
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

func CheckJWT(tokenString string) (*jwt.StandardClaims, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, jwt.NewValidationError("token contains an invalid number of segments", jwt.ValidationErrorMalformed)
	}

	claimBytes, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return nil, jwt.ValidationError{Inner: err, Errors: jwt.ValidationErrorMalformed}
	}

	claim := &jwt.StandardClaims{}
	dec := json.NewDecoder(bytes.NewBuffer(claimBytes))
	err = dec.Decode(claim)
	if err != nil {
		return nil, jwt.ValidationError{Inner: err, Errors: jwt.ValidationErrorMalformed}
	}

	return claim, nil
}
