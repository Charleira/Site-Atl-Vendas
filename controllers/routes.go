package controllers

import (
	"atlanta-site/controllers/cart_controllers"
	"atlanta-site/controllers/order_controllers"
	payment_controller "atlanta-site/controllers/payments_controller"
	product_controllers "atlanta-site/controllers/product_controller"
	"atlanta-site/controllers/shipping_controller"
	"atlanta-site/controllers/user_controllers"
	"atlanta-site/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Rotas administrativas
	admin := router.Group("/admin")
	router.GET("/auth/discord", auth_controller.RedirectToDiscord)

	admin.Use(middlewares.AdminMiddleware())
	{
		
		admin.GET("/orders", order_controllers.ListOrders)
		admin.PUT("/orders/:id", order_controllers.TrackOrder) // Corrigido método para update
		admin.PUT("/users/:id/promote", user_controllers.UpdateUserDetails)
	}
	// Rotas de carrinho
	cart := router.Group("/cart")
	{
		cart.GET("/:user_id/", cart_controllers.GetCart)                                     // Recupera o carrinho do usuário
		cart.POST("/:user_id/add/:product_id/", cart_controllers.AddProductToCart)           // Adiciona produto ao carrinho
		cart.DELETE("/:user_id/clear", cart_controllers.ClearCart)                           // Limpa o carrinho
		cart.DELETE("/:user_id/remove/:product_id/", cart_controllers.RemoveProductFromCart) // Remove produto do carrinho
	}

	// Rotas de pedidos
	orders := router.Group("/orders")
	{
		orders.POST("/", order_controllers.CreateOrder)                 // Cria um novo pedido
		orders.GET("/:order_id", order_controllers.GetOrderDetails)     // Detalhes de um pedido específico
		orders.POST("/:order_id/cancel", order_controllers.CancelOrder) // Cancela um pedido específico
		orders.GET("/:order_id/track", order_controllers.TrackOrder)    // Rastreia um pedido específico
	}

	// Rotas de pagamentos
	payments := router.Group("/payments")
	{
		payments.POST("/create", payment_controller.CreatePaymentIntent) // Webhook para notificações do Stripe
	}

	// Rotas de produtos
	products := router.Group("/products")
	{
		products.GET("/", product_controllers.ListProducts)   // Lista todos os produtos
		products.POST("/", product_controllers.CreateProduct) // Cria um novo produto
	}

	// Rotas de produtos com ID
	productsWithID := router.Group("/products/:product_id")
	{
		productsWithID.GET("/", product_controllers.GetProductById)   // Recupera um produto pelo ID
		productsWithID.PUT("/", product_controllers.UpdateProduct)    // Atualiza as informações de um produto
		productsWithID.DELETE("/", product_controllers.RemoveProduct) // Deleta um produto pelo ID
	}

	// Rotas de envio
	shipping := router.Group("/shipping")
	{
		shipping.POST("/create", shipping_controller.CreateShipping)     // Cria um novo pedido de envio
		shipping.GET("/options", shipping_controller.GetShippingOptions) // Obtém as opções de envio disponíveis
	}

	// Rotas de pedidos de um usuário
	userOrders := router.Group("/users/orders/:user_id")
	{
		userOrders.GET("/", user_controllers.GetUserDetails) // Lista pedidos do usuário
	}

	// Rotas de usuários
	users := router.Group("/users")
	{
		// Alterações nas rotas de usuários para evitar conflito
		users.GET("/:user_id", user_controllers.GetUserDetails)    // Retorna informações do usuário
		users.PUT("/:user_id", user_controllers.UpdateUserDetails) // Atualiza informações do usuário
		users.DELETE("/:user_id", user_controllers.DeleteUser)     // Deleta usuário

		// Alteração de senha
		users.PUT("/:user_id/password", user_controllers.ChangePassword) // Altera a senha de um usuário
	}
}
