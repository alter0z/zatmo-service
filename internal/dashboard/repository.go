package dashboard

import (
	"context"
	"time"

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
	GetSummaryCounts(ctx context.Context, villagerID string, date *time.Time) (Summary, error)
	GetReceiverCounts(ctx context.Context) (Receiver, error)
}

type pgRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &pgRepository{db: db}
}

func (r *pgRepository) GetSummaryCounts(ctx context.Context, villagerID string, date *time.Time) (Summary, error) {
	var s Summary
	var args [2]any

	// $1 :: uuid
	if villagerID == "" {
			args[0] = nil
	} else {
			args[0] = villagerID
	}

	// $2 :: date
	if date == nil {
			args[1] = nil
	} else {
			args[1] = date.Format("2006-01-02")
	}
	err := r.db.QueryRow(ctx, GetSummaryCounts, args[0], args[1]).Scan(
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
