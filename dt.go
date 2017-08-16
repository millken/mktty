package main

import (
	"fmt"
	"net/url"

	"github.com/millken/mktty/ts"

	"github.com/gin-gonic/gin"
)

func dtInit(c *gin.Context) {
	get := url.Values{}
	c.Header("Access-Control-Allow-Origin", "*")
	get, _ = url.ParseQuery(string(c.Request.URL.RawQuery))

	v := c.DefaultQuery("v", "1.0")
	ver, ok := ts.Versions[v]
	if !ok {
		c.JSON(200, gin.H{"status": 404, "error": fmt.Sprintf("v not found: %s", v)})
		return
	}
	tsparam := ts.Param{
		Url:     get,
		Content: c,
		Db:      db,
		Red:     redisclient,
	}
	av, _ := ver(tsparam)
	av.Handler()
}
