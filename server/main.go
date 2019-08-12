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
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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

	store := cookie.NewStore([]byte("s3-web-browser"))
	router.Use(sessions.Sessions("s3-web-browser", store))
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
		gpagenologin.GET("/logout", page.LogoutGET)
		gpagenologin.Static("/static", "./static")
	}
	gpagelogin := router.Group("/")
	{
		gpagelogin.Use(loginFilterPage())
		gpagelogin.GET("/browser", page.BrowserGET)
	}
	// api
	gapinologin := router.Group("/api")
	{
		gapinologin.POST("/login", api.LoginPOST)
		gapinologin.GET("/serverstatus", api.ServerstatusGET)
	}
	gapilogin := router.Group("/api")
	{
		gapilogin.Use(loginFilterAPI())
		gapilogin.GET("/profiles", api.ProfilesGET)
		gapilogin.POST("/profile", api.ProfilePOST)
		gapilogin.PUT("/profile/:id", api.ProfilePUT)
		gapilogin.GET("/s3dir/:profileid/*path", api.S3dirGET)
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

func loginFilterAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginInfo := session.Get(loginsession.SessionKey)

		if loginInfo == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func loginFilterPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginInfo := session.Get(loginsession.SessionKey)

		if loginInfo == nil {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		c.Next()
	}
}
