package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetUserConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	meta_gin.SetupModelRoutes[User, UserRequestDTO, UserResponseDTO](
		meta_gin.SetupConfig[User, UserRequestDTO, UserResponseDTO]{
			DB:          db,
			Router:      router,
			Config:      config,
			DTOHandler:  NewUserDTOHandler(),
			Version:     "v1",
			GroupName:   "users",
			Middlewares: []gin.HandlerFunc{meta_gin.AuthMiddleware(config, "editor")},
		},
	)
}
