package post

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetPostConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	repository := meta_gin.NewRepository[Post](db)
	service := meta_gin.NewService[Post](repository)

	meta_gin.SetupGenericRouteForModel(
		meta_gin.GenericConfig[Post, CreateProductRequestDTO, CreateProductResponseDTO]{
			SetupConfig: meta_gin.SetupConfig[Post, CreateProductRequestDTO, CreateProductResponseDTO]{
				Router:      router,
				Config:      config,
				Version:     "v1",
				GroupName:   "posts",
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
				Decorators:  []meta_gin.Decorator{meta_gin.NewPermissionDecorator()},
			},
			Handlers: []meta_gin.Handler{
				meta_gin.NewCreateHandler(db, service, NewCreateHandler()),
				meta_gin.NewReadHandler(db, service, NewReadHandler()),
				meta_gin.NewUpdateHandler(db, service, NewUpdateHandler()),
				meta_gin.NewDeleteHandler(db, service, NewDeleteHandler()),
			},
		},
	)
}
