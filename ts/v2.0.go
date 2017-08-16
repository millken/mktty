package ts

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/millken/mktty/dt"
)

func init() {
	TsRegister("2.0", NewVersion2)
}

type Version2 struct {
	Param
}

func NewVersion2(param Param) (Version, error) {
	return &Version2{
		param,
	}, nil
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (v *Version2) Handler() {
	var data, response gin.H
	var appKey string
	var err error
	c := v.Content
	db := v.Db
	if c.Request.Method != "POST" {
		c.JSON(200, gin.H{"status": 403, "error": "only accept POST"})
		return
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	postData, _ := url.ParseQuery(string(body))
	token := postData.Get("token")
	if token == "" {
		c.JSON(200, gin.H{"status": 403, "error": "token empty"})
		return
	}

	appId := strToInt(c.DefaultQuery("appid", "0"))

	sql := fmt.Sprintf("select key from app.users where id=%d and expire>now()", appId)
	db.QueryRow(sql).Scan(&appKey)

	if len(appKey) != 32 {
		c.JSON(200, gin.H{"status": 403, "error": "The length of the key should be 32"})
		return
	}

	postData.Del("token")

	postRaw := postData.Encode()
	hmac256 := ComputeHmac256(postRaw, appKey)

	if hmac256 != token {
		c.JSON(200, gin.H{"status": 504, "error": "token error"})
		return
	}

	action := postData.Get("action")
	act, ok := dt.Actions[action]
	if !ok {
		log.Printf("[ERROR] %s action not found", action)
		c.JSON(200, gin.H{"status": 404, "error": fmt.Sprintf("action not found: %s", action)})
		return
	}
	param := dt.Param{
		Get: postData,
		//Content:   c,
		Db: db,
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
