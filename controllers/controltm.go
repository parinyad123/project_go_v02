package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/csv"


	// "reflect"

	"project_go_v02/models"

	"github.com/gin-gonic/gin"

	"github.com/go-pg/pg/v10"
)

var dbtmConnect *pg.DB

func InitiateDB_tm(dbtm *pg.DB) *pg.DB {
	dbtmConnect = dbtm

	fmt.Println("dbtm connect = ", dbtmConnect)
	return dbtmConnect
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
		Where("lost_state=?", 0).
		Select(&utc_tm)

	// avg
	var avg_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.avg").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Select(&avg_tm)

	// STD
	var std_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.std").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Select(&std_tm)

	// MAX
	var max_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.max").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Select(&max_tm)

	// MIN
	var min_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.min").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Select(&min_tm)

	// Q1
	var q1_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.q1").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Select(&q1_tm)

	// Q2
	var q2_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.q2").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Select(&q2_tm)

	// Q3
	var q3_tm []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.q3").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Select(&q3_tm)

	// utc anomaly 1
	var utc_ano1 []string
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Where("anomaly_state=?", 2).
		Select(&utc_ano1)

	// anomaly 1
	var ano1 []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.avg").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Where("anomaly_state=?", 2).
		Select(&ano1)

	// utc anomaly 2
	var utc_ano2 []string
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Where("anomaly_state=?", 4).
		Select(&utc_ano2)

	// anomaly 2
	var ano2 []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.avg").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Where("anomaly_state=?", 4).
		Select(&ano2)

	// utc anomaly 3
	var utc_ano3 []string
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Where("anomaly_state=?", 8).
		Select(&utc_ano3)

	// anomaly 3
	var ano3 []float32
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.avg").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 0).
		Where("anomaly_state=?", 8).
		Select(&ano3)

	// utc lost
	var utc_lost []string
	err = dbtmConnect.Model().
		TableExpr(idTM+" AS t").
		Column("t.utc").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?", 1).
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
		"tm_utc":      utc_tm,
		"tm_avg":      avg_tm,
		"tm_std":      std_tm,
		"tm_min":      min_tm,
		"tm_max":      max_tm,
		"tm_q1":       q1_tm,
		"tm_q2":       q2_tm,
		"tm_q3":       q3_tm,
		"tm_utc_ano1": utc_ano1,
		"tm_ano1":     ano1,
		"tm_utc_ano2": utc_ano2,
		"tm_ano2":     ano2,
		"tm_utc_ano3": utc_ano3,
		"tm_ano3":     ano3,
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
		"tm_utc":      utc_tm,
		"tm_avg":      avg_tm,
		"tm_std":      std_tm,
		"tm_min":      min_tm,
		"tm_max":      max_tm,
		"tm_q1":       q1_tm,
		"tm_q2":       q2_tm,
		"tm_q3":       q3_tm,
		"tm_utc_ano1": utc_ano1,
		"tm_ano1":     ano1,
		"tm_utc_ano2": utc_ano2,
		"tm_ano2":     ano2,
		"tm_utc_ano3": utc_ano3,
		"tm_ano3":     ano3,
		"tm_utc_lost": utc_lost,
	})

}

