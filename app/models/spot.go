package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

type Spot struct {
	gorm.Model
	TripID string `json:"tripid"  binding:"required"`
	Name string `json:"name" binding:"required"`
	Date time.Time `json:"date" binding:"required"`
	StartTime time.Time `json:"starttime" binding:"required"`
	EndTime time.Time `json:"endtime" binding:"required"`
	Cost int `json:"cost" binding:"required"`
	Memo string `json:"memo" binding:"required"`
	ImagePath string `json:"imagepath"`
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