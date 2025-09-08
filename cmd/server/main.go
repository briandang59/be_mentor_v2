package main

import (
	"mentors/config"
	"mentors/internal/app/attachment"
	"mentors/internal/app/contract"
	coverletter "mentors/internal/app/cover_letter"
	"mentors/internal/app/education"
	"mentors/internal/app/language"
	"mentors/internal/app/portfolio"
	"mentors/internal/app/post"
	"mentors/internal/app/tag"
	"mentors/internal/app/user"
	"mentors/internal/routes"
	"mentors/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	database.DB.AutoMigrate(
		&user.User{},
		&tag.Tag{},
		&post.Post{},
		&coverletter.CoverLetter{},
		&attachment.Attachment{},
		&contract.Contract{},
		&education.Education{},
		&language.Language{},
		&portfolio.Portfolio{},
	)

	r := gin.Default()
	routes.Setup(r, cfg)

	r.Run(`:` + cfg.BackendPort)
}
