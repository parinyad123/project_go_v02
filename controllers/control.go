package controllers

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"project_go_v02/models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetTMdata(c *gin.Context) {

	var tms []models.Telemetry
	err := dbConnect.Model(&tms).Select()
	if err != nil {
		log.Panicf("Error while getting all TM, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		// "status":  http.StatusOK,
		// "massage": "All data",
		"data":    tms,
	})
	return
}

func GETtmTHEOS(c *gin.Context) {
	// var tmTHEOSs []models.Telemetry
	var tmid []struct {
		Id string `json:"id"`
		TM_name string `json:"tm_name"`	
	}
	fmt.Println("Start..........................")
	// tmTHEOSs := new(models.Telemetry)
	// fmt.Println("tmTHEOS = ",tmTHEOSs)
	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name","id").
		Where("satellite_name=?", "THEOS").
		Select(&tmid)
	if err != nil {
		log.Panicf("Error while getting all TM, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}
	// fmt.Println("tmTHEOS = ",tmTHEOSs)
	fmt.Println("tmid = ",tmid)
	c.JSON(http.StatusOK, gin.H{
		// "status": http.StatusOK,
		// "massage": "All data",
		"data": tmid,
	})

	

	return
}

	

func GETtmTHEOS_sub(c *gin.Context) {
	
	type Telem struct {
		Id string `json:"id"`
		TM_name string `json:"tm_name"`
	}
	type TelemSub struct {
		Subsystem_name string `json:"subsystem_name"`
		Telem []Telem `json:"telem"`
	}
	type Sattel struct {
		Satellite_name string `่json:"satellite_name"`
		TelemSub []TelemSub `json:"telemsub"`
	}

	var tmid []struct {
		Id string `json:"id"`
		Satellite_name string `json:"satellite_name"`
		Subsystem_name string `json:"subsystem_name"`
		TM_name string `json:"tm_name"`
	}

	// value := []interface{} {
	// 	(*Telem)(nil),
	// }
	// marshalStruct, _ := json.MarshalIndent(value, "",)

	fmt.Println("Start..........................")
	// tmTHEOSs := new(models.Telemetry)
	// fmt.Println("tmTHEOS = ",tmTHEOSs)
	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name","id","satellite_name","subsystem_name").
		// Where("satellite_name=?", "THEOS").
		// Group("subsystem_name").
		Select(&tmid)

	// err = dbConnect.Model()

	if err != nil {
		log.Panicf("Error while getting all TM, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}
	// fmt.Println("tmTHEOS = ",tmTHEOSs)
	fmt.Println("tmid = ",tmid)
	c.JSON(http.StatusOK, gin.H{
		// "status": http.StatusOK,
		// "massage": "All data",
		"data": tmid,
	})

	for i := range tmid {
		fmt.Println(tmid[i].Id)
	}

	

	return
}


// func SelectDBPost(pg *pg.DB, tmTHEOSs []models.Telemetry) ([]models.Telemetry_id, error) {
//    err := dbConnect.Model(tmTHEOSs).
// 		Column("id","tm_name", "satellite_name").
// 		Where("satellite_name=?", "THEOS").
// 		Select()
//    return tmTHEOSs, err
// }



func GETtm_onlyid(c *gin.Context) {
	// var tmIds []models.Telemetry
	var id, tm_name, satellite_name string
	tmIds := new(models.Telemetry)
	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name", "satellite_name","id").
		Where("id=?", "0020010001").
		Select(&id, &tm_name, &satellite_name)
	if err != nil {
		log.Panicf("Error while getting all TM, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		// "status": http.StatusOK,
		// "massage": "All data",
		"data": tmIds,
	})
	return
}


type Tele struct {
	Id string `json:"id"`
	TM_name string `json:"tm_name"`
	Satellite_name string `json:"satellite_name"`
}

func GetTele(c *gin.Context) {
	var id, tm_name, satellite_name string
	tele := new(Tele)
	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name", "satellite_name","id").
		Where("id=?", "0020010001").
		Select(&id, &tm_name, &satellite_name)
	if err != nil {
		log.Panicf("Error while getting all TM, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		// "status": http.StatusOK,
		// "massage": "All data",
		"data": tele,
	})
	return

}

