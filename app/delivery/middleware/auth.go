package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wdwiramadhan/kanban-board-api/app/helper"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Set("user", verifyToken)
		c.Next()
	}
}

func Authorization(roles []string) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		userAuth := ctx.MustGet("user").(jwt.MapClaims)
		userRole := userAuth["role"]
		for _,role := range roles {
			if role  == userRole {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "you don't have access",
		})
	}
}