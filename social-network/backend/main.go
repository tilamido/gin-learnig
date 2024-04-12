package main

import (
	"social-network/queue"
	"social-network/routes"
)

func main() {
	go queue.SyncData()
	r := routes.Router()

	r.Run(":8888")
}
