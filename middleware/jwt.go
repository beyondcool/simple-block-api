package middleware

import (
	"simple-block-api/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var whiteList = map[string]bool{
	"/api/v1/login":    true,
	"/api/v1/register": true,
	"/api/v1/health":   true,
}

func GenerateToken(userID uint, username string, hours int) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(hours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.JWT.SecretKey))
}

func MyJWT(c *gin.Context) {
	if whiteList[c.Request.URL.Path] {
		// do nothing, just continue to next handler
	} else {
		// get token from Authorization header
		auth := c.GetHeader("Authorization")
		if auth == "" {
			panic("Authorization header is missing")
		}
		// check token format: "Bearer <token>"
		authParts := strings.SplitN(auth, " ", 2)
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			panic("Invalid Authorization header format")
		}
		tokenStr := authParts[1]
		// parse and validate token
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.JWT.SecretKey), nil
		})
		if err != nil || !token.Valid {
			panic("Invalid token")
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			panic("Invalid token claims")
		}

		// set user info to context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
	}

	// continue to next handler
	c.Next()

}
