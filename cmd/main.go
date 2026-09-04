package main

import (
	"encore-be/pkg/network"
	"encore-be/pkg/startup"
)

func main() {
	app := network.NewApp()
	startup.StartServer(app)
}