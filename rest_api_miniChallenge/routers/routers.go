package routers

import (
	"rest_api_mini_challenge/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("order", controllers.CreateOrder)
		v1.GET("order", controllers.GetAllOrders)
		v1.GET("order/:id", controllers.GetOrderById)
		v1.PUT("order/:id", controllers.UpdateOrder)
		v1.DELETE("order/:id", controllers.DeleteOrder)
	}

	return r
}
