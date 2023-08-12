package models

type FinanceReport struct {
	ID         uint `gorm:"primaryKey"`
	ProductID  uint `json:"product_id"`
	SoldAmount int
	AmountSold float64
}
