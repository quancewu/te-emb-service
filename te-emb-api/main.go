package main

import (
	"te-emb-api/controllers"
	"te-emb-api/initalizers"
	"te-emb-api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvVariables()
	// initalizers.ConnectToDb()
	initalizers.ConnectToSQLITE()
	initalizers.ConnectToSqliteTimeseries()
	initalizers.SyncDatabase()
	initalizers.ConnectToRedis()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Singup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/ams/:id/latest", controllers.TeAmsData)
	r.GET("/ams/:id", controllers.TeAmsDatas)
	r.POST("/ams/:id", controllers.TeAmsDataInsert)
	r.Run()
}
