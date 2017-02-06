package main

import (
	"log"

	"gopkg.in/redis.v3"

	"github.com/millken/mktty/common"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB
var redisclient *redis.Client
var session *common.Session
var cf *common.Config
var err error

func initDb() {
	db, err = sqlx.Connect("postgres", cf.Server.Db)
	if err != nil {
		log.Fatalf("connect db server error: %s", err)
	}

	redisclient = redis.NewClient(&redis.Options{
		Addr:     cf.Server.Redis,
		Password: "",
		DB:       0,
	})
	_, err = redisclient.Ping().Result()
	if err != nil {
		log.Fatalf("connect redis server error: %s", err)
	}
}
