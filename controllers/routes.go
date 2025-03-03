package controllers

import (
	"atlanta-site/controllers/order_controllers"
	"atlanta-site/controllers/user_controllers"
	"atlanta-site/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	admin.Use(middlewares.AdminMiddleware())
	{
		admin.GET("/orders", middlewares.AdminMiddleware(), order_controllers.ListOrders)
		admin.PUT("/orders/:id", middlewares.AdminMiddleware(), order_controllers.UpdateOrderStatus)
		admin.PUT("/users/:id/promote", middlewares.AdminMiddleware(), user_controllers.PromoteUserToAdmin)
	}
	router.POST("/create_order/", order_controllers.CreateOrder)
	router.POST("/process-payment/", ProcessPayment)
	router.PUT("/update-order-status/:order_id/", order_controllers.UpdateOrderStatus)
}
