package routers

import (
	"bengkel-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.GET("/cars", controllers.GetAllCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	router.DELETE("/cars/:carID", controllers.DeleteCar)

	router.GET("/cars/:carID", controllers.GetCar)

	return router
}
