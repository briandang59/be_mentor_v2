package language

import (
	"mentors/config"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func LanguageRoutes(r *gin.RouterGroup, cfg *config.Config) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	controller := NewController(service, cfg)

	r.GET("/languages", controller.GetAll)
	r.POST("/languages", controller.Create)
	r.PATCH("/languages/:id", controller.UpdatePartial)
	r.DELETE("/languages/:id", controller.Delete)

}
