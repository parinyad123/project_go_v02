package main

import (
	"fmt"
	"log"
	configs "project_go_v02/configs"
	"project_go_v02/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.Connect()
	fmt.Println("db main = ", db)
	// defer configs.Connect().Close()
	defer db.Close()

	router := gin.Default()

	routes.Routes(router)
	log.Fatal(router.Run(":3030"))
}