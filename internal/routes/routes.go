package routes

import (
	"mentors/config"
	contactinformation "mentors/internal/app/contact_information"
	"mentors/internal/app/education"
	"mentors/internal/app/experience"
	"mentors/internal/app/language"
	"mentors/internal/app/portfolio"
	"mentors/internal/app/post"
	"mentors/internal/app/system"
	"mentors/internal/app/tag"
	"mentors/internal/app/user"
	"mentors/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, cfg *config.Config) {
	system.RegisterRoutes(r.Group("/api"), cfg)

	public := r.Group("/api")
	user.RegisterPublicRoutes(public, cfg)

	auth := r.Group("/api", middlewares.AuthMiddleware(cfg))
	user.RegisterProtectedRoutes(auth, cfg)
	tag.TagRoutes(auth, cfg)
	education.EducationRoutes(auth, cfg)
	post.PostRoutes(auth, cfg)
	language.LanguageRoutes(auth, cfg)
	portfolio.PortfolioRoutes(auth, cfg)
	contactinformation.ContactInformationRoutes(auth, cfg)
	experience.ExperienceRoutes(auth, cfg)
}
