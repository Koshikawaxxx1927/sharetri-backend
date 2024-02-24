package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

type Trip struct {
	gorm.Model
	UserID string
	PrefectureID string `json:"prefectureid"`
	Spots []Spot `gorm:"fossreignKey:TripID;constraint:OnUpdate:CASCADE"`
	Title string `json:"title"`
	StartDate time.Time `json:"startdate"`
	EndDate time.Time `json:"enddate"`
	Memo string `json:"memo"`
	ImagePath string `json:"imagepath"`
	IsPublic bool `json:"ispublic"`
}

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