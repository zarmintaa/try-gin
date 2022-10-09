package routers

import (
	"belajar-golang-restful-api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controllers.GetCar)
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)
	return router
}
