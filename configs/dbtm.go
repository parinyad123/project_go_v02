package configs

import (
	"log"
	"os"
	controllers "project_go_v02/controllers"

	"github.com/go-pg/pg/v10"
	// "project_go_v02/controllers"
	// models "project_go_v02/models"
)

func Connect_tm() *pg.DB {

	optstm := &pg.Options{
		User: "postgres",
		Password: "1150",
		Addr: ":5432",
		Database: "AnomalyDetection",
	}

	var dbtm *pg.DB = pg.Connect(optstm)

	if dbtm == nil {
		log.Printf("Failed to connect Telemetry db....")
		os.Exit(100)
	}

	log.Printf("Connect to Telemetry DB success ....")

	controllers.InitiateDB_tm(dbtm)


	return dbtm

}