package main

import (
	"h2c/client"
	"h2c/server"
	"time"
)


func main()  {
	go server.StartServer()
	client.Client()
	time.Sleep(10000000000000)
}









