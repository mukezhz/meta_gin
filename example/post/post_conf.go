package post

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
)

func SetupPost(m *meta_gin.MetaGin) {
	repository := meta_gin.NewRepository[Post](m.DB)
	service := meta_gin.NewService[Post](repository)

	createHandler := meta_gin.NewCreateHandler[Post, CreateProductRequestDTO, CreateProductResponseDTO](m.DB, service, NewCreateHandler())
	readHandler := meta_gin.NewReadHandler[Post, ReadProductRequestDTO, ReadProductResponseDTO](m.DB, service, NewReadHandler())
	createHandler.AddServiceExecuter(NewPostAddExecutor())
	readHandler.AddServiceExecuter(NewPostAddExecutor())
	meta_gin.SetupGenericRouteForModel(
		meta_gin.GenericConfig[Post]{
			SetupConfig: meta_gin.SetupConfig[Post]{
				Router:      m.Router,
				Config:      m.Config,
				Version:     "v1",
				GroupName:   "posts",
				Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(m.Config, "editor")},
				Decorators:  []meta_gin.Decorator{meta_gin.NewPermissionDecorator()},
			},
			Handlers: []meta_gin.Handler[Post]{
				createHandler,
				readHandler,
				meta_gin.NewUpdateHandler[Post, UpdateProductRequestDTO, UpdateProductResponseDTO](m.DB, service, NewUpdateHandler()),
				meta_gin.NewDeleteHandler[Post, DeleteProductRequestDTO, DeleteProductResponseDTO](m.DB, service, NewDeleteHandler()),
			},
		},
	)
}
