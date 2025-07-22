package routers

import (
	"aplikasi-adakost-be/common"
	"aplikasi-adakost-be/modules/user/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/hello", HelloWorld)

		// LOGIN DAN REGISTER
		api.POST("/signup", controller.SaveRegister)
	}
	url := ginSwagger.URL("http://localhost:8080/v3/api-docs") // endpoint JSON swagger kamu
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/v3/api-docs", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	return r
}

// HelloWorld godoc
// @Summary Menampilkan pesan Hello World
// @Description API ini mengembalikan pesan Hello World
// @Tags kost-controller
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func HelloWorld(ctx *gin.Context) {
	common.GenerateSuccessResponseWithData(ctx, "hello world", nil)
}
