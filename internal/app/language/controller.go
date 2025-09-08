package language

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
	var req dto.LanguageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}
	lang, err := ctl.service.CreateLanguage(&Language{
		Name: req.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, dto.Success(lang, "Language created successfully"))
}

func (ctl *Controller) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	langs, total, err := ctl.service.GetLanguagesWithPagination(limit, offset)
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
	c.JSON(http.StatusOK, dto.SuccessWithMeta(langs, "Languages retrieved successfully", meta))
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
	lang, err := ctl.service.UpdateLanguagePartial(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.Success(lang, "Language updated successfully"))
}

func (ctl *Controller) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail("Invalid ID"))
		return
	}
	if err := ctl.service.DeleteLanguage(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil, "Language deleted successfully"))
}
