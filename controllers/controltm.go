package controllers

import (
	"fmt"
	"log"
	"net/http"
	// "reflect"

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
			"status":  http.StatusInternalServerError,
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
		Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").Where("epoch_ten<?", 1476554900).
		Select(&tm_anomaly_epoch)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
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
		Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").Where("epoch_ten>?", 1476554900).Where("epoch_ten<?", 1476565700).
		Select(&tm_anomaly_epoch)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
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
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Select(&tm_anomaly_epoch)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data_detail": tmdetail,
		"data_tm":     tm_anomaly_epoch,
	})

}

// func MODEL_table(idtable string) (tm_anomaly_epoch , model_query interface{}) {
func MODEL_table(idtable string, epochstart uint32, epochend uint32) (err error, tm_anomaly_epoch interface{}) {
	switch idtable {
	case "tm0010010001":
		var tm_anomaly_epoch []models.Tm0010010001Theos
		err := dbtmConnect.Model((*models.Tm0010010001Theos)(nil)).
			Column("id", "avg", "max", "min", "std", "q1", "q2", "q3", "lost_state", "anomaly_state", "utc", "epoch_ten").
			Where("epoch_ten>?", epochstart).
			Where("epoch_ten<?", epochend).
			Select(&tm_anomaly_epoch)
		return err, tm_anomaly_epoch

		// case "tm0010010002":
		// 	var tm_anomaly_epoch []models.Tm0010010001Theos
		// 	model_query := dbtmConnect.Model((*models.Tm0010010001Theos)(nil))
		// return  tm_anomaly_epoch, model_query
	}

	return err, tm_anomaly_epoch
}

func POST_request_v2(c *gin.Context) {

	var param models.ParamInput
	// fmt.Println(param)
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd
	// fmt.Println(param)
	fmt.Println("Param: ", idTM, epochstart, epochend)
	fmt.Println("dynamic")

	switch idTM {
	case "tm0010010001":
		tm0010010001(idTM, epochstart, epochend, c)

	}
}



func POST_request_dynamic_str(c *gin.Context) {

	var param models.ParamInput
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd

	fmt.Println("Param: ", idTM, epochstart, epochend)

	// TM DETAIL
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

	// TM DATA
	var tm_anomalys []models.TabletmString

	err := dbtmConnect.Model().
		TableExpr(idTM+" AS tmodel").
		Column("tmodel.id", "tmodel.avg", "tmodel.max", "tmodel.min", "tmodel.std", "tmodel.q1", "tmodel.q2", "tmodel.q3", "tmodel.lost_state", "tmodel.anomaly_state", "tmodel.utc", "tmodel.epoch_ten").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Select(&tm_anomalys)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data_detail": tmdetail,
		"data_tm":     tm_anomalys,
	})

}

func POST_request_dynamic_float(c *gin.Context) {

	var param models.ParamInput
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd

	fmt.Println("Param: ", idTM, epochstart, epochend)

	// TM DETAIL
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

	// TM DATA
	var tm_anomalys []models.TabletmFload

	err := dbtmConnect.Model().
		TableExpr(idTM+" AS tmodel").
		Column("tmodel.id", "tmodel.avg", "tmodel.max", "tmodel.min", "tmodel.std", "tmodel.q1", "tmodel.q2", "tmodel.q3", "tmodel.lost_state", "tmodel.anomaly_state", "tmodel.utc", "tmodel.epoch_ten").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Select(&tm_anomalys)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	// fmt.Println("tm => ", tm_anomalys)
	// fmt.Println("tm0 => ", tm_anomalys[0].Avg)
	// fmt.Println("type > ",reflect.TypeOf(tm_anomalys[0]))

	

	c.JSON(http.StatusOK, gin.H{
		"data_detail": tmdetail,
		"data_tm":     tm_anomalys,
	})

}

