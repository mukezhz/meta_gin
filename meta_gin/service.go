package meta_gin

type Service[M any] struct {
	Repository *Repository[M]
}

func NewService[M any](repository *Repository[M]) *Service[M] {
	return &Service[M]{Repository: repository}
}

func (s *Service[M]) Create(model M) (M, error) {
	return s.Repository.Create(model)
}

func (s *Service[M]) FindOrCreate(model M) error {
	return s.Repository.FindOrCreate(model)
}

func (s *Service[M]) FindByID(id string) (M, error) {
	return s.Repository.FindByID(id)
}

func (s *Service[M]) Update(model M) error {
	return s.Repository.Update(model)
}

func (s *Service[M]) Delete(model M) error {
	return s.Repository.Delete(model)
}

func (s *Service[M]) DeleteByID(id string) error {
	return s.Repository.DeleteByID(id)
}

func (s *Service[M]) FindWithPagination(page, pageSize int) ([]M, int64, error) {
	var count int64 = 0
	res, err := s.Repository.Find(WithPagination(page, pageSize), WithCount(&count))
	return res, count, err
}

func (s *Service[M]) FindWithCondition(condition string, args ...interface{}) ([]M, error) {
	return s.Repository.Find(WithCondition(condition, args...))
}

func (s *Service[M]) FindWithOrder(order string) ([]M, error) {
	return s.Repository.Find(WithOrder(order))
}

func (s *Service[M]) FindWithSort(field string, order SortOrder) ([]M, error) {
	return s.Repository.Find(WithSort(field, order))
}

func (s *Service[M]) FindWithCount(count *int64) ([]M, error) {
	return s.Repository.Find(WithCount(count))
}

func (s *Service[M]) FindWithPaginationAndCondition(page, pageSize int, condition string, args ...interface{}) ([]M, int64, error) {
	var count int64 = 0
	res, err := s.Repository.Find(WithPagination(page, pageSize), WithCondition(condition, args...), WithCount(&count))
	return res, count, err
}

func (s *Service[M]) FindWithPaginationAndOrder(page, pageSize int, order string) ([]M, int64, error) {
	var count int64 = 0
	res, err := s.Repository.Find(WithPagination(page, pageSize), WithOrder(order), WithCount(&count))
	return res, count, err
}

func (s *Service[M]) FindWithPaginationAndSort(page, pageSize int, field string, order SortOrder) ([]M, int64, error) {
	var count int64 = 0
	res, err := s.Repository.Find(WithPagination(page, pageSize), WithSort(field, order), WithCount(&count))
	return res, count, err
}
