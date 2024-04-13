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
}

func SetupModelRoutes[M Model, ReqDTO, ResDTO any](
	setupConfig SetupConfig[M, ReqDTO, ResDTO],
) {
	if setupConfig.Config != nil && (setupConfig.Config.Roles != nil && len(setupConfig.Middlewares) == 0) {
		setupConfig.Middlewares = append(setupConfig.Middlewares, AuthMiddleware(setupConfig.Config, "editor"))
	}
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
				Handler:     CheckPermissionDecorator(userHandler.Create()),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        "/",
				Method:      "GET",
				Handler:     CheckPermissionDecorator(userHandler.List()),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        ":id/",
				Method:      "GET",
				Handler:     CheckPermissionDecorator(userHandler.Get()),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        ":id/",
				Method:      "PUT",
				Handler:     CheckPermissionDecorator(userHandler.Update()),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        ":id/",
				Method:      "DELETE",
				Handler:     CheckPermissionDecorator(userHandler.DeleteByID()),
				Middlewares: setupConfig.Middlewares,
			},
			{
				Path:        "/",
				Method:      "DELETE",
				Handler:     CheckPermissionDecorator(userHandler.Delete()),
				Middlewares: setupConfig.Middlewares,
			},
		},
	}
	RegisterRoutes[M, ReqDTO, ResDTO](userRouter, userConfig)
}
