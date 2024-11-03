package middleware

import (
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/response"

	"github.com/gin-gonic/gin"
)

func (m *Middlewares) NotLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if helpers.IsLogin(c) {
			response.NewResponseUnauthorized(c)
			c.Abort()
			return
		}
	}

}
