package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitub.com/Jidetireni/events-restapi/utils"
)

func Auth(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userid, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userid", userid)
	context.Next()
}
