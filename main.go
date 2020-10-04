package main

import (
	"quasarFire/cache"
	"quasarFire/server"
)

func main(){
	cache.InitCache()
	server.Init()
}
