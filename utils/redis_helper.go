package utils

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

var RedisC *redis.Client
func init(){
	RedisC = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               "127.0.0.1:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})
}

func RedisSet(key string, value interface{}, expiration time.Duration) {
	statusCmd := RedisC.Set(key,value,expiration)
	log.Printf("%v", statusCmd)
	//log.Println("设置值", key)
}

func RedisGet(key string) *redis.StringCmd{
	strCmd := RedisC.Get(key)
	//log.Printf("获取redis: key %v   ,value %v", key, strCmd)
	return strCmd
}

func RedisDel(key string) *redis.IntCmd{
	intCmd := RedisC.Del(key)
	//log.Printf("删除key %v", intCmd)
	return intCmd
}