func POST_request_dynamic_float_slice_struct(c *gin.Context) {

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

	// var l []models.VerticalLine

	// err = dbtmConnect.Model().
	// 	TableExpr(idTM+" AS tmodel").
	// 	Column("tmodel.utc").
	// 	Where("epoch_ten>?", epochstart).
	// 	Where("epoch_ten<?", epochend).
	// 	Where("lost_state=?", 1).
	// 	Select(&l)

	// for i,_ := range l {
	// 	l[i].Tyte = "line"
	// 	l[i].X1 = l[i].Utc
	// 	l[i].Yref = "paper"
	// 	l[i].Y1 = 1
	// 	l[i].Opacity = 0.01
	// 	l[i].Line.Color = "rgb(0, 255, 153)"
	// 	l[i].Line.Width = 1.0
	// }

	var ds models.DataSlice
	var ano_state []float32

	for _, s := range tm_anomalys {
		ds.Utc_tm = append(ds.Utc_tm, s.UTC)
		ds.Avg_tm = append(ds.Avg_tm, s.Avg)
		ds.Std_tm = append(ds.Std_tm, s.Std)
		ds.Min_tm = append(ds.Min_tm, s.Min)
		ds.Max_tm = append(ds.Max_tm, s.Max)
		ds.Q1_tm = append(ds.Q1_tm, s.Q1)
		ds.Q2_tm = append(ds.Q2_tm, s.Q2)
		ds.Q3_tm = append(ds.Q3_tm, s.Q3)

		if s.LostState == 0 {
			ano_state = append(ano_state, s.AnomalyState)
			// ds.Utc_tm = append(ds.Utc_tm, s.UTC)
			// ds.Avg_tm = append(ds.Avg_tm, s.Avg)
			// ds.Std_tm = append(ds.Std_tm, s.Std)
			// ds.Min_tm = append(ds.Min_tm, s.Min)
			// ds.Max_tm = append(ds.Max_tm, s.Max)
			// ds.Q1_tm = append(ds.Q1_tm, s.Q1)
			// ds.Q2_tm = append(ds.Q2_tm, s.Q2)
			// ds.Q3_tm = append(ds.Q3_tm, s.Q3)
			if s.AnomalyState == 2 {
				ds.Utc_ano1 = append(ds.Utc_ano1, s.UTC)
				ds.Ano1 = append(ds.Ano1, s.Avg)
			} else if s.AnomalyState == 4 {
				ds.Utc_ano2 = append(ds.Utc_ano2, s.UTC)
				ds.Ano2 = append(ds.Ano2, s.Avg)
			} else if s.AnomalyState == 8 {
				ds.Utc_ano3 = append(ds.Utc_ano3, s.UTC)
				ds.Ano3 = append(ds.Ano3, s.Avg)
			}

		} else if s.LostState == 1 {
			ano_state = append(ano_state, 0)
		}
	}

	// ds.Line_lost = l
	fmt.Println("len ano = ", len(ano_state))
	fmt.Println("len date = ", len(ds.Utc_tm))
	// Create Anomaly bar
	// Constant
	NoCount_Condition := 2
	AnoCount_Condition := 3
	// Variation
	NoCount := 0
	Ano_1 := 0
	var start string = ""
	var end string = ""
	s_collect := []string{}
	e_collect := []string{}

	for k := 0; k < AnoCount_Condition; k++ {
		ano_state = append(ano_state, 0)
	}

	for i, t := range ano_state {
		// Find Start date begin
		if t != 0 && NoCount == 0 && Ano_1 == 0 && start == "" && end == "" {
			start = ds.Utc_tm[i]

			Ano_1 = 1
		} else if t != 0 && start != "" {
			Ano_1 += 1
			NoCount = 0
		} else if t == 0 && start != "" {
			NoCount += 1
		}

		if t == 0 && NoCount > NoCount_Condition && Ano_1 >= AnoCount_Condition {
			// fmt.Println("-------------",i,"------------------")
			// fmt.Println("Ano Count = ", Ano_1)
			// fmt.Println("No Count = ",NoCount)
			end = ds.Utc_tm[i-3]
			s_collect = append(s_collect, start)
			e_collect = append(e_collect, end)
			NoCount = 0
			Ano_1 = 0
			start = ""
			end = ""

		} else if t == 0 && NoCount > NoCount_Condition && Ano_1 < AnoCount_Condition {
			NoCount = 0
			Ano_1 = 0
			start = ""
			end = ""
		}
	}

	// fmt.Println("Start = ", s_collect)
	// fmt.Println("End = ", e_collect)
	// fmt.Println("Start = ", len(s_collect))
	// fmt.Println("End = ", len(e_collect))

	var collect []models.VerticalLine

	// colect[0].Tyte = "Hello"
	fmt.Println(collect)
	if len(s_collect) == len(e_collect) {
		for i := 0; i < len(s_collect); i++ {
			// fmt.Println(i)
			col := new(models.VerticalLine)
			col.Tyte = "rect"
			col.X0 = s_collect[i]
			col.Y0 = 0
			col.X1 = e_collect[i]
			col.Y1 = 1
			col.Xref = "x"
			col.Yref = "paper"
			col.Opacity = 0.7
			col.Fillcolor = "rgb(255, 128, 255)"
			col.Layer = "below"
			col.Line.Width = 0
			collect = append(collect, *col)
		}
	} else {
		fmt.Println("s_collect's length is not equal e_collect's length")
	}
	// fmt.Println("Coll= ",collect)
	ds.Ano_bar = collect

	c.JSON(http.StatusOK, gin.H{
		"data_detail": tmdetail,
		"data_tm":     ds,
	})

}

