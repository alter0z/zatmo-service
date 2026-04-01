package zakat

import (
	"context"
	"time"
)

type Service interface {
	List(ctx context.Context, villagerID string, name string, category *bool, date *time.Time) ([]Zakat, error)
	// Create(ctx context.Context, input CreateZakatInput) (Zakat, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

// type CreateZakatInput struct {
// 	VillagerID string  `json:"villager_id" binding:"required,uuid4"`
// 	Name       string  `json:"name" binding:"required"`
// 	TotalPeople int16  `json:"total_people" binding:"gte=1"`
// 	Amount     float32 `json:"amount" binding:"gte=0"`
// 	Charity    float32 `json:"charity" binding:"gte=0"`
// 	Category   bool    `json:"category"`
// }

func (s *service) List(ctx context.Context, villagerID string, name string, category *bool, date *time.Time) ([]Zakat, error) {
	return s.repo.List(ctx, villagerID, name, category, date)
}

// func (s *service) Create(ctx context.Context, input CreateZakatInput) (Zakat, error) {
//     // example implementation
//     z := Zakat{
//         // fill from input…
//     }
//     if err := s.repo.Create(ctx, &z); err != nil {
//         return Zakat{}, err
//     }
//     return z, nil
// }
