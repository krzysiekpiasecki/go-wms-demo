package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kpiasecki/wms/internal/logger"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error().
					Interface("panic", err).
					Msg("panic recovered")

				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					gin.H{
						"error": "internal server error",
					},
				)
			}
		}()

		c.Next()
	}
}
