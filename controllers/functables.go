package controllers

import (
	"github.com/gin-gonic/gin"
	"project_go_v02/models"

	"log"
	"net/http"
)

func tm0010010001(idTM string, epochstart, epochend uint32, c *gin.Context) {
	var tmdetail []models.Telemetry
	errdetail := dbConnect.Model(&tmdetail).Where("id=?", idTM).Select()

	if errdetail != nil {
		log.Panicf("Error while getting Detail of tm, Reason: %v\n", errdetail)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}

	var tm_anomaly_epoch []models.Tm0010010001Theos
	err := dbtmConnect.Model(&tm_anomaly_epoch).
	Where("epoch_ten>?",epochstart).
	Where("epoch_ten<?",epochend).
	Select()

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data_detail": tmdetail,
		"data_tm": tm_anomaly_epoch,
	})

}