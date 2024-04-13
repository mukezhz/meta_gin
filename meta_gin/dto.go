package meta_gin

type DTOHandler[M any, ReqDTO any, ResDTO any] struct {
	ToModel   func(ReqDTO) M
	FromModel func(M) ResDTO
}
