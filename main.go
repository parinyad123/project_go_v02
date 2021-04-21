package main

import (
	"fmt"
	configs "project_go_v02/configs"
)

func main() {
	db := configs.Connect()
	fmt.Println("db main = ", db)
	// defer configs.Connect().Close()
	defer db.Close()

}