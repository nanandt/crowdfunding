package main

import (
	"crowdfunding/handler"
	"crowdfunding/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main(){
	dsn := "root:root@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	input := user.LoginInput{
		Email:    "diki@gmail.com",
		Password: "passwordddd",
	}
	user,err := userService.Login(input)
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(user.Name)
	fmt.Println(user.Email)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()

}

