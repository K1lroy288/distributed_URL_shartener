package handler

import (
	"log"
	"net/http"
	"redirect-service/service"
	"redirect-service/utils"

	"github.com/gin-gonic/gin"
)

type RedirectHandler struct {
	service *service.RedirectService
}

func NewRedirectHandler(service *service.RedirectService) *RedirectHandler {
	return &RedirectHandler{service: service}
}

func (h *RedirectHandler) Resolve(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	claims, err := utils.ValidateJWT(token)
	if err != nil {
		log.Printf("Invalid token: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	code := ctx.Param("shortCode")
	ownerIdFloat := claims["user_id"].(float64)
	url, err := h.service.Resolve(ctx, code, int(ownerIdFloat))
	if err != nil {
		log.Printf("error git link: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url)
}
