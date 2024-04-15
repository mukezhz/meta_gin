package meta_gin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateHandler[M Model, ReqDTO any, ResDTO any] struct {
	DB               *gorm.DB
	DTOHandler       DTOHandler[M, ReqDTO, ResDTO]
	Service          *Service[M]
	ServiceExecuters []ServiceExecutor[M]
}

func NewUpdateHandler[M Model, ReqDTO any, ResDTO any](
	db *gorm.DB,
	service *Service[M],
	dtoHandler DTOHandler[M, ReqDTO, ResDTO],
) *UpdateHandler[M, ReqDTO, ResDTO] {
	return &UpdateHandler[M, ReqDTO, ResDTO]{
		DB:         db,
		Service:    service,
		DTOHandler: dtoHandler,
	}
}

func (h *UpdateHandler[M, ReqDTO, ResDTO]) AddServiceExecuter(
	serviceExecuter ServiceExecutor[M],
) {
	h.ServiceExecuters = append(h.ServiceExecuters, serviceExecuter)
}

func (h *UpdateHandler[M, ReqDTO, ResDTO]) GetName() string {
	return "update_handler"
}

func (h *UpdateHandler[M, ReqDTO, ResDTO]) Method() string {
	return http.MethodPut
}

func (h *UpdateHandler[M, ReqDTO, ResDTO]) Handlers() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"/": h.Update(),
	}
}

func (h *UpdateHandler[M, ReqDTO, ResDTO]) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		model, err := h.Service.FindByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c := context.WithValue(ctx.Request.Context(), resID, id)
		for i, service := range h.ServiceExecuters {
			if service != nil {
				service.Execute(c, &model)
			}
			if i == 0 {
				break
			}
		}

		var dto ReqDTO
		if err := ctx.ShouldBindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		model = h.DTOHandler.ToModel(dto)
		c = context.WithValue(ctx.Request.Context(), dtoKey, dto)
		for i, service := range h.ServiceExecuters {
			if i == 0 {
				continue
			}
			if service != nil {
				service.Execute(c, &model)
			}
		}
		err = h.Service.Update(model)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resDTO := h.DTOHandler.FromModel(model)
		ctx.JSON(http.StatusOK, gin.H{"data": resDTO})
	}
}
