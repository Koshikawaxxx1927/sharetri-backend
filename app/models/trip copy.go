package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
	"github.com/Koshikawaxxx1927/sharetri-backend/database/types"
)

type Trip struct {
	gorm.Model
	UserID string `json:"userid"`
	PrefectureID string `json:"prefectureid" binding:"required"`
	Spots []Spot `gorm:"fossreignKey:TripID;constraint:OnUpdate:CASCADE"`
	Title string `json:"title" binding:"required"`
	StartDate types.Date `json:"startdate" binding:"required"`
	EndDate types.Date `json:"enddate" binding:"required"`
	Memo string `json:"memo" binding:"required"`
	ImagePath string `json:"imagepath"`
	IsPublic bool `json:"ispublic"`
}

type Trips []Trip

func (trip *Trip) CreateTrip() (err error) {
	db := config.GetDB()
	return db.Create(trip).Error
}

func (trip *Trip) FindTripByID(tripid string) (err error) {
	db := config.GetDB()
	return db.First(trip, tripid).Error
}

func (trip *Trip) UpdateTripByID() (err error) {
	db := config.GetDB()
	return db.Save(&trip).Error
}

func (trip *Trip) DeleteTripByID(tripid string) (err error) {
	db := config.GetDB()
	err = db.Delete(trip, tripid).Error
	return err
}

func (trips *Trips) GetAllTrips() (err error) {
	db := config.GetDB()
	return db.Find(&trips).Error
}

func (trips *Trips) GetTrips(offset, limit int) (err error) {
	db := config.GetDB()
	return db.Offset(offset).Limit(limit).Find(&trips).Error
}