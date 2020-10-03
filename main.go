package main

import (
	"goCore/config"
	"goCore/db"
	"goCore/server"
)

func main(){
	config.InitFlagConfig()
	db.InitConnection(&config.Fconfig)
	server.Init()

}
