package models

import (
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

type User struct {
	Uid string `json:"uid" binding:"required" gorm:"primaryKey"`
	Trips []Trip `gorm:"foreignKey:Uid;constraint:OnUpdate:CASCADE"`
	FavoriteTrips []string `gorm:"type:text" json:"favoriteTrips"  binding:"required"`
}

func (user *User) CreateUser() (err error) {
	db := config.GetDB()
	return db.Create(user).Error
}

func (user *User) FindUserByID(userid string) (err error) {
	db := config.GetDB()
	return db.First(&user, "uid = ?",userid).Error
}

func (user *User) UpdateUserByID(userid string) (err error) {
	db := config.GetDB()
	// return db.Where("uid = ?", userid).Updates(user).Error
	return db.Save(&user).Error
}

func (user *User) DeleteUserByID() (err error) {
	db := config.GetDB()
	err = db.Delete(user).Error
	return err
}