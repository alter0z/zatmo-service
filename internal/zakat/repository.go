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
	// VillagerID  uuid.UUID `json:"villager_id"`
	Villager  	string 		`json:"villager"`
	Name        string    `json:"name"`
	TotalPeople int16     `json:"total_people"`
	Amount      float32   `json:"amount"`
	Charity     float32   `json:"charity"`
	Category    bool      `json:"category"`
}

type Repository interface {
	List(ctx context.Context, villagerID string, name string, category *bool, date *time.Time) ([]Zakat, error)
	// Create(ctx context.Context, z *Zakat) error
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
	var args [4]interface{}

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

// func (r *pgRepository) Create(ctx context.Context, z *Zakat) error {
// 	err := r.db.QueryRow(ctx, `
// 		INSERT INTO zakat (villager, name, total_people, amount, charity, category)
// 		VALUES ($1, $2, $3, $4, $5, $6)
// 		RETURNING id, created_at, updated_at`,
// 		z.Villager, z.Name, z.TotalPeople, z.Amount, z.Charity, z.Category,
// 	).Scan(&z.ID, &z.CreatedAt, &z.UpdatedAt)
// 	return err
// }