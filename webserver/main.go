package main

import "webserver/internal/model/manager"

func main() {
	connManager := manager.NewConcurrentConnectionManger(8081)
	connManager.Start()
}
