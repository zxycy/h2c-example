package main

import (
	"h2c-example/server"
)

func main() {
	go server.StartServer("http")
	select {}
}
