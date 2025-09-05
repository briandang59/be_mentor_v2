package routes

import (
	"mentors/config"
	"mentors/internal/app/user"
	"mentors/internal/jobs"
	"mentors/internal/middleware"
	"mentors/internal/utils"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, cfg *config.Config) {
	// Health check
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"db":     database.DB != nil,
		})
	})

	// Test email
	r.GET("/api/test-email", func(c *gin.Context) {
		to := c.Query("to")
		body, err := utils.RenderTemplate("internal/templates/verify_email.html", map[string]string{
			"VerifyURL": "http://localhost:8080/verify-email?token=dummytoken",
		})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		sender := jobs.NewEmailSender(cfg)
		if err := sender.Send(to, "Test Email from Mentors", body); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "✅ Email sent to " + to})
	})

	userRepo := user.NewRepository(database.DB)
	userService := user.NewService(userRepo)
	userController := user.NewController(userService, cfg)

	// 📌 Public routes (không cần JWT)
	public := r.Group("/api")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
		public.GET("/verify-email", userController.VerifyEmail)
		public.POST("/forgot-password", userController.ForgotPassword)
		public.POST("/reset-password", userController.ResetPassword)
	}

	// 📌 Protected routes (cần JWT)
	auth := r.Group("/api", middleware.AuthMiddleware(cfg))
	{
		auth.GET("/me", func(c *gin.Context) {
			userID := c.GetUint("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})

		auth.POST("/change-password", userController.ChangePassword)
	}
}
