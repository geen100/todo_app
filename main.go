package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "example@a.com"
	u.PassWord = "test"
	fmt.Println(u)

	u.CreateUser()
}
