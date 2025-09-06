package routes

import (
	"mentors/config"
	"mentors/internal/app/system"
	"mentors/internal/app/tag"
	"mentors/internal/app/user"
	"mentors/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, cfg *config.Config) {
	// system routes
	system.RegisterRoutes(r.Group("/api"), cfg)

	// public routes
	public := r.Group("/api")
	user.RegisterPublicRoutes(public, cfg)

	// protected routes
	auth := r.Group("/api", middleware.AuthMiddleware(cfg))
	user.RegisterProtectedRoutes(auth, cfg)
	tag.RegisterRoutes(auth, cfg)
}
