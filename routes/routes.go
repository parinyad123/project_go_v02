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
	router.GET("/gettmanomaly_epoch", controllers.GetTM_Anomaly_epochQuery)
	router.GET("/gettmanomaly_epochbewteen", controllers.GetTM_Anomaly_epochbetweenQuery)
	router.POST("/postparam", controllers.POST_request)
	router.POST("/postparamv2", controllers.POST_request_v2)
	router.POST("/postsqltmstring", controllers.POST_pgTMstring)
	router.POST("/postdynamicstr",controllers.POST_request_dynamic_str)
	router.POST("/postdynamicfloat",controllers.POST_request_dynamic_float)
	router.POST("/post_data", controllers.POST_request_dynamic_slice)
	router.POST("/postfloatslice", controllers.POST_request_dynamic_float_slice)
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



