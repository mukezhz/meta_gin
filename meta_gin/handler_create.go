package meta_gin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateHandler[M Model, ReqDTO any, ResDTO any] struct {
	DB         *gorm.DB
	DTOHandler DTOHandler[M, ReqDTO, ResDTO]
	Service    *Service[M]
}

func NewCreateHandler[M Model, ReqDTO any, ResDTO any](
	db *gorm.DB,
	service *Service[M],
	dtoHandler DTOHandler[M, ReqDTO, ResDTO],
) *CreateHandler[M, ReqDTO, ResDTO] {
	return &CreateHandler[M, ReqDTO, ResDTO]{
		DB:         db,
		Service:    service,
		DTOHandler: dtoHandler,
	}
}

func (h *CreateHandler[M, ReqDTO, ResDTO]) GetName() string {
	return "create_handler"
}

func (h *CreateHandler[M, ReqDTO, ResDTO]) Method() string {
	return http.MethodPost
}

func (h *CreateHandler[M, ReqDTO, ResDTO]) Handlers() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"/": h.Create(nil),
	}
}

func (h *CreateHandler[M, ReqDTO, ResDTO]) Create(
	services ...ServiceExecutor[M],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto ReqDTO
		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		model := h.DTOHandler.ToModel(dto)
		c := context.WithValue(ctx.Request.Context(), dtoKey, dto)
		for _, service := range services {
			if service != nil {
				m, err := service.Execute(c, &model)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				model = m
			}
		}
		m, err := h.Service.Create(model)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resDTO := h.DTOHandler.FromModel(m)
		ctx.JSON(http.StatusCreated, gin.H{"data": resDTO})
	}
}
