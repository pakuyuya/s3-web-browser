package main

import (
	"fmt"
	"net/http"
	"time"
	api "s3-web-browser/server/go/controller/api"
	page "s3-web-browser/server/go/controller/page"

	loginsession "s3-web-browser/server/go/domain/loginsession"
	"s3-web-browser/server/go/setting"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.setFlags(log.Lshortfile)

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
	gpagenologin := router.Group("/")
	{
		gpagenologin.GET("/", page.IndexGET)
		gpagenologin.Static("/static", "./static")
		gpagenologin.GET("/logout", api.LogoutGET)
	}
	gpagelogin := router.Group("/")
	{
		gpagelogin.Middleware(loginFilterMiddleware())
		gpagelogin.GET("/s3", page.IndexGET)
	}
	// api
	gapinologin := router.Group("/api")
	{
		gapinologin.POST("/login", api.LoginPOST)
		gapinologin.GET("/serverstatus", api.ServerstatusGET)
	}
	gapilogin := router.Group("/api")
	{
		gpagelogin.Middleware(loginFilterMiddleware())
		gpai.GET("/profiles", api.ProfilesGET)
		gpai.POST("/profile", api.ProfilesPOST)
		gpai.PUT("/profile", api.ProfilesPUT)
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

func loginFilterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginInfo := session.Get(loginsession.SessionKey)

		if loginInfo == nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		c.Next()
	}
}
