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