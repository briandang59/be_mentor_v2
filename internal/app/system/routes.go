package system

import (
	"mentors/config"
	"mentors/internal/jobs"
	"mentors/internal/utils"
	"mentors/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, cfg *config.Config) {
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"db":     database.DB != nil,
		})
	})

	// Test email
	r.GET("/test-email", func(c *gin.Context) {
		to := c.Query("to")
		body, err := utils.RenderTemplate("internal/templates/verify_email.html", map[string]string{
			"VerifyURL": "http://localhost:8080/verify-email?token=dummytoken",
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		sender := jobs.NewEmailSender(cfg)
		if err := sender.Send(to, "Test Email from Mentors", body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "✅ Email sent to " + to})
	})
}
