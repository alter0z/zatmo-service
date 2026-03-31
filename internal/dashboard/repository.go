package dashboard

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Summary struct {
	TotalZakat  			int 	`json:"total_zakat"`
	TotalRiceZakat 		float32 `json:"total_rice_zakat"`
	TotalCashZakat 		int   `json:"total_cash_zakat"`
	TotalRiceCharity 	float32 `json:"total_rice_charity"`
	TotalCashCharity	int 	`json:"total_cash_charity"`
	TotalRiceAmount 	float32	`json:"total_rice_amount"`
	TotalCashAmount 	int  	`json:"total_cash_amount"`
}

type Repository interface {
	Data(ctx context.Context) (Summary, error)
}

type pgRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &pgRepository{db: db}
}

func (r *pgRepository) Data(ctx context.Context) (Summary, error) {
	rows, err := r.db.Query(ctx, GetSummaryCounts)
	if err != nil {
		return Summary{}, err
	}
	defer rows.Close()

	var result Summary
	for rows.Next() {
		var s Summary
		if err := rows.Scan(
			&s.TotalZakat,
			&s.TotalRiceZakat,
			&s.TotalCashZakat,
			&s.TotalRiceCharity,
			&s.TotalCashCharity,
			&s.TotalRiceAmount,
			&s.TotalCashAmount,
		); err != nil {
			return Summary{}, err
		}
		result = s
	}
	return result, rows.Err()
}