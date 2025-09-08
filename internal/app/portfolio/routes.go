package portfolio

import (
	"mentors/config"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func PortfolioRoutes(r *gin.RouterGroup, cfg *config.Config) {
	repo := NewRepository(database.DB)
	service := NewService(repo)
	controller := NewController(service, cfg)

	r.GET("/portfolios", controller.GetAll)
	r.POST("/portfolios", controller.Create)
	r.PATCH("/portfolios/:id", controller.UpdatePartial)
	r.DELETE("/portfolios/:id", controller.Delete)

}
