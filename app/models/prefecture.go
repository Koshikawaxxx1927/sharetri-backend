package models

import (
	// "gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

// type Prefecture struct {
//     ID   int    `gorm:"foreignKey:PrefectureID;constraint:OnUpdate:CASCADE"`
//     Name string `gorm:"not null"`
//     Kana string `gorm:"not null"`
// }

type Prefecture struct {
    ID   int    `gorm:"primaryKey";
	json:"id"`
	Trips []Trip `gorm:"foreignKey:PrefectureID;constraint:OnUpdate:CASCADE"`
    Name string `json:"name"`
    Kana string `json:"kana"`
}

func (prefecture *Prefecture) FindPrefectureByID(id string) (err error) {
	db := config.GetDB()
	return db.First(prefecture, id).Error
}

// func  (prefecture *Prefecture) ()

func CreatePrefecturesBatches(prefectures []Prefecture) (err error) {
	db := config.GetDB()
	return db.CreateInBatches(prefectures, 10).Error
}