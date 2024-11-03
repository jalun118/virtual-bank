package middleware

import (
	"context"
	"time"
	"virtual-bank/internal/constants"
	"virtual-bank/internal/helpers"
	"virtual-bank/internal/response"

	"github.com/gin-gonic/gin"
)

func (m *Middlewares) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if !helpers.IsLogin(c) {
			response.NewResponseUnauthorized(c)
			c.Abort()
			return
		}

		sessionId := helpers.GetCookieId(c)

		session, resNum := m.sessionUsecase.FindByIdAndNotExpired(ctx, sessionId)
		if resNum != constants.SuccessRead {
			if resNum == constants.RecordNotFound {
				response.NewResponseUnauthorized(c)
				c.Abort()
				return
			}
			response.NewResponseInternalServerError(c)
			c.Abort()
			return
		}

		if _, resNum := m.userUsecase.GetById(ctx, session.UserId); resNum != constants.SuccessRead {
			if resNum == constants.RecordNotFound {
				response.NewResponseUnauthorized(c)
				c.Abort()
				return
			}
			response.NewResponseInternalServerError(c)
			c.Abort()
			return
		}
	}
}
