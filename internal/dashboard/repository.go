package dashboard

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Summary struct {
	TotalZakat  			int 		`json:"total_zakat"`
	TotalRiceZakat 		float32	`json:"total_rice_zakat"`
	TotalCashZakat 		int   	`json:"total_cash_zakat"`
	TotalRiceCharity 	float32 `json:"total_rice_charity"`
	TotalCashCharity	int 		`json:"total_cash_charity"`
	TotalRiceAmount 	float32	`json:"total_rice_amount"`
	TotalCashAmount 	int  		`json:"total_cash_amount"`
}

type Receiver struct {
	TotalNeedy     int `json:"total_needy"`
	TotalDestitute int `json:"total_destitute"`
}

type Repository interface {
	GetSummaryCounts(ctx context.Context, villagerID string) (Summary, error)
	GetReceiverCounts(ctx context.Context) (Receiver, error)
}

type pgRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &pgRepository{db: db}
}

func (r *pgRepository) GetSummaryCounts(ctx context.Context, villagerID string) (Summary, error) {
	var s Summary
	var arg any

	if villagerID == "" {
		arg = nil
	} else {
		arg = villagerID
	}

	err := r.db.QueryRow(ctx, GetSummaryCounts, arg).Scan(
		&s.TotalZakat,
		&s.TotalRiceZakat,
		&s.TotalCashZakat,
		&s.TotalRiceCharity,
		&s.TotalCashCharity,
		&s.TotalRiceAmount,
		&s.TotalCashAmount,
	)
	return s, err
}

func (r *pgRepository) GetReceiverCounts(ctx context.Context) (Receiver, error) {
	var receiver Receiver

	err := r.db.QueryRow(ctx, GetReceiverCounts).Scan(
		&receiver.TotalNeedy,
		&receiver.TotalDestitute,
	)
	return receiver, err
}
