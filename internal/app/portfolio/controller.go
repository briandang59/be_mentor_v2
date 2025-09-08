package portfolio

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
	var req dto.PortfolioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	pf := &Portfolio{
		Title:       req.Title,
		Description: req.Description,
		URL:         req.Url,
		JobTitle:    req.JobTitle,
		TimePeriod:  req.TimePeriod,
		IsDraft:     req.IsDraft,
		ThumbnailID: req.ThumbnailID,
	}

	createdPf, err := ctl.service.CreatePortfolio(pf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.Success(createdPf, "Portfolio created successfully"))
}

func (ctl *Controller) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	pfs, total, err := ctl.service.GetPortfoliosWithPagination(limit, offset)
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
	c.JSON(http.StatusOK, dto.SuccessWithMeta(pfs, "Portfolios retrieved successfully", meta))
}

func (ctl *Controller) UpdatePartial(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}
	pf, err := ctl.service.UpdatePortfolioPartial(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.Success(pf, "Portfolio updated successfully"))
}

func (ctl *Controller) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctl.service.DeletePortfolio(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, dto.Success(nil, "Portfolio deleted successfully"))
}

func (ctl *Controller) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pf, err := ctl.service.GetPortfolioByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Fail("Portfolio not found"))
		return
	}
	c.JSON(http.StatusOK, dto.Success(pf, "Portfolio retrieved successfully"))
}
