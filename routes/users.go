package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sep.com/eventapi/models"
	"sep.com/eventapi/utils"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user "})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful!", "token": token})

}
