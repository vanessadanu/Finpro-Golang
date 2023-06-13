package models

// Struct untuk menyimpan data tagihan
type Bill struct {
	Id       int     `json:"id" gorm:"primary_key"`
	Order_id int     `json:"order_id"`
	Order    Order   `json:"order"`
	Amount   float64 `json:"amount"`
	Paid     bool    `json:"paid "`
}

type BillInput struct {
	Order_id int     `json:"order_id"`
	Amount   float64 `json:"amount"`
	Paid     bool    `json:"paid "`
}
