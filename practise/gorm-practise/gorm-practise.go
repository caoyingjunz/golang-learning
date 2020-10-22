package main

import (
	"fmt"
	"log"

	"golang-learning/practise/gorm-practise/db"
)

// refer to http://gorm.book.jasperxu.com

var (
	id   uint = 123456
	name      = "caoyingjun"
)

func main() {

	u := db.NewUserDB()

	//  创建 user
	if err := u.CreateUser(id, name); err != nil {
		log.Printf(err.Error())
		return
	}

	user, err := u.GetUser(name)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println(user)

}
