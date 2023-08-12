package models

type Product struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Price    float64
	Stock    int
	Category Category
}
