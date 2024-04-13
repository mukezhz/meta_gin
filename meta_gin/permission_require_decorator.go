package meta_gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPermissionDecorator(handler gin.HandlerFunc) gin.HandlerFunc {
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Permissions not found"})
			return
		}
		perms := permissions.(map[string]bool)
		log.Println(perms, requiredPermission)
		if !perms[requiredPermission] {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}
		handler(c)
	}
}
