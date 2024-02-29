package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

type Trip struct {
	gorm.Model
	UserID string `json:"userid"`
	PrefectureID string `json:"prefectureid" binding:"required"`
	Spots []Spot `gorm:"fossreignKey:TripID;constraint:OnUpdate:CASCADE"`
	Title string `json:"title" binding:"required"`
	StartDate time.Time `json:"startdate" binding:"required"`
	EndDate time.Time `json:"enddate" binding:"required"`
	Memo string `json:"memo" binding:"required"`
	ImagePath string `json:"imagepath"`
	IsPublic bool `json:"ispublic"`
}

type Trips []Trip

func (trip *Trip) CreateTrip() (err error) {
	db := config.GetDB()
	return db.Create(trip).Error
}

func (trip *Trip) FindTripByID(id string) (err error) {
	db := config.GetDB()
	return db.First(trip, id).Error
}

func (trip *Trip) UpdateTripByID() (err error) {
	db := config.GetDB()
	return db.Save(&trip).Error
}

func (trip *Trip) DeleteTripByID(id string) (err error) {
	db := config.GetDB()
	err = db.Delete(trip, id).Error
	return err
}

func (trips *Trips) GetAllTrips() (err error) {
	db := config.GetDB()
	return db.Find(&trips).Error
}