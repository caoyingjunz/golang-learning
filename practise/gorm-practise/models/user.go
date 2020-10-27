package models

import "time"

//CREATE TABLE `user` (
//	`id` int(128) NOT NULL AUTO_INCREMENT,
//	`name` varchar(255) DEFAULT NULL,
//	`age` int(11) DEFAULT NULL,
//	PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=123456;

type User struct {
	ID        int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Name      string `gorm:"column:name;type:varchar(128);index:idx_name;not null" json:"name"` // index 需要在看看
	Age       int    `gorm:"column:age;not null" json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "user"
}
