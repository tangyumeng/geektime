package main

import "example.com/hystrix-demo/server"

func main() {
	server.NewDownStreamServer(0.2).Run(":8000")
}
