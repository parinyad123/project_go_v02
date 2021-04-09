package configs

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"

	models "project_go_v02/models"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Password: "1150",
		Addr: ":5432",
		Database: "project_go_v2",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect ....")
		os.Exit(100)
	}

	log.Printf("Connect to DB success ....")

	err := models.CreateSchema(db)
	if err != nil {
		panic(err)
	}

	return db
} 