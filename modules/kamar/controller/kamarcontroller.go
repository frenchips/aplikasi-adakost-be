package controller

import (
	"aplikasi-adakost-be/common"
	"aplikasi-adakost-be/databases/connection"
	"aplikasi-adakost-be/middleware"
	"aplikasi-adakost-be/modules/kamar/repository"
	"aplikasi-adakost-be/modules/kamar/request"
	"aplikasi-adakost-be/modules/kamar/response"
	"aplikasi-adakost-be/modules/kamar/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags kamar-controller
// @Accept json
// @Produce json
// @Param body body request.KamarRequest true "Data Kost"
// @Success 200 {object} common.APIResponse{data=response.KamarResponse}
// @Router /kamar [post]
// @Security BearerAuth
func AddKamar(ctx *gin.Context) {
	var input request.KamarRequest

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	// Ambil user dari context
	claimsInterface, exists := ctx.Get("user")
	if !exists {
		common.GenerateErrorResponse(ctx, "Unauthorized: token not found")
		return
	}

	claims, ok := claimsInterface.(*middleware.Claims)
	if !ok {
		common.GenerateErrorResponse(ctx, "Invalid token data")
		return
	}

	// Buat repo dan service
	repo := repository.NewKamarRepository(connection.DBConnections)
	service := service.NewKamarService(repo)

	responseData, err := service.InsertKamar(input, claims.Username)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully insert kamar ", responseData)
}

// @Tags kamar-controller
// @Accept json
// @Produce json
// @Param id path int true "ID Kamar"
// @Param body body request.UpdateKamarRequest true "Data Kamar"
// @Success 200 {object} common.APIResponse{data=response.KamarResponse}
// @Router /kamar/{id} [put]
// @Security BearerAuth
func UpdateKamar(ctx *gin.Context) {
	var input request.UpdateKamarRequest

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	// Ambil user dari context
	claimsInterface, exists := ctx.Get("user")
	if !exists {
		common.GenerateErrorResponse(ctx, "Unauthorized: token not found")
		return
	}

	claims, ok := claimsInterface.(*middleware.Claims)
	if !ok {
		common.GenerateErrorResponse(ctx, "Invalid token data")
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	// Buat repo dan service
	repo := repository.NewKamarRepository(connection.DBConnections)
	service := service.NewKamarService(repo)

	responseData, err := service.UpdateKamar(input, id, claims.Username)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully update kamar ", responseData)
}

// @Tags kamar-controller
// @Accept json
// @Produce json
// @Param id path int true "ID Kamar"
// @Success 200 {object} common.APIResponse
// @Router /kamar/{id} [delete]
// @Security BearerAuth
func DeleteKamar(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	// Buat repo dan service
	repo := repository.NewKamarRepository(connection.DBConnections)
	service := service.NewKamarService(repo)

	err := service.DeleteKamar(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully delete kamar ", nil)
}

// @Tags kamar-controller
// @Accept json
// @Produce json
// @Success 200 {object} common.APIResponse
// @Router /kamar [get]
// @Security BearerAuth
func GetAllKamar(ctx *gin.Context) {

	var input response.GetKamarResponse

	// Buat repo dan service
	repo := repository.NewKamarRepository(connection.DBConnections)
	service := service.NewKamarService(repo)

	responseData, err := service.GetAllKamar(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully get all kamar ", responseData)
}
