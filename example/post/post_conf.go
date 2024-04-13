package post

import (
	"github.com/gin-gonic/gin"
	"github.com/mukezhz/meta_gin/meta_gin"
	"gorm.io/gorm"
)

func GetPostConfig(db *gorm.DB, config *meta_gin.Config, router *gin.Engine) {
	meta_gin.SetupModelRoutes[Post, PostRequestDTO, PostResponseDTO](
		db,
		router,
		config,
		NewPostDTOHandler(),
		"v1",
		"posts",
	)
}
