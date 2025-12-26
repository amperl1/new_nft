package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

func StatusUnauthorized(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}
