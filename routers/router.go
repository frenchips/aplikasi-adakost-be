package routers

import (
	kamarcontroller "aplikasi-adakost-be/modules/kamar/controller"
	kostcontroller "aplikasi-adakost-be/modules/kost/controller"
	usercontroller "aplikasi-adakost-be/modules/user/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// LOGIN DAN REGISTER
		api.POST("/signup", usercontroller.SaveRegister)

		// KOST
		api.POST("/kost", kostcontroller.AddKost)
		api.PUT("/kost/:id", kostcontroller.UpdateKost)
		api.GET("/kost", kostcontroller.GetAllKost)
		api.DELETE("/kost/:id", kostcontroller.DeleteKost)

		// KAMAR
		api.POST("/kamar", kamarcontroller.AddKamar)
		api.PUT("/kamar/:id", kamarcontroller.UpdateKamar)
		api.DELETE("/kamar/:id", kamarcontroller.DeleteKamar)
		api.GET("/kamar", kamarcontroller.GetAllKamar)
	}
	url := ginSwagger.URL("http://localhost:8080/v3/api-docs")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/v3/api-docs", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	return r
}
