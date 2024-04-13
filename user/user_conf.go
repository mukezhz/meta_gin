package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetUserConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	repository := meta_gin.NewRepository[User](db)
	service := meta_gin.NewService[User](repository)
	userDTOHandler := NewUserDTOHandler()
	userHandler := meta_gin.NewCRUDHandler[User, UserRequestDTO, UserResponseDTO](db, service, userDTOHandler)
	userRouter := meta_gin.NewRouteHandler[User, UserRequestDTO, UserResponseDTO](userHandler, router)
	userConfig := meta_gin.RouteConfig{
		Version:   "v1",
		GroupName: "users",
		Routes: []meta_gin.RouteInfo{
			{
				Path:        "/",
				Method:      "POST",
				Handler:     meta_gin.CheckPermissionDecorator(userHandler.Create()),
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			},
			{
				Path:        "/",
				Method:      "GET",
				Handler:     meta_gin.CheckPermissionDecorator(userHandler.List()),
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			},
			{
				Path:        ":id/",
				Method:      "GET",
				Handler:     meta_gin.CheckPermissionDecorator(userHandler.Get()),
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			},
			{
				Path:        ":id/",
				Method:      "PUT",
				Handler:     meta_gin.CheckPermissionDecorator(userHandler.Update()),
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			},
			{
				Path:        ":id/",
				Method:      "DELETE",
				Handler:     meta_gin.CheckPermissionDecorator(userHandler.DeleteByID()),
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			},
			{
				Path:        "/",
				Method:      "DELETE",
				Handler:     meta_gin.CheckPermissionDecorator(userHandler.Delete()),
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
			},
		},
	}
	meta_gin.RegisterRoutes[User, UserRequestDTO, UserResponseDTO](userRouter, userConfig)
}
