package main

import (
	"mentors/config"
	"mentors/internal/app/user"
	"mentors/internal/routes"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// migrate user model
	database.DB.AutoMigrate(&user.User{})

	r := gin.Default()
	routes.Setup(r, cfg)

	r.Run(":8080")
}
