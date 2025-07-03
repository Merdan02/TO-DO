package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/handler"
	"todo-app/internal/middleware"
)

func SetupRoutes(user *handler.UserHandler, task *handler.TaskHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/Login", user.Login)
	router.POST("/Register", user.CreateUser)

	userRouter := router.Group("/users", middleware.Middleware())
	{

		userRouter.GET("/", user.GetAllUser)
		userRouter.PUT("/:id", user.UpdateUser)
		userRouter.GET("/:id", user.GetUserByID)
		userRouter.GET("/name/:name", user.GetUserByName)
	}
	adminGroup := router.Group("/admin/users", middleware.Middleware(), middleware.AdminOnly())
	{
		adminGroup.DELETE("/:id", user.DeleteUser)

	}

	taskGroup := router.Group("/tasks", middleware.Middleware())
	{
		taskGroup.POST("/", task.CreateTask)
		taskGroup.GET("/", task.GetAllTasks)
		taskGroup.GET("/:id", task.GetByID)
		taskGroup.PUT("/:id", task.UpdateTask)
		taskGroup.DELETE(":id", task.DeleteByID)
	}
	return router

}
