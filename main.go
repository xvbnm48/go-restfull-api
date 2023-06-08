package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/xvbnm48/go-restfull-api/config"
	"github.com/xvbnm48/go-restfull-api/handler"
	"github.com/xvbnm48/go-restfull-api/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	// load .env file
	err := godotenv.Load(".env") // Load file .env
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConfig := config.NewDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.GetUsername(), dbConfig.GetPassword(), dbConfig.GetHost(),
		dbConfig.GetPort(), dbConfig.GetDatabaseName())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(dsn)
		log.Fatal(err.Error())
	}

	// user repo
	userRepository := user.NewRepository(db)

	// user service
	userService := user.NewService(userRepository)

	// user handler
	userHandler := handler.NewUserHandler(userService)

	// router
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailIsAvailability)
	router.Run()
}
