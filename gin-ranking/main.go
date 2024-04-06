package main

import (
	"gin-ranking/router"
)

func main() {
	r := router.Router()

	r.Run(":8080")
}
