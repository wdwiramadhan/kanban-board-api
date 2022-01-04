package helper

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wdwiramadhan/kanban-board-api/domain"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(id int64, role string) string {
	claims := jwt.MapClaims{
		"id": id,
		"role": role,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, domain.ErrUnauthorized
	}
	if len(headerToken) <= 6 {
		return nil, domain.ErrUnauthorized
	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnauthorized
		}
		return []byte(secretKey), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, domain.ErrUnauthorized
	}
	return token.Claims.(jwt.MapClaims), nil
}