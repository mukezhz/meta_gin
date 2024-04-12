package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Permissions struct {
	CanView   bool `toml:"can_view"`
	CanEdit   bool `toml:"can_edit"`
	CanDelete bool `toml:"can_delete"`
}

func AuthMiddleware(config *Config, userRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const userID = 1
		if permissions, ok := config.Roles[userRole]; ok {
			userPermissions := make(map[string]bool)
			for _, p := range permissions {
				userPermissions[p] = true
			}
			c.Set("permissions", userPermissions)
			c.Set("userID", userID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role not found"})
		}
	}
}
