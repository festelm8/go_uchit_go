package main

import (
	"log"

	"go_uchit_go/milestone5/app"
	"go_uchit_go/milestone5/conf"
	"go_uchit_go/milestone5/utils"
)


func main() {
	// load config
	config, err := conf.NewConfig("config.yaml").Load()
	utils.CheckError(err)

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
