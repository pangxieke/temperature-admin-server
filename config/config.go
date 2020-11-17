package config

import (
	"github.com/astaxie/beego"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type LogsConfig struct {
	File string
}

type OSSConfig struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
}

type MySQLConfig struct {
	URI   string
	Debug bool
}

type RedisConfig struct {
	Address string
}

var Env string
var MySQL MySQLConfig
var OSS OSSConfig
var LogConf LogsConfig
var Redis RedisConfig

func Init(configPaths ...string) (err error) {
	if err := setup(configPaths...); err != nil {
		return err
	}
	if err := initMySQL(); err != nil {
		return err
	}
	if err := initLog(); err != nil {
		return err
	}
	//if err := initOSS(); err != nil {
	//	return err
	//}
	//if err := initRedis(); err != nil {
	//	return err
	//}
	return
}

func setup(paths ...string) (err error) {
	Env = os.Getenv("GO_ENV")
	if "" == Env {
		Env = "test"
	}
	godotenv.Load(".env." + Env)
	//godotenv.Load()
	beego.BConfig.RunMode = Env
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Failed to read config file (but environment config still affected), err = %+v\n", err)
		err = nil
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return
}

func initMySQL() (err error) {
	MySQL.URI = viper.GetString("mysql.uri")
	if MySQL.URI == "" {
		return errors.New("mysql.db should not be empty")
	}
	MySQL.Debug = viper.GetBool("mysql.debug")
	return
}

func initLog() (err error) {
	LogConf.File = viper.GetString("logs.file")
	if LogConf.File == "" {
		return errors.New("logs.file should not be empty")
	}
	return
}

func initOSS() (err error) {
	OSS.Endpoint = viper.GetString("oss.endpoint")
	if OSS.Endpoint == "" {
		return errors.New("oss.endpoint should not be empty")
	}
	OSS.AccessKeyId = viper.GetString("oss.access_key_id")
	if OSS.AccessKeyId == "" {
		return errors.New("oss.access_key_id should not be empty")
	}
	OSS.AccessKeySecret = viper.GetString("oss.access_key_secret")
	if OSS.AccessKeySecret == "" {
		return errors.New("oss.access_key_secret should not be empty")
	}
	OSS.Bucket = viper.GetString("oss.bucket")
	if OSS.Bucket == "" {
		return errors.New("oss.bucket should not be empty")
	}
	return
}

func initRedis() (err error) {
	Redis.Address = viper.GetString("redis.address")
	if Redis.Address == "" {
		return errors.New("redis.address should not be empty")
	}
	return
}
