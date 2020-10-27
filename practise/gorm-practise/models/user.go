package models

import "time"

type User struct {
	ID        uint   `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Name      string `gorm:"column:name;type:varchar(128);not null" json:"name"`
	Age       int    `gorm:"column:age;not null" json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "user"
}
