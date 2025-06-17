package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/handler"
)

func SetupRoutes(handler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	userRouter.POST("/", handler.CreateUser)
	userRouter.GET("/", handler.GetAllUser)
	userRouter.PUT("/:id", handler.UpdateUser)
	userRouter.GET("/:id", handler.GetUserByID)
	userRouter.GET("/name/:name", handler.GetUserByName)
	userRouter.DELETE("/:id", handler.DeleteUser)

	return router

}
