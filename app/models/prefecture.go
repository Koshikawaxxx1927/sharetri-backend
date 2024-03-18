package models

import (
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

type Prefecture struct {
    ID   int    `gorm:"primaryKey";
	json:"id"`
	Trips []Trip `gorm:"foreignKey:PrefectureID;constraint:OnUpdate:CASCADE"`
    Name string `json:"name"`
    Kana string `json:"kana"`
}

type Prefectures []Prefecture

func (prefecture *Prefecture) FindPrefectureByID(prefectureid string) (err error) {
	db := config.GetDB()
	return db.First(prefecture, prefectureid).Error
}

func CreatePrefecturesBatches(prefectures []Prefecture) (err error) {
	db := config.GetDB()
	return db.CreateInBatches(prefectures, 47).Error
}

func (prefectures *Prefectures) GetAllPrefectures() (err error) {
	db := config.GetDB()
	return db.Find(&prefectures).Error
}