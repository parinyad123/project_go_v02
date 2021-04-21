package routes

import (
	"net/http"
	"project_go_v02/controllers"

	"github.com/gin-gonic/gin"
)


func Routes(r *gin.Engine) {
	r.GET("/", apitest)
	r.GET("/tm", controllers.GetTMdata)
	r.NoRoute(notFound)
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