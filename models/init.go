package models

import (
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"log"
	"strings"
	"temp-admin/config"

	"github.com/go-redis/redis"
)

var secret []byte
var engine *xorm.Engine
var RedisCluster *redis.ClusterClient

const PageSize = 15

func Init() (err error) {
	err = initMySQL()
	if err != nil {
		return
	}
	secret, _ = hex.DecodeString("dc2dc8a96e7050c54ee5267363f9cd803912ea81")
	//err = initMongo(config.Mongo)
	//if err != nil {
	//	return
	//}
	//err = initOSS(config.OSS)
	//if err != nil {
	//	return
	//}
	//if err = initRedis(); err != nil {
	//	return
	//}
	return
}

//func initMongo(cfg config.MongoConfig) (err error) {
//	log.Printf("mongo config:%+v", cfg)
//	err = mongoInit(cfg)
//	return
//}

func initMySQL() (err error) {
	//host := config.MySQL.Host
	//
	//port := config.MySQL.Port
	//username := config.MySQL.User
	//password := config.MySQL.Password
	//name := config.MySQL.DB
	//
	//source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&autocommit=%s&timeout=%s",
	//	username, password, host, port, name, "utf8", "true", "3s")
	engine, err = xorm.NewEngine("mysql", config.MySQL.URI)
	if err != nil {
		panic(fmt.Errorf("mysql error: %v", err))
	}
	engine.SetMaxIdleConns(3)
	engine.SetMaxOpenConns(20)
	engine.ShowSQL(config.MySQL.Debug)
	log.Printf("mysql, init end ......")
	//共用层DBEngine注册
	//dal.RegDbEngine(engine, engine)
	return
}

func initRedis() (err error) {
	opt := redis.ClusterOptions{
		Addrs:    strings.Split(config.Redis.Address, ","),
		Password: "",
		PoolSize: 10,
	}
	RedisCluster = redis.NewClusterClient(&opt)

	//RedisClient = redis.NewClient(&redis.Options{
	//	Addr:     config.Redis.Address,
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})

	pong, err := RedisCluster.Ping().Result()
	if err != nil {
		return errors.Wrapf(err, "Failed to connect redis")
	}
	if pong != "PONG" {
		return errors.Wrapf(err, "Failed to ping redis")
	}
	return
}