func POST_request_dynamic_slice(c *gin.Context) {
	var param models.ParamInput
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd

	fmt.Println("Param: ", idTM, epochstart, epochend)

	// TM DETAIL
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

	// TM DATA
	// utc
	var utc_tm []string
	err := dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&utc_tm)

	// avg
	var avg_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.avg").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&avg_tm)

	// STD
	var std_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.std").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&std_tm)

	// MAX
	var max_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.max").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&max_tm)

	// MIN
	var min_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.min").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&min_tm)

	// Q1
	var q1_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.q1").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&q1_tm)

	// Q2
	var q2_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.q2").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&q2_tm)

	// Q3
	var q3_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.q3").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&q3_tm)

	// utc anomaly 1
	var utc_ano1 []string
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Where("anomaly_state=?",2).
		Select(&utc_ano1)

	// anomaly 1
	var ano1 []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.avg").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Where("anomaly_state=?",2).
		Select(&ano1)

	// utc anomaly 2
	var utc_ano2 []string
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Where("anomaly_state=?",4).
		Select(&utc_ano2)

	// anomaly 2
	var ano2 []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.avg").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Where("anomaly_state=?",4).
		Select(&ano2)

	// utc anomaly 3
	var utc_ano3 []string
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Where("anomaly_state=?",8).
		Select(&utc_ano3)

	// anomaly 3
	var ano3 []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.avg").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Where("anomaly_state=?",8).
		Select(&ano3)

	// utc lost
	var utc_lost []string
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",1).
		Select(&utc_lost)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data_detail": tmdetail,
		"tm_utc": utc_tm,
		"tm_avg": avg_tm,
		"tm_std": std_tm,
		"tm_min": min_tm,
		"tm_max": max_tm,
		"tm_q1": q1_tm,
		"tm_q2": q2_tm,
		"tm_q3": q3_tm,
		"tm_utc_ano1": utc_ano1,
		"tm_ano1": ano1,
		"tm_utc_ano2": utc_ano2,
		"tm_ano2": ano2,
		"tm_utc_ano3": utc_ano3,
		"tm_ano3": ano3,
		"tm_utc_lost": utc_lost,
	})

}

func POST_request_dynamic_float_slice(c *gin.Context) {

	var param models.ParamInput
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd

	fmt.Println("Param: ", idTM, epochstart, epochend)

	// TM DETAIL
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

	// TM DATA
	var tm_anomalys []models.TabletmFload

	err := dbtmConnect.Model().
		TableExpr(idTM+" AS tmodel").
		Column("tmodel.id", "tmodel.avg", "tmodel.max", "tmodel.min", "tmodel.std", "tmodel.q1", "tmodel.q2", "tmodel.q3", "tmodel.lost_state", "tmodel.anomaly_state", "tmodel.utc", "tmodel.epoch_ten").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Select(&tm_anomalys)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	// fmt.Println("tm => ", tm_anomalys)
	// fmt.Println("tm0 => ", tm_anomalys[0].Avg)
	// fmt.Println("type > ",reflect.TypeOf(tm_anomalys[0]))
	var utc_tm []string
	var avg_tm []float32
	var std_tm []float32
	var min_tm []float32
	var max_tm []float32
	var q1_tm []float32
	var q2_tm []float32
	var q3_tm []float32
	var utc_ano1 []string
	var ano1 []float32
	var utc_ano2 []string
	var ano2 []float32
	var utc_ano3 []string
	var ano3 []float32
	var utc_lost []string

	for _, s := range tm_anomalys {
		// fmt.Println(s.EpochTen)
		if s.LostState == 0 {
			utc_tm = append(utc_tm, s.UTC)
			avg_tm = append(avg_tm, s.Avg)
			std_tm = append(std_tm, s.Std)
			min_tm = append(min_tm, s.Min)
			max_tm = append(max_tm, s.Max)
			q1_tm = append(q1_tm, s.Q1)
			q2_tm = append(q2_tm, s.Q2)
			q3_tm = append(q3_tm, s.Q3)
			if s.AnomalyState == 2 {
				utc_ano1 = append(utc_ano1, s.UTC)
				ano1 = append(ano1, s.Avg)
			} else if s.AnomalyState == 4 {
				utc_ano2 = append(utc_ano2, s.UTC)
				ano2 = append(ano2, s.Avg)
			} else if s.AnomalyState == 8 {
				utc_ano3 = append(utc_ano3, s.UTC)
				ano3 = append(ano3, s.Avg)
			}

		} else if s.LostState == 1 {
			utc_lost = append(utc_lost, s.UTC)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data_detail": tmdetail,
		"tm_utc": utc_tm,
		"tm_avg": avg_tm,
		"tm_std": std_tm,
		"tm_min": min_tm,
		"tm_max": max_tm,
		"tm_q1": q1_tm,
		"tm_q2": q2_tm,
		"tm_q3": q3_tm,
		"tm_utc_ano1": utc_ano1,
		"tm_ano1": ano1,
		"tm_utc_ano2": utc_ano2,
		"tm_ano2": ano2,
		"tm_utc_ano3": utc_ano3,
		"tm_ano3": ano3,
		"tm_utc_lost": utc_lost,
	})

}

