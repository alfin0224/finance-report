package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Stock		uint
	Price      float64
	CategoryID uint
}

type Category struct {
	gorm.Model
	Name string
}
