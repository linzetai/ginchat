package main

import (
	app "ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := app.Router()

	r.Run(":8080")
}
