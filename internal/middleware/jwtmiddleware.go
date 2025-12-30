package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/dto"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/service"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware saves user ID to Gin context
func JWTMiddleware(jwtService *service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwtService.ValidateToken(dto.JWTPayload{Token: tokenString})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims, ok := token.Claims.(*service.TokenPayload)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}
		c.Set("id_user", claims.UserID)
		c.Next()
	}
}

// JWTMiddlewareWithToken saves user ID and token to Gin context (for registration)
func JWTMiddlewareWithToken(jwtService *service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		fmt.Println(authHeader)

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwtService.ValidateToken(dto.JWTPayload{Token: tokenString})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims, ok := token.Claims.(*service.TokenPayload)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}
		c.Set("id_user", claims.UserID)
		c.Set("token", tokenString)
		c.Next()
	}
}
