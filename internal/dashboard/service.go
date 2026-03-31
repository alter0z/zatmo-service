package dashboard

import "context"

type Service interface {
	List(ctx context.Context) (Summary, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) List(ctx context.Context) (Summary, error) {
	return s.repo.Data(ctx)
}
