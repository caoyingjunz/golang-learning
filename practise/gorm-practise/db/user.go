package db

import (
	"time"

	"golang-learning/practise/gorm-practise/models"
)

type UserInterface interface {
	GetUser(name string) (user models.User, err error)
	CreateUser(name string) (err error)
	UpdateUser(id string) (err error)
	DeleleUser(id string) (err error)

	// 获取多条记录
	GetUsers(names []string) (users []models.User, err error)

	// 通过Raw方式执行
	GetRawUsers(names []string) (users []models.User, err error)
}

func NewUserDB() UserInterface {
	return &UserDB{}
}

type UserDB struct{}

func (u *UserDB) GetUser(name string) (user models.User, err error) {
	// Frist 获取第一个
	// Find 获取满足条件，如果只有一个返回，返回最后一个
	if err = DB.Where("name = ?", name).First(&user).Error; err != nil {
		return
	}
	return
}

func (u *UserDB) GetRawUsers(names []string) (user []models.User, err error) {
	// 如果需要，强制使用索引
	if err = DB.Raw("select * from user force index(idx_user_name_age) where name in (?)", names).Scan(&user).Error; err != nil {
		return
	}
	return
}

// names 应该是 Slices
func (u *UserDB) GetUsers(names []string) (users []models.User, err error) {
	if err = DB.Where("name in (?)", names).Find(&users).Error; err != nil {
		return
	}
	return
}

func (u *UserDB) CreateUser(name string) (err error) {
	user := models.User{
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
	if err = DB.Model(&models.User{}).
		Where("name = ?", name).
		Update(&models.User{
			UpdatedAt: time.Now(),
			Age:       19,
		}).Error; err != nil {
		return
	}
	return
}

func (u *UserDB) DeleleUser(name string) (err error) {
	if err = DB.Where("name = ?", name).Delete(&models.User{}).Error; err != nil {
		return
	}
	return
}
