package middleware

import (
	"time"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/response"

	"github.com/gin-gonic/gin"
)

func (m *Middlewares) TokenGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !helpers.IsLogin(c) {
			response.NewResponseUnauthorized(c)
			c.Abort()
			return
		}

		tokenAuth := helpers.GetCookieRefreshToken(c)

		payload, err := helpers.DecodeRefreshToken(tokenAuth)
		if err != nil {
			response.NewResponseUnauthorized(c)
			c.Abort()
			return
		}

		if time.Now().After(payload.ExpiredAt) {
			response.NewResponseUnauthorized(c)
			c.Abort()
			return
		}
	}

}
