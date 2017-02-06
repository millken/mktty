package main

import (
	"log"

	"github.com/millken/mktty/common"
	"github.com/millken/mktty/dt"

	"github.com/gin-gonic/gin"
)

func dtInit(c *gin.Context) {
	var response gin.H
	c.Header("Access-Control-Allow-Origin", "*")
	v := c.DefaultQuery("v", "1.0")
	appKey := c.DefaultQuery("appkey", "")
	action := c.DefaultQuery("action", "")

	log.Printf("v: %s, appKey: %s, action: %s",
		v, appKey, action)

	session, err = common.NewSession(appKey)
	if err != nil {
		log.Printf("[ERROR] create session err: %s", err)
		c.JSON(200, gin.H{"status": 500})
		return
	}
	session.SetRedis(redisclient)

	act, ok := dt.Actions[action]
	if !ok {
		log.Printf("[ERROR] %s action not found", action)
		return
	}
	param := dt.Param{
		Db:      db,
		Session: session,
	}
	a, _ := act(param)
	response, _ = a.Response()
	data := gin.H{
		"status":  200,
		"data":    response,
		"cookies": []gin.H{},
	}

	c.JSON(200, data)
}
