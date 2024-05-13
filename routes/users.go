package routes

import (
	"net/http"
	"strconv"

	"github.com/alerebal/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

}

func getUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "User data could not be parse"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func getUser(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	user, err := models.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// func deleteUser(ctx *gin.Context) {
// 	user, err :=
// }
