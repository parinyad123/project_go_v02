package controllers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project_go_v02/models"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"

	"math/rand"
	"time"
)

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
	fmt.Println("db connect = ", dbConnect)
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
		"data": tms,
	})
	return
}

func GETtmTHEOS(c *gin.Context) {
	// var tmTHEOSs []models.Telemetry
	var tmid []struct {
		Id      string `json:"id"`
		TM_name string `json:"tm_name"`
	}
	fmt.Println("Start..........................")
	// tmTHEOSs := new(models.Telemetry)
	// fmt.Println("tmTHEOS = ",tmTHEOSs)
	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name", "id").
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
	fmt.Println("tmid = ", tmid)
	fmt.Println("tmid = ", tmid)
	c.JSON(http.StatusOK, gin.H{
		// "status": http.StatusOK,
		// "massage": "All data",
		"data": tmid,
	})

	return
}

func GETtmTHEOS_sub(c *gin.Context) {

	type Telem struct {
		Id      string `json:"id"`
		TM_name string `json:"tm_name"`
	}
	type TelemSub struct {
		Subsystem_name string  `json:"subsystem_name"`
		Telem          []Telem `json:"telem"`
	}
	type Sattel struct {
		Satellite_name string     `่json:"satellite_name"`
		TelemSub       []TelemSub `json:"telemsub"`
	}

	var tmid []struct {
		Id             string `json:"id"`
		Satellite_name string `json:"satellite_name"`
		Subsystem_name string `json:"subsystem_name"`
		TM_name        string `json:"tm_name"`
	}

	fmt.Println("Start..........................")
	// tmTHEOSs := new(models.Telemetry)
	// fmt.Println("tmTHEOS = ",tmTHEOSs)
	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name", "id", "satellite_name", "subsystem_name").
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
	fmt.Println("tmid = ", tmid)
	c.JSON(http.StatusOK, gin.H{
		// "status": http.StatusOK,
		// "massage": "All data",
		"data": tmid,
	})

	// for i := range tmid {
	// 	fmt.Println(tmid[i].Id)
	// }

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

	var tmonlyid []struct {
		Id             string `json:"id"`
		Satellite_name string `json:"satellite_name"`
		TM_name        string `json:"tm_name"`
	}
	// tmIds := new(models.Telemetry)
	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name", "satellite_name", "id").
		Where("id=?", "0020010001").
		// Select(&id, &tm_name, &satellite_name)
		Select(&tmonlyid)
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
		"data": tmonlyid,
	})
	return
}

func GET_count_sat(c *gin.Context) {
	var satcount []struct {
		Satellite_name   string `json:"satellite_name"`
		TelemetryCount int    `json:"sat_count"`
	}

	err := dbConnect.Model((*models.Telemetry)(nil)).
		// Column("Satellite_name").
		ColumnExpr("count(*) AS sat_count").
		// Group("satellite_name").
		Select()

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
		"data": satcount,
	})
	return
}

type Tele struct {
	Id             string `json:"id"`
	TM_name        string `json:"tm_name"`
	Satellite_name string `json:"satellite_name"`
}

func GETtmTHEOS_sub_relation(c *gin.Context) {

	type Telem struct {
		Id      string `json:"id"`
		TM_name string `json:"tm_name"`
	}
	type TelemSub struct {
		Subsystem_name string  `json:"subsystem_name"`
		Telem          []Telem `json:"telem"`
	}
	// type Sattel struct {
	// 	Satellite_name string `่json:"satellite_name"`
	// 	TelemSub []TelemSub `json:"telemsub"`
	// }

	var subtele []TelemSub
	err := dbConnect.Model((*models.SubSystem)(nil)).Relation("Telemetry").
		Column("Subsystem_name", "Id", "TM_name").
		Select(&subtele)

	// raw := dbConnect.Query('')



	if err != nil {
		log.Panicf("Error while getting all TM, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}

	fmt.Println("tmid = ", subtele)
	c.JSON(http.StatusOK, gin.H{
		"data": subtele,
	})

	return
}


func GET_json(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": `"THEOS2m_sub_3":[{"tmName":"THEOS2m_sub_30","tmid":"020300"},{"tmName":"THEOS2m_sub_31","tmid":"020301"}]`,
	})
}

