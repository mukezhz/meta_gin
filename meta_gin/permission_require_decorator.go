package meta_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PermissionDecorator struct {
}

func NewPermissionDecorator() *PermissionDecorator {
	return &PermissionDecorator{}
}

func (d *PermissionDecorator) Decorate(handler gin.HandlerFunc) gin.HandlerFunc {
	requiredPermission := ""
	return func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			requiredPermission = "can_view"
		case http.MethodPost:
			requiredPermission = "can_edit"
		case http.MethodDelete:
			requiredPermission = "can_delete"
		case http.MethodPut:
			requiredPermission = "can_edit"
		case http.MethodPatch:
			requiredPermission = "can_edit"
		case http.MethodOptions:
			requiredPermission = "can_view"
		default:
			requiredPermission = "can_view"
		}

		permissions, exists := c.Get("permissions")
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Permissions not found"})
			return
		}
		perms := permissions.(map[string]bool)

		if !perms[requiredPermission] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}
		handler(c)
	}
}
