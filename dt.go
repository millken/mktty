package main

import (
	"fmt"
	"log"

	"github.com/millken/mktty/common"
	"github.com/millken/mktty/dt"

	"github.com/gin-gonic/gin"
)

func dtInit(c *gin.Context) {
	var data, response gin.H
	c.Header("Access-Control-Allow-Origin", "*")
	v := c.DefaultQuery("v", "1.0")
	appKey := c.DefaultQuery("appkey", "")
	requestId := strToInt(c.DefaultQuery("requestid", "1"))
	action := c.DefaultQuery("action", "")

	log.Printf("[DEBUG] v: %s, appKey: %s, action: %s, requestId: %d", v, appKey, action, requestId)

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
		c.JSON(200, gin.H{"status": 404})
		return
	}
	key := fmt.Sprintf("dns_requestid:%s", appKey)
	if requestId == 0 {
		redisclient.Set(key, 0, 0)
	} else {
		redisclient.Incr(key)
	}
	n, err := redisclient.Get(key).Int64()
	if err != nil || int64(requestId) < n {
		c.JSON(200, gin.H{"status": 403})
		return
	}
	param := dt.Param{
		RequestId: requestId,
		AppKey:    appKey,
		Content:   c,
		Db:        db,
		Session:   session,
	}
	a, _ := act(param)
	response, err = a.Response()
	if err != nil {
		data = gin.H{
			"status": 501,
			"error":  err,
		}
	} else {
		data = gin.H{
			"status": 200,
			"data":   response,
		}
	}
	c.JSON(200, data)
}
