package meta_gin

type DTOHandler[M any, ReqDTO any, ResDTO any] interface {
	ToModel(ReqDTO) M
	FromModel(M) ResDTO
}
