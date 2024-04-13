package meta_gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CRUDHandler[M Model, ReqDTO any, ResDTO any] struct {
	DB         *gorm.DB
	DTOHandler *DTOHandler[M, ReqDTO, ResDTO]
	Service    *Service[M]
}

func NewCRUDHandler[M Model, ReqDTO any, ResDTO any](
	db *gorm.DB,
	service *Service[M],
	dtoHandler *DTOHandler[M, ReqDTO, ResDTO],
) *CRUDHandler[M, ReqDTO, ResDTO] {
	return &CRUDHandler[M, ReqDTO, ResDTO]{
		DB:         db,
		Service:    service,
		DTOHandler: dtoHandler,
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto ReqDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		model := h.DTOHandler.ToModel(dto)
		m, err := h.Service.Create(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resDTO := h.DTOHandler.FromModel(m)
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
		dtos := []ResDTO{}
		for _, model := range models {
			dtos = append(dtos, h.DTOHandler.FromModel(model))
		}
		c.JSON(http.StatusOK, gin.H{"data": dtos})
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		model, err := h.Service.FindByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		dto := h.DTOHandler.FromModel(model)
		c.JSON(http.StatusOK, dto)
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		model, err := h.Service.FindByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		var dto ReqDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		model = h.DTOHandler.ToModel(dto)
		err = h.Service.Update(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resDTO := h.DTOHandler.FromModel(model)
		c.JSON(http.StatusOK, resDTO)
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) DeleteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := h.Service.DeleteByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto ReqDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		model := h.DTOHandler.ToModel(dto)
		err := h.Service.Delete(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
