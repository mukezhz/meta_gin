package meta_gin

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReadHandler[M Model, ReqDTO any, ResDTO any] struct {
	DB               *gorm.DB
	DTOHandler       DTOHandler[M, ReqDTO, ResDTO]
	Service          *Service[M]
	ServiceExecuters []ServiceExecutor[M]
}

func NewReadHandler[M Model, ReqDTO any, ResDTO any](
	db *gorm.DB,
	service *Service[M],
	dtoHandler DTOHandler[M, ReqDTO, ResDTO],
) *ReadHandler[M, ReqDTO, ResDTO] {
	return &ReadHandler[M, ReqDTO, ResDTO]{
		DB:         db,
		Service:    service,
		DTOHandler: dtoHandler,
	}
}

func (h *ReadHandler[M, ReqDTO, ResDTO]) AddServiceExecuter(
	serviceExecuter ServiceExecutor[M],
) {
	h.ServiceExecuters = append(h.ServiceExecuters, serviceExecuter)
}

func (h *ReadHandler[M, ReqDTO, ResDTO]) GetName() string {
	return "read_handler"
}

func (h *ReadHandler[M, ReqDTO, ResDTO]) Method() string {
	return http.MethodGet
}

func (h *ReadHandler[M, ReqDTO, ResDTO]) Handlers() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"/all": h.List(),
		"/":    h.ListPagination(),
		"/:id": h.Get(),
	}
}
func (h *ReadHandler[M, ReqDTO, ResDTO]) List() gin.HandlerFunc {
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

func (h *ReadHandler[M, ReqDTO, ResDTO]) ListPagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := 1
		limit := 10
		if l, err := strconv.ParseInt(ctx.Query("limit"), 10, 0); err == nil {
			limit = int(l)
		}

		if p, err := strconv.ParseInt(ctx.Query("page"), 10, 0); err == nil {
			page = int(p)
		}
		log.Println("page", page, "limit", limit, "Executors:::", len(h.ServiceExecuters))
		for _, queryExecutor := range h.ServiceExecuters {
			if queryExecutor != nil {
				c := context.WithValue(ctx.Request.Context(), pageLimit, PaginationRequest{
					Page:  page,
					Limit: limit,
				})
				queryExecutor.Execute(c, nil)
			}
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

func (h *ReadHandler[M, ReqDTO, ResDTO]) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		c := context.WithValue(ctx.Request.Context(), resID, id)
		for _, service := range h.ServiceExecuters {
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
