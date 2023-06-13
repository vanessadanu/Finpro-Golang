package models

// Struct untuk menyimpan data pengambilan
type Pickup struct {
	ID      int    `json:"id " gorm:"primary_key"`
	OrderID int    `json:"order_id"`
	Order   Order  `json:"order"`
	Date    string `json:"date"`
}

type PickupInput struct {
	OrderID int    `json:"order_id"`
	Date    string `json:"date"`
}