func GET_json02(c * gin.Context) {

	type WeightData struct {
	Date time.Time `json:"Weight_date"`
	Kgs  float64   `json:"Weight_kg"`
}

	type AutoGenerated struct {
	StatusMessage string       `json:"Status_message"`
	Weights       []WeightData `json:"Weight_data"`
}


	mainStruct := AutoGenerated{StatusMessage: "Success"}
	
	for i := 0; i<5; i++ {
		w := WeightData{time.Now(), rand.Float64()*25}
		mainStruct.Weights = append(mainStruct.Weights, w)

		js, _ := json.MarshalIndent(mainStruct, "", "  ")
		fmt.Printf("%s\n", js)
	}

	c.JSON(http.StatusOK, gin.H{"data": mainStruct})

}


func GET_json03(c *gin.Context) {
	type Telem struct {
		Id      string `json:"id"`
		TM_name string `json:"tm_name"`
	}
	type TelemSub struct {
		Subsystem_name string  `json:"subsystem_name"`
		Telem          []Telem `json:"telem"`
	}
	type Sattel struct {
		Satellite_name string     `่json:"satellite_name"`
		TelemSub       []TelemSub `json:"telemsub"`
	}

	var tmid []struct {
		Id             string `json:"id"`
		Satellite_name string `json:"satellite_name"`
		Subsystem_name string `json:"subsystem_name"`
		TM_name        string `json:"tm_name"`
	}

	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name", "id", "satellite_name", "subsystem_name").
		Select(&tmid)

	if err != nil {
		log.Panicf("Error while getting all TM, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"data": tmid,
	// })

	sat_name := []string{}

	sat_name = append(sat_name, "THEOS")
	fmt.Println(sat_name)

	var satName []struct {
		Satellite_name string `json:"satellite_name"`
	}

	err = dbConnect.Model((*models.Satellite)(nil)).
		Column( "satellite_name").
		Select(&satName)

	fmt.Println(satName)
	fmt.Println(satName[1].Satellite_name)

	SatSub := Sattel{Satellite_name: satName[1].Satellite_name}
	// var 
	// SatSub.Satellite_name = append(SatSub.Satellite_name, satName[0].Satellite_name)
	SatSub = Sattel{Satellite_name: satName[0].Satellite_name}
	fmt.Println("SatSub = ",SatSub)


	// Add Sattellite name into ssatsub 
	var ssatsub []Sattel
	// satName is set of Sattellite name form go-pg
	for sb := range satName {
		subb := Sattel{Satellite_name: satName[sb].Satellite_name}
		fmt.Println("subb =: ",subb)
		ssatsub = append(ssatsub, subb)
	}
	fmt.Println(ssatsub)
	fmt.Println(ssatsub[1].TelemSub)
	fmt.Println(ssatsub[0].Satellite_name)

	for sbb := range ssatsub {
		fmt.Println(sbb, ssatsub[sbb].Satellite_name)
	}

	ssatsub_clone := ssatsub
	fmt.Println("ssatsub_clone = ", ssatsub_clone)
	
	c.JSON(http.StatusOK, gin.H{
		"data": ssatsub,
	})
	
	// SSub := Sattel{}
	// fmt.Println("SSub = ",SSub)

	// for sn := range satName {
	// 	fmt.Println(sn, satName[sn].Satellite_name)
	// 	SatSub.Satellite_name := append(SatSub.Satellite_name, satName[sn].Satellite_name) 
	// }
	
	// SatSub = append(SatSub.Satellite_name, )

	
	for t := range tmid {
		// SatSub.Satellite_name = 
		for s := range satName {
			if tmid[t].Satellite_name != satName[s].Satellite_name {
				fmt.Println("ok", tmid[t].Satellite_name, tmid[t].Subsystem_name)
				// sub_name := TelemSub{tmid[t].Subsystem_name}
				// SatSub.Satellite_name = append(SatSub.Satellite_name, sub_name)
			}
		}
	}

	// for i := range tmid {
	// 	fmt.Println(tmid[i].Satellite_name)
	// 	for s := range sat_name {
	// 		if 
	// 	}

	// }

	return
}

func GET_json04(c *gin.Context) {
	type Telem struct {
		Id      string `json:"id"`
		TM_name string `json:"tm_name"`
	}
	type TelemSub struct {
		// Subsystem_name string  `json:"subsystem_name"`
		Telem          []Telem `json:"telem"`
	}
	type Sattel struct {
		// Satellite_name string     `่json:"satellite_name"`
		TelemSub       []TelemSub `json:"telemsub"`
	}


	shoppingList := make(map[string]map[string]int)

	vaggiesMap := map[string]int{"onion":2, "orka":3}
	var v_name string = "veggies"
	fmt.Println(v_name)
	shoppingList[v_name] = vaggiesMap

	fruitsMap := map[string]int{"banana": 12, "apples": 5, "oranges": 3}
	shoppingList["fruits"] = fruitsMap

	fmt.Println("Shopping list categories:")
	for key := range shoppingList {
		fmt.Println("Category:", key)
		fmt.Println("Category Details:", shoppingList[key])
	}
	
	fmt.Println("Map :", shoppingList)

	js, err := json.Marshal(shoppingList)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(js))
	}

	


	c.JSON(http.StatusOK, gin.H{
		"data": shoppingList,
	})

}


