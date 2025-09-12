package experience

import (
	"net/http"
	"strconv"

	"mentors/config"
	"mentors/internal/dto"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
	cfg     *config.Config
}

func NewController(s Service, cfg *config.Config) *Controller {
	return &Controller{service: s, cfg: cfg}
}

// POST /experiences
func (ctl *Controller) Create(c *gin.Context) {
	var req dto.ExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	experience := &Experience{
		Title:       req.Title,
		Description: req.Description,
		Company:     req.Company,
		Location:    req.Location,
		From:        req.From,
		To:          req.To,
		IsCurrent:   req.IsCurrent,
	}

	if err := ctl.service.CreateExperience(experience); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.Success(experience, "experience created successfully"))
}

// GET /experiences?offset=0&limit=10
func (ctl *Controller) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	experiences, total, err := ctl.service.GetExperiencesWithPagination(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(gin.H{
		"data":  experiences,
		"total": total,
	}, "experiences fetched successfully"))
}

// PATCH /experiences/:id
func (ctl *Controller) UpdatePartial(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("invalid id"))
		return
	}

	fields := make(map[string]interface{})
	if err := c.ShouldBindJSON(&fields); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	experience, err := ctl.service.UpdateExperiencePartial(uint(id), fields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(experience, "experience updated successfully"))
}

// DELETE /experiences/:id
func (ctl *Controller) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("invalid id"))
		return
	}

	exp, err := ctl.service.GetExperienceByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Fail("experience not found"))
		return
	}

	if err := ctl.service.DeleteExperience(exp); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success("deleted", "experience deleted successfully"))
}
