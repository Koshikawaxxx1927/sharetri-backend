package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

type Spot struct {
	gorm.Model
	TripID string
	Name string
	Date time.Time
	StartTime time.Time
	EndTime time.Time
	Cost int
	Memo string
}

func (spot *Spot) CreateSpot() (err error) {
	db := config.GetDB()
	return db.Create(spot).Error
}

func (spot *Spot) FindSpotByID(id string) (err error) {
	db := config.GetDB()
	return db.First(spot, id).Error
}

func (spot *Spot) UpdateSpotByID() (err error) {
	db := config.GetDB()
	return db.Save(&spot).Error
}

func (spot *Spot) DeleteSpotByID(id string) (err error) {
	db := config.GetDB()
	err = db.Delete(spot, id).Error
	return err
}