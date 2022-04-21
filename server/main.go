package main

import (
	"authapp/packages/api"
	"authapp/packages/config"
)

func main() {
	config.InitilizeConfig()
	api.StartServer()
}