package user

import (
	"mentors/config"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.RouterGroup, cfg *config.Config) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	controller := NewController(service, cfg)

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/verify-email", controller.VerifyEmail)
	r.POST("/forgot-password", controller.ForgotPassword)
	r.POST("/reset-password", controller.ResetPassword)
}

func RegisterProtectedRoutes(r *gin.RouterGroup, cfg *config.Config) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	controller := NewController(service, cfg)

	r.GET("/me", controller.Me)
	r.POST("/change-password", controller.ChangePassword)
}
