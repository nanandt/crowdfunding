package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var secretKey = []byte("SECRET_KEY")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID

	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := withClaims.SignedString(secretKey)
	if err != nil{
		return signedString, err
	}
	return signedString, nil

}
