package tag

import (
	"math"
	"mentors/config"
	"mentors/internal/dto"
	"net/http"
	"strconv"

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	tags, total, err := ctl.service.GetTagsWithPagination(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	meta := &dto.Meta{
		Page:       page,
		Limit:      limit,
		Total:      int(total),
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, dto.SuccessWithMeta(tags, "Tags retrieved successfully", meta))
}
