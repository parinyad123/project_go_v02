package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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



func POST_request(c *gin.Context) {

	var param models.ParamInput
	// fmt.Println(param)
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd
	// fmt.Println(param)
	fmt.Println("Param: ", idTM, epochstart, epochend)

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

	// fmt.Println("Type : ", reflect.TypeOf(tm_anomaly_epoch), " and ", reflect.TypeOf(mod), reflect.TypeOf(tm_ano))
	// fmt.Println(MODEL_table(idTM))
	err := dbtmConnect.Model((*models.Tm0010010001Theos)(nil)).
	Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").
	Where("epoch_ten>?",epochstart).
	Where("epoch_ten<?",epochend).
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
		"data_detail": tmdetail,
		"data_tm": tm_anomaly_epoch,
	})

	
}



// func MODEL_table(idtable string) (tm_anomaly_epoch , model_query interface{}) {
func MODEL_table(idtable string, epochstart uint32, epochend uint32) (err error, tm_anomaly_epoch interface{}) {
	switch idtable {
	case "tm0010010001":
		var tm_anomaly_epoch []models.Tm0010010001Theos
		err := dbtmConnect.Model((*models.Tm0010010001Theos)(nil)).
		Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").
		Where("epoch_ten>?",epochstart).
		Where("epoch_ten<?",epochend).
		Select(&tm_anomaly_epoch)
	return err, tm_anomaly_epoch

	// case "tm0010010002":
	// 	var tm_anomaly_epoch []models.Tm0010010001Theos
	// 	model_query := dbtmConnect.Model((*models.Tm0010010001Theos)(nil))
	// return  tm_anomaly_epoch, model_query
}

	return err, tm_anomaly_epoch
}






func POST_request_version2(c *gin.Context) {

	var param models.ParamInput
	// fmt.Println(param)
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd
	// fmt.Println(param)
	fmt.Println("Param: ", idTM, epochstart, epochend)

	// if idTM == "tm0010010001" {
	// 	tm0010010001(idTM, epochstart, epochend, c)
	// }

	switch idTM {
	case "tm0010010001":
		tm0010010001(idTM, epochstart, epochend, c)
	

	}

	

	
}