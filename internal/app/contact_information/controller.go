package contactinformation

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

func NewController(service Service, cfg *config.Config) *Controller {
	return &Controller{service: service, cfg: cfg}
}

func (ctl *Controller) Create(c *gin.Context) {
	var req dto.ContactInformationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	ci := &ContactInformation{
		Phone:    req.Phone,
		LinkedIn: req.LinkedIn,
		Facebook: req.Facebook,
		Zalo:     req.Zalo,
		Telegram: req.Telegram,
		X:        req.X,
		GitHub:   req.GitHub,
		Twitter:  req.Twitter,
		Website:  req.Website,
		Location: req.Location,
	}

	created, err := ctl.service.Create(ci)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.Success(created, "Contact information created successfully"))
}

func (ctl *Controller) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	items, total, err := ctl.service.FindWithPagination(limit, offset)
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
	c.JSON(http.StatusOK, dto.SuccessWithMeta(items, "Contact informations retrieved successfully", meta))
}

func (ctl *Controller) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ci, err := ctl.service.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Fail("Contact information not found"))
		return
	}
	c.JSON(http.StatusOK, dto.Success(ci, "Contact information retrieved successfully"))
}

func (ctl *Controller) UpdatePartial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}
	ci, err := ctl.service.UpdateFields(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.Success(ci, "Contact information updated successfully"))
}

func (ctl *Controller) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctl.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil, "Contact information deleted successfully"))
}
