package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	id   int
	name string
	DB   *gorm.DB
)

type User struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(16);not null;"`
}

func init() {
	user := "kubeloader"
	password := "kubeloader"
	ip := "172.16.50.151"
	port := 3306
	database := "kubeloader"
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

	if !DB.HasTable(&User{}) {
		if err := DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			CreateTable(&User{}).
			Error; err != nil {
			panic(err)
		}
	}
}

func wait() {
	time.Sleep(time.Second * 2)
}

func main() {

	defer DB.Close()

	// insert data into DB
	user := &User{
		ID:   1111111,
		Name: "test",
	}
	if err := DB.Create(user).Error; err != nil {
		panic(err)
	}
	log.Print("Create")

	wait()
	// Get data from DB
	row := DB.Model(&User{}).Where("id = ?", 1111111).Select("id, name").Row()
	row.Scan(&id, &name)
	log.Println(fmt.Sprintf("%d  %s", id, name))

	wait()
	DB.Model(&User{}).Updates(User{ID: 1111111, Name: "xiaoming"})
	log.Print("Updates")

	// Get data from DB
	row = DB.Model(&User{}).Where("id = ?", 2222222).Select("id, name").Row()
	row.Scan(&id, &name)
	log.Println(fmt.Sprintf("%d  %s", id, name))
	wait()

	// delete data from DB
	if err := DB.Where(&User{ID: 1111111}).Delete(User{}).Error; err != nil {
		panic(err)
	}
	log.Print("Delete")
}
