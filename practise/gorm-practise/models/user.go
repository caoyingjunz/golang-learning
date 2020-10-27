package models

import "time"

//CREATE TABLE `user` (
//	`id` INT NOT NULL AUTO_INCREMENT,
//	`name` VARCHAR(16) NOT NULL,
//	`age` TINYINT NOT NULL,
//	PRIMARY KEY (`id`)
//) ENGINE=InnoDB CHARSET=utf8 AUTO_INCREMENT=20110000;

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
