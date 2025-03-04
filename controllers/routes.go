package controllers

import (
	"atlanta-site/controllers/order_controllers"
	payment_controller "atlanta-site/controllers/payments_controller"
	"atlanta-site/controllers/user_controllers"
	"atlanta-site/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	admin.Use(middlewares.AdminMiddleware())
	{
		admin.GET("/orders", middlewares.AdminMiddleware(), order_controllers.ListOrders)
		admin.PUT("/orders/:id", middlewares.AdminMiddleware(), order_controllers.ListOrders)
		admin.PUT("/users/:id/promote", middlewares.AdminMiddleware(), user_controllers.UpdateUserDetails)
	}
	router.POST("/create_order/", order_controllers.CreateOrder)
	router.POST("/process-payment/", payment_controller.CreatePaymentIntent)
	router.PUT("/update-order-status/:order_id/", order_controllers.GetOrderDetails)
}
