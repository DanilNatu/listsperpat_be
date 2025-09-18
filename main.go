package main

import (
	"fmt"
	config "listsperpat/db"
	"listsperpat/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.Routes()
	fmt.Println("🚀 ListSperpart API berjalan di http://localhost:8080 ")
	r.Run(":8080")
}
