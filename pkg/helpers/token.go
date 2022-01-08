package helpers

import (
	"clean-micro/internal/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type AuthTokenClaim struct {
	ID uuid.UUID
	jwt.StandardClaims
}

func GenerateAccessToken(id uuid.UUID) (accessToken string, err error) {

	clog := log.WithFields(log.Fields{
		"method": "GenerateAccessToken",
	})

	expiresAt := time.Now().Add(time.Minute * time.Duration(config.Conf.Jwt.AccessTokenExpiry)).Unix()

	claims := &AuthTokenClaim{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err = token.SignedString([]byte(config.Conf.Jwt.Secret))
	if err != nil {
		eMsg := "error in token.SignetString"
		clog.WithError(err).Error(eMsg)
		return
	}

	return
}

func VerifyAccessToken(token string) (id uuid.UUID, err error) {

	clog := log.WithFields(log.Fields{
		"method": "VerifyAccessToken",
	})

	claims := jwt.MapClaims{}
	_, err1 := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.Jwt.Secret), nil
	})

	if err1 != nil {
		eMsg := "error in jwt.ParseWithClaims"
		clog.WithError(err).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	idStr := fmt.Sprintf("%v", claims["ID"])

	id, err1 = ConvertStringToUUID(idStr)
	if err1 != nil {
		eMsg := "error in uuid.FromString"
		clog.WithError(err1).Error(eMsg)
		err = status.Error(codes.Code(500), eMsg)
		return
	}

	return
}

func GenerateRefreshToken() (refreshToken uuid.UUID, err error) {
	refreshToken, err = uuid.NewV4()
	if err != nil {
		fmt.Println("error RefreshToken")
		return
	}

	return
}
