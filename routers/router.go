package routers

import (
	"aplikasi-adakost-be/middleware"
	kamarcontroller "aplikasi-adakost-be/modules/kamar/controller"
	kostcontroller "aplikasi-adakost-be/modules/kost/controller"
	transcationcontroller "aplikasi-adakost-be/modules/transaction/controller"
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
		api.POST("/login", usercontroller.Login)

		// KOST
		api.Use(middleware.JwtMiddleware())
		{
			api.POST("/kost", middleware.RoleMiddleware("pemilik"), kostcontroller.AddKost)
		}

		api.PUT("/kost/:id", kostcontroller.UpdateKost)
		api.GET("/kost", kostcontroller.GetAllKost)
		api.DELETE("/kost/:id", kostcontroller.DeleteKost)
		api.GET("/kost/kamar", kostcontroller.GetKamarKost)

		// KAMAR
		api.POST("/kamar", kamarcontroller.AddKamar)
		api.PUT("/kamar/:id", kamarcontroller.UpdateKamar)
		api.DELETE("/kamar/:id", kamarcontroller.DeleteKamar)
		api.GET("/kamar", kamarcontroller.GetAllKamar)

		// TRANSACTION
		api.POST("/transaction-booking/:id", transcationcontroller.SaveOrderBooking)
		api.PUT("/transaction-booking-cancel/:id", transcationcontroller.CancelOrderBooking)
		api.GET("/transaction-booking/:id", transcationcontroller.GetBookingList)
		api.GET("/transaction-booking/users/:id", transcationcontroller.GetUsersBookingList)
	}
	url := ginSwagger.URL("http://localhost:8080/v3/api-docs")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/v3/api-docs", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})
	return r
}
