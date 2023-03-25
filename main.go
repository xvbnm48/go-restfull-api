package main

import (
	"github.com/xvbnm48/go-restfull-api/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userInput := user.RegisterUserInput{}
	userInput.Name = "Nabirra gusmarlia"
	userInput.Email = "nabirra@gmail.com"
	userInput.Occupation = "idol"
	userInput.Password = "1234"
	userService.RegisterUser(userInput)
}
