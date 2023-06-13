package models

// Struct untuk menyimpan data pesanan
type Order struct {
	ID         int      `json:"id " gorm:"primary_key"`
	CustomerID int      `json:"customer_id"`
	Customer   Customer `json:"customer"`
	Item       string   `json:"item"`
	Quantity   int      `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
}

type OrderInput struct {
	CustomerID int     `json:"customer_id"`
	Item       string  `json:"item"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
