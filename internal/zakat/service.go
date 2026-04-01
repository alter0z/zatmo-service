package zakat

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	List(ctx context.Context, villagerID string, name string, category *bool, date *time.Time) ([]Zakat, error)
	GetVillager(ctx context.Context) ([]Villager, error)
	Create(ctx context.Context, input CreateZakatInput) (Zakat, error)
	Update(ctx context.Context, input UpdateZakatInput) (Zakat, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

type CreateZakatInput struct {
	VillagerID 	uuid.UUID  `json:"villager_id" binding:"required,uuid4"`
	Name       	string  `json:"name" binding:"required"`
	TotalPeople int16  	`json:"total_people" binding:"gte=1"`
	Charity    	float32 `json:"charity" binding:"gte=0"`
	Category   	bool    `json:"category"`
}

type UpdateZakatInput struct {
	ID          uuid.UUID  `json:"-"`
	VillagerID 	uuid.UUID  `json:"villager_id" binding:"required,uuid4"`
	Name       	string  `json:"name" binding:"required"`
	TotalPeople int16  	`json:"total_people" binding:"gte=1"`
	Charity    	float32 `json:"charity" binding:"gte=0"`
	Category   	bool    `json:"category"`
}

func (s *service) List(ctx context.Context, villagerID string, name string, category *bool, date *time.Time) ([]Zakat, error) {
	return s.repo.List(ctx, villagerID, name, category, date)
}

func (s *service) GetVillager(ctx context.Context) ([]Villager, error) {
	return s.repo.GetVillager(ctx)
}

func (s *service) Create(ctx context.Context, input CreateZakatInput) (Zakat, error) {
    z := Zakat{
		VillagerID: input.VillagerID,
		Name:       input.Name,
		TotalPeople:input.TotalPeople,
		Charity:  	input.Charity,
		Category: 	input.Category,
	}

	if err := s.repo.Create(ctx, &z); err != nil {
		return Zakat{}, err
	}

	return z, nil
}

func (s *service) Update(ctx context.Context, input UpdateZakatInput) (Zakat, error) {
	z := Zakat{
		ID:          input.ID,
		VillagerID:  input.VillagerID,
		Name:        input.Name,
		TotalPeople: input.TotalPeople,
		Charity:     input.Charity,
		Category:    input.Category,
	}

	if err := s.repo.Update(ctx, &z); err != nil {
		return Zakat{}, err
	}
	return z, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
