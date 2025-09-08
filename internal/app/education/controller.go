package education

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
	var req dto.EducationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}
	education, err := ctl.service.CreateEducation(&Education{
		Name:   req.Name,
		Object: req.Object,
		From:   req.From,
		To:     req.To,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, dto.Success(education, "Education created successfully"))
}

func (ctl *Controller) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	educations, total, err := ctl.service.GetEducationsWithPagination(limit, offset)
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
	c.JSON(http.StatusOK, dto.SuccessWithMeta(educations, "Educations retrieved successfully", meta))
}
func (ctl *Controller) UpdatePartial(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("Invalid ID"))
		return
	}
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}
	education, err := ctl.service.UpdateEducationPartial(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.Success(education, "Education updated successfully"))
}
func (ctl *Controller) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("Invalid ID"))
		return
	}
	if err := ctl.service.DeleteEducation(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil, "Education deleted successfully"))
}
