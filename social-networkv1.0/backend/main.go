package main

import (
	"social-network/routes"
)

func main() {

	r := routes.Router()

	r.Run(":8888")

}
