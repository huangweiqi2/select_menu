package main

import "select_menu/router"

func main() {
	r := router.Router()
	r.Run(":8080")

}
