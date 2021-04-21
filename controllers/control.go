package controllers

import (
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
			"status": http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"massage": "All data",
		"data": tms,
	})
	return
}
