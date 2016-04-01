package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var r *gin.Engine
var err error

func main() {

	readConfig()
	if err = checkAndCreateDir(config.Dir); err != nil {
		log.Fatalln(err.Error())
	}
	if config.Debug {
		r = gin.Default()
	} else {
		r = gin.New()
		r.Use(gin.Recovery())
		gin.SetMode(gin.ReleaseMode)
	}

	xmlAPI := r.Group("/")

	xmlAPI.Use(checkPost()).Use(checkMime("text/xml"))
	{
		xmlAPI.POST("/notifier_api/v2/notices", createV2Notice)
	}

	bindAddress := config.Host + ":" + strconv.Itoa(config.Port)
	log.Println("Starting " + appName + " on " +
		bindAddress)

	if err := r.Run(bindAddress); err != nil {
		log.Fatalln("Error starting Gin: ", err)
	}
}
