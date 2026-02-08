package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetTodayReport() (*models.TodayReport, error) {
	report := &models.TodayReport{}

	// total revenue + total transactions today
	err := r.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&report.TotalRevenue, &report.TotalTransactions)

	if err != nil {
		return nil, err
	}

	// best selling product today
	err = r.db.QueryRow(`
		SELECT p.name, COALESCE(SUM(td.quantity), 0) AS qty
		FROM transaction_details td
		JOIN products p ON p.id = td.product_id
		JOIN transactions t ON t.id = td.transaction_id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`).Scan(&report.BestProduct.Name, &report.BestProduct.QtySold)

	if err == sql.ErrNoRows {
		report.BestProduct = models.BestSellingProduct{}
		return report, nil
	}

	if err != nil {
		return nil, err
	}

	return report, nil
}

func (r *ReportRepository) GetReportByDateRange(startDate, endDate string) (*models.TodayReport, error) {
	report := &models.TodayReport{}

	err := r.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE DATE(created_at) BETWEEN $1 AND $2
	`, startDate, endDate).Scan(
		&report.TotalRevenue,
		&report.TotalTransactions,
	)

	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`
		SELECT p.name, COALESCE(SUM(td.quantity), 0) AS qty
		FROM transaction_details td
		JOIN products p ON p.id = td.product_id
		JOIN transactions t ON t.id = td.transaction_id
		WHERE DATE(t.created_at) BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`, startDate, endDate).Scan(
		&report.BestProduct.Name,
		&report.BestProduct.QtySold,
	)

	if err == sql.ErrNoRows {
		return report, nil
	}
	if err != nil {
		return nil, err
	}

	return report, nil
}
