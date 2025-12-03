package main

import (
	"exc8/client"
	"exc8/server"
	"time"
)

func main() {
	go func() {
		// start server
		err := server.StartGrpcServer()
		if err != nil {
			println(err.Error())
		}
	}()
	time.Sleep(1 * time.Second)
	// start client
	c, err := client.NewGrpcClient()
	if err != nil {
		println(err.Error())
	}
	err = c.Run()
	if err != nil {
		println(err.Error())
	}
	println("Orders complete!")
}
