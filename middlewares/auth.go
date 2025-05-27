package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sep.com/eventapi/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized!"}) //abort the  current request and send the message
		return
	}
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
		return
	}

	context.Set("userId", userId) //enable us to add some data to this context

	context.Next() //make sure next request handler will getexecuted
}
