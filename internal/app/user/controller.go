package user

import (
	"mentors/config"
	"mentors/internal/dto"
	"mentors/internal/jobs"
	"mentors/internal/utils"
	"mentors/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

	// 🔑 Tạo token verify email
	token, _ := utils.GenerateEmailVerifyToken(u.ID, ctl.cfg.JWTSecret)
	verifyURL := "http://localhost:8080/verify-email?token=" + token

	// 🔑 Render template email
	body, err := utils.RenderTemplate("internal/templates/verify_email.html", map[string]string{
		"Username":  u.Username,
		"VerifyURL": verifyURL,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail("failed to render email template"))
		return
	}
	// 🔑 Gửi email
	emailSender := jobs.NewEmailSender(ctl.cfg)
	if err := emailSender.Send(u.Email, "Verify your Mentors account", body); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail("User created but failed to send email"))
		return
	}

	c.JSON(http.StatusCreated, dto.Success(u, "User registered successfully, please check your email to verify your account"))
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

func (ctl *Controller) VerifyEmail(c *gin.Context) {
	tokenStr := c.Query("token")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(ctl.cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusBadRequest, dto.Fail("invalid or expired token"))
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["purpose"] != "verify_email" {
		c.JSON(http.StatusBadRequest, dto.Fail("wrong token type"))
		return
	}

	userID := uint(claims["user_id"].(float64))
	if err := database.DB.Model(&User{}).
		Where("id = ?", userID).
		Update("is_verified", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail("failed to verify email"))
		return
	}

	c.JSON(http.StatusOK, dto.Success(nil, "Email verified successfully"))
}

func (ctl *Controller) ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	var u User
	if err := database.DB.Where("email = ?", req.Email).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Fail("email not found"))
		return
	}

	token, _ := utils.GenerateResetToken(u.ID, ctl.cfg.JWTSecret)
	resetURL := "http://localhost:8080/reset-password?token=" + token

	body, _ := utils.RenderTemplate("internal/templates/reset_password.html", map[string]string{
		"ResetURL": resetURL,
	})

	emailSender := jobs.NewEmailSender(ctl.cfg)
	emailSender.Send(u.Email, "Reset your Mentors account password", body)

	c.JSON(http.StatusOK, dto.Success(nil, "Reset link sent to email"))
}

func (ctl *Controller) ResetPassword(c *gin.Context) {
	var req struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(err.Error()))
		return
	}

	token, err := jwt.Parse(req.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(ctl.cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusBadRequest, dto.Fail("invalid or expired token"))
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["purpose"] != "reset_password" {
		c.JSON(http.StatusBadRequest, dto.Fail("wrong token type"))
		return
	}

	userID := uint(claims["user_id"].(float64))
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 12)
	if err := database.DB.Model(&User{}).
		Where("id = ?", userID).
		Update("password", string(hash)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail("failed to reset password"))
		return
	}

	c.JSON(http.StatusOK, dto.Success(nil, "Password reset successfully"))
}
