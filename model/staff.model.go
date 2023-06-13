package models

// Struct untuk menyimpan data staff
type Staff struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type StaffInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
