package middleware

import (
	"explorekg/config"
	"explorekg/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthRequired(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authedHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.AccessToken, config.Auth.RefreshToken).ParseAccessToken(token)
		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		c.Set("ContextEmailKey", data.Email)

		c.Next()
	}
}
