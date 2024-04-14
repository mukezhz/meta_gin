package meta_gin

import "github.com/gin-gonic/gin"

type Handler interface {
	GetName() string
	Method() string
	Handlers() map[string]gin.HandlerFunc
}
