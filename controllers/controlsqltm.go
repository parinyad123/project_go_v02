package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	// "reflect"

	// "strconv"

	"project_go_v02/models"
	// "project_go_v02/configs"

	"github.com/gin-gonic/gin"
	// "github.com/go-pg/pg/v10"
)


var dbsqlCon *sql.DB
// ทำการเรียกใช้ค่า dbsqlpg จาก main.go ผ่าน InitiateDB_dbsql function 
// เนื่องจาก ไม่สามารถเรียกใช้ dbsqlpg ได้โดยตรง
func InitiateDB_dbsql(dbsqlpg *sql.DB) {
	fmt.Println("func init : ",dbsqlpg)
	dbsqlCon = dbsqlpg
}

func POST_pgTMstring(c *gin.Context) {

	//  Input parameters ======================================
	var param models.ParamPost

	c.BindJSON(&param)
	idTM := param.Idtm
	epochstart := param.EpochTenStart
	epochend := param.EpochTenEnd

	fmt.Println("Param: ", idTM, epochstart, epochend)

	// Telametry detail =======================================
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
	fmt.Println("Parameter detial finish")

	//  Telemetry Query ======================================
	var (
		// sql.Null... ใช้เพื่อให้สามารถดึงข้อมูล ใน row ที่มีค่า null ออกมาได้ ถ้าไม่ใส่จะ error
		// id string
		// avg sql.NullFloat64
		// max sql.NullFloat64
		// min sql.NullFloat64
		// std sql.NullFloat64
		// q1 sql.NullFloat64
		// q2 sql.NullFloat64
		// q3 sql.NullFloat64
		// lost_state string
		// anomaly_state sql.NullInt32
		// utc time.Time
		// epoch_ten string

		id string
		avg sql.NullString
		max sql.NullString
		min sql.NullString
		std sql.NullString
		q1 sql.NullString
		q2 sql.NullString
		q3 sql.NullString
		lost_state string
		anomaly_state sql.NullString
		utc time.Time
		epoch_ten string

	)
	raw_query := "Select id, avg, max, min, std, q1, q2, q3, lost_state, anomaly_state, utc, epoch_ten from " + idTM + " where epoch_ten between " + fmt.Sprint(epochstart) + " and " + fmt.Sprint(epochend)
	// raw_query := "Select id, avg, max, min, std, q1, q2, q3, lost_state, anomaly_state, utc, epoch_ten from " + idTM + " where epoch_ten between " + fmt.Sprint(epochstart) + " and " + fmt.Sprint(epochend) + " AND lost_state = 0"
	// fmt.Println("Raw Query: ", raw_query)
	rows, err := dbsqlCon.Query(raw_query)
	// fmt.Println("pppppppppppppppppppppppppppppppp")
	// fmt.Println("func pg ", dbsqlCon)
	// rows, err := dbsqlCon.Query("Select id, avg, max, min from tm0010010001_theos")
	// fmt.Println("4444444444444")
	// rows, err := dbsqlCon.Query("Select id, avg, max, min, std from "+tablename+" where id < 50")
	if err != nil {
	log.Fatal(err)
}

	T := models.TabletmString{}
	var tmS []models.TabletmString

	for rows.Next() {
		err := rows.Scan(&id, &avg, &max, &min, &std, &q1, &q2, &q3, &lost_state, &anomaly_state, &utc, &epoch_ten)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(id, avg, max, min, std, q1, q2, q3, lost_state, anomaly_state, utc, epoch_ten)
		// fmt.Println(reflect.TypeOf(utc))
		// fmt.Println(reflect.TypeOf(anomaly_state))
		// fmt.Println(utc)
	
	// T.Id = id
	// T.Avg = fmt.Sprintf("%f",avg.Float64)
	// T.Max = fmt.Sprintf("%f",max.Float64)
	// T.Min = fmt.Sprintf("%f",min.Float64)
	// T.Std = fmt.Sprintf("%f",std.Float64)
	// T.Q1 = fmt.Sprintf("%f", q1.Float64)
	// T.Q2 = fmt.Sprintf("%f", q2.Float64)
	// T.Q3 = fmt.Sprintf("%f", q3.Float64)
	// T.LostState = lost_state
	// T.AnomalyState = strconv.Itoa(int(anomaly_state.Int32))
	// T.UTC = utc.Format("2006-01-02 15:04:05")
	// T.EpochTen = epoch_ten

	T.Id = id
	T.Avg = avg.String
	T.Max = max.String
	T.Min = min.String
	T.Std = std.String
	T.Q1 = q1.String
	T.Q2 = q2.String
	T.Q3 = q3.String
	T.LostState = lost_state
	T.AnomalyState = anomaly_state.String
	T.UTC = utc.Format("2006-01-02 15:04:05")
	T.EpochTen = epoch_ten

	tmS = append(tmS, T)
	}


	c.JSON(http.StatusOK, gin.H{
		"data_detail": tmdetail,
		"data_tm":     tmS,
	})
}
