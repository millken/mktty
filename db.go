package main

import (
	"log"

	"gopkg.in/redis.v3"

	"github.com/millken/mktty/common"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbdns, dbcdn, dbapp *sqlx.DB
var redisclient *redis.Client
var session *common.Session
var cf *common.Config
var err error

func initDb() {
	dbdns, err = sqlx.Connect("postgres", cf.Db.Dns)
	if err != nil {
		log.Fatalf("[ERROR] connect dns database error: %s", err)
	}

	dbapp, err = sqlx.Connect("postgres", cf.Db.App)
	if err != nil {
		log.Fatalf("[ERROR] connect app database error: %s", err)
	}

	redisclient = redis.NewClient(&redis.Options{
		Addr:     cf.Server.Redis,
		Password: "",
		DB:       0,
	})
	_, err = redisclient.Ping().Result()
	if err != nil {
		log.Fatalf("[ERROR] connect redis server error: %s", err)
	}
}
