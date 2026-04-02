package zakat

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Zakat struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	VillagerID  uuid.UUID `json:"villager_id"`
	Villager  	string 		`json:"villager"`
	Name        string    `json:"name"`
	TotalPeople int16     `json:"total_people"`
	Amount      float32   `json:"amount"`
	Charity     float32   `json:"charity"`
	Category    bool      `json:"category"`
}

type Villager struct {
	ID        uuid.UUID `json:"id"`
	Villager	string 		`json:"villager"`

}

type Repository interface {
	List(ctx context.Context, villagerID string, name string, category *bool, date *time.Time) ([]Zakat, error)
	GetVillager(ctx context.Context) ([]Villager, error)
	Create(ctx context.Context, z *Zakat) error
	Update(ctx context.Context, z *Zakat) error
	Delete(ctx context.Context, id string) error
}

type pgRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &pgRepository{db: db}
}

func (r *pgRepository) List(
	ctx context.Context,
	villagerID string,
	name string,
	category *bool,
	date *time.Time) ([]Zakat, error) {
	var args [4]any

	// $1 :: uuid
	if villagerID == "" {
			args[0] = nil
	} else {
			args[0] = villagerID
	}

	// $2 :: text
	if name == "" {
			args[1] = nil
	} else {
			args[1] = name
	}

	// $3 :: boolean
	if category == nil {
			args[2] = nil
	} else {
			args[2] = *category
	}

	// $4 :: date
	if date == nil {
			args[3] = nil
	} else {
			args[3] = date.Format("2006-01-02")
	}

	rows, err := r.db.Query(ctx, GetZakat,
			args[0], args[1], args[2], args[3],
	)
	if err != nil {
			return nil, err
	}
	defer rows.Close()

	var result []Zakat
	for rows.Next() {
			var z Zakat
			if err := rows.Scan(
					&z.ID,
					&z.CreatedAt,
					&z.UpdatedAt,
					&z.Villager,
					&z.Name,
					&z.TotalPeople,
					&z.Amount,
					&z.Charity,
					&z.Category,
			); err != nil {
					return nil, err
			}
			result = append(result, z)
	}
	return result, rows.Err()
}

func (r *pgRepository) GetVillager(ctx context.Context) ([]Villager, error) {
	rows, err := r.db.Query(ctx, GetVillager)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var villagers []Villager
	for rows.Next() {
		var v Villager
		if err := rows.Scan(&v.ID, &v.Villager); err != nil {
			return nil, err
		}
		villagers = append(villagers, v)
	}
	if len(villagers) == 0 {
		return nil, nil
	}
	return villagers, rows.Err()
}

func (r *pgRepository) Create(ctx context.Context, z *Zakat) error {
	err := r.db.QueryRow(ctx, CreateZakat,
		z.VillagerID,
		z.Name,
		z.TotalPeople,
		z.TotalPeople,           
		z.Charity,
		z.Category,
	).Scan(&z.ID, &z.CreatedAt, &z.UpdatedAt)
	return err
}

func (r *pgRepository) Update(ctx context.Context, z *Zakat) error {
	return r.db.QueryRow(ctx, UpdateZakat,
		z.ID,
		z.Name,
		z.TotalPeople,
		z.TotalPeople,
		z.Charity,
		z.Category,
	).Scan(&z.CreatedAt, &z.UpdatedAt)
}

func (r *pgRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, DeleteZakat, id)
	return err
}