package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CRUDHandler[M any, ReqDTO any, ResDTO any] struct {
	DB         *gorm.DB
	DTOHandler *DTOHandler[M, ReqDTO, ResDTO]
	Service    *Service[M]
}

func NewCRUDHandler[M any, ReqDTO any, ResDTO any](
	db *gorm.DB,
	service *Service[M],
	dtoHandler *DTOHandler[M, ReqDTO, ResDTO],
) *CRUDHandler[M, ReqDTO, ResDTO] {
	return &CRUDHandler[M, ReqDTO, ResDTO]{DB: db, Service: service, DTOHandler: dtoHandler}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto ReqDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		model := h.DTOHandler.ToModel(dto)
		h.Service.Create(model)
		resDTO := h.DTOHandler.FromModel(model)
		c.JSON(http.StatusCreated, resDTO)
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		models, count, err := h.Service.FindWithPagination(1, 10)
		log.Println("COUNT:", count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var dtos []ResDTO
		for _, model := range models {
			dtos = append(dtos, h.DTOHandler.FromModel(model))
		}
		c.JSON(http.StatusOK, dtos)
	}
}
