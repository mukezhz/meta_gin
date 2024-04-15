package meta_gin

import "github.com/gin-gonic/gin"

type Decorator interface {
	Decorate(handler gin.HandlerFunc) gin.HandlerFunc
}
