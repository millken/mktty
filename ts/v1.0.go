package ts

import (
	"crypto/rc4"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/millken/mktty/dt"
)

func init() {
	TsRegister("1.0", NewVersion1)
}

type Version1 struct {
	Param
}

func NewVersion1(param Param) (Version, error) {
	return &Version1{
		param,
	}, nil
}

func (v *Version1) Handler() {
	var data, response gin.H
	var appKey string
	var err error
	c := v.Content
	get := v.Url
	db := v.Db
	appId := strToInt(c.DefaultQuery("appid", "0"))
	appen := get.Get("appen")
	redisclient := v.Red
	if appen == "" {
		appen, _ = c.GetPostForm("appen")
	}
	appen = strings.Replace(appen, " ", "+", -1)

	sql := fmt.Sprintf("select key from app.users where id=%d and expire>now()", appId)
	db.QueryRow(sql).Scan(&appKey)

	if len(appKey) != 32 {
		c.JSON(200, gin.H{"status": 403, "error": "The length of the key should be 32"})
		return
	}

	if appen != "" {
		key := []byte(appKey)
		cipher, err := rc4.NewCipher(key)
		if err != nil {
			c.JSON(200, gin.H{"status": 405, "error": err})
			return

		}
		data, err := base64.StdEncoding.DecodeString(appen)
		if err != nil {
			log.Printf("[ERROR]  err: %s", err)
			c.JSON(200, gin.H{"status": 500, "error": err})
			return

		}
		decryptedText := make([]byte, len(data))
		cipher.XORKeyStream(decryptedText, data)
		//log.Printf("[DEBUG] data=%s", decryptedText)
		get, _ = url.ParseQuery(string(decryptedText))
	}

	action := get.Get("action")
	act, ok := dt.Actions[action]
	if !ok {
		log.Printf("[ERROR] %s action not found", action)
		c.JSON(200, gin.H{"status": 404, "error": fmt.Sprintf("action not found: %s", action)})
		return
	}
	key := fmt.Sprintf("dns_requestid:%s", appKey)

	requestId := strToInt(get.Get("requestid"))
	if requestId == 0 {
		redisclient.Set(key, 0, 0)
	} else {
		redisclient.Incr(key)
	}
	n, err := redisclient.Get(key).Int64()
	if err != nil || int64(requestId) != n {
		c.JSON(200, gin.H{"status": 403, "requestid": n})
		return
	}
	param := dt.Param{
		Get: get,
		//Content:   c,
		Db: db,
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
