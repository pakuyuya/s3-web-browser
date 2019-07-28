package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"s3-web-brawser/server/go/controller/api"
	"s3-web-brawser/server/go/controller/page"

	"s3-web-brawser/server/go/setting"
)

func main() {
	var err error
	err = setting.LoadSetting()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if setting.ServerSetting.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))
	router.LoadHTMLGlob("templates/*")

	// static
	gpage := router.Group("/")
	{
		gpage.GET("/", page.IndexGET)
		gpage.Static("/static", "./static")
	}
	// api
	gapi := router.Group("/api")
	{
		gapi.GET("/serverstatus", api.ServerstatusGET)
		// TODO:
	}

	server := &http.Server{
		Addr:           ":" + setting.ServerSetting.Port,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 18,
	}
	if setting.ServerSetting.Debug {
		fmt.Println("run http server at " + server.Addr)
	}
	server.ListenAndServe()
}
