package dashboard

import (
	"context"
	"time"
)

type Service interface {
	GetSummaryCounts(ctx context.Context, villagerID string, date *time.Time) (Summary, error)
	GetReceiverCounts(ctx context.Context) (Receiver, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetSummaryCounts(ctx context.Context, villagerID string, date *time.Time) (Summary, error) {
	return s.repo.GetSummaryCounts(ctx, villagerID, date)
}

func (s *service) GetReceiverCounts(ctx context.Context) (Receiver, error) {
	return s.repo.GetReceiverCounts(ctx)
}