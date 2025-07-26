package controller

import (
	"aplikasi-adakost-be/common"
	"aplikasi-adakost-be/databases/connection"
	"aplikasi-adakost-be/middleware"
	"aplikasi-adakost-be/modules/kost/repository"
	"aplikasi-adakost-be/modules/kost/request"
	"aplikasi-adakost-be/modules/kost/response"
	"aplikasi-adakost-be/modules/kost/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags kost-controller
// @Accept json
// @Produce json
// @Param body body request.AddKostRequest true "Data Kost"
// @Success 200 {object} common.APIResponse{data=response.KostResponse}
// @Router /kost [post]
// @Security BearerAuth
func AddKost(ctx *gin.Context) {
	var input request.AddKostRequest

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
	repo := repository.NewKostRepository(connection.DBConnections, nil)
	service := service.NewKostService(repo)

	responseData, err := service.InsertKost(input, claims.Username, claims.UserID)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully insert kost ", responseData)
}

// @Tags kost-controller
// @Accept json
// @Produce json
// @Param id path int true "ID Kost"
// @Param body body request.UpdateKostRequest true "Data Kost"
// @Success 200 {object} common.APIResponse{data=response.KostResponse}
// @Router /kost/{id} [put]
// @Security BearerAuth
func UpdateKost(ctx *gin.Context) {
	var input request.UpdateKostRequest

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
	repo := repository.NewKostRepository(connection.DBConnections, nil)
	service := service.NewKostService(repo)

	responseData, err := service.UpdateKost(input, id, claims.Username)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully update kost ", responseData)
}

// @Tags kost-controller
// @Accept json
// @Produce json
// @Success 200 {object} common.APIResponse{data=response.ViewKostResponse}
// @Router /kost [GET]
// @Security BearerAuth
func GetAllKost(ctx *gin.Context) {
	var input response.ViewKostResponse

	// Buat repo dan service
	repo := repository.NewKostRepository(connection.DBConnections, nil)
	service := service.NewKostService(repo)

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

	responseData, err := service.GetAllKost(input, claims.UserID)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully get all kost ", responseData)
}

// @Tags kost-controller
// @Accept json
// @Produce json
// @Param id path int true "ID Kost"
// @Success 200 {object} common.APIResponse
// @Router /kost/{id} [delete]
// @Security BearerAuth
func DeleteKost(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	// Buat repo dan service
	repo := repository.NewKostRepository(connection.DBConnections, nil)
	service := service.NewKostService(repo)

	err := service.DeleteKost(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully Delete kost ", nil)
}

// @Tags kost-controller
// @Accept json
// @Produce json
// @Success 200 {object} common.APIResponse{data=response.KamarKostReponse}
// @Router /kost/kamar [get]
// @Security BearerAuth
func GetKamarKost(ctx *gin.Context) {

	// Buat repo dan service
	repo := repository.NewKostRepository(connection.DBConnections, nil)
	service := service.NewKostService(repo)

	responseData, err := service.GetKostKamar()
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully Get Kost ", responseData)
}
