package Routers

import (
	"GinGo/Controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", Controller.CreateCar)

	router.PUT("/cars/:carID", Controller.UpdateCar)

	router.GET("/cars/:carID", Controller.GetCar)

	router.DELETE("/cars/:carID", Controller.DeleteCar)

	return router
}
