package main

import (
	"gogogo/config"
	"gogogo/routes"
)

func main() {
	config.Init()
	r := routes.NewRouter()
	r.Run(config.HttpPort)
}
