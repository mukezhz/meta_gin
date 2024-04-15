package meta_gin

import "github.com/gin-gonic/gin"

type Handler[M Model] interface {
	GetName() string
	Method() string
	Handlers() map[string]gin.HandlerFunc
}
