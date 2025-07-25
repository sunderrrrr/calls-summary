package handlers

import (
	"api/models"
	"api/pkg/utils/logger"
	"api/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Ping(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "failed to get user id")
		return
	}
	c.JSON(200, gin.H{"pong": id})
}

func (h *Handler) getUserInfo(c *gin.Context) {
	user, err := getUserInfo(c)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "failed to get user info")
	}
	c.JSON(200, gin.H{"user": user})
}

func (h *Handler) forgotPassword(c *gin.Context) {
	var input models.ResetRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "field validation fail")
		return
	}
	if err := h.service.User.ForgotPassword(input); err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "failed to process forgot password request")
		logger.Log.Errorf("failed to forgot password request: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "reset link sent to your email"})

}

func (h *Handler) resetPassword(c *gin.Context) {
	var input models.UserReset
	if err := c.ShouldBindJSON(&input); err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "field validation fail")
		return
	}
	if err := h.service.User.ResetPassword(input); err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "failed to reset password")
		logger.Log.Errorf("failed to reset password: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})
}
