package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetUserConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	meta_gin.SetupModelRoutes[User, UserRequestDTO, UserResponseDTO](
		db,
		router,
		config,
		NewUserDTOHandler(),
		"v1",
		"users",
	)
}
