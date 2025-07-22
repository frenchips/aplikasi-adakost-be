package controller

import (
	"aplikasi-adakost-be/common"
	"aplikasi-adakost-be/databases/connection"
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
func AddKamar(ctx *gin.Context) {
	var input request.KamarRequest

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	// Buat repo dan service
	repo := repository.NewKamarRepository(connection.DBConnections)
	service := service.NewKamarService(repo)

	responseData, err := service.InsertKamar(input)
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
func UpdateKamar(ctx *gin.Context) {
	var input request.UpdateKamarRequest

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	// Buat repo dan service
	repo := repository.NewKamarRepository(connection.DBConnections)
	service := service.NewKamarService(repo)

	responseData, err := service.UpdateKamar(input, id)
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
