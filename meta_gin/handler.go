package meta_gin

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServiceExecutor[M any] interface {
	Execute(context context.Context, model *M) (M, error)
}

type QueryExecutor[M any] interface {
	Execute(context.Context) ([]M, error)
}

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

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Create(
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

func (h *CRUDHandler[M, ReqDTO, ResDTO]) List(
	queryExecutor QueryExecutor[PaginationRequest],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		models, err := h.Service.Find()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		dtos := []ResDTO{}
		for _, model := range models {
			dtos = append(dtos, h.DTOHandler.FromModel(model))
		}
		ctx.JSON(http.StatusOK, gin.H{"datas": dtos})
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) ListPagination(
	queryExecutor QueryExecutor[PaginationRequest],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := 1
		limit := 10
		if l, err := strconv.ParseInt(ctx.Query("limit"), 10, 0); err == nil {
			limit = int(l)
		}

		if p, err := strconv.ParseInt(ctx.Query("page"), 10, 0); err == nil {
			page = int(p)
		}

		if queryExecutor != nil {
			c := context.WithValue(ctx.Request.Context(), pageLimit, PaginationRequest{
				Page:  page,
				Limit: limit,
			})
			queryExecutor.Execute(c)
		}
		paginatedResponse, err := h.Service.FindWithPagination(page, limit)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		dtos := []ResDTO{}
		for _, model := range paginatedResponse.Items {
			dtos = append(dtos, h.DTOHandler.FromModel(model))
		}
		ctx.JSON(http.StatusOK, gin.H{"datas": dtos, "total": paginatedResponse.Total, "has_next": paginatedResponse.HasNext})
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Get(services ...ServiceExecutor[M]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		c := context.WithValue(ctx.Request.Context(), resID, id)
		for _, service := range services {
			if service != nil {
				service.Execute(c, nil)
			}
		}

		model, err := h.Service.FindByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		resDTO := h.DTOHandler.FromModel(model)
		ctx.JSON(http.StatusOK, gin.H{"data": resDTO})
	}
}

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Update(services ...ServiceExecutor[M]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		model, err := h.Service.FindByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c := context.WithValue(ctx.Request.Context(), resID, id)
		for i, service := range services {
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
		for i, service := range services {
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

func (h *CRUDHandler[M, ReqDTO, ResDTO]) DeleteByID(services ...ServiceExecutor[M]) gin.HandlerFunc {
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

func (h *CRUDHandler[M, ReqDTO, ResDTO]) Delete(services ...ServiceExecutor[M]) gin.HandlerFunc {
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
