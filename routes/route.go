package routes

import (
	"golang-rest-api/controllers"
	_ "golang-rest-api/docs"
	"golang-rest-api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

	func SetupRouter() {
		r := gin.Default()

		// Prefix untuk API
		api := r.Group("/api/v1")

		// Auth
		auth := api.Group("/auth")
		{
			authController := controllers.NewAuthController()
			auth.POST("/signup", authController.SignUp)
			auth.POST("/login", authController.Login)
			auth.POST("/logout", middleware.RequireAuth, authController.Logout)
		}

		// Merchant
		merchant := api.Group("/merchants")
		{
			merchant.POST("/", middleware.RequireAuth, controllers.CreateMerchant)
			merchant.GET("/", middleware.RequireAuth, controllers.GetAllMerchants)
			merchant.GET("/:id", middleware.RequireAuth, controllers.GetMerchantByID)
			merchant.PUT("/:id", middleware.RequireAuth, controllers.UpdateMerchant) 
			merchant.DELETE("/:id", middleware.RequireAuth, controllers.DeleteMerchant)
		}

		// Transaction
		transaction := api.Group("/transactions")
		{
			transactionController := controllers.NewTransactionController()
			transaction.POST("/", middleware.RequireAuth, transactionController.CreateTransaction)
			transaction.GET("/:id", transactionController.GetTransactionByID)
			transaction.GET("/", transactionController.GetAllTransactions)
			transaction.DELETE("/:id", transactionController.DeleteTransaction)
		}

		// Swagger
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		r.Run()
	}
