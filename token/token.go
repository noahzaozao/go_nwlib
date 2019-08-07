package token

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/noahzaozao/go_nwlib/cache"
	gerror "github.com/noahzaozao/go_nwlib/error"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
	"sync"
	"time"
)

const UserSessionPrefix = "u_session:"

type User interface {
	GetIdStr() string
	GetUuid() string
}

type TokenManager struct{}

var instance *TokenManager
var once sync.Once

func TokenMgr() *TokenManager {
	once.Do(func() {
		instance = &TokenManager{}
	})
	return instance
}

func (tokenMgr *TokenManager) Encode(user User, issuer string, subject string, secretKey string, d time.Duration) (string, error) {
	expireToken := time.Now().Add(d).Unix()
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

func (tokenMgr *TokenManager) Decode(tokenString string, secretKey string) (*jwt.StandardClaims, error) {
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

// Generate JWT from Uuid
func (tokenMgr *TokenManager) GenerateJWT(user User) (string, string, error) {
	client, err := cache.CacheMgr().Conn()
	if err != nil {
		return "", "", err
	}
	userJWTSecret := uuid.NewV4().String()

	client.HSet(UserSessionPrefix+user.GetUuid(), "jwt_token_secret", userJWTSecret)
	client.Expire(UserSessionPrefix+user.GetUuid(), time.Minute*20)

	client.HSet(UserSessionPrefix+user.GetUuid(), "refresh_token_secret", userJWTSecret)
	client.Expire(UserSessionPrefix+user.GetUuid(), time.Hour*24*5)

	jwtToken, err := tokenMgr.Encode(user, "user.srv", "normal", userJWTSecret, time.Minute*15)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := tokenMgr.Encode(user, "user.srv", "normal", userJWTSecret, time.Hour*24*5)
	if err != nil {
		return "", "", err
	}
	return jwtToken, refreshToken, nil
}

// Get JWT from ctx
func (tokenMgr *TokenManager) GetJWT(ctx context.Context) (string, error) {
	if _, ok := metadata.FromIncomingContext(ctx); !ok {
		log.Println("WARNING: No MetaData")
		return "", gerror.GeneralRaiseError("WARNING: No MetaData")
	}
	metaData, _ := metadata.FromIncomingContext(ctx)
	if _, ok := metaData["authorization"]; !ok {
		log.Println("WARNING: No Auth	Information")
		return "", gerror.GeneralRaiseError("WARNING: No Auth Information")
	}
	jwtToken := metaData["authorization"]
	return strings.Join(jwtToken, ""), nil
}

// Get Claims from tokenString
func (tokenMgr *TokenManager) GetClaims(tokenString string) (*jwt.StandardClaims, error) {
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

// Check JWT from tokenString
func (tokenMgr *TokenManager) CheckJWT(tokenString string) (string, error) {
	client, err := cache.CacheMgr().Conn()
	if err != nil {
		return "", err
	}
	claim, err := tokenMgr.GetClaims(tokenString)
	if err != nil {
		return "", err
	}
	secretKey, err := client.HGet(UserSessionPrefix+claim.Audience, "jwt_token_secret").Result()
	if err != nil {
		return "", err
	}
	authClaims, err := tokenMgr.Decode(tokenString, secretKey)
	if err != nil {
		return "", err
	}
	return authClaims.Id, nil
}

// Check JWT from tokenString
func (tokenMgr *TokenManager) CheckRefreshToken(tokenString string) (string, error) {
	client, err := cache.CacheMgr().Conn()
	if err != nil {
		return "", err
	}
	claim, err := tokenMgr.GetClaims(tokenString)
	if err != nil {
		return "", err
	}
	secretKey, err := client.HGet(UserSessionPrefix+claim.Audience, "refresh_token_secret").Result()
	if err != nil {
		return "", err
	}
	authClaims, err := tokenMgr.Decode(tokenString, secretKey)
	if err != nil {
		return "", err
	}
	return authClaims.Id, nil
}

// Clean JWT from ctx
func (tokenMgr *TokenManager) CleanJWT(ctx context.Context) error {
	client, err := cache.CacheMgr().Conn()
	if err != nil {
		return err
	}
	jwtToken, err := tokenMgr.GetJWT(ctx)
	if err != nil {
		return err
	}
	claim, err := tokenMgr.GetClaims(jwtToken)
	if err != nil {
		return err
	}
	err = client.HDel(UserSessionPrefix+claim.Audience, "jwt_token_secret").Err()
	if err != nil {
		return err
	}
	return nil
}
