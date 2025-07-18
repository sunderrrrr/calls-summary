package responser

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Result string `json:"result"`
}

func NewErrorResponse(c *gin.Context, errorCode int, message string) {
	c.AbortWithStatusJSON(errorCode, ErrorResponse{Result: message})
}
