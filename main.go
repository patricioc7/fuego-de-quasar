package main

import (
	"goCore/cache"
	"goCore/server"
)

func main(){
	cache.InitCache()
	server.Init()
}
