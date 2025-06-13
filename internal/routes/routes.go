package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/handler"
)

func SetupRoutes(handler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	userRouter.POST("/", handler.CreateUser)

	return router

}
