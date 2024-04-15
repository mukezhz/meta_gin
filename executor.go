package meta_gin

import (
	"context"
)

type ServiceExecutor[M any] interface {
	Execute(context context.Context, model *M) (*M, error)
}

type QueryExecutor[M any] interface {
	Execute(context.Context) ([]M, error)
}
