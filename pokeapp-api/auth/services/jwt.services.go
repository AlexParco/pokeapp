package services

import (
	"strconv"
	"strings"
	"time"

	"github.com/alexparco/pokeapp-api/model"
	"github.com/golang-jwt/jwt/v4"
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
	Username string `json:"username"`
	UserId   string `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *jwtService) GeneratedToken(user *model.User) (string, error) {
	customClaims := Claim{
		user.Username,
		strconv.FormatUint(uint64(user.UserId), 10),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
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
