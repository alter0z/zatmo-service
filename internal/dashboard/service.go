package dashboard

import "context"

type Service interface {
	GetSummaryCounts(ctx context.Context, villagerID string) (Summary, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetSummaryCounts(ctx context.Context, villagerID string) (Summary, error) {
	return s.repo.GetSummaryCounts(ctx, villagerID)
}
