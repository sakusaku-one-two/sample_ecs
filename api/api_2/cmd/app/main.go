package main

import (
	root "module/handler"
	"module/infra"
)

func main() {
	server := infra.FactoryServer()
	root.SetupRoutes(server)
	server.Logger.Fatal(server.Start(infra.SERVER_PORT))
}
