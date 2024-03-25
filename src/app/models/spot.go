package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/config"
)

type Spot struct {
	gorm.Model
	TripID string `json:"tripid"  binding:"required"`
	Name string `json:"name" binding:"required"`
	StartTime time.Time `json:"starttime" binding:"required"`
	EndTime time.Time `json:"endtime" binding:"required"`
	Cost int `json:"cost" binding:"required"`
	Memo string `json:"memo" binding:"required"`
	ImagePath string `json:"imagepath"`
}

type Spots []Spot

func (spot *Spot) CreateSpot() (err error) {
	db := config.GetDB()
	return db.Create(spot).Error
}

func (spot *Spot) FindSpotByID(spotid string) (err error) {
	db := config.GetDB()
	return db.First(spot, spotid).Error
}

func (spot *Spot) UpdateSpotByID() (err error) {
	db := config.GetDB()
	return db.Save(&spot).Error
}

func (spot *Spot) DeleteSpotByID(spotid string) (err error) {
	db := config.GetDB()
	err = db.Delete(spot, spotid).Error
	return err
}

func (spots *Spots) GetSpotsByTripID(tripid string) (err error) {
	db := config.GetDB()
	return db.Where("trip_id = ?", tripid).Find(&spots).Error
}

func (spots *Spots) GetAllSpots() (err error) {
	db := config.GetDB()
	return db.Find(&spots).Error
}

func (spots *Spots) GetSpots(offset, limit int) (err error) {
	db := config.GetDB()
	return db.Offset(offset).Limit(limit).Find(&spots).Error
}

func (spots *Spots) GetSpotsListByTripID(tripid, offset, limit int) (err error) {
	db := config.GetDB()
	return db.Where("trip_id = ?", tripid).Offset(offset).Limit(limit).Find(&spots).Error
}