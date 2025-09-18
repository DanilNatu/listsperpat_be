package main

import (
	"fmt"
	config "listSparepart/db"
	"listSparepart/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.Routes()
	fmt.Println("🚀 ListSparepart API berjalan di http://localhost:8080 ")
	r.Run(":8080")
}
