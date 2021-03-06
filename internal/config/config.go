package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Jwt struct {
	AccessTokenExpiry  int    `json:"access_token_expiry"`
	RefreshTokenExpiry int    `json:"refresh_token_expiry"`
	Secret             string `json:"secret"`
}

type Config struct {
	DbConn        string `json:"db_conn"`
	DbMaxConn     int    `json:"db_max_conn"`
	RedisConn     string `json:"redis_conn"`
	RedisDB       int    `json:"redis_db"`
	EndpointUrl   string `json:"endpoint_url"`
	ListenAddress string `json:"listen_address"`
	Jwt           *Jwt   `json:"jwt"`
}

var Conf *Config
var isFirst bool = true

func ReadConfig(source string) (err error) {
	err = godotenv.Load("../pkg/.env")
	if err != nil {
		return
	}

	var raw []byte
	raw, err = ioutil.ReadFile(source)
	if err != nil {
		wMsg := "error reading config from file, creating new sample"
		log.Warn(wMsg)

		err = createDefaultConfig(source)
		if err != nil {
			eMsg := "error creating config file"
			log.WithError(err).Error(eMsg)
			err = errors.Wrap(err, eMsg)
			return
		}

		raw, err = ioutil.ReadFile(source)
		if err != nil {
			eMsg := "error reading config from file"
			log.WithError(err).Error(eMsg)
			err = errors.Wrap(err, eMsg)
			return
		}
	}

	err = json.Unmarshal(raw, &Conf)
	if err != nil {
		eMsg := "error parsing config from json"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		Conf = nil
		return
	}

	m := make(map[string]string)

	m["HOST"] = os.Getenv("DB_HOST")
	m["PORT"] = os.Getenv("DB_PORT")
	m["USER"] = os.Getenv("DB_USER")
	m["PASSWORD"] = os.Getenv("DB_PASSWORD")
	m["DATABASE"] = os.Getenv("DB_DATABASE")
	m["LISTEN_ADDR"] = os.Getenv("LISTEN_ADDR")
	m["REDIS_HOST"] = os.Getenv("REDIS_HOST")
	m["REDIS_PORT"] = os.Getenv("REDIS_PORT")

	Conf.DbConn = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", m["HOST"], m["PORT"], m["USER"], m["PASSWORD"], m["DATABASE"])
	Conf.ListenAddress = m["LISTEN_ADDR"]
	Conf.RedisConn = fmt.Sprintf("%v:%v", m["REDIS_HOST"], m["REDIS_PORT"])

	return
}

func createDefaultConfig(source string) (err error) {
	//static := filepath.Join(filepath.Dir(source), "static")
	c := Config{
		DbConn:        "user=test password=test dbname=test sslmode=disable",
		DbMaxConn:     20,
		RedisConn:     "127.0.0.1:6379",
		RedisDB:       0,
		EndpointUrl:   "http://127.0.0.1:8081",
		ListenAddress: "127.0.0.1:8888",
		Jwt: &Jwt{
			AccessTokenExpiry:  10,
			RefreshTokenExpiry: 20,
			Secret:             "hello",
		},
	}

	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		eMsg := "error marshall config file"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		return
	}

	err = ioutil.WriteFile(source, b, 0644)
	if err != nil {
		eMsg := "error creating config file"
		log.WithError(err).Error(eMsg)
		err = errors.Wrap(err, eMsg)
		return
	}
	return
}
