package db

import (
    "github.com/go-redis/redis"
)

var RedisCli *redis.Client

func InitRdsDB(cnf string) {
    RedisCli = redis.NewClient(&redis.Options{
        Addr: cnf,
        DB:   0,
    })

    _, err := RedisCli.Ping().Result()
    if err != nil {
        panic(err)
    }
}
