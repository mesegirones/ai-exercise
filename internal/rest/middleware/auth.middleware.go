package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type OAuthProxy interface {
	AddAuthToContext(ctx context.Context, authHeader string) context.Context
}

// AuthContext sets the middleware for auth context
func AuthContext(oAuthProxy OAuthProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		ctx := oAuthProxy.AddAuthToContext(c.Request.Context(), authHeader)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
