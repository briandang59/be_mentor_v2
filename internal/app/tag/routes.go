package tag

import (
	"mentors/config"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, cfg *config.Config) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	controller := NewController(service, cfg)

	r.GET("/tags", controller.GetAll)
	r.POST("/tags", controller.Create)
	r.PATCH("/tags/:id", controller.UpdatePartial)
	r.DELETE("/tags/:id", controller.Delete)

}
