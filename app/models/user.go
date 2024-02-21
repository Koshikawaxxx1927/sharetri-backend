package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/Koshikawaxxx1927/sharetri-backend/config"
)

type User struct {
	gorm.Model
	Trips []Trip `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
	Name string `json:"name"`
	IconPath string `json:"iconPath"`
	LastLoginTime time.Time `json:lastLoginTate`
}

func (user *User) CreateUser() (err error) {
	db := config.GetDB()
	return db.Create(user).Error
}

func (user *User) FindUserByID(id string) (err error) {
	db := config.GetDB()
	return db.First(user, id).Error
}

func (user *User) DeleteUserByID(id string) (err error) {
	db := config.GetDB()
	err = db.Delete(user, id).Error
	return err
}