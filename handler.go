package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CRUDHandler[M any, ReqDTO any, ResDTO any] struct {
	DB         *gorm.DB
	DTOHandler *DTOHandler[M, ReqDTO, ResDTO]
}

func NewCRUDHandler[M any, ReqDTO any, ResDTO any](db *gorm.DB, dtoHandler *DTOHandler[M, ReqDTO, ResDTO]) *CRUDHandler[M, ReqDTO, ResDTO] {
	return &CRUDHandler[M, ReqDTO, ResDTO]{DB: db, DTOHandler: dtoHandler}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto ReqDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		model := h.DTOHandler.ToModel(dto)
		h.DB.Create(&model)
		resDTO := h.DTOHandler.FromModel(model)
		c.JSON(http.StatusCreated, resDTO)
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var models []M
		h.DB.Find(&models)
		var dtos []ResDTO
		for _, model := range models {
			dtos = append(dtos, h.DTOHandler.FromModel(model))
		}
		c.JSON(http.StatusOK, dtos)
		return
	}
}
