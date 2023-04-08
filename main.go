package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-restfull-api/handler"
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

	// user repo
	userRepository := user.NewRepository(db)

	// user service
	userService := user.NewService(userRepository)
	inputUser := user.LoginInput{
		Email:    "riko@nogi.com",
		Password: "rikocantik",
	}
	userLogin, err := userService.LoginUser(inputUser)
	if err != nil {
		fmt.Println("terjadi error")
		fmt.Println(err.Error())
	}
	fmt.Println(userLogin.Name)
	fmt.Println(userLogin.Email)
	// user handler
	userHandler := handler.NewUserHandler(userService)

	// router
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
}
