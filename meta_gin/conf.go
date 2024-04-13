package meta_gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SetupConfig[M Model, ReqDTO, ResDTO any] struct {
	DB          *gorm.DB
	Router      *gin.Engine
	Config      *Config
	DTOHandler  *DTOHandler[M, ReqDTO, ResDTO]
	Version     string
	GroupName   string
	Middlewares []gin.HandlerFunc
	Decorators  []Decorator
	routeConfig *RouteConfig
}

func SetupModelRoutes[M Model, ReqDTO, ResDTO any](
	setupConfig SetupConfig[M, ReqDTO, ResDTO],
) {
	repository := NewRepository[M](setupConfig.DB)
	service := NewService[M](repository)
	crudHandler := NewCRUDHandler[M, ReqDTO, ResDTO](setupConfig.DB, service, setupConfig.DTOHandler)
	routeHandler := NewRouteHandler[M, ReqDTO, ResDTO](crudHandler, setupConfig.Router)
	if setupConfig.routeConfig == nil {
		setupConfig.routeConfig = &RouteConfig{
			Version:   setupConfig.Version,
			GroupName: setupConfig.GroupName,
			Routes: []RouteInfo{
				{
					Path:        "/",
					Method:      "POST",
					Handler:     AddDecorators(crudHandler.Create(nil), setupConfig.Decorators),
					Middlewares: setupConfig.Middlewares,
				},
				{
					Path:        "/",
					Method:      "GET",
					Handler:     AddDecorators(crudHandler.ListPagination(nil), setupConfig.Decorators),
					Middlewares: setupConfig.Middlewares,
				},
				{
					Path:        ":id/",
					Method:      "GET",
					Handler:     AddDecorators(crudHandler.Get(nil), setupConfig.Decorators),
					Middlewares: setupConfig.Middlewares,
				},
				{
					Path:        "/all/",
					Method:      "GET",
					Handler:     AddDecorators(crudHandler.List(nil), setupConfig.Decorators),
					Middlewares: setupConfig.Middlewares,
				},
				{
					Path:        ":id/",
					Method:      "PUT",
					Handler:     AddDecorators(crudHandler.Update(nil), setupConfig.Decorators),
					Middlewares: setupConfig.Middlewares,
				},
				{
					Path:        ":id/",
					Method:      "DELETE",
					Handler:     AddDecorators(crudHandler.DeleteByID(nil), setupConfig.Decorators),
					Middlewares: setupConfig.Middlewares,
				},
				{
					Path:        "/",
					Method:      "DELETE",
					Handler:     AddDecorators(crudHandler.Delete(nil), setupConfig.Decorators),
					Middlewares: setupConfig.Middlewares,
				},
			},
		}
	}

	RegisterRoutes[M, ReqDTO, ResDTO](routeHandler, *setupConfig.routeConfig)
}

func AddDecorators(handler gin.HandlerFunc, decorators []Decorator) gin.HandlerFunc {
	for _, decorator := range decorators {
		handler = decorator.Decorate(handler)
	}
	return handler
}
