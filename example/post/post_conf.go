package post

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetPostConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	repository := meta_gin.NewRepository[Post](db)
	service := meta_gin.NewService[Post](repository)

	createHandler := meta_gin.NewCreateHandler[Post, CreateProductRequestDTO, CreateProductResponseDTO](db, service, NewCreateHandler())
	readHandler := meta_gin.NewReadHandler[Post, ReadProductRequestDTO, ReadProductResponseDTO](db, service, NewReadHandler())
	createHandler.AddServiceExecuter(NewPostAddExecutor())
	readHandler.AddServiceExecuter(NewPostAddExecutor())
	meta_gin.SetupGenericRouteForModel(
		meta_gin.GenericConfig[Post]{
			SetupConfig: meta_gin.SetupConfig[Post]{
				Router:      router,
				Config:      config,
				Version:     "v1",
				GroupName:   "posts",
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
				Decorators:  []meta_gin.Decorator{meta_gin.NewPermissionDecorator()},
			},
			Handlers: []meta_gin.Handler[Post]{
				createHandler,
				readHandler,
				meta_gin.NewUpdateHandler[Post, UpdateProductRequestDTO, UpdateProductResponseDTO](db, service, NewUpdateHandler()),
				meta_gin.NewDeleteHandler[Post, DeleteProductRequestDTO, DeleteProductResponseDTO](db, service, NewDeleteHandler()),
			},
		},
	)
}
