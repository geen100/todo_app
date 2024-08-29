package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "example@a.com"
	// u.PassWord = "test"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(1)

	fmt.Println(u)

	u.Name = "yugen"
	u.Email = "yugen@test"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
