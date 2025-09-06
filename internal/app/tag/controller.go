package tag

import (
	"mentors/config"
	"mentors/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
	cfg     *config.Config
}

func NewController(s Service, cfg *config.Config) *Controller {
	return &Controller{service: s, cfg: cfg}
}

func (ctl *Controller) Create(c *gin.Context) {
	var req dto.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	t, err := ctl.service.CreateTag(req.TagName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.Success(t, "Tag created successfully"))
}

func (ctl *Controller) GetAll(c *gin.Context) {
	tags, err := ctl.service.GetAllTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(tags, "Tags retrieved successfully"))
}
