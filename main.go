package main

import (
	configs "project_go_v02/configs"
	
)

func main() {
	db := configs.Connect()
	// defer configs.Connect().Close()
	defer db.Close()

}