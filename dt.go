package main

import (
	"crypto/rc4"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/millken/mktty/common"
	"github.com/millken/mktty/dt"

	"github.com/gin-gonic/gin"
)

func dtInit(c *gin.Context) {
	var data, response gin.H
	var appKey string
	c.Header("Access-Control-Allow-Origin", "*")
	v := c.DefaultQuery("v", "1.0")
	appId := strToInt(c.DefaultQuery("appid", "0"))
	requestId := strToInt(c.DefaultQuery("requestid", "1"))
	action := c.DefaultQuery("action", "")
	appen := strings.Replace(c.DefaultQuery("appen", ""), " ", "+", -1)

	sql := fmt.Sprintf("select key from users where id=%d and expire>now()", appId)
	dbapp.QueryRow(sql).Scan(&appKey)
	log.Printf("[DEBUG] v: %s, appId: %d(%s), action: %s, requestId: %d, appen: %s", v, appId, appKey, action, requestId, appen)

	if len(appKey) != 32 {
		c.JSON(200, gin.H{"status": 403})
		return
	}

	if appen != "" {
		key := []byte(appKey)
		cipher, err := rc4.NewCipher(key)
		if err != nil {
			c.JSON(200, gin.H{"status": 405})
			return

		}
		data, err := base64.StdEncoding.DecodeString(appen)
		if err != nil {
			log.Printf("[ERROR]  err: %s", err)
			c.JSON(200, gin.H{"status": 500})
			return

		}
		decryptedText := make([]byte, len(data))
		cipher.XORKeyStream(decryptedText, data)
		m, _ := url.ParseQuery(string(decryptedText))
		log.Printf("%q", m)
	}

	session, err = common.NewSession(appId)
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
		c.JSON(200, gin.H{"status": 403, "requestid": n})
		return
	}
	param := dt.Param{
		RequestId: requestId,
		AppKey:    appKey,
		Content:   c,
		Dns:       dbdns,
		Cdn:       dbcdn,
		Session:   session,
	}
	a, _ := act(param)
	response, err = a.Response()
	if err != nil {
		data = gin.H{
			"status":    501,
			"error":     err,
			"requestid": n,
		}
	} else {
		data = gin.H{
			"status":    200,
			"data":      response,
			"requestid": n,
		}
	}
	c.JSON(200, data)
}
