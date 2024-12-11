package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitub.com/Jidetireni/events-restapi/models"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user. Try again later."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCred()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!"})

}
