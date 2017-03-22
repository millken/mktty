package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/millken/mktty/common"
)

var mode string

func main() {
	c := flag.String("c", "config.toml", "config path")
	flag.Parse()
	cf, err = common.LoadConfig(*c)
	if err != nil {
		log.Fatalln("read config failed, err:", err)
	}
	switch cf.Server.Mode {
	case "release":
		mode = gin.ReleaseMode
	case "debug":
		mode = gin.DebugMode
	case "test":
		mode = gin.TestMode
	default:
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	dt := gin.Default()
	dt.GET("/", dtInit)
	dt.POST("/", dtInit)

	sdt := &http.Server{
		Addr:           ":6020",
		Handler:        dt,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}
	initDb()
	sdt.ListenAndServe()
}
