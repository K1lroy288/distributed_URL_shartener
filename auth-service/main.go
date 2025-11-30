package main

import (
	"auth-service/config"
	"auth-service/handler"
	"auth-service/model"
	"auth-service/repository"
	"auth-service/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	repo := repository.NewUserRepository(db)
	service := service.NewAuthService(repo)
	handler := handler.NewAuthHandler(service)

	r := gin.Default()

	api := r.Group("/auth")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Auth service is up!")
		})

		api.POST("/login", handler.Login)

		api.POST("/register", handler.Register)
	}

	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
