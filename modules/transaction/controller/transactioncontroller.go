package controller

import (
	"aplikasi-adakost-be/common"
	"aplikasi-adakost-be/databases/connection"
	kamarrepository "aplikasi-adakost-be/modules/kamar/repository"
	transactionrepository "aplikasi-adakost-be/modules/transaction/repository"
	"aplikasi-adakost-be/modules/transaction/request"
	"aplikasi-adakost-be/modules/transaction/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags transaction-controller
// @Accept json
// @Produce json
// @Param id path int true "ID Kost"
// @Param body body request.BookingSaveRequest true "Data Kost"
// @Success 200 {object} common.APIResponse{data=response.BookingSaveResponse}
// @Router /transaction-booking/{id} [post]
func SaveOrderBooking(ctx *gin.Context) {
	var input request.BookingSaveRequest

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	// Buat repo dan service
	repoTransaction := transactionrepository.NewTransactionRepository(connection.DBConnections)
	repoKamar := kamarrepository.NewKamarRepository(connection.DBConnections)
	service := service.NewTransactionService(repoTransaction, repoKamar)

	responseData, err := service.SaveOrderBooking(input, id)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully save booking ", responseData)
}

// @Tags transaction-controller
// @Accept json
// @Produce json
// @Param id path int true "ID Booking"
// @Success 200 {object} common.APIResponse{data=response.BookingSaveResponse}
// @Router /transaction-booking-cancel/{id} [put]
func CancelOrderBooking(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	// Buat repo dan service
	repoTransaction := transactionrepository.NewTransactionRepository(connection.DBConnections)
	repoKamar := kamarrepository.NewKamarRepository(connection.DBConnections)
	service := service.NewTransactionService(repoTransaction, repoKamar)

	err := service.CancelOrderBooking(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully cancel booking ", nil)
}
