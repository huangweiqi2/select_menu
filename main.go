package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"select_menu/router"
)

func main() {

	r := gin.Default()
	router.Router(r)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Run err:", err)
		panic(err)
	}
}
