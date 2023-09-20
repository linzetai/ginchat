package main

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"time"
)

func main() {

	// 迁移 schema
	utils.DB.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.LoginTime = time.Now()
	user.HeartbeatTime = time.Now()
	user.LoginOutTime = time.Now()
	user.Name = "神专"

	utils.DB.Create(user)

	fmt.Println(utils.DB.First(user, 1))

	utils.DB.Model(user).Update("PassWord", 1234)
}
