package meta_gin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MetaGin struct {
	// Gin Router
	Router *gin.Engine
	// Gin Router
	DB *gorm.DB
	// Gin Config
	Config *Config
}

func NewMetaGin(
	engine *gin.Engine,
	database *gorm.DB,
	config *Config,
) *MetaGin {
	return &MetaGin{
		Router: engine,
		DB:     database,
		Config: config,
	}
}

func (m *MetaGin) Run(port string) {
	if port == "" {
		port = ":8888"
	}
	err := m.Router.Run(port)
	if err != nil {
		panic("failed to run server")
	}
}
