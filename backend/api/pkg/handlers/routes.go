package handlers

import (
	"api/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	r.RedirectTrailingSlash = false
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Baggage", "Sentry-Trace"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
			report := v1.Group("/report", h.userIdentify)
			{
				report.POST("/", h.reportOfCall)

			}
		}

	}
	return r
}
