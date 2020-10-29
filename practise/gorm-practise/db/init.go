package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"golang-learning/practise/gorm-practise/models"
)

var DB *gorm.DB

func init() {
	user := "kubez"
	password := "kubez"
	ip := "103.39.211.122"
	port := 3306
	database := "kubez"
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, ip, port, database)
	// must declare the err to aviod panic: runtime error: invalid memory address or nil pointer dereferences
	var err error
	DB, err = gorm.Open("mysql", dbConnection)
	if err != nil {
		panic(err)
	}

	// set the connect pools
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	// 设置非复数表名
	//DB.SingularTable(true)

	// 检查表是否存在
	if !DB.HasTable(&models.User{}) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			CreateTable(&models.User{}).
			Error; err != nil {
			panic(err)
		}
	}

	// 增加索引
	//DB.Model(&User{}).AddIndex("idx_user_name", "name")
	// 为`name`, `age`列添加索引`idx_user_name_age`
	DB.Model(&models.User{}).AddIndex("idx_user_name_age", "name", "age")
}
