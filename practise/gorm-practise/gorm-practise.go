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
	//if err := u.CreateUser(id, name); err != nil {
	//	log.Println("CreateUser", err.Error())
	//	return
	//}

	user, err := u.GetUser(name)
	if err != nil {
		log.Println("GetUser", err.Error())
		return
	}
	fmt.Println("GetUser:", user)

	users, err := u.GetUsers([]string{name})
	if err != nil {
		log.Println("GetUsers", err.Error())
		return
	}
	fmt.Println("GetUsers:", users)

	if err := u.UpdateUser(name); err != nil {
		log.Println("UpdateUser", err.Error())
		return
	}

	//if err := u.DeleleUser(name); err != nil {
	//	log.Println("DeleleUser", err.Error())
	//	return
	//}

}
