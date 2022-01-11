package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wdwiramadhan/kanban-board-api/app/helper"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helper.VerifyToken(ctx)
		_ = verifyToken

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.Set("user", verifyToken)
		ctx.Next()
	}
}

func Authorization(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userAuth := ctx.MustGet("user").(jwt.MapClaims)
		userRole := userAuth["role"]
		for _, role := range roles {
			if role == userRole {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "you don't have access",
		})
	}
}
