package db

import (
	"golang-learning/practise/gorm-practise/models"
)

type UserInterface interface {
	Get(id string)
	Create(id string) (err error)
	Update(id string) (err error)
	Delele(id string) (err error)
}

func NewUserDB() *UserInterface {
	return &UserDB{}
}

type UserDB struct{}

func (u *UserDB) Get(id string) {

}

func (u *UserDB) Create(id string) (err error) {
	return
}

func (u *UserDB) Update(id string) (err error) {
	return
}

func (u *UserDB) Delele(id string) (err error) {
	return
}
