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
		api.GET("/kost/kamar", kostcontroller.GetKamarKost)

		// KOST ROLE PEMILIK
		api.Use(middleware.JwtMiddleware())
		{
			api.POST("/kost", middleware.RoleMiddleware("pemilik"), kostcontroller.AddKost)
			api.PUT("/kost/:id", middleware.RoleMiddleware("pemilik"), kostcontroller.UpdateKost)
			api.GET("/kost", middleware.RoleMiddleware("pemilik"), kostcontroller.GetAllKost)
			api.DELETE("/kost/:id", middleware.RoleMiddleware("pemilik"), kostcontroller.DeleteKost)

		}

		// KAMAR
		api.Use(middleware.JwtMiddleware())
		{
			api.POST("/kamar", middleware.RoleMiddleware("pemilik"), kamarcontroller.AddKamar)
			api.PUT("/kamar/:id", middleware.RoleMiddleware("pemilik"), kamarcontroller.UpdateKamar)
			api.DELETE("/kamar/:id", middleware.RoleMiddleware("pemilik"), kamarcontroller.DeleteKamar)
			api.GET("/kamar", middleware.RoleMiddleware("pemilik"), kamarcontroller.GetAllKamar)
		}

		// TRANSACTION
		api.Use(middleware.JwtMiddleware())
		{
			api.POST("/transaction-booking/:id", middleware.RoleMiddleware("penyewa"), transcationcontroller.SaveOrderBooking)
			api.PUT("/transaction-booking-cancel/:id", middleware.RoleMiddleware("penyewa"), transcationcontroller.CancelOrderBooking)
			api.GET("/transaction-booking-history", middleware.RoleMiddleware("pemilik", "penyewa"), transcationcontroller.GetHistoryBookingList)
		}

	}

	//
	// url := ginSwagger.URL("http://localhost:8080/v3/api-docs")

	// DEPLOYMENT
	url := ginSwagger.URL("/v3/api-docs")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/v3/api-docs", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "API is running!",
	// 	})
	// })
	return r
}
