package handlers

import (
	"api/pkg/utils/responser"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	requests     = make(map[string]int)
	blockedUntil = make(map[string]time.Time)
)

const (
	maxRequests   = 5                // Сколько запросов можно сделать
	duration      = 5 * time.Second  // За какой период
	blockDuration = 10 * time.Second // Бан при превышении
)

func (h *Handler) requestRateLimit(c *gin.Context) {
	ip := c.ClientIP()
	now := time.Now()

	if until, ok := blockedUntil[ip]; ok && now.Before(until) {
		responser.NewErrorResponse(c, http.StatusBadRequest, "temp ban, wait")
		return
	}
	requests[ip]++
	if requests[ip] == 1 {
		go func(ip string) {
			time.Sleep(duration)
			delete(requests, ip)
		}(ip)
	}
	if requests[ip] > maxRequests {
		blockedUntil[ip] = now.Add(blockDuration)
		delete(requests, ip) // сбрасываем счётчик
		responser.NewErrorResponse(c, http.StatusTooManyRequests, "too many requests")
		return
	}

	c.Next()
}
