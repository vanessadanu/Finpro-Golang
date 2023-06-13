package models

// Struct untuk menyimpan data pelanggan
type Customer struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Email   string `json: "email`
	Address string `json:"address"`
}

type CustomerInput struct {
	Name    string `json:"name"`
	Email   string `json: "email`
	Address string `json:"address"`
}