func GET_json05(c *gin.Context) {

	type tm struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

	var slicemap []tm
	tmd := tm{"1","tm1"}
	slicemap = append(slicemap, tmd)
	tmd2 := tm{"2","tm2"}
	slicemap = append(slicemap, tmd2)

	// sm := make([]map[string]string)
	fmt.Println(slicemap[1].Id)

	datatmlist := []tm{{"3","tm3"},{"4","tm4"}}
	fmt.Println(datatmlist[0].Id)
	// fmt.Println(datatmlist[0].Id)
	for eachtm := range datatmlist {
		sl := tm{datatmlist[eachtm].Id,datatmlist[eachtm].Name}
		slicemap = append(slicemap, sl)
		
}

	fmt.Println(slicemap)


	// var mapslice map[string]map[string][]tm
	mapslice := make(map[string]map[string][]tm)
	fmt.Println(mapslice)

	submap := make(map[string][]tm)
	submap["subsat"] = slicemap
	fmt.Println(submap)
	fmt.Println(submap["subsat"][1])

	mapslice["sat"] = submap
	fmt.Println(mapslice)

	jsonString, err := json.Marshal(mapslice)

	if err != nil {
		fmt.Println("Error : ", err)
	}

	fmt.Println(string(jsonString))

	c.JSON(http.StatusOK, gin.H{
		"data": mapslice,
	})


}



func GET_json06(c *gin.Context) {

	type Telem struct {
		Id      string `json:"tmId"`
		TM_name string `json:"tmName"`
	}

	var datatmid []struct {
		Id             string `json:"id"`
		Satellite_name string `json:"satellite_name"`
		Subsystem_name string `json:"subsystem_name"`
		TM_name        string `json:"tm_name"`
	}

	err := dbConnect.Model((*models.Telemetry)(nil)).
		Column("tm_name", "id", "satellite_name", "subsystem_name").
		Select(&datatmid)

	if err != nil {
		log.Panicf("Error while getting all TM, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massage": "Someting went wrong",
		})
		return
	}
	

	//  create empty data from map and slice
	SatMapSlice := make(map[string]map[string][]Telem)
	// SubMapSlice := make(map[string][]Telem)

	var tmslice []Telem

	// Create key for Sattellite name
	for num_dt := range datatmid {
		Sat_Name :=  datatmid[num_dt].Satellite_name
		SatMapSlice[Sat_Name] = make(map[string][]Telem)
	}

	for num_dt := range datatmid {
		Sat_Name :=  datatmid[num_dt].Satellite_name
		Sub_name := datatmid[num_dt].Subsystem_name
		SatMapSlice[Sat_Name][Sub_name] = tmslice
	}

	for num_dt := range datatmid {
		tm := Telem{datatmid[num_dt].Id, datatmid[num_dt].TM_name}
		Sat_Name :=  datatmid[num_dt].Satellite_name
		Sub_name := datatmid[num_dt].Subsystem_name
		SatMapSlice[Sat_Name][Sub_name] = append(SatMapSlice[Sat_Name][Sub_name], tm)

	}

	// fmt.Println("SatMapSlice = ", SatMapSlice)


	// for 

	
	// for num_datatmid := range datatmid {
	// 	tmid := Telem{datatmid[num_datatmid].Id, datatmid[num_datatmid].TM_name}
	// 	tmslice := append(tmslice, tmid)
	// 	Sub_Name := datatmid[num_datatmid].Subsystem_name
	// 	SubMapSlice[Sub_Name] = tmslice
			
	// }
		// fmt.Println(num)
	


	// fmt.Println(tmslice)
	// fmt.Println(SubMapSlice)


	// fmt.Println("Map = ", SatMapSlice)


	c.JSON(http.StatusOK, gin.H{

		// "data": datatmid,
		"Data": SatMapSlice,
	})


	return
}
