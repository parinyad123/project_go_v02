package configs

import (
	"fmt"
	"reflect"

	"database/sql"
	_ "github.com/lib/pq"
	controllers "project_go_v02/controllers"
)


func Connect_sqlpg() *sql.DB {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1150"
		dbname   = "AnomalyDetection"
	)

	psqInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
fmt.Println("Pg: ", psqInfo)
dbsqlpg, err := sql.Open("postgres", psqInfo)

if err != nil {
	panic(err)
}

err = dbsqlpg.Ping()
if err != nil {
	panic(err)
}

fmt.Println("Successfully connected!")

fmt.Println(reflect.TypeOf(dbsqlpg))

controllers.InitiateDB_dbsql(dbsqlpg)
return dbsqlpg

}


