package models

import (
	"fmt"
	// "log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	// configs "project_go_v02/configs"
)

type Satellite struct {
	Id string `json:"id"`
	Satellite_name string `่json:"satellite_name"`
	SubSystem []*SubSystem `pg:"rel:has-many"`
}

type SubSystem struct {
	Id string `json:"id"`
	SubSystem_name string `json:"subsystem_name"`
	Telemetry []*Telemetry `pg:"rel:has-many"`	
}

type Telemetry struct {
	Id string `json:"id"`
	Satellite_name string `json:"satellite_name"`
	Subsystem_name string `json:"subsystem_name"`
	TM_name string `json:"tm_name"`
	Description string `json:"description"`
	Image string `json:"image"`
	TM_package string `json:"tm_package"`
}

// type Telemetry struct {
// 	Id string 
// 	Satellite_name string 
// 	Subsystem_name string 
// 	TM_name string 
// 	Description string 
// 	Image string 
// 	TM_package string
// }

func CreateSchema(db *pg.DB) error {
	models := []interface{} {
		(*Satellite)(nil),
		(*SubSystem)(nil),
		(*Telemetry)(nil),
	}

	fmt.Println("in create table func")

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil { 
			fmt.Println("Error create table func")
			return err
			
		}
		fmt.Println("Done")
	}
	return nil
}