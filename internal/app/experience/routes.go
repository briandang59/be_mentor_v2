package experience

import (
	"mentors/config"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func ExperienceRoutes(r *gin.RouterGroup, cfg *config.Config) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	controller := NewController(service, cfg)

	r.GET("/experiences", controller.GetAll)
	r.POST("/experiences", controller.Create)
	r.PATCH("/experiences/:id", controller.UpdatePartial)
	r.DELETE("/experiences/:id", controller.Delete)

}
