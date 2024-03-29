package main

import (
	"fmt"
	"net/http"
	api "s3-web-browser/server/go/controller/api"
	page "s3-web-browser/server/go/controller/page"
	"time"

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
		gpagenologin.GET("/login", page.LoginGET)
		gpagenologin.GET("/logout", page.LogoutGET)
		gpagenologin.Static("/js", "./static/js")
		gpagenologin.Static("/css", "./static/css")
		gpagenologin.Static("/favicon.ico", "./static/favicon.ico")
	}
	gpagelogin := router.Group("/")
	{
		if !setting.ServerSetting.AuthDisabled {
			gpagelogin.Use(loginFilterPage())
		}
		gpagelogin.GET("/browser", page.BrowserGET)
		gpagelogin.GET("/browser/*subpath", page.BrowserGET)
	}
	// api
	gapinologin := router.Group("/api")
	{
		gapinologin.POST("/login", api.LoginPOST)
		gapinologin.GET("/serverstatus", api.ServerstatusGET)
	}
	gapilogin := router.Group("/api")
	{
		if !setting.ServerSetting.AuthDisabled {
			gapilogin.Use(loginFilterAPI())
		}
		gapilogin.GET("/logininfo", api.LogininfoGET)
		gapilogin.GET("/profiles", api.ProfilesGET)
		gapilogin.GET("/s3dir/:profileid/*path", api.S3dirGET)
		gapilogin.GET("/s3download/:profileid/*path", api.S3downloadGET)
	}
	gapiadmin := router.Group("/api")
	{
		if !setting.ServerSetting.AuthDisabled {
			gapiadmin.Use(loginFilterAPI())
			gapiadmin.Use(permissionFilterAPI("admin"))
		}
		gapiadmin.POST("/profile", api.ProfilePOST)
		gapiadmin.PUT("/profile/:id", api.ProfilePUT)
		gapiadmin.DELETE("/profile/:id", api.ProfileDELETE)
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
		return
	}
}
func permissionFilterAPI(allowPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginInfo := session.Get(loginsession.SessionKey)

		if loginInfo == nil {
			c.Next()
			return
		}

		for _, key := range allowPermissions {
			if _, ok := loginInfo.(*loginsession.Logininfo).Permissions[key]; ok {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
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