func CSVdownload(c *gin.Context) {
	var param models.ParamInput
	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd

	var file []models.CSVstruct
	err := dbtmConnect.Model().
		TableExpr(idTM+" AS tmodel").
		Column("tmodel.utc", "tmodel.epoch_ten", "tmodel.avg", "tmodel.max", "tmodel.min", "tmodel.std", "tmodel.q1", "tmodel.q2", "tmodel.q3", "tmodel.anomaly_state").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&file)

	if err != nil {
		log.Panicf("Error getting CSV , Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	var data [][]string

	header := []string{"utc", "epoch", "average", "maximum", "minimum", "standard_deviation",
	 "quartile1", "quartile2", "quartile3", "anomaly_state"}

	data = append(data, header)

	for _, rows := range file {
		row := []string{}
		row = append(row, rows.UTC, rows.Epoch_ten, rows.Avg, rows.Max, rows.Min, rows.Std, rows.Q1, rows.Q2, rows.Q3, rows.AnomalyState)
		data = append(data, row)
	}

	pathName := "D:/Backend/project_go_v02/api/"+idTM+".csv"

	csvFile, err := os.Create(pathName)

	if err != nil {
		log.Fatalf("failed creating csv file : %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()

	//fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	// header "Content-Disposition" is content to downloading file by browser 
	// Its value is attachment
	// filename is downloaded file renamed
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "csvfile.csv"))
    // c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Type", "text/csv")
	c.File(pathName)

	
}

func GET_CSVdownload(c *gin.Context) {

	idTM := c.Query("idtm")
	epochstart := c.Query("start")
	epochend := c.Query("end")

	fmt.Println(idTM, epochstart, epochend)
	var tmdetail models.Telemetry
	errName := dbConnect.Model(&tmdetail).Column("satellite_name", "tm_name").Where("id=?", idTM).Select()
	if errName != nil {
		log.Panicf("Error getting CSV detail name , Reason: %v\n", errName)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	fmt.Println(tmdetail)
	csvName := tmdetail.Satellite_name+"-"+tmdetail.TM_name+".csv"
	// fmt.Println(csvName)
	var file []models.CSVstruct
	err := dbtmConnect.Model().
		TableExpr(idTM+" AS tmodel").
		Column("tmodel.utc", "tmodel.epoch_ten", "tmodel.avg", "tmodel.max", "tmodel.min", "tmodel.std", "tmodel.q1", "tmodel.q2", "tmodel.q3", "tmodel.anomaly_state").
		Where("epoch_ten>?", epochstart).
		Where("epoch_ten<?", epochend).
		Where("lost_state=?",0).
		Select(&file)

	if err != nil {
		log.Panicf("Error getting CSV , Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	var data [][]string

	header := []string{"satellite_name","telemetry_name", "utc", "epoch", "average", "maximum", "minimum", "standard_deviation",
	 "quartile1", "quartile2", "quartile3", "anomaly_state"}

	data = append(data, header)

	for _, rows := range file {
		row := []string{}
		row = append(row, tmdetail.Satellite_name, tmdetail.TM_name, rows.UTC, rows.Epoch_ten, rows.Avg, rows.Max, rows.Min, rows.Std, rows.Q1, rows.Q2, rows.Q3, rows.AnomalyState)
		data = append(data, row)
	}

	pathName := "D:/Backend/project_go_v02/api/"+idTM+".csv"

	csvFile, err := os.Create(pathName)

	if err != nil {
		log.Fatalf("failed creating csv file : %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()

	//fmt.Sprintf("attachment; filename=%s", filename) Downloaded file renamed
	// header "Content-Disposition" is content to downloading file by browser 
	// Its value is attachment
	// filename is downloaded file renamed
	// c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "csvfile.csv"))
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(csvName))
    c.Writer.Header().Add("Content-Type", "application/octet-stream")
	// c.Writer.Header().Add("Content-Type", "text/csv")
	c.File(pathName)

	
}