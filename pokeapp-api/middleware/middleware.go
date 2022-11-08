package middleware

import (
	"fmt"
	"strings"

	"github.com/alexparco/pokeapp-api/auth/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type middleware struct {
	jwtSesion services.JwtService
}

func NewMiddleware(jwtSesion services.JwtService) *middleware {
	return &middleware{
		jwtSesion,
	}
}

func (m *middleware) AuthSessionMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		fmt.Println(bearer)
		if bearer != "" {
			token := strings.TrimPrefix(bearer, "Bearer ")
			fmt.Println(token)
			claims := m.jwtSesion.ValidateToken(token)
			if claims == nil {
				c.AbortWithStatus(401)
				return
			}

			mapClaims := claims.Claims.(jwt.MapClaims)
			c.Set("user_id", mapClaims["user_id"].(string))
		}
	}
}
