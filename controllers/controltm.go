package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/locales/id_ID"

	// "project_go_v02/configs"
	"project_go_v02/models"

	"github.com/go-pg/pg/v10"
)

var dbtmConnect *pg.DB

func InitiateDB_tm(dbtm *pg.DB) {
	dbtmConnect = dbtm

	fmt.Println("dbtm connect = ", dbtmConnect)
}

func GetTM_Anomaly(c *gin.Context) {

	var tm_anomaly_data []models.TmTest02Tsurvobs

	// fmt.Println(tm_anomaly_data)
	// fmt.Println(dbtmConnect)

	// err := dbtmConnect.Model(&tm_anomaly_data).Select()
	err := dbtmConnect.Model((*models.TmTest02Tsurvobs)(nil)).
	Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").
	Select(&tm_anomaly_data)

	fmt.Println(tm_anomaly_data)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tm_anomaly_data,
	})
}

func GetTM_Anomaly_epochQuery(c *gin.Context) {
	var tm_anomaly_epoch []models.TmTest02Tsurvobs

	
	err := dbtmConnect.Model((*models.TmTest02Tsurvobs)(nil)).
	Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").Where("epoch_ten<?",1476554900).
	Select(&tm_anomaly_epoch)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tm_anomaly_epoch,
	})
}

func GetTM_Anomaly_epochbetweenQuery(c *gin.Context) {
	var tm_anomaly_epoch []models.TmTest02Tsurvobs

	
	err := dbtmConnect.Model((*models.TmTest02Tsurvobs)(nil)).
	Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").Where("epoch_ten>?",1476554900).Where("epoch_ten<?",1476565700).
	Select(&tm_anomaly_epoch)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tm_anomaly_epoch,
	})
}
type ParamInput struct {
	Idtm string `json:"tm_id" binding:"required"`
	EpochTenStart uint32 `json:"Epoch_start" binding:"required"`
	EpochTenEnd uint32 `json:"Epoch_end" binding:"required"`
}
func POST_requet(c *gin.Context) {

	var param ParamInput
	// fmt.Println(param)
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd
	// fmt.Println(param)
	fmt.Println("Param: ", idTM, epochstart, epochend)

	var tm_anomaly_epoch []models.TmTest02Tsurvobs

	
	err := dbtmConnect.Model((*models.TmTest02Tsurvobs)(nil)).
	Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").Where("epoch_ten>?",epochstart).Where("epoch_ten<?",epochend).
	Select(&tm_anomaly_epoch)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tm_anomaly_epoch,
	})

	
}