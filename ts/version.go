package ts

import (
	"net/url"
	"strconv"

	"gopkg.in/redis.v3"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Version interface {
	Handler()
}

type Param struct {
	Url     url.Values
	Content *gin.Context
	Db      *sqlx.DB
	Red     *redis.Client
}

var Versions = map[string]func(Param) (Version, error){}

func TsRegister(name string, versionFactory func(Param) (Version, error)) {
	if versionFactory == nil {
		panic(" versionFactory is nil")
	}
	if _, dup := Versions[name]; dup {
		panic(" Register called twice for " + name)
	}
	Versions[name] = versionFactory
}

func strToInt(str string) int {
	var i int
	i, err := strconv.Atoi(str)
	if err != nil {
		i = 0
	}
	return i
}
