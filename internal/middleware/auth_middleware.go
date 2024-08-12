package middleware

import (
	"net/http"
	"server/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.String(http.StatusForbidden, "Forbidden")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.String(http.StatusForbidden, "Forbidden")
			return
		}

		tokenStr := parts[1]
		claims, err := auth.ValidateToken(tokenStr)
		if err != nil {
			c.String(http.StatusForbidden, "Forbidden")
			return
		}

		c.Set("id", claims.ID)

		c.Next()
	}
}
