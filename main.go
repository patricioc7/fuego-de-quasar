package main

import (
	"goCore/cache"
	"goCore/config"
	"goCore/db"
	"goCore/server"
)

func main(){
	config.InitFlagConfig()
	db.InitConnection(&config.Fconfig)
	cache.InitCache()
	server.Init()

}
