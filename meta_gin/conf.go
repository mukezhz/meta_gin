package meta_gin

import (
	"github.com/gin-gonic/gin"
)

type SetupConfig[M Model, ReqDTO, ResDTO any] struct {
	Router      *gin.Engine
	Config      *Config
	Version     string
	GroupName   string
	Middlewares []gin.HandlerFunc
	Decorators  []Decorator
	routeConfig *RouteConfig
}

// CRUD
// type CRUDConfig[M Model, ReqDTO, ResDTO any] struct {
// 	SetupConfig[M, ReqDTO, ResDTO]
// 	Handler *CRUDHandler[M, ReqDTO, ResDTO]
// }

// func SetupCRUDRoutesForModel[M Model, ReqDTO, ResDTO any](
// 	setupConfig CRUDConfig[M, ReqDTO, ResDTO],
// ) {
// 	routeHandler := NewRouteHandler[M, ReqDTO, ResDTO](setupConfig.Handler, setupConfig.Router)
// 	handler := setupConfig.Handler
// 	if setupConfig.routeConfig == nil {
// 		setupConfig.routeConfig = &RouteConfig{
// 			Version:   setupConfig.Version,
// 			GroupName: setupConfig.GroupName,
// 			Routes: []RouteInfo{
// 				{
// 					Path:        "/",
// 					Method:      "POST",
// 					Handler:     AddDecorators(handler.Create(nil), setupConfig.Decorators),
// 					Middlewares: setupConfig.Middlewares,
// 				},

// 				{
// 					Path:        "/",
// 					Method:      "GET",
// 					Handler:     AddDecorators(handler.ListPagination(nil), setupConfig.Decorators),
// 					Middlewares: setupConfig.Middlewares,
// 				},
// 				{
// 					Path:        ":id/",
// 					Method:      "GET",
// 					Handler:     AddDecorators(handler.Get(nil), setupConfig.Decorators),
// 					Middlewares: setupConfig.Middlewares,
// 				},
// 				{
// 					Path:        "/all/",
// 					Method:      "GET",
// 					Handler:     AddDecorators(handler.List(nil), setupConfig.Decorators),
// 					Middlewares: setupConfig.Middlewares,
// 				},

// 				{
// 					Path:        ":id/",
// 					Method:      "PUT",
// 					Handler:     AddDecorators(handler.Update(nil), setupConfig.Decorators),
// 					Middlewares: setupConfig.Middlewares,
// 				},
// 				{
// 					Path:        ":id/",
// 					Method:      "DELETE",
// 					Handler:     AddDecorators(handler.DeleteByID(nil), setupConfig.Decorators),
// 					Middlewares: setupConfig.Middlewares,
// 				},
// 				{
// 					Path:        "/",
// 					Method:      "DELETE",
// 					Handler:     AddDecorators(handler.Delete(nil), setupConfig.Decorators),
// 					Middlewares: setupConfig.Middlewares,
// 				},
// 			},
// 		}
// 	}

// 	RegisterRoutes[M, ReqDTO, ResDTO](routeHandler, *setupConfig.routeConfig)
// }

// CREATE
type GenericConfig[M Model, ReqDTO, ResDTO any] struct {
	SetupConfig[M, ReqDTO, ResDTO]
	Handlers []Handler
}

func SetupGenericRouteForModel[M Model, ReqDTO, ResDTO any](
	setupConfig GenericConfig[M, ReqDTO, ResDTO],
) {
	routes := []RouteInfo{}
	routeHandler := &RouteHandler[M, ReqDTO, ResDTO]{}
	for _, handler := range setupConfig.Handlers {
		routeHandler = NewRouteHandler[M, ReqDTO, ResDTO](handler, setupConfig.Router)
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
