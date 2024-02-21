package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

type Trip struct {
	gorm.Model
	UserID string
	PrefectureID string
	Spots []Spot `gorm:"foreignKey:TripID;constraint:OnUpdate:CASCADE"`
	Title string
	StartDate time.Time
	EndDate time.Time
	Memo string
	ImagePath string
	IsPublic bool
}

func (trip *Trip) CreateTrip() (err error) {
	db := config.GetDB()
	return db.Create(trip).Error
}

func (trip *Trip) FindTripByID(id string) (err error) {
	db := config.GetDB()
	return db.First(trip, id).Error
}

func (trip *Trip) DeleteTripByID(id string) (err error) {
	db := config.GetDB()
	err = db.Delete(trip, id).Error
	return err
}