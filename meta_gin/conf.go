package meta_gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupModelRoutes[M Model, ReqDTO, ResDTO any](
	db *gorm.DB,
	router *gin.Engine,
	config *Config,
	dtoHandler *DTOHandler[M, ReqDTO, ResDTO],
	version string,
	groupName string,
) {
	repository := NewRepository[M](db)
	service := NewService[M](repository)
	userHandler := NewCRUDHandler[M, ReqDTO, ResDTO](db, service, dtoHandler)
	userRouter := NewRouteHandler[M, ReqDTO, ResDTO](userHandler, router)
	userConfig := RouteConfig{
		Version:   version,
		GroupName: groupName,
		Routes: []RouteInfo{
			{
				Path:        "/",
				Method:      "POST",
				Handler:     CheckPermissionDecorator(userHandler.Create()),
				Middlewares: []gin.HandlerFunc{AuthMiddleware(config, "editor")},
			},
			{
				Path:        "/",
				Method:      "GET",
				Handler:     CheckPermissionDecorator(userHandler.List()),
				Middlewares: []gin.HandlerFunc{AuthMiddleware(config, "editor")},
			},
			{
				Path:        ":id/",
				Method:      "GET",
				Handler:     CheckPermissionDecorator(userHandler.Get()),
				Middlewares: []gin.HandlerFunc{AuthMiddleware(config, "editor")},
			},
			{
				Path:        ":id/",
				Method:      "PUT",
				Handler:     CheckPermissionDecorator(userHandler.Update()),
				Middlewares: []gin.HandlerFunc{AuthMiddleware(config, "editor")},
			},
			{
				Path:        ":id/",
				Method:      "DELETE",
				Handler:     CheckPermissionDecorator(userHandler.DeleteByID()),
				Middlewares: []gin.HandlerFunc{AuthMiddleware(config, "editor")},
			},
			{
				Path:        "/",
				Method:      "DELETE",
				Handler:     CheckPermissionDecorator(userHandler.Delete()),
				Middlewares: []gin.HandlerFunc{AuthMiddleware(config, "editor")},
			},
		},
	}
	RegisterRoutes[M, ReqDTO, ResDTO](userRouter, userConfig)
}
