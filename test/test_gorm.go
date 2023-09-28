package main

import (
	// "fmt"
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "time"
)

func main() {
	db, err := gorm.Open(mysql.Open("linzetai:000521@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}
	// 迁移 schema
	// db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Message{})

	// user := &models.UserBasic{}
	// user.LoginTime = time.Now()
	// user.HeartbeatTime = time.Now()
	// user.LoginOutTime = time.Now()
	// user.Name = "神专"

	// utils.DB.Create(user)

	// fmt.Println(utils.DB.First(user, 1))

	// utils.DB.Model(user).Update("PassWord", 1234)
}
