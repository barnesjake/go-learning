package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		//AbortWithStatusJSON calls `Abort()` and then `JSON` internally. This method stops the chain, writes the status code and return a JSON body
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization token found"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized."})
	}

	context.Set("userId", userId)

	//Next() just allows chain to continue
	context.Next()
}
