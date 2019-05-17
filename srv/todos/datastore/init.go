package datastore

import (
	"fmt"
	"github.com/micro/go-log"
	"gopkg.in/mgo.v2"
	"os"
)

var sessionInstance *mgo.Session

func CreateSession(host string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	sessionInstance = session
	return session, nil

}

func GetSession() *mgo.Session {
	return sessionInstance
}

var redisClient *RedisClient

func CreateRedisClient() {

	defaultRedisHost := "redis"
	defaultRedisPort := "6379"

	host := os.Getenv("REDIS_HOST")

	if host == "" {
		host = defaultRedisHost
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = defaultRedisPort
	}

	db := "0"

	jsonConf := fmt.Sprintf(`{"host": "%s","port": "%s","db": %s}`, host, port, db)

	log.Log("jsonCOnf", jsonConf)

	redisClient = NewRedisClient(jsonConf)
}

func GetRedisClient() *RedisClient {
	return redisClient
}

const KEY_ID = "key_id"

func GetKeyId() int {
	id, err := redisClient.Int(redisClient.Do("incr", KEY_ID))
	if err != nil {
		id = 1000
		redisClient.Do("set", KEY_ID, id)
	}
	if id < 1000 {
		id = 1000
		redisClient.Do("set", KEY_ID, id)
	}

	return id
}
