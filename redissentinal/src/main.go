package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	// 不是直接连接 redis，而是连接 sentinel
	// 这里使用 go 编写客户端，不使用 go 可以不关心细节
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"redis-sentinel1:26379", "redis-sentinel2:26379"},
	})

	for {
		client.Set("test", "jaychen", 0)
		fmt.Println(client.Get("test"))
		time.Sleep(time.Second * 3)
	}
}
