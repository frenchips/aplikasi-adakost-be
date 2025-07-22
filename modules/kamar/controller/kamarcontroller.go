package controller

import (
	"aplikasi-adakost-be/common"

	"github.com/gin-gonic/gin"
)

// HelloWorld godoc
// @Summary Menampilkan pesan Hello World
// @Description API ini mengembalikan pesan Hello World
// @Tags kost-controller
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /kamarHello [get]
func KamarController(ctx *gin.Context) {
	common.GenerateSuccessResponseWithData(ctx, "hello world", nil)
}
