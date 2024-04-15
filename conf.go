package meta_gin

import (
	"github.com/gin-gonic/gin"
)

type SetupConfig[M Model] struct {
	Router      *gin.Engine
	Config      *Config
	Version     string
	GroupName   string
	Middlewares []gin.HandlerFunc
	Decorators  []Decorator
	routeConfig *RouteConfig
}

type GenericConfig[M Model] struct {
	SetupConfig[M]
	Handlers []Handler[M]
}

func SetupGenericRouteForModel[M Model](
	setupConfig GenericConfig[M],
) {
	routes := []RouteInfo{}
	routeHandler := &RouteHandler[M]{}
	for _, handler := range setupConfig.Handlers {
		routeHandler = NewRouteHandler[M](handler, setupConfig.Router)
		for p, h := range handler.Handlers() {
			routes = append(routes, RouteInfo{
				Path:        p,
				Method:      handler.Method(),
				Handler:     AddDecorators(h, setupConfig.Decorators),
				Middlewares: setupConfig.Middlewares,
			})
		}
	}
	if setupConfig.routeConfig == nil {
		setupConfig.routeConfig = &RouteConfig{
			Version:   setupConfig.Version,
			GroupName: setupConfig.GroupName,
			Routes:    routes,
		}
	}
	RegisterRoutes(routeHandler, *setupConfig.routeConfig)
}

func AddDecorators(handler gin.HandlerFunc, decorators []Decorator) gin.HandlerFunc {
	for _, decorator := range decorators {
		handler = decorator.Decorate(handler)
	}
	return handler
}
