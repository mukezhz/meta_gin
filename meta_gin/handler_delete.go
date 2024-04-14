package meta_gin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteHandler[M Model, ReqDTO any, ResDTO any] struct {
	DB         *gorm.DB
	DTOHandler DTOHandler[M, ReqDTO, ResDTO]
	Service    *Service[M]
}

func NewDeleteHandler[M Model, ReqDTO any, ResDTO any](
	db *gorm.DB,
	service *Service[M],
	dtoHandler DTOHandler[M, ReqDTO, ResDTO],
) *DeleteHandler[M, ReqDTO, ResDTO] {
	return &DeleteHandler[M, ReqDTO, ResDTO]{
		DB:         db,
		Service:    service,
		DTOHandler: dtoHandler,
	}
}

func (h *DeleteHandler[M, ReqDTO, ResDTO]) GetName() string {
	return "delete_handler"
}

func (h *DeleteHandler[M, ReqDTO, ResDTO]) Method() string {
	return http.MethodDelete
}

func (h *DeleteHandler[M, ReqDTO, ResDTO]) Handlers() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"/":    h.Delete(nil),
		"/:id": h.DeleteByID(nil),
	}
}

func (h *DeleteHandler[M, ReqDTO, ResDTO]) DeleteByID(services ...ServiceExecutor[M]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		c := context.WithValue(ctx.Request.Context(), resID, id)
		for _, service := range services {
			if service != nil {
				service.Execute(c, nil)
			}
		}
		err := h.Service.DeleteByID(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}

func (h *DeleteHandler[M, ReqDTO, ResDTO]) Delete(services ...ServiceExecutor[M]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto ReqDTO
		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		model := h.DTOHandler.ToModel(dto)
		c := context.WithValue(ctx.Request.Context(), dtoKey, dto)
		c = context.WithValue(c, modelKey, model)
		for _, service := range services {
			if service != nil {
				service.Execute(c, &model)
			}
		}

		err := h.Service.Delete(model)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
