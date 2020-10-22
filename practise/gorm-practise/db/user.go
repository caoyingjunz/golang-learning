package db

import (
	"time"

	"golang-learning/practise/gorm-practise/models"
)

type UserInterface interface {
	GetUser(name string) (user models.User, err error)
	CreateUser(id uint, name string) (err error)
	UpdateUser(id string) (err error)
	DeleleUser(id string) (err error)
}

func NewUserDB() UserInterface {
	return &UserDB{}
}

type UserDB struct{}

func (u *UserDB) GetUser(name string) (user models.User, err error) {
	if err = DB.Where("name = ?", name).First(&user).Error; err != nil {
		return
	}
	return
}

func (u *UserDB) CreateUser(id uint, name string) (err error) {
	user := models.User{
		ID:        id,
		Name:      name,
		Age:       18,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err = DB.Create(&user).Error; err != nil {
		return
	}
	return
}

func (u *UserDB) UpdateUser(name string) (err error) {
	if err = DB.Where("name = %s", name).Update(&models.User{UpdatedAt: time.Now(), Age: 19}).Error; err != nil {
		return
	}
	return
}

func (u *UserDB) DeleleUser(id string) (err error) {
	if err = DB.Where("name = %s", "caoyingjun").Delete(&models.User{}).Error; err != nil {
		return
	}
	return
}
