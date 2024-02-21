package models

import (
	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"

)

type Product struct {
	gorm.Model
	pid string `json:"pid" gorm:"unique;not null"`
	Code string
	Price uint
}

func (product *Product) Create() (err error) {
	db := config.GetDB()
	return db.Create(product).Error
}