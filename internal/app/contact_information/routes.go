package contactinformation

import (
	"mentors/config"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func ContactInformationRoutes(r *gin.RouterGroup, cfg *config.Config) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	controller := NewController(service, cfg)

	r.GET("/contact-information", controller.GetAll)
	r.POST("/contact-information", controller.Create)
	r.PATCH("/contact-information/:id", controller.UpdatePartial)
	r.DELETE("/contact-information/:id", controller.Delete)

}
