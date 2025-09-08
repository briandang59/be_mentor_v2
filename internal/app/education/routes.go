package education

import (
	"mentors/config"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func EducationRoutes(r *gin.RouterGroup, cfg *config.Config) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	controller := NewController(service, cfg)

	r.GET("/educations", controller.GetAll)
	r.POST("/educations", controller.Create)
	r.PATCH("/educations/:id", controller.UpdatePartial)
	r.DELETE("/educations/:id", controller.Delete)

}
