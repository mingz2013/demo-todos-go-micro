package datastore

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type RedisClient struct {
	Pool *redis.Pool
	host string
	port string
	db   int // select
}

func NewRedisClient(conf string) *RedisClient {
	client := &RedisClient{}

	client.Init(conf)
	return client

}

func (c *RedisClient) Init(conf string) {
	log.Println("init", conf)
	var confJs map[string]interface{}
	err := json.Unmarshal([]byte(conf), &confJs)

	if err != nil {
		log.Println(err)
		log.Fatal(err)
		return
	}

	log.Println(confJs)

	c.host = confJs["host"].(string)
	c.port = confJs["port"].(string)
	c.db = int(confJs["db"].(float64))

	log.Println(c.host, c.port, c.db)

	c.Pool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   1024,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {

			conn, err := redis.Dial("tcp", c.host+":"+c.port)
			if err != nil {
				return nil, err
			}

			log.Println("conn success...")

			// select db

			conn.Do("SELECT", c.db)

			return conn, nil
		},
	}

}

func (c *RedisClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := c.Pool.Get()
	defer conn.Close()

	return conn.Do(commandName, args...)
}

func (c *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

func (c *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

func (c *RedisClient) Close() {
	c.Pool.Close()
}
