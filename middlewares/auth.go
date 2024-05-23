package middlewares

import (
	"net/http"

	"github.com/alerebal/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized user"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"messsage": "Not authorized user"})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
