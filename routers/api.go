package routers

import (
	"bengkel-api/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "bengkel-api/docs"
)

// @title Car API
// @version 1.0
// @description This is a simple service for managing car
// @termOfService https://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.GET("/cars", controllers.GetAllCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	// router.DELETE("/cars/:carID", controllers.DeleteCar)

	router.GET("/cars/:carID", controllers.GetCar)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
