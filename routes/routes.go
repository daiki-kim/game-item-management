package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"game-item-management/controllers"
	"game-item-management/middlewares"
	"game-item-management/repositories"
	"game-item-management/services"
)

func SetupRouter(db *gorm.DB, router *gin.Engine) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	userRoutes := router.Group("/user")
	userRoutesWithAuth := router.Group("/user", middlewares.AuthMiddleware(userService))
	userRoutes.POST("/signup", userController.Signup)
	userRoutes.POST("/login", userController.Login)
	userRoutesWithAuth.POST("/profile", userController.GetUsersProfile)
	userRoutesWithAuth.GET("/profile/:id", userController.GetUserById)
	userRoutesWithAuth.PUT("/profile/update", userController.UpdateUserProfile)
	userRoutesWithAuth.DELETE("/delete", userController.DeleteUser)

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)
	itemRoutes := router.Group("/item")
	itemRoutesWithAuth := router.Group("/item", middlewares.AuthMiddleware(userService))
	itemRoutes.GET("/all", itemController.FindAllItems)
	itemRoutes.GET("/find/:id", itemController.FindItemById)
	itemRoutesWithAuth.POST("/create", itemController.CreateItem)
	itemRoutesWithAuth.PUT("/update/:id", itemController.UpdateItem)
	itemRoutesWithAuth.DELETE("/delete/:id", itemController.DeleteItem)
	itemRoutesWithAuth.GET("/user/all", itemController.FindMyAllItems)

	emailService := services.NewEmailService()
	tradeRepository := repositories.NewTradeRepository(db)
	tradeService := services.NewTradeService(itemRepository, tradeRepository, userRepository, emailService)
	tradeController := controllers.NewTradeController(tradeService)
	tradeRoutesWithAuth := router.Group("/trade", middlewares.AuthMiddleware(userService))
	tradeRoutesWithAuth.GET("/find/:id", tradeController.FindTradeByTradeId)
	tradeRoutesWithAuth.GET("/all/by/item/:id", tradeController.FindAllTradesByItemId)
	tradeRoutesWithAuth.POST("/create/:id", tradeController.CreateNewTrade)
	tradeRoutesWithAuth.PUT("/update/status/:id", tradeController.UpdateTradeStatus)
	tradeRoutesWithAuth.GET("/all/by/user/:id", tradeController.FindAllTradesByUserId)
}
