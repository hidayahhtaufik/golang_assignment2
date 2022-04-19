package main

import (
	"assignment2/routers"
)

func main() {
	port := ":8881"

	server := routers.StartServer()
	server.Run(port)
}
