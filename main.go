package main

import (
	app "ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := app.Router()

	r.Run(":10086")
}
