package main

import (
	"fmt"
	"net/http"
	"shortener-service/config"
	"shortener-service/handler"
	"shortener-service/model"
	"shortener-service/repository"
	"shortener-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Url{})

	repo := repository.NewShortenerRepository(db)
	service := service.NewShortenerService(repo)
	handler := handler.NewShortenerHandler(service)

	r := gin.Default()

	api := r.Group("/short")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Shortener service is up!")
		})

		api.POST("/shortLink", handler.SaveCode)

		api.POST("/:shortCode", handler.GetLink)
	}

	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
