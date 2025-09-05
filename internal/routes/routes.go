package routes

import (
	"mentors/config"
	"mentors/internal/app/user"
	"mentors/internal/middleware"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, cfg *config.Config) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"db":     database.DB != nil,
		})
	})

	userRepo := user.NewRepository(database.DB)
	userService := user.NewService(userRepo)
	userController := user.NewController(userService, cfg)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	auth := r.Group("/api", middleware.AuthMiddleware(cfg))
	{
		auth.GET("/me", func(c *gin.Context) {
			userID := c.GetUint("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})

		auth.POST("/change-password", userController.ChangePassword)
	}
}
