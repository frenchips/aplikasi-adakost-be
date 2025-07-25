package controller

import (
	"aplikasi-adakost-be/common"
	"aplikasi-adakost-be/databases/connection"
	"aplikasi-adakost-be/modules/user/repository"
	"aplikasi-adakost-be/modules/user/request"
	"aplikasi-adakost-be/modules/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags user-controller
// @Accept json
// @Produce json
// @Param user body request.RegisterRequest true "Data user"
// @Success 200 {object} common.APIResponse{data=response.SignUpResponse}
// @Router /signup [post]
func SaveRegister(ctx *gin.Context) {
	var input request.RegisterRequest

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	// Buat repo dan service
	repo := repository.NewUsersRepo(connection.DBConnections)
	service := service.NewUserService(repo)

	responseData, err := service.SaveRegisterUser(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully signup", responseData)
}

// @Tags user-controller
// @Accept json
// @Produce json
// @Param user body request.Login true "Data user"
// @Success 200 {object} common.APIResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	repo := repository.NewUsersRepo(connection.DBConnections)
	service := service.NewUserService(repo)

	token, err := service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
