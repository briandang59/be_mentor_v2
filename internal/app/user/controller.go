package user

import (
	"mentors/config"
	"mentors/internal/dto"
	"mentors/internal/utils"
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

func (ctl *Controller) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	u, err := ctl.service.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.Success(u, "User registered successfully"))
}

func (ctl *Controller) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	u, err := ctl.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.Fail("invalid credentials"))
		return
	}

	token, _ := utils.GenerateJWT(u.ID, ctl.cfg)

	resp := gin.H{
		"token": token,
		"user":  u,
	}

	c.JSON(http.StatusOK, dto.Success(resp, "Login successful"))
}

func (ctl *Controller) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	// Lấy user_id từ middleware JWT
	userID := c.GetUint("user_id")

	if err := ctl.service.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(nil, "Password changed successfully"))
}
