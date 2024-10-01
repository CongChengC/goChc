package service

import (
	//"github.com/garyburd/redigo/redis"
	"time"

	"github.com/gomodule/redigo/redis" //这个是自己安装的。
)

// 定义一个全局的pool
var RedisPool *redis.Pool

func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {

	RedisPool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: idleTimeout, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip的redis
			return redis.Dial("tcp", address)
		},
	}
}
