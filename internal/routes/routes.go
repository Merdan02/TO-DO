package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/handler"
	"todo-app/internal/middleware"
)

func SetupRoutes(handler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/Login", handler.Login)

	userRouter := router.Group("/users", middleware.Middleware())
	{
		userRouter.POST("/", handler.CreateUser)
		userRouter.GET("/", handler.GetAllUser)
		userRouter.PUT("/:id", handler.UpdateUser)
		userRouter.GET("/:id", handler.GetUserByID)
		userRouter.GET("/name/:name", handler.GetUserByName)
	}
	adminGroup := router.Group("/admin/users", middleware.Middleware(), middleware.AdminOnly())
	{
		adminGroup.DELETE("/:id", handler.DeleteUser)

	}
	return router

}
