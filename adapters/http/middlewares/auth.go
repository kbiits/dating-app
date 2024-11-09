package middlewares

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kbiits/dealls-take-home-test/config"
	"github.com/kbiits/dealls-take-home-test/domain/entity"
	jwt_util "github.com/kbiits/dealls-take-home-test/utils/jwt"
)

func RequireUserAuth(
	jwtUtil *jwt_util.JwtUtil,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
		jwtParsed, isValid := jwtUtil.VerifyToken(ctx, bearerToken)
		if !isValid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userId, _ := jwtParsed.Claims.GetSubject()
		c.Set("user_id", userId)

		ctxCopy := c.Request.Context()
		ctxCopy = context.WithValue(ctxCopy, entity.CtxUserID, userId)
		c.Request = c.Request.WithContext(ctxCopy)

		c.Next()
	}
}

func RequireInternalAuth(
	cfg config.InternalConfig,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("X-API-KEY")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "X-API-KEY is required"})
			c.Abort()
			return
		}

		if authHeader != cfg.APIKey {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
