package routes

import (
	"database/sql"
	"jwt-try/internal/handler"
	"jwt-try/internal/repository"
	"jwt-try/internal/usecase"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(database *sql.DB) *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	userRepo := repository.NewUserRepo(database)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", userHandler.RegisterUser)
		userRoutes.POST("/login", userHandler.CheckCredential)
		userRoutes.GET("/view", userHandler.ViewWeb)
	}

	return router
}
