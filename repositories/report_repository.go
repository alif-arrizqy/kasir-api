package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository interface {
	GetReport(start, end *time.Time) (*models.ReportData, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (repo *reportRepository) GetReport(start, end *time.Time) (*models.ReportData, error) {
	// defaults: today in WIB
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	var s, e time.Time
	if start == nil || end == nil {
		s = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).UTC()
		e = s.Add(24 * time.Hour).UTC()
	} else {
		s = start.In(loc).UTC()
		e = end.In(loc).UTC()
	}

	var totalRevenue sql.NullInt64
	var totalTrans sql.NullInt64
	// total revenue and transactions
	err := repo.db.QueryRow("SELECT COALESCE(SUM(total_amount),0), COUNT(*) FROM transactions WHERE created_at >= $1 AND created_at < $2", s, e).Scan(&totalRevenue, &totalTrans)
	if err != nil {
		return nil, err
	}

	// top product
	var topName sql.NullString
	var topQty sql.NullInt64
	err = repo.db.QueryRow(`SELECT p.name, COALESCE(SUM(td.quantity),0) as qty FROM transaction_details td JOIN products p ON p.id = td.product_id JOIN transactions t ON t.id = td.transaction_id WHERE t.created_at >= $1 AND t.created_at < $2 GROUP BY p.name ORDER BY qty DESC LIMIT 1`, s, e).Scan(&topName, &topQty)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	name := ""
	qty := 0
	if topName.Valid {
		name = topName.String
	}
	if topQty.Valid {
		qty = int(topQty.Int64)
	}

	return &models.ReportData{
		TotalRevenue:   int(totalRevenue.Int64),
		TotalTransaksi: int(totalTrans.Int64),
		ProdukTerlaris: models.TopProduct{
			Nama:       name,
			QtyTerjual: qty,
		},
	}, nil
}
