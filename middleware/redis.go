package middleware

import (
	"fmt"
	"time"

	"github.com/daigd/v-mall-go/config"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	// 初始化Redis 连接池
	redisURL := config.Config.RedisConfig.Host + ":" + config.Config.RedisConfig.Port
	fmt.Printf("Redis URL:%q,DB:%q\n", redisURL, config.Config.RedisConfig.DB)
	pool = &redis.Pool{
		MaxIdle:     8,                 // 初始连接数
		MaxActive:   0,                 // 最大连接数，不确定可用0，按需分配
		IdleTimeout: 300 * time.Second, // 连接持续时间，300秒后关闭
		Wait:        false,
		Dial: func() (conn redis.Conn, e error) { // 要连接的数据库
			c, err := redis.Dial("tcp", redisURL)
			if err != nil {
				return nil, err
			}
			// 指定对应数据库
			if _, err := c.Do("SELECT", config.Config.RedisConfig.DB); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

// Get 获取 key 对应的字符串
func Get(key string) string {
	c := pool.Get()
	defer c.Close()
	r, err := redis.String(c.Do("GET", key))
	if err != nil {
		return ""
	}
	return r
}

// Set 设置 key 对应的值 val
func Set(key string, val string) error {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("SET", key, val)
	return err
}
