package routes

import (
	"net/http"
	"project_go_v02/controllers"

	"github.com/gin-gonic/gin"
)


func Routes(router *gin.Engine) {
	router.GET("/", apitest)
	router.GET("/tm", controllers.GetTMdata)
	router.GET("/tmtheos", controllers.GETtmTHEOS)
	router.GET("/tmtheosub", controllers.GETtmTHEOS_sub)
	router.GET("/tmid", controllers.GETtm_onlyid)
	router.GET("/gettele", controllers.GETtmTHEOS)
	router.GET("/sub_relation", controllers.GETtmTHEOS_sub_relation)
	router.GET("/countsat", controllers.GET_count_sat)
	router.GET("/getjson", controllers.GET_json)
	router.GET("/getjson02", controllers.GET_json02)
	router.GET("/getjson03", controllers.GET_json03)
	router.GET("/getjson04", controllers.GET_json04)
	router.GET("/getjson05", controllers.GET_json05)
	router.GET("/getjson06", controllers.GET_json06)
	router.GET("/gettmanomaly", controllers.GetTM_Anomaly)
	router.NoRoute(notFound)
}

func apitest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"messege": "API success",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"messege": "Route Not Found",
	})
	return
}