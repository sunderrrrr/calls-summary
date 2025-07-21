package handlers

import (
	"api/models"
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

}

func (h *Handler) resetPassword(c *gin.Context) {
	var input models.UserReset
	if err := c.ShouldBindJSON(&input); err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "field validation fail")
		return
	}
}
