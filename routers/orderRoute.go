package routers

import (
	"assignment2/config"
	"assignment2/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	db := config.NewPostgres()
	if db == nil {
		fmt.Println("running db is fail")
	}

	orderController := controllers.NewOrderController(db)
	router := gin.Default()

	router.GET("/orders", orderController.GetOrder)
	router.POST("/orders", orderController.CreateOrder)
	router.PUT("/orders/:id", orderController.UpdateOrderByID)
	router.DELETE("/orders/:id", orderController.DeleteOrderByID)
	// router.GET("/orders/:id", orderController.FindOrderByID)order

	return router
}
