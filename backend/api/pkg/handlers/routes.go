package handlers

import (
	"api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	api := r.Group("/api", h.requestRateLimit)
	{
		v1 := api.Group("/v1")
		{
			health := v1.Group("/health")
			{
				health.GET("/", h.checkHealth)
			}
			auth := v1.Group("/auth")
			{
				auth.POST("/sign-up", h.signUp)
				auth.POST("/sign-in", h.signIn)
				auth.POST("/forgot", h.forgotPassword)
				auth.POST("/reset", h.resetPassword)
			}
			user := v1.Group("/user", h.userIdentify)
			{
				user.GET("/ping", h.Ping)
				user.GET("/", h.getUserInfo)
			}

		}

	}
	return r
}
