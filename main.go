package main

import (
	"fmt"
	"log"
	configs "project_go_v02/configs"
	// configsdb "project_go_v02/configs/dbtm"
	"project_go_v02/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.Connect()
	fmt.Println("db main = ", db)
	// fmt.Println("db tm =", dbtm)
	// defer configs.Connect().Close()
	defer db.Close()

	dbtm := configs.Connect_tm()
	// fmt.Println("db tm =", dbtm)
	defer dbtm.Close()

	dbsqlpg := configs.Connect_sqlpg()
	fmt.Println("dbsqlpg =", dbsqlpg)
	defer dbsqlpg.Close()
	
	router := gin.Default()

	routes.Routes(router)
	log.Fatal(router.Run(":3030"))
}