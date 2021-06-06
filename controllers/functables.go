package controllers

import (
	// "fmt"
	"project_go_v02/models"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"
	// "reflect"
	
)

type tmStruct struct {
	Id int `json:"id"`
	Avg float32 `json:"avg"`
	Max float32 `json:"max"`
	Min float32 `json:"min"`
	Std float32 `json:"std"`
	Q1 float32 `json:"q1"`
	Q2 float32 `json:"q2"`
	Q3 float32 `json:"q3"`
	LostState float32 `json:"lost_state"`
	AnomalyState float32 `json:"anomaly_state"`
	UTC string `json:"utc"`
	EpochTen uint32 `json:"epoch_ten"`
}

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

// 	type tmStruct struct {
// 	Id int `json:"id"`
// 	Avg float32 `json:"avg"`
// 	Max float32 `json:"max"`
// 	Min float32 `json:"min"`
// 	Std float32 `json:"std"`
// 	Q1 float32 `json:"q1"`
// 	Q2 float32 `json:"q2"`
// 	Q3 float32 `json:"q3"`
// 	LostState float32 `json:"lost_state"`
// 	AnomalyState float32 `json:"anomaly_state"`
// 	UTC string `json:"utc"`
// 	EpochTen uint32 `json:"epoch_ten"`
// }

	var tm_anomaly_epoch []models.Tm0010010001Theos
	// type Tm0010010001Theos struct {
	// 	tmStruct
	// }
	// var tm_anomaly_epoch []models.tmStruct
	// fmt.Println("Type : ", reflect.TypeOf(tm_anomaly_epoch))
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