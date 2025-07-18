package handlers

import (
	"api/models"
	"api/pkg/utils/logger"
	"api/pkg/utils/responser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		responser.NewErrorResponse(c, 400, "invalid input data")
		return
	}
	id, err := h.service.Auth.SignUp(input)
	if err != nil {
		responser.NewErrorResponse(c, 500, "sign up failed")
		logger.Log.Errorf("sign up error: %v", err)
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		responser.NewErrorResponse(c, 400, "invalid input data")
		return
	}
	token, err := h.service.Auth.GenerateToken(input)
	if err != nil {
		responser.NewErrorResponse(c, 500, "sign in failed")
		logger.Log.Errorf("sign in error: %v", err)
		return
	}
	c.JSON(200, gin.H{"token": token})
}
