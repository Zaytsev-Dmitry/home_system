package middleware

import (
	"context"
	apikitErr "github.com/Zaytsev-Dmitry/apikit/custom_errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"userService/configs"
	"userService/internal/app/ports/out/keycloak"
)

var allowedPaths = []string{
	"/docs",
	"/public",
	"/spec",
}

func isAllowed(path string) bool {
	for _, allowed := range allowedPaths {
		if strings.Contains(path, allowed) {
			return true
		}
	}
	return false
}

func TokenIntrospectionMiddleware(keycloakClient *keycloak.KeycloakClient, config *configs.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isAllowed(c.Request.URL.Path) {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, apikitErr.GetErrorDto(c, apikitErr.Unauthorized))
				return
			}
			token := strings.TrimPrefix(authHeader, "Bearer ")
			ctx := context.Background()
			introspect, err := keycloakClient.Introspect(ctx, config, token)
			if err != nil || introspect == nil || !*introspect.Active {
				c.AbortWithStatusJSON(http.StatusUnauthorized, apikitErr.GetErrorDto(c, apikitErr.Unauthorized))
			}
		}
	}
}
