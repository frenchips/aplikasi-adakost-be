package controller

import (
	"aplikasi-adakost-be/common"
	"aplikasi-adakost-be/databases/connection"
	"aplikasi-adakost-be/middleware"
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
// @Security BearerAuth
func SaveOrderBooking(ctx *gin.Context) {
	var input request.BookingSaveRequest

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

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
	repoTransaction := transactionrepository.NewTransactionRepository(connection.DBConnections)
	repoKamar := kamarrepository.NewKamarRepository(connection.DBConnections)
	service := service.NewTransactionService(repoTransaction, repoKamar)

	responseData, err := service.SaveOrderBooking(input, id, claims.Username)
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
// @Security BearerAuth
func CancelOrderBooking(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

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
	repoTransaction := transactionrepository.NewTransactionRepository(connection.DBConnections)
	repoKamar := kamarrepository.NewKamarRepository(connection.DBConnections)
	service := service.NewTransactionService(repoTransaction, repoKamar)

	err := service.CancelOrderBooking(id, claims.Username)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully cancel booking ", nil)
}

// @Tags transaction-controller
// @Accept json
// @Produce json
// @Param id path int true "ID Booking"
// @Success 200 {object} common.APIResponse{data=response.BookingResponse}
// @Router /transaction-booking/{id} [get]
// @Security BearerAuth
func GetBookingList(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.GenerateErrorResponse(ctx, "ID tidak valid")
		return
	}

	repo := transactionrepository.NewTransactionRepository(connection.DBConnections)
	service := service.NewTransactionService(repo, nil)

	bookings, err := service.GetDetailBooking(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Berhasil mengambil data booking", bookings)
}

// @Tags transaction-controller
// @Accept json
// @Produce json
// @Param id path int true "ID Booking"
// @Success 200 {object} common.APIResponse{data=response.BookingResponse}
// @Router /transaction-booking/users/{id} [get]
// @Security BearerAuth
func GetUsersBookingList(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.GenerateErrorResponse(ctx, "ID tidak valid")
		return
	}

	repo := transactionrepository.NewTransactionRepository(connection.DBConnections)
	service := service.NewTransactionService(repo, nil)

	bookings, err := service.GetDetailUserBooking(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Berhasil mengambil data booking", bookings)
}
