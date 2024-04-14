package meta_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteInfo struct {
	Path        string
	Method      string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

type RouteConfig struct {
	Version   string
	GroupName string
	Routes    []RouteInfo
}

func RegisterRoutes[M Model](router *RouteHandler[M], config RouteConfig) {
	api := router.Router.Group("/api")
	if config.Version != "" {
		api = api.Group(config.Version)
	}
	if config.GroupName != "" {
		api = api.Group(config.GroupName)
	}
	for _, route := range config.Routes {
		handlers := make([]gin.HandlerFunc, 0, len(route.Middlewares)+1)
		handlers = append(handlers, route.Middlewares...)
		handlers = append(handlers, route.Handler)
		switch route.Method {
		case http.MethodGet:
			api.GET(route.Path, handlers...)
		case http.MethodPost:
			api.POST(route.Path, handlers...)
		case http.MethodPut:
			api.PUT(route.Path, handlers...)
		case http.MethodPatch:
			api.PATCH(route.Path, handlers...)
		case http.MethodDelete:
			api.DELETE(route.Path, handlers...)
		}
	}
}

type RouteHandler[M Model] struct {
	Handler Handler[M]
	Router  *gin.Engine
}

func NewRouteHandler[M Model](
	handler Handler[M],
	router *gin.Engine,
) *RouteHandler[M] {
	return &RouteHandler[M]{
		Handler: handler,
		Router:  router,
	}
}
