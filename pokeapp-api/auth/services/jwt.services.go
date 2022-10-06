package services

import (
	"strings"
	"time"

	"github.com/alexparco/pokeapp-api/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtService interface {
	ValidateToken(token string) *jwt.Token
	GeneratedToken(user *model.User) (string, error)
}

type jwtService struct {
	key string
}

func NewJwtService(key string) JwtService {
	return &jwtService{key}
}

type Claim struct {
	Email  string    `json:"email"`
	UserId uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *jwtService) GeneratedToken(user *model.User) (string, error) {
	customClaims := Claim{
		user.Email,
		user.UserId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	token, err := t.SignedString([]byte(j.key))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *jwtService) ValidateToken(token string) *jwt.Token {

	splitToken := strings.Split(token, " ")
	tokenString := splitToken[0]

	parseToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.key), nil
	})

	if err != nil {
		return nil
	}
	return parseToken
}
