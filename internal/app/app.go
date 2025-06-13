package app

import (
	"database/sql"
	"go.uber.org/zap"
	"log"
	"todo-app/internal/database"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/routes"
	"todo-app/internal/service"
)

func InitApp(db *sql.DB, log *zap.Logger) *handler.UserHandler {
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserServ(userRepo, log)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}

func Run() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database", zap.Error(err))
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("failed to close db", zap.Error(err))
		}
	}(db)

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("failed to initialize logger", zap.Error(err))
		return
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal("failed to sync logger", zap.Error(err))
		}
	}(logger)

	userHandler := InitApp(db, logger)

	r := routes.SetupRoutes(userHandler)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server", zap.Error(err))
		return
	}

}
