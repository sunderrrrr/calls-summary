package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) signUp(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Sign up successful"})

}
func (h *Handler) signIn(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Sign in successful"})
}
