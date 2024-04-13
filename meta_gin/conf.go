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
}

func SetupModelRoutes[M Model, ReqDTO, ResDTO any](
	setupConfig SetupConfig[M, ReqDTO, ResDTO],
) {
	repository := NewRepository[M](setupConfig.DB)
	service := NewService[M](repository)
	userHandler := NewCRUDHandler[M, ReqDTO, ResDTO](setupConfig.DB, service, setupConfig.DTOHandler)
	userRouter := NewRouteHandler[M, ReqDTO, ResDTO](userHandler, setupConfig.Router)
	userConfig := RouteConfig{
		Version:   setupConfig.Version,
		GroupName: setupConfig.GroupName,
		Routes: []RouteInfo{
			{
				Path:        "/",
				Method:      "POST",
				Handler:     addDecorators(userHandler.Create(), setupConfig.Decorators),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        "/",
				Method:      "GET",
				Handler:     addDecorators(userHandler.List(), setupConfig.Decorators),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        ":id/",
				Method:      "GET",
				Handler:     addDecorators(userHandler.Get(), setupConfig.Decorators),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        ":id/",
				Method:      "PUT",
				Handler:     addDecorators(userHandler.Update(), setupConfig.Decorators),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        ":id/",
				Method:      "DELETE",
				Handler:     addDecorators(userHandler.DeleteByID(), setupConfig.Decorators),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        "/",
				Method:      "DELETE",
				Handler:     addDecorators(userHandler.Delete(), setupConfig.Decorators),
				Middlewares: setupConfig.Middlewares,
			},
		},
	}
	RegisterRoutes[M, ReqDTO, ResDTO](userRouter, userConfig)
}

func addDecorators(handler gin.HandlerFunc, decorators []Decorator) gin.HandlerFunc {
	for _, decorator := range decorators {
		handler = decorator.Decorate(handler)
	}
	return handler
}
