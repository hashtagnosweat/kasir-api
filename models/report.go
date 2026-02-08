package models

type BestSellingProduct struct {
	Name    string `json:"name"`
	QtySold int    `json:"qty_sold"`
}

type TodayReport struct {
	TotalRevenue      int                `json:"total_revenue"`
	TotalTransactions int                `json:"total_transactions"`
	BestProduct       BestSellingProduct `json:"best_product"`
}